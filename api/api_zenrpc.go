// Code generated by zenrpc; DO NOT EDIT.

package api

import (
	"context"
	"encoding/json"

	"github.com/vmkteam/zenrpc/v2"
	"github.com/vmkteam/zenrpc/v2/smd"

	"github.com/vmkteam/mfd-generator/mfd"
)

var RPC = struct {
	ProjectService struct{ Open, Current, Update, Save, Tables string }
	PublicService  struct{ GoPGVersions, Modes, SearchTypes, Types, DBTypes, Ping string }
	XMLService     struct{ GenerateEntity, LoadEntity, UpdateEntity string }
	XMLLangService struct{ LoadTranslation, TranslateEntity string }
	XMLVTService   struct{ GenerateEntity, LoadEntity, UpdateEntity string }
}{
	ProjectService: struct{ Open, Current, Update, Save, Tables string }{
		Open:    "open",
		Current: "current",
		Update:  "update",
		Save:    "save",
		Tables:  "tables",
	},
	PublicService: struct{ GoPGVersions, Modes, SearchTypes, Types, DBTypes, Ping string }{
		GoPGVersions: "gopgversions",
		Modes:        "modes",
		SearchTypes:  "searchtypes",
		Types:        "types",
		DBTypes:      "dbtypes",
		Ping:         "ping",
	},
	XMLService: struct{ GenerateEntity, LoadEntity, UpdateEntity string }{
		GenerateEntity: "generateentity",
		LoadEntity:     "loadentity",
		UpdateEntity:   "updateentity",
	},
	XMLLangService: struct{ LoadTranslation, TranslateEntity string }{
		LoadTranslation: "loadtranslation",
		TranslateEntity: "translateentity",
	},
	XMLVTService: struct{ GenerateEntity, LoadEntity, UpdateEntity string }{
		GenerateEntity: "generateentity",
		LoadEntity:     "loadentity",
		UpdateEntity:   "updateentity",
	},
}

