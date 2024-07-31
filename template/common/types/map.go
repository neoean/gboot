package types

import (
	"reflect"
)

// SliceToMap 将结构体切片转换为映射，使用标签指定键的字段
func SliceToMap(slice interface{}, key string) map[interface{}]interface{} {
	resultMap := make(map[interface{}]interface{})

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return resultMap
	}

	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)

		itemOrigin := item

		// 如果字段是指针类型，则获取其指向的类型
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		fieldName := getKeyField(item, key)
		field := item.FieldByName(fieldName)
		if reflect.ValueOf(field).IsZero() {
			continue
		}
		fieldValue := field.Interface()

		resultMap[fieldValue] = itemOrigin.Interface()
	}

	return resultMap
}

// getKeyField 获取结构体的键字段名
func getKeyField(item reflect.Value, k string) string {
	itemType := item.Type()

	// 如果字段是指针类型，则获取其指向的类型
	if itemType.Kind() == reflect.Ptr {
		itemType = itemType.Elem()
	}

	for i := 0; i < itemType.NumField(); i++ {
		field := itemType.Field(i)
		if tagValue, ok := field.Tag.Lookup("json"); ok && tagValue == k {
			return field.Name
		}
	}
	// 如果没有指定json标签，默认使用结构体的第一个字段作为键
	return ""
}
