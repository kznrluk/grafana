package dashboard

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"

	"github.com/grafana/grafana/pkg/apimachinery/identity"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/storage/unified/resource"

	dashboardv1alpha1 "github.com/grafana/grafana/pkg/apis/dashboard/v1alpha1"
)

// The DTO returns everything the UI needs in a single request
type SearchConnector struct {
	newFunc func() runtime.Object
	client  resource.ResourceIndexClient
	log     log.Logger
}

func NewSearchConnector(
	client resource.ResourceIndexClient,
	newFunc func() runtime.Object,
) (rest.Storage, error) {
	v := &SearchConnector{
		client:  client,
		newFunc: newFunc,
		log:     log.New("grafana-apiserver.dashboards.search"),
	}
	return v, nil
}

var (
	_ rest.Connecter            = (*SearchConnector)(nil)
	_ rest.StorageMetadata      = (*SearchConnector)(nil)
	_ rest.Scoper               = (*SearchConnector)(nil)
	_ rest.SingularNameProvider = (*SearchConnector)(nil)
)

func (s *SearchConnector) New() runtime.Object {
	return s.newFunc()
}

func (s *SearchConnector) Destroy() {
}

func (s *SearchConnector) NamespaceScoped() bool {
	return true // namespace == org
}

func (s *SearchConnector) GetSingularName() string {
	return "Search"
}

func (s *SearchConnector) ConnectMethods() []string {
	return []string{"GET"}
}

func (s *SearchConnector) NewConnectOptions() (runtime.Object, bool, string) {
	return nil, false, ""
}

func (s *SearchConnector) ProducesMIMETypes(verb string) []string {
	return nil
}

func (s *SearchConnector) ProducesObject(verb string) interface{} {
	return s.newFunc()
}

func (s *SearchConnector) Connect(ctx context.Context, name string, opts runtime.Object, responder rest.Responder) (http.Handler, error) {
	user, err := identity.GetRequester(ctx)
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queryParams, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			responder.Error(err)
			return
		}

		// get limit and offset from query params
		limit := 50
		offset := 0
		if queryParams.Has("limit") {
			limit, _ = strconv.Atoi(queryParams.Get("limit"))
		}
		if queryParams.Has("offset") {
			offset, _ = strconv.Atoi(queryParams.Get("offset"))
		}

		searchRequest := &resource.ResourceSearchRequest{
			Options: &resource.ListOptions{
				Key: &resource.ResourceKey{
					Namespace: user.GetNamespace(),
					Group:     dashboardv1alpha1.GROUP,
					Resource:  "dashboards",
				},
			},
			Query:  queryParams.Get("query"),
			Limit:  int64(limit),
			Offset: int64(offset),
		}

		// TODO... actually query
		result, err := s.client.Search(r.Context(), searchRequest)
		if err != nil {
			responder.Error(err)
			return
		}

		t, err := result.Results.ToK8s()
		if err != nil {
			responder.Error(err)
			return
		}

		responder.Object(200, &t)

		// jj, err := json.Marshal(result)
		// if err != nil {
		// 	responder.Error(err)
		// 	return
		// }
		// _, _ = w.Write(jj)
	}), nil
}
