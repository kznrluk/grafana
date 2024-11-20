package sql

import (
	"context"
	"errors"
	"fmt"
	"io"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	mysql "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var dbName = "mydb"

type DB struct {
	inMemoryDb *memory.Database
}

func (db *DB) TablesList(rawSQL string) ([]string, error) {
	return nil, errors.New("not implemented")
}

func (db *DB) RunCommands(commands []string) (string, error) {
	return "", errors.New("not implemented")
}

// TODO: Should this accept a row limit and converters, like sqlutil.FrameFromRows?
func convertToDataFrame(iter mysql.RowIter, schema mysql.Schema, f *data.Frame) error {
	// Create fields based on the schema
	for _, col := range schema {
		var field *data.Field
		// switch col.Type.Type() {
		switch colType := col.Type.(type) {
		// NumberType represents all integer and floating point types
		// TODO: branch between int and float
		case mysql.NumberType:
			field = data.NewField(col.Name, nil, []int64{})
		// StringType represents all string types, including VARCHAR and BLOB.
		case mysql.StringType:
			field = data.NewField(col.Name, nil, []string{})
		// TODO: Implement the following types
		// DatetimeType represents DATE, DATETIME, and TIMESTAMP.
		// YearType represents the YEAR type.
		// SetType represents the SET type.
		// EnumType represents the ENUM type.
		// DecimalType represents the DECIMAL type.
		// Also the NullType (and DeferredType) ?

		// case int8:
		// 	field = data.NewField(col.Name, nil, []int64{})
		default:
			return fmt.Errorf("unsupported type for column %s: %v", col.Name, colType)
		}
		f.Fields = append(f.Fields, field)
	}

	// Iterate through the rows and append data to fields
	for {
		// TODO: Use a more appropriate context
		row, err := iter.Next(mysql.NewEmptyContext())
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading row: %v", err)
		}

		for i, val := range row {
			switch v := val.(type) {
			// TODO: The types listed here should be the same as that
			// used when creating the fields. Am I using the wrong fields
			// from the schema instance?
			case int8:
				f.Fields[i].Append(int64(v))
			case int64:
				f.Fields[i].Append(v)
			case float64:
				f.Fields[i].Append(v)
			case string:
				f.Fields[i].Append(v)
			case bool:
				f.Fields[i].Append(v)
			// Add more types as needed
			default:
				return fmt.Errorf("unsupported value type for column %s: %T", schema[i].Name, v)
			}
		}
	}

	return nil
}

// func (db *DB) writeDataframeToDb(name string, frame *data.Frame) error {
// 	// TODO: Check these details:
// 	// - Do we need a primary key?
// 	// - Can we omit `Nullable` and `Source`?
// 	table := memory.NewTable(db.inMemoryDb, name, gomysql.NewPrimaryKeySchema(gomysql.Schema{
// 		// https://pkg.go.dev/github.com/dolthub/go-mysql-server/sql#Column
// 		{Name: "name", Type: types.Text, Nullable: false, Source: name},
// 		{Name: "profession", Type: types.Text, Nullable: false, Source: name},
// 	}), nil)

// 	db.inMemoryDb.AddTable(name, table)
// 	// TODO: Use a more appropriate context
// 	err := table.Insert(gomysql.NewEmptyContext(), gomysql.NewRow("sam", "engineer"))
// 	if err != nil {
// 		return fmt.Errorf("error inserting row: %v", err)
// 	}

// 	return nil
// }

// TODO: Check if it really makes sense to receive a shared context here, rather than creating a new one
func (db *DB) writeDataframeToDb(ctx *mysql.Context, tableName string, frame *data.Frame) error {
	if frame == nil {
		return fmt.Errorf("input frame is nil")
	}

	// Create schema based on frame fields
	schema := make(mysql.Schema, len(frame.Fields))
	for i, field := range frame.Fields {
		schema[i] = &mysql.Column{
			Name:     field.Name,
			Type:     convertDataType(field.Type()),
			Nullable: true,
			Source:   tableName,
		}
	}

	// Create table with the dynamic schema
	table := memory.NewTable(db.inMemoryDb, tableName, mysql.NewPrimaryKeySchema(schema), nil)
	db.inMemoryDb.AddTable(tableName, table)

	// Insert data from the frame
	for i := 0; i < frame.Rows(); i++ {
		row := make(mysql.Row, len(frame.Fields))
		for j, field := range frame.Fields {
			row[j] = field.At(i)
		}
		err := table.Insert(ctx, row)
		if err != nil {
			return fmt.Errorf("error inserting row %d: %v", i, err)
		}
	}

	return nil
}

// Helper function to convert data.FieldType to types.Type
func convertDataType(fieldType data.FieldType) mysql.Type {
	switch fieldType {
	case data.FieldTypeInt8, data.FieldTypeInt16, data.FieldTypeInt32, data.FieldTypeInt64:
		return types.Int64
	case data.FieldTypeUint8, data.FieldTypeUint16, data.FieldTypeUint32, data.FieldTypeUint64:
		return types.Uint64
	case data.FieldTypeFloat32, data.FieldTypeFloat64:
		return types.Float64
	case data.FieldTypeString:
		return types.Text
	case data.FieldTypeBool:
		return types.Boolean
	case data.FieldTypeTime:
		return types.Timestamp
	default:
		return types.JSON
	}
}

func (db *DB) QueryFramesInto(tableName string, query string, frames []*data.Frame, f *data.Frame) error {
	// TODO: Consider if this should be moved outside of this function
	// or indeed into convertToDataFrame
	f.Name = tableName

	pro := memory.NewDBProvider(db.inMemoryDb)
	session := memory.NewSession(mysql.NewBaseSession(), pro)
	ctx := mysql.NewContext(context.Background(), mysql.WithSession(session))

	for _, frame := range frames {
		err := db.writeDataframeToDb(ctx, tableName, frame)
		if err != nil {
			return err
		}
	}

	// Select the database in the context
	ctx.SetCurrentDatabase(dbName)

	// TODO: Check if it's wise to reuse the existing provider, rather than creating a new one
	engine := sqle.NewDefault(pro)
	// engine := sqle.NewDefault(
	// 	mysql.NewDatabaseProvider(
	// 		db.inMemoryDb,
	// 	))

	schema, iter, _, err := engine.Query(ctx, query)
	if err != nil {
		return err
	}

	// rowLimit := int64(1000) // TODO - set the row limit

	// // converters := sqlutil.ConvertersFromSchema(f.RefID, f.Fields)
	// // Use nil converters for now
	// var converters []sqlutil.Converter

	// rows := sqlutil.NewRowIter(mysqlRows, nil)
	// frame, err := sqlutil.FrameFromRows(rows, rowLimit, converters...)

	err = convertToDataFrame(iter, schema, f)
	if err != nil {
		return err
	}

	return nil
}

func NewInMemoryDB() *DB { // TODO - name the function. The InMemoryDB name is now used on line 13
	return &DB{
		inMemoryDb: memory.NewDatabase(dbName), // TODO - change the name of the database
	}
}
