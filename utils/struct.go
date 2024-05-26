package utils

import (
	"fmt"
	"reflect"
)

func MapToStruct(m map[string]interface{}, s interface{}) error {
	val := reflect.ValueOf(s).Elem()
	for k, v := range m {
		fieldName := CamelToPascal(k)
		field := val.FieldByName(fieldName)
		if !field.IsValid() {
			continue
		}
		if !field.CanSet() {
			continue
		}

		fieldType := field.Type()
		val := reflect.ValueOf(v)

		if fieldType.Kind() == reflect.Struct && val.Kind() == reflect.Map {
			fieldValuePointer := field.Addr().Interface()
			mapValue, ok := v.(map[string]interface{})
			if !ok {
				return fmt.Errorf("field %s is not a map, it is a %T", fieldName, v)
			}
			if err := MapToStruct(mapValue, fieldValuePointer); err != nil {
				return err
			}
		} else if val.Type().ConvertibleTo(fieldType) {
			field.Set(val.Convert(fieldType))
		} else {
			return fmt.Errorf("field %s of type %s is not convertible to %s", fieldName, val.Type(), fieldType)
		}
	}

	return nil
}
