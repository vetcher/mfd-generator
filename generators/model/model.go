package model

import (
	"fmt"
	"html/template"

	"github.com/vmkteam/mfd-generator/mfd"

	"github.com/dizzyfool/genna/model"
	"github.com/dizzyfool/genna/util"
)

// this code is used to pack mdf to template

// NamespaceData stores package info
type NamespaceData struct {
	Package string

	HasImports bool
	Imports    []string

	GoPGVer string

	Entities []EntityData
}

// PackNamespace creates a package for template
func PackNamespace(namespaces []*mfd.Namespace, options Options) NamespaceData {
	imports := mfd.NewSet()

	var models []EntityData
	for _, namespace := range namespaces {
		for _, entity := range namespace.Entities {
			// creating entity for template
			mdl := PackEntity(*entity, options)
			models = append(models, mdl)
			// adding imports to uniq set
			for _, imp := range mdl.Imports {
				imports.Add(imp)
			}

		}
	}

	goPGVer := ""
	if options.GoPGVer == mfd.GoPG9 {
		goPGVer = "/v9"
	}

	return NamespaceData{
		Package: options.Package,

		HasImports: imports.Len() > 0,
		Imports:    imports.Elements(),

		GoPGVer: goPGVer,

		Entities: models,
	}
}

// EntityData stores struct info
type EntityData struct {
	mfd.Entity

	ShortVarName string

	Tag template.HTML

	NoAlias bool
	Alias   string

	Imports []string

	Columns []AttributeData

	HasRelations bool
	Relations    []RelationData
}

// PackEntity creates an entity for template
func PackEntity(entity mfd.Entity, options Options) EntityData {
	imports := mfd.NewSet()
	var columns []AttributeData
	var relations []RelationData

	// adding columns
	for _, attribute := range entity.Attributes {
		column := PackAttribute(entity, *attribute, options)
		columns = append(columns, column)

		// adding imports to uniq set
		if column.Import != "" {
			imports.Add(column.Import)
		}

		// adding relation from column
		if attribute.ForeignKey != "" && !attribute.IsArray {
			relations = append(relations, PackRelation(*attribute, options))
		}
	}

	// adding annotations for go-pg to column
	tagName := tagName(options)
	tags := util.NewAnnotation()
	tags.AddTag(tagName, util.Quoted(entity.Table, true))
	tags.AddTag(tagName, fmt.Sprintf("alias:%s", util.DefaultAlias))
	tags.AddTag("pg", ",discard_unknown_columns")

	return EntityData{
		Entity: entity,

		ShortVarName: mfd.ShortVarName(entity.Name),

		// avoid escaping
		Tag:   template.HTML(fmt.Sprintf("`%s`", tags.String())),
		Alias: util.DefaultAlias,

		Imports: imports.Elements(),

		Columns: columns,

		HasRelations: len(relations) > 0,
		Relations:    relations,
	}
}

// AttributeData stores column info
type AttributeData struct {
	mfd.Attribute

	Name   string
	Type   string
	Import string

	Tag     template.HTML
	Comment template.HTML
}

// PackAttribute creates a column for template
func PackAttribute(entity mfd.Entity, attribute mfd.Attribute, options Options) AttributeData {
	comment := ""
	tagName := tagName(options)
	tags := util.NewAnnotation()
	tags.AddTag(tagName, attribute.DBName)

	// pk tag
	if attribute.PrimaryKey {
		tags.AddTag(tagName, "pk")
	}

	// types tag
	if attribute.DBType == model.TypePGHstore {
		tags.AddTag(tagName, "hstore")
	} else if attribute.IsArray {
		tags.AddTag(tagName, "array")
	}
	if attribute.DBType == model.TypePGUuid {
		tags.AddTag(tagName, "type:uuid")
	}

	// nullable tag
	if !attribute.Nullable() && !attribute.PrimaryKey {
		if options.GoPGVer == mfd.GoPG9 {
			tags.AddTag(tagName, "use_zero")
		} else {
			tags.AddTag(tagName, "notnull")
		}
	}

	// go type
	goType, err := model.GoType(attribute.DBType)
	if err != nil {
		goType = model.TypeInterface
	}

	// mark unknown types as interface & unsupported
	if goType == model.TypeInterface {
		comment = "// unsupported"
		tags = util.NewAnnotation().AddTag(tagName, "-")
	}

	// fix pointer in case of inconsistency
	attribute.GoType = fixPointer(attribute)

	return AttributeData{
		Attribute: attribute,

		Type:   goType,
		Name:   util.ColumnName(attribute.Name),
		Import: model.GoImport(attribute.DBType, attribute.Nullable(), false, options.GoPGVer),

		Tag:     template.HTML(fmt.Sprintf("`%s`", tags.String())),
		Comment: template.HTML(comment),
	}
}

// RelationData stores relation info
type RelationData struct {
	mfd.Attribute

	Name string
	Type string

	Tag     template.HTML
	Comment template.HTML
}

// PackRelation creates relation for template
func PackRelation(relation mfd.Attribute, options Options) RelationData {
	// adding go-pg's fk annotation
	tags := util.NewAnnotation().AddTag("pg", "fk:"+relation.DBName)
	comment := ""

	// getting pk in foreign table
	var fkFields []string
	if relation.ForeignEntity != nil {
		pks := relation.ForeignEntity.PKs()
		for _, pk := range pks {
			fkFields = append(fkFields, pk.DBName)
		}
	}

	if len(fkFields) > 1 {
		tagName := tagName(options)
		tags.AddTag(tagName, "-")
		comment = "// unsupported"
	}

	return RelationData{
		Attribute: relation,

		// ObjectID -> Object, UserID -> User
		Name: util.ReplaceSuffix(util.ColumnName(relation.DBName), util.ID, ""),
		Type: relation.ForeignKey,

		Tag:     template.HTML(fmt.Sprintf("`%s`", tags.String())),
		Comment: template.HTML(comment),
	}
}

func tagName(options Options) string {
	if options.GoPGVer == mfd.GoPG9 {
		return "pg"
	}
	return "sql"
}

func fixPointer(attribute mfd.Attribute) string {
	// basic
	if !attribute.Nullable() || attribute.PrimaryKey {
		return attribute.GoType
	}

	// type opts
	if attribute.IsMap() || attribute.IsArray /*|| attribute.IsJSON()*/ {
		return attribute.GoType
	}

	if attribute.GoType[0] != '*' {
		return "*" + attribute.GoType
	}

	return attribute.GoType
}
