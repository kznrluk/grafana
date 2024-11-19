package sql

import (
	"testing"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/stretchr/testify/require"
)

func TestQueryFramesInto(t *testing.T) {
	db := NewInMemoryDB()

	tests := []struct {
		name         string
		query        string
		input_frames []*data.Frame
		expected     *data.Frame
	}{
		{
			name:         "valid query with no input frames, one row one column",
			query:        `SELECT '1' AS 'n';`,
			input_frames: []*data.Frame{},
			expected: data.NewFrame(
				"sqlExpressionRefId",
				data.NewField("n", nil, []string{"1"}),
			),
		},
		{
			name:         "valid query with no input frames, one row two columns",
			query:        `SELECT 'sam' AS 'name', 40 AS 'age';`,
			input_frames: []*data.Frame{},
			expected: data.NewFrame(
				"sqlExpressionRefId",
				data.NewField("name", nil, []string{"sam"}),
				data.NewField("age", nil, []int64{40}),
			),
		},
		{
			// TODO: Also ORDER BY to ensure the order is preserved
			name:  "query all rows from single input frame",
			query: `SELECT * FROM inputFrameRefId LIMIT 1;`,
			input_frames: []*data.Frame{
				data.NewFrame(
					"inputFrameRefId",
					//nolint:misspell
					data.NewField("OSS Projects with Typos", nil, []string{"Garfana", "Pormetheus"}),
				),
			},
			expected: data.NewFrame(
				"sqlExpressionRefId",
				data.NewField("OSS Projects with Typos", nil, []string{"Garfana"}),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var frame data.Frame
			err := db.QueryFramesInto("sqlExpressionRefId", tt.query, tt.input_frames, &frame)
			require.NoError(t, err)
			require.NotNil(t, frame.Fields)

			require.Equal(t, tt.expected.Name, frame.Name)
			require.Equal(t, len(tt.expected.Fields), len(frame.Fields))
			for i := range tt.expected.Fields {
				require.Equal(t, tt.expected.Fields[i].Name, frame.Fields[i].Name)
				require.Equal(t, tt.expected.Fields[i].At(0), frame.Fields[i].At(0))
			}
		})
	}
}