func (ProjectService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"Open": {
				Description: `Loads project from file`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "filePath",
						Description: `the path to mfd file`,
						Type:        smd.String,
					},
					{
						Name:        "connection",
						Description: `connection string`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `Project`,
					Optional:    true,
					Type:        smd.Object,
					Properties: smd.PropertyList{
						{
							Name: "name",
							Type: smd.String,
						},
						{
							Name: "languages",
							Type: smd.Array,
							Items: map[string]string{
								"type": smd.String,
							},
						},
						{
							Name: "goPGVer",
							Type: smd.Integer,
						},
						{
							Name: "customTypes",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.CustomTypes",
							},
						},
						{
							Name: "namespaces",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.NSMapping",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.CustomTypes": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "dbType",
									Type: smd.String,
								},
								{
									Name: "goType",
									Type: smd.String,
								},
								{
									Name: "goImport",
									Type: smd.String,
								},
							},
						},
						"mfd.NSMapping": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "namespace",
									Type: smd.String,
								},
								{
									Name: "entity",
									Type: smd.String,
								},
							},
						},
					},
				},
			},
			"Current": {
				Description: `Returns currently open project`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `Project`,
					Optional:    true,
					Type:        smd.Object,
					Properties: smd.PropertyList{
						{
							Name: "name",
							Type: smd.String,
						},
						{
							Name: "languages",
							Type: smd.Array,
							Items: map[string]string{
								"type": smd.String,
							},
						},
						{
							Name: "goPGVer",
							Type: smd.Integer,
						},
						{
							Name: "customTypes",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.CustomTypes",
							},
						},
						{
							Name: "namespaces",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.NSMapping",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.CustomTypes": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "dbType",
									Type: smd.String,
								},
								{
									Name: "goType",
									Type: smd.String,
								},
								{
									Name: "goImport",
									Type: smd.String,
								},
							},
						},
						"mfd.NSMapping": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "namespace",
									Type: smd.String,
								},
								{
									Name: "entity",
									Type: smd.String,
								},
							},
						},
					},
				},
			},
			"Update": {
				Description: `Updates project in memory`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "project",
						Description: `Project`,
						Type:        smd.Object,
						Properties: smd.PropertyList{
							{
								Name: "name",
								Type: smd.String,
							},
							{
								Name: "languages",
								Type: smd.Array,
								Items: map[string]string{
									"type": smd.String,
								},
							},
							{
								Name: "goPGVer",
								Type: smd.Integer,
							},
							{
								Name: "customTypes",
								Type: smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.CustomTypes",
								},
							},
							{
								Name: "namespaces",
								Type: smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.NSMapping",
								},
							},
						},
						Definitions: map[string]smd.Definition{
							"mfd.CustomTypes": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "dbType",
										Type: smd.String,
									},
									{
										Name: "goType",
										Type: smd.String,
									},
									{
										Name: "goImport",
										Type: smd.String,
									},
								},
							},
							"mfd.NSMapping": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "namespace",
										Type: smd.String,
									},
									{
										Name: "entity",
										Type: smd.String,
									},
								},
							},
						},
					},
				},
			},
			"Save": {
				Description: `Saves project from memory to disk`,
				Parameters:  []smd.JSONSchema{},
			},
			"Tables": {
				Description: `Gets all tables from database`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of tables`,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.String,
					},
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s ProjectService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.ProjectService.Open:
		var args = struct {
			FilePath   string `json:"filePath"`
			Connection string `json:"connection"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"filePath", "connection"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Open(args.FilePath, args.Connection))

	case RPC.ProjectService.Current:
		resp.Set(s.Current())

	case RPC.ProjectService.Update:
		var args = struct {
			Project mfd.Project `json:"project"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"project"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Update(args.Project))

	case RPC.ProjectService.Save:
		resp.Set(s.Save())

	case RPC.ProjectService.Tables:
		resp.Set(s.Tables())

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (PublicService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"GoPGVersions": {
				Description: `Gets all supported go-pg versions`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of versions`,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.Integer,
					},
				},
			},
			"Modes": {
				Description: `Gets all available entity modes`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of modes`,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.String,
					},
				},
			},
			"SearchTypes": {
				Description: `Gets all available search types`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of search types`,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.String,
					},
				},
			},
			"Types": {
				Description: `Gets std types`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of types`,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.String,
					},
				},
			},
			"DBTypes": {
				Description: `Gets postgres types`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of types`,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.String,
					},
				},
			},
			"Ping": {
				Parameters: []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Type: smd.String,
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s PublicService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}

	switch method {
	case RPC.PublicService.GoPGVersions:
		resp.Set(s.GoPGVersions())

	case RPC.PublicService.Modes:
		resp.Set(s.Modes())

	case RPC.PublicService.SearchTypes:
		resp.Set(s.SearchTypes())

	case RPC.PublicService.Types:
		resp.Set(s.Types())

	case RPC.PublicService.DBTypes:
		resp.Set(s.DBTypes())

	case RPC.PublicService.Ping:
		resp.Set(s.Ping())

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (XMLService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"GenerateEntity": {
				Description: `Gets entity for selected table`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "table",
						Description: `selected table name`,
						Type:        smd.String,
					},
					{
						Name:        "namespace",
						Description: `namespace of the new entity`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `Entity`,
					Optional:    true,
					Type:        smd.Object,
					Properties: smd.PropertyList{
						{
							Name: "name",
							Type: smd.String,
						},
						{
							Name: "namespace",
							Type: smd.String,
						},
						{
							Name: "table",
							Type: smd.String,
						},
						{
							Name: "attributes",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.Attributes",
							},
						},
						{
							Name: "searches",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.Searches",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.Attributes": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "dbName",
									Type: smd.String,
								},
								{
									Name: "isArray",
									Type: smd.Boolean,
								},
								{
									Name: "disablePointer",
									Type: smd.Boolean,
								},
								{
									Name: "dbType",
									Type: smd.String,
								},
								{
									Name: "goType",
									Type: smd.String,
								},
								{
									Name: "pk",
									Type: smd.Boolean,
								},
								{
									Name: "fk",
									Type: smd.String,
								},
								{
									Name: "nullable",
									Type: smd.String,
								},
								{
									Name:     "addable",
									Optional: true,
									Type:     smd.Boolean,
								},
								{
									Name:     "updatable",
									Optional: true,
									Type:     smd.Boolean,
								},
								{
									Name: "min",
									Type: smd.Integer,
								},
								{
									Name: "max",
									Type: smd.Integer,
								},
								{
									Name: "defaultVal",
									Type: smd.String,
								},
							},
						},
						"mfd.Searches": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "attrName",
									Type: smd.String,
								},
								{
									Name: "searchType",
									Type: smd.String,
								},
								{
									Name: "goType",
									Type: smd.String,
								},
							},
						},
					},
				},
			},
			"LoadEntity": {
				Description: `Gets selected entity from project`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Description: `namespace of the entity`,
						Type:        smd.String,
					},
					{
						Name:        "entity",
						Description: `the name of the entity`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `Entity`,
					Optional:    true,
					Type:        smd.Object,
					Properties: smd.PropertyList{
						{
							Name: "name",
							Type: smd.String,
						},
						{
							Name: "namespace",
							Type: smd.String,
						},
						{
							Name: "table",
							Type: smd.String,
						},
						{
							Name: "attributes",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.Attributes",
							},
						},
						{
							Name: "searches",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.Searches",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.Attributes": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "dbName",
									Type: smd.String,
								},
								{
									Name: "isArray",
									Type: smd.Boolean,
								},
								{
									Name: "disablePointer",
									Type: smd.Boolean,
								},
								{
									Name: "dbType",
									Type: smd.String,
								},
								{
									Name: "goType",
									Type: smd.String,
								},
								{
									Name: "pk",
									Type: smd.Boolean,
								},
								{
									Name: "fk",
									Type: smd.String,
								},
								{
									Name: "nullable",
									Type: smd.String,
								},
								{
									Name:     "addable",
									Optional: true,
									Type:     smd.Boolean,
								},
								{
									Name:     "updatable",
									Optional: true,
									Type:     smd.Boolean,
								},
								{
									Name: "min",
									Type: smd.Integer,
								},
								{
									Name: "max",
									Type: smd.Integer,
								},
								{
									Name: "defaultVal",
									Type: smd.String,
								},
							},
						},
						"mfd.Searches": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "attrName",
									Type: smd.String,
								},
								{
									Name: "searchType",
									Type: smd.String,
								},
								{
									Name: "goType",
									Type: smd.String,
								},
							},
						},
					},
				},
			},
			"UpdateEntity": {
				Description: `Saves selected entity in project`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "entity",
						Optional:    true,
						Description: `Entity`,
						Type:        smd.Object,
						Properties: smd.PropertyList{
							{
								Name: "name",
								Type: smd.String,
							},
							{
								Name: "namespace",
								Type: smd.String,
							},
							{
								Name: "table",
								Type: smd.String,
							},
							{
								Name: "attributes",
								Type: smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.Attributes",
								},
							},
							{
								Name: "searches",
								Type: smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.Searches",
								},
							},
						},
						Definitions: map[string]smd.Definition{
							"mfd.Attributes": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "name",
										Type: smd.String,
									},
									{
										Name: "dbName",
										Type: smd.String,
									},
									{
										Name: "isArray",
										Type: smd.Boolean,
									},
									{
										Name: "disablePointer",
										Type: smd.Boolean,
									},
									{
										Name: "dbType",
										Type: smd.String,
									},
									{
										Name: "goType",
										Type: smd.String,
									},
									{
										Name: "pk",
										Type: smd.Boolean,
									},
									{
										Name: "fk",
										Type: smd.String,
									},
									{
										Name: "nullable",
										Type: smd.String,
									},
									{
										Name:     "addable",
										Optional: true,
										Type:     smd.Boolean,
									},
									{
										Name:     "updatable",
										Optional: true,
										Type:     smd.Boolean,
									},
									{
										Name: "min",
										Type: smd.Integer,
									},
									{
										Name: "max",
										Type: smd.Integer,
									},
									{
										Name: "defaultVal",
										Type: smd.String,
									},
								},
							},
							"mfd.Searches": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "name",
										Type: smd.String,
									},
									{
										Name: "attrName",
										Type: smd.String,
									},
									{
										Name: "searchType",
										Type: smd.String,
									},
									{
										Name: "goType",
										Type: smd.String,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s XMLService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.XMLService.GenerateEntity:
		var args = struct {
			Table     string `json:"table"`
			Namespace string `json:"namespace"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"table", "namespace"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.GenerateEntity(args.Table, args.Namespace))

	case RPC.XMLService.LoadEntity:
		var args = struct {
			Namespace string `json:"namespace"`
			Entity    string `json:"entity"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "entity"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.LoadEntity(args.Namespace, args.Entity))

	case RPC.XMLService.UpdateEntity:
		var args = struct {
			Entity *mfd.Entity `json:"entity"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"entity"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.UpdateEntity(args.Entity))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (XMLLangService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"LoadTranslation": {
				Description: `Loads full translation of project`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "language",
						Description: `language`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `Translation`,
					Optional:    true,
					Type:        smd.Object,
					Properties: smd.PropertyList{
						{
							Name: "language",
							Type: smd.String,
						},
						{
							Name: "namespaces",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.TranslationNamespace",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.TranslationNamespace": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "entities",
									Type: smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/mfd.TranslationEntity",
									},
								},
							},
						},
						"mfd.TranslationEntity": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "key",
									Type: smd.String,
								},
								{
									Name:     "crumbs",
									Optional: true,
									Ref:      "#/definitions/mfd.XMLMap",
									Type:     smd.Object,
								},
								{
									Name:     "form",
									Optional: true,
									Ref:      "#/definitions/mfd.XMLMap",
									Type:     smd.Object,
								},
								{
									Name:     "list",
									Optional: true,
									Ref:      "#/definitions/mfd.TranslationList",
									Type:     smd.Object,
								},
							},
						},
					},
				},
			},
			"TranslateEntity": {
				Description: `Translates entity`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Description: `namespace of the vt entity`,
						Type:        smd.String,
					},
					{
						Name:        "entity",
						Description: `vt entity from vt.xml`,
						Type:        smd.String,
					},
					{
						Name:        "language",
						Description: `language`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `TranslationEntity`,
					Optional:    true,
					Type:        smd.Object,
					Properties: smd.PropertyList{
						{
							Name: "name",
							Type: smd.String,
						},
						{
							Name: "key",
							Type: smd.String,
						},
						{
							Name:     "crumbs",
							Optional: true,
							Ref:      "#/definitions/mfd.XMLMap",
							Type:     smd.Object,
						},
						{
							Name:     "form",
							Optional: true,
							Ref:      "#/definitions/mfd.XMLMap",
							Type:     smd.Object,
						},
						{
							Name:     "list",
							Optional: true,
							Ref:      "#/definitions/mfd.TranslationList",
							Type:     smd.Object,
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.XMLMap": {
							Type:       "object",
							Properties: smd.PropertyList{},
						},
						"mfd.TranslationList": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "title",
									Type: smd.String,
								},
								{
									Name:     "filter",
									Optional: true,
									Ref:      "#/definitions/mfd.XMLMap",
									Type:     smd.Object,
								},
								{
									Name:     "headers",
									Optional: true,
									Ref:      "#/definitions/mfd.XMLMap",
									Type:     smd.Object,
								},
							},
						},
					},
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s XMLLangService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.XMLLangService.LoadTranslation:
		var args = struct {
			Language string `json:"language"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"language"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.LoadTranslation(args.Language))

	case RPC.XMLLangService.TranslateEntity:
		var args = struct {
			Namespace string `json:"namespace"`
			Entity    string `json:"entity"`
			Language  string `json:"language"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "entity", "language"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.TranslateEntity(args.Namespace, args.Entity, args.Language))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (XMLVTService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"GenerateEntity": {
				Description: `Gets vt entity for selected base entity`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Description: `namespace of the base entity`,
						Type:        smd.String,
					},
					{
						Name:        "entity",
						Description: `base entity from namespace.xml`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `VTEntity`,
					Optional:    true,
					Type:        smd.Object,
					Properties: smd.PropertyList{
						{
							Name: "name",
							Type: smd.String,
						},
						{
							Name: "terminalPath",
							Type: smd.String,
						},
						{
							Name: "attributes",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.VTAttributes",
							},
						},
						{
							Name: "template",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.TmplAttributes",
							},
						},
						{
							Name: "mode",
							Type: smd.String,
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.VTAttributes": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "attrName",
									Type: smd.String,
								},
								{
									Name: "searchName",
									Type: smd.String,
								},
								{
									Name: "summary",
									Type: smd.Boolean,
								},
								{
									Name: "search",
									Type: smd.Boolean,
								},
								{
									Name: "max",
									Type: smd.Integer,
								},
								{
									Name: "min",
									Type: smd.Integer,
								},
								{
									Name: "required",
									Type: smd.Boolean,
								},
								{
									Name: "validate",
									Type: smd.String,
								},
							},
						},
						"mfd.TmplAttributes": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "vtAttrName",
									Type: smd.String,
								},
								{
									Name:        "list",
									Description: `show in list`,
									Type:        smd.Boolean,
								},
								{
									Name:        "fkOpts",
									Description: `how to show fks`,
									Type:        smd.String,
								},
								{
									Name:        "form",
									Description: `show in object editor`,
									Type:        smd.String,
								},
								{
									Name:        "search",
									Description: `input type in search`,
									Type:        smd.String,
								},
							},
						},
					},
				},
			},
			"LoadEntity": {
				Description: `Gets vt entity for selected entity from project`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Description: `namespace of the vt entity`,
						Type:        smd.String,
					},
					{
						Name:        "entity",
						Description: `the name of the vt entity`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `VTEntity`,
					Optional:    true,
					Type:        smd.Object,
					Properties: smd.PropertyList{
						{
							Name: "name",
							Type: smd.String,
						},
						{
							Name: "terminalPath",
							Type: smd.String,
						},
						{
							Name: "attributes",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.VTAttributes",
							},
						},
						{
							Name: "template",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.TmplAttributes",
							},
						},
						{
							Name: "mode",
							Type: smd.String,
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.VTAttributes": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "attrName",
									Type: smd.String,
								},
								{
									Name: "searchName",
									Type: smd.String,
								},
								{
									Name: "summary",
									Type: smd.Boolean,
								},
								{
									Name: "search",
									Type: smd.Boolean,
								},
								{
									Name: "max",
									Type: smd.Integer,
								},
								{
									Name: "min",
									Type: smd.Integer,
								},
								{
									Name: "required",
									Type: smd.Boolean,
								},
								{
									Name: "validate",
									Type: smd.String,
								},
							},
						},
						"mfd.TmplAttributes": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "name",
									Type: smd.String,
								},
								{
									Name: "vtAttrName",
									Type: smd.String,
								},
								{
									Name:        "list",
									Description: `show in list`,
									Type:        smd.Boolean,
								},
								{
									Name:        "fkOpts",
									Description: `how to show fks`,
									Type:        smd.String,
								},
								{
									Name:        "form",
									Description: `show in object editor`,
									Type:        smd.String,
								},
								{
									Name:        "search",
									Description: `input type in search`,
									Type:        smd.String,
								},
							},
						},
					},
				},
			},
			"UpdateEntity": {
				Description: `Saves vt entity in project`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Description: `namespace of the vt entity`,
						Type:        smd.String,
					},
					{
						Name:        "entity",
						Optional:    true,
						Description: `VTEntity`,
						Type:        smd.Object,
						Properties: smd.PropertyList{
							{
								Name: "name",
								Type: smd.String,
							},
							{
								Name: "terminalPath",
								Type: smd.String,
							},
							{
								Name: "attributes",
								Type: smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.VTAttributes",
								},
							},
							{
								Name: "template",
								Type: smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.TmplAttributes",
								},
							},
							{
								Name: "mode",
								Type: smd.String,
							},
						},
						Definitions: map[string]smd.Definition{
							"mfd.VTAttributes": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "name",
										Type: smd.String,
									},
									{
										Name: "attrName",
										Type: smd.String,
									},
									{
										Name: "searchName",
										Type: smd.String,
									},
									{
										Name: "summary",
										Type: smd.Boolean,
									},
									{
										Name: "search",
										Type: smd.Boolean,
									},
									{
										Name: "max",
										Type: smd.Integer,
									},
									{
										Name: "min",
										Type: smd.Integer,
									},
									{
										Name: "required",
										Type: smd.Boolean,
									},
									{
										Name: "validate",
										Type: smd.String,
									},
								},
							},
							"mfd.TmplAttributes": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "name",
										Type: smd.String,
									},
									{
										Name: "vtAttrName",
										Type: smd.String,
									},
									{
										Name:        "list",
										Description: `show in list`,
										Type:        smd.Boolean,
									},
									{
										Name:        "fkOpts",
										Description: `how to show fks`,
										Type:        smd.String,
									},
									{
										Name:        "form",
										Description: `show in object editor`,
										Type:        smd.String,
									},
									{
										Name:        "search",
										Description: `input type in search`,
										Type:        smd.String,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s XMLVTService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.XMLVTService.GenerateEntity:
		var args = struct {
			Namespace string `json:"namespace"`
			Entity    string `json:"entity"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "entity"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.GenerateEntity(args.Namespace, args.Entity))

	case RPC.XMLVTService.LoadEntity:
		var args = struct {
			Namespace string `json:"namespace"`
			Entity    string `json:"entity"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "entity"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.LoadEntity(args.Namespace, args.Entity))

	case RPC.XMLVTService.UpdateEntity:
		var args = struct {
			Namespace string        `json:"namespace"`
			Entity    *mfd.VTEntity `json:"entity"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "entity"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.UpdateEntity(args.Namespace, args.Entity))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}
