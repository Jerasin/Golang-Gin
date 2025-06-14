package service

import (
	"reflect"
	"strings"

	"github.com/goforj/godump"
	log "github.com/sirupsen/logrus"
)

func DbHandleSelectField(field any) map[string]interface{} {
	fields := reflect.TypeOf(field)
	result := make(map[string]interface{})
	for i := 0; i < fields.NumField(); i++ {
		// Get the field
		field := fields.Field(i)

		// Get the json tag value
		var key string
		jsonTag := field.Tag.Get("json")
		jsonGormTag := field.Tag.Get("gorm")

		if jsonGormTag != "" {
			strArr := strings.Split(jsonGormTag, ":")
			godump.Dump(jsonGormTag)
			if len(strArr) == 2 {
				key = strArr[1]
			}
		} else {
			key = jsonTag
		}

		// Print the json tag value
		log.Infof("Field %d: %s\n", i+1, key)
		result[key] = ""
	}

	return result
}
