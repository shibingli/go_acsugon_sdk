package api

import (
	ct "acsugon_sdk/constant"
	"acsugon_sdk/internal/fasttemplate"
	"net/http"
)

type API struct {
	Path        string
	Method      string
	ContentType string
}

var (
	Tokens      = &API{Path: "/ac/openapi/v2/tokens", Method: http.MethodPost, ContentType: ct.ContentTypeJSON}
	TokensNext  = &API{Path: "/ac/openapi/v2/tokens/next", Method: http.MethodPost, ContentType: ct.ContentTypeJSON}
	TokensState = &API{Path: "/ac/openapi/v2/tokens/state", Method: http.MethodGet}

	Center = &API{Path: "/ac/openapi/v2/center", Method: http.MethodGet}

	User            = &API{Path: "/ac/openapi/v2/user", Method: http.MethodGet}
	UserMember      = &API{Path: "/ac/openapi/v2/user/member", Method: http.MethodPost}
	QuotaUserMember = &API{Path: "/ac/openapi/v2/quota/user-member", Method: http.MethodPost}
	GroupMembers    = &API{Path: "/ac/openapi/v2/groupmembers", Method: http.MethodGet}

	Cluster          = &API{Path: "/hpc/openapi/v2/cluster", Method: http.MethodGet}
	ClusterGroupJobs = &API{Path: "/ac/openapi/v2/clusters/{clusterId}/groups/{groupId}/clusterUserNames/{clusterUserName}/jobs", Method: http.MethodGet}

	AppTemplatesJob = &API{Path: "/hpc/openapi/v2/apptemplates/{apptype}/{appname}/job", Method: http.MethodPost, ContentType: ct.ContentTypeJSON}
	Jobs            = &API{Path: "/hpc/openapi/v2/jobs", Method: http.MethodGet}
	HistoryJobs     = &API{Path: "/hpc/openapi/v2/historyjobs", Method: http.MethodGet}
	Job             = &API{Path: "/hpc/openapi/v2/jobs/{jobId}", Method: http.MethodGet}
	HistoryJob      = &API{Path: "/hpc/openapi/v2/historyjobs/{jobmanagerId}/{jobId}", Method: http.MethodGet}
)

func ParseTemplate(tpl string, params map[string]interface{}) string {
	t := fasttemplate.New(tpl, "{", "}")

	return t.ExecuteString(params)
}
