package service

import (
	"reflect"
	"strings"
)

func DbHandleSelectField(field any) map[string]interface{} {
	result := make(map[string]any)
	buildSelectField(reflect.TypeOf(field), result)
	return result
}

func buildSelectField(t reflect.Type, result map[string]interface{}) {
	// ถ้าเป็น pointer ให้ดึง element
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		// ถ้าเป็น embedded struct ให้ทำ recursive
		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			buildSelectField(f.Type, result)
			continue
		}

		// ถ้าไม่มี tag ข้าม
		jsonTag := f.Tag.Get("json")
		gormTag := f.Tag.Get("gorm")

		var key string

		if gormTag != "" {
			// แยกตาม `;` และหา "column:xxx"
			parts := strings.Split(gormTag, ";")
			for _, part := range parts {
				if strings.HasPrefix(part, "column:") {
					key = strings.TrimPrefix(part, "column:")
					break
				}
			}
		}

		if key == "" && jsonTag != "" {
			key = strings.Split(jsonTag, ",")[0] // ตัด off omitempty
		}

		if key == "" {
			continue
		}

		result[key] = ""
	}
}
