package request

import (
	"reflect"
	"strings"
)

func Sanitize(input any) error {
	value := reflect.ValueOf(input)

	if value.Kind() != reflect.Pointer || value.IsNil() {
		return nil
	}

	return sanitizeValue(value.Elem())
}

func sanitizeValue(value reflect.Value) error {
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil
		}
		return sanitizeValue(value.Elem())
	}

	if value.Kind() != reflect.Struct {
		return nil
	}

	valueType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		structField := valueType.Field(i)

		if !field.CanSet() {
			continue
		}

		if field.Kind() == reflect.Struct {
			_ = sanitizeValue(field)
			continue
		}

		if field.Kind() == reflect.Slice {
			tag := structField.Tag.Get("sanitize")
			for j := 0; j < field.Len(); j++ {
				item := field.Index(j)
				if item.Kind() == reflect.String && tag != "" {
					item.SetString(applyStringSanitizers(item.String(), tag))
					continue
				}
				if item.Kind() == reflect.Struct {
					_ = sanitizeValue(item)
				}
				if item.Kind() == reflect.Pointer && !item.IsNil() {
					_ = sanitizeValue(item.Elem())
				}
			}
			continue
		}

		if field.Kind() != reflect.String {
			continue
		}

		tag := structField.Tag.Get("sanitize")
		if tag == "" {
			continue
		}

		field.SetString(applyStringSanitizers(field.String(), tag))
	}

	return nil
}

func applyStringSanitizers(value string, tag string) string {
	rules := strings.Split(tag, ",")

	for _, rule := range rules {
		switch strings.TrimSpace(rule) {
		case "trim":
			value = strings.TrimSpace(value)
		case "lower":
			value = strings.ToLower(value)
		case "upper":
			value = strings.ToUpper(value)
		}
	}

	return value
}
