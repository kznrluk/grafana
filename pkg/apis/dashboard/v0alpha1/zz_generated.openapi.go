//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by openapi-gen. DO NOT EDIT.

package v0alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationActions":       schema_pkg_apis_dashboard_v0alpha1_AnnotationActions(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationPermission":    schema_pkg_apis_dashboard_v0alpha1_AnnotationPermission(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.Dashboard":               schema_pkg_apis_dashboard_v0alpha1_Dashboard(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardAccess":         schema_pkg_apis_dashboard_v0alpha1_DashboardAccess(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardHit":            schema_pkg_apis_dashboard_v0alpha1_DashboardHit(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardList":           schema_pkg_apis_dashboard_v0alpha1_DashboardList(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardVersionInfo":    schema_pkg_apis_dashboard_v0alpha1_DashboardVersionInfo(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardVersionList":    schema_pkg_apis_dashboard_v0alpha1_DashboardVersionList(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardWithAccessInfo": schema_pkg_apis_dashboard_v0alpha1_DashboardWithAccessInfo(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.FacetResult":             schema_pkg_apis_dashboard_v0alpha1_FacetResult(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.LibraryPanel":            schema_pkg_apis_dashboard_v0alpha1_LibraryPanel(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.LibraryPanelList":        schema_pkg_apis_dashboard_v0alpha1_LibraryPanelList(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.LibraryPanelSpec":        schema_pkg_apis_dashboard_v0alpha1_LibraryPanelSpec(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.LibraryPanelStatus":      schema_pkg_apis_dashboard_v0alpha1_LibraryPanelStatus(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.SearchResults":           schema_pkg_apis_dashboard_v0alpha1_SearchResults(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.SortBy":                  schema_pkg_apis_dashboard_v0alpha1_SortBy(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.SortableFields":          schema_pkg_apis_dashboard_v0alpha1_SortableFields(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.TermFacet":               schema_pkg_apis_dashboard_v0alpha1_TermFacet(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.VersionsQueryOptions":    schema_pkg_apis_dashboard_v0alpha1_VersionsQueryOptions(ref),
	}
}

func schema_pkg_apis_dashboard_v0alpha1_AnnotationActions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"canAdd": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canEdit": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canDelete": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
				},
				Required: []string{"canAdd", "canEdit", "canDelete"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_AnnotationPermission(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"dashboard": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationActions"),
						},
					},
					"organization": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationActions"),
						},
					},
				},
				Required: []string{"dashboard", "organization"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationActions"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_Dashboard(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "The dashboard body (unstructured for now)",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardAccess(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Information about how the requesting user can use a given dashboard",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"slug": {
						SchemaProps: spec.SchemaProps{
							Description: "Metadata fields",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"canSave": {
						SchemaProps: spec.SchemaProps{
							Description: "The permissions part",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"canEdit": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canAdmin": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canStar": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canDelete": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"annotationsPermissions": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationPermission"),
						},
					},
				},
				Required: []string{"canSave", "canEdit", "canAdmin", "canStar", "canDelete", "annotationsPermissions"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationPermission"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardHit(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Dashboard or folder\n\nPossible enum values:\n - `\"dash\"`\n - `\"folder\"`",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"dash", "folder"},
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "The UID",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"title": {
						SchemaProps: spec.SchemaProps{
							Description: "The display nam",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tags": {
						SchemaProps: spec.SchemaProps{
							Description: "Filter tags",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"folder": {
						SchemaProps: spec.SchemaProps{
							Description: "The UID/name for the folder",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"sorted": {
						SchemaProps: spec.SchemaProps{
							Description: "Current sorting supports sort by name, stats and date Name does not need to be returned, and the others can be numbers",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"score": {
						SchemaProps: spec.SchemaProps{
							Description: "When using \"real\" search, this is the score",
							Type:        []string{"number"},
							Format:      "double",
						},
					},
					"extra": {
						SchemaProps: spec.SchemaProps{
							Description: "Untyped extra fields/values, useful for dynamic development, but do not count on them",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
					"explain": {
						SchemaProps: spec.SchemaProps{
							Description: "Explain the score (if possible)",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
				},
				Required: []string{"type", "name", "title"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.Dashboard"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.Dashboard", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardVersionInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "The internal ID for this version (will be replaced with resourceVersion)",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"parentVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "If the dashboard came from a previous version, it is set here",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"created": {
						SchemaProps: spec.SchemaProps{
							Description: "The creation timestamp for this version",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"createdBy": {
						SchemaProps: spec.SchemaProps{
							Description: "The user who created this version",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message passed while saving the version",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"version", "created"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardVersionList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardVersionInfo"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardVersionInfo", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardWithAccessInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "This is like the legacy DTO where access and metadata are all returned in a single call",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "The dashboard body (unstructured for now)",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
					"access": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardAccess"),
						},
					},
				},
				Required: []string{"spec", "access"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured", "github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardAccess", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_FacetResult(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"field": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"total": {
						SchemaProps: spec.SchemaProps{
							Description: "The distinct terms",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"missing": {
						SchemaProps: spec.SchemaProps{
							Description: "The number of documents that do *not* have this field",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"terms": {
						SchemaProps: spec.SchemaProps{
							Description: "Term facets",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.TermFacet"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.TermFacet"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_LibraryPanel(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Panel properties",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.LibraryPanelSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status will show errors",
							Ref:         ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.LibraryPanelStatus"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.LibraryPanelSpec", "github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.LibraryPanelStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_LibraryPanelList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.LibraryPanel"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.LibraryPanel", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_LibraryPanelSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "The panel type",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pluginVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "The panel type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"title": {
						SchemaProps: spec.SchemaProps{
							Description: "The panel title",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Library panel description",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"options": {
						SchemaProps: spec.SchemaProps{
							Description: "The options schema depends on the panel type",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
					"fieldConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "The fieldConfig schema depends on the panel type",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
					"datasource": {
						SchemaProps: spec.SchemaProps{
							Description: "The default datasource type",
							Ref:         ref("github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataSourceRef"),
						},
					},
					"targets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "The datasource queries",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataQuery"),
									},
								},
							},
						},
					},
				},
				Required: []string{"type", "options", "fieldConfig"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataQuery", "github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataSourceRef", "github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_LibraryPanelStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"warnings": {
						SchemaProps: spec.SchemaProps{
							Description: "Translation warnings (mostly things that were in SQL columns but not found in the saved body)",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"missing": {
						SchemaProps: spec.SchemaProps{
							Description: "The properties previously stored in SQL that are not included in this model",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_SearchResults(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"offset": {
						SchemaProps: spec.SchemaProps{
							Description: "Where the query started from",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"totalHits": {
						SchemaProps: spec.SchemaProps{
							Description: "The number of matching results",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"hits": {
						SchemaProps: spec.SchemaProps{
							Description: "The dashboard body (unstructured for now)",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardHit"),
									},
								},
							},
						},
					},
					"queryCost": {
						SchemaProps: spec.SchemaProps{
							Description: "Cost of running the query",
							Type:        []string{"number"},
							Format:      "double",
						},
					},
					"maxScore": {
						SchemaProps: spec.SchemaProps{
							Description: "Max score",
							Type:        []string{"number"},
							Format:      "double",
						},
					},
					"sortBy": {
						SchemaProps: spec.SchemaProps{
							Description: "How are the results sorted",
							Ref:         ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.SortBy"),
						},
					},
					"facets": {
						SchemaProps: spec.SchemaProps{
							Description: "Facet results",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.FacetResult"),
									},
								},
							},
						},
					},
				},
				Required: []string{"totalHits", "hits"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardHit", "github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.FacetResult", "github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.SortBy"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_SortBy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"field": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"desc": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"field"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_SortableFields(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"fields": {
						SchemaProps: spec.SchemaProps{
							Description: "Sortable fields (depends on backend support)",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.TableColumnDefinition"),
									},
								},
							},
						},
					},
				},
				Required: []string{"fields"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.TableColumnDefinition"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_TermFacet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"term": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"count": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_VersionsQueryOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"path": {
						SchemaProps: spec.SchemaProps{
							Description: "Path is the URL path",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
			},
		},
	}
}
