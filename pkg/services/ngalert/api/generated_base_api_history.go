/*Package api contains base API implementation of unified alerting
 *
 *Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 *
 *Do not manually edit these files, please find ngalert/api/swagger-codegen/ for commands on how to generate them.
 */
package api

import (
	"net/http"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/middleware/requestmeta"
	contextmodel "github.com/grafana/grafana/pkg/services/contexthandler/model"
	"github.com/grafana/grafana/pkg/services/ngalert/metrics"
)

type HistoryApi interface {
	RouteGetStateHistory(*contextmodel.ReqContext) response.Response
}

func (f *HistoryApiHandler) RouteGetStateHistory(ctx *contextmodel.ReqContext) response.Response {
	return f.handleRouteGetStateHistory(ctx)
}

func (api *API) RegisterHistoryApiEndpoints(srv HistoryApi, m *metrics.API) {
	api.RouteRegister.Group("", func(group routing.RouteRegister) {
		group.Get(
			toMacaronPath("/api/v1/rules/history"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodGet, "/api/v1/rules/history"),
			api.block(http.MethodGet, "/api/v1/rules/history"),
			metrics.Instrument(
				http.MethodGet,
				"/api/v1/rules/history",
				api.Hooks.Wrap(srv.RouteGetStateHistory),
				m,
			),
		)
	}, middleware.ReqSignedIn)
}
