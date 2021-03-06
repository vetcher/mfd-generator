package mfd

import (
	"fmt"

	"github.com/dizzyfool/genna/model"
)

func MakeSearchType(typ, searchType string) string {
	if typ[0] == '*' {
		typ = typ[1:]
	}

	switch searchType {
	case SearchArray, SearchNotArray, SearchTypeArrayContained, SearchTypeArrayIntersect:
		if _, ok := IsArray(typ); ok {
			return typ
		}
		return "[]" + typ
	case SearchTypeArrayContains, SearchTypeArrayNotContains:
		if el, ok := IsArray(typ); ok {
			return "*" + el
		}
		return "*" + typ
	case SearchNull, SearchNotNull:
		return "*" + model.TypeBool
	}

	if _, ok := IsArray(typ); ok {
		return typ
	}
	return "*" + typ
}

func IsArray(typ string) (string, bool) {
	if typ != "" && typ[0] == '[' {
		return typ[2:], true
	}

	return typ, false
}

func MakeZeroValue(typ string) string {
	switch typ {
	case model.TypeInt, model.TypeInt32, model.TypeInt64, model.TypeFloat32, model.TypeFloat64:
		return "0"
	case model.TypeByte, model.TypeDuration:
		return "0"
	case model.TypeString:
		return `""`
	case model.TypeBool:
		return "false"
	case model.TypeTime:
		return "time.Time{}"
	case model.TypeIP:
		return "new.IP{}"
	case model.TypeIPNet:
		return "new.IPNet{}"
	}

	return "nil"
}

func MakeJSType(typ string, isArray bool) string {
	if typ[0] == '*' {
		typ = typ[1:]
	}

	jsType := ""
	switch typ {
	case model.TypeInt, model.TypeInt32, model.TypeInt64, model.TypeFloat32, model.TypeFloat64:
		jsType = "number"
	case model.TypeBool:
		jsType = "boolean"
	default:
		jsType = "string"
	}

	if isArray {
		return fmt.Sprintf("Array<%s>", jsType)
	}

	return jsType
}

func MakeJSZero(typ string, isArray bool) string {
	if typ[0] == '*' {
		typ = typ[1:]
	}

	jsZero := ""
	switch typ {
	case model.TypeInt, model.TypeInt32, model.TypeInt64, model.TypeFloat32, model.TypeFloat64:
		jsZero = "0"
	case model.TypeBool:
		jsZero = "false"
	default:
		jsZero = `""`
	}

	if isArray {
		return fmt.Sprintf("[%s]", jsZero)
	}

	return jsZero
}
