package dynamic_utils

import "reflect"

func OptionalStructFieldsToMap(s any) map[string]any {
	structType := reflect.TypeOf(s)
	structValue := reflect.ValueOf(s)

	if structType.Kind() == reflect.Pointer {
		structValue = structValue.Elem()
		structType = structType.Elem()
	}

	r := make(map[string]any)
	for i := 0; i < structValue.NumField(); i++ {
		fieldType := structType.Field(i)
		fieldValue := structValue.Field(i)

		if fieldValue.IsNil() {
			continue
		}

		fieldValue = reflect.Indirect(fieldValue)
		r[fieldType.Name] = fieldValue.Interface()
	}

	return r
}
