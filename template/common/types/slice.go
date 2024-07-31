package types

import "reflect"

func Contains(arr interface{}, target interface{}) bool {
	arrValue := reflect.ValueOf(arr)
	if arrValue.Kind() != reflect.Slice {
		panic("not a slice")
	}

	targetValue := reflect.ValueOf(target)
	for i := 0; i < arrValue.Len(); i++ {
		if reflect.DeepEqual(arrValue.Index(i).Interface(), targetValue.Interface()) {
			return true
		}
	}

	return false
}

// SliceExtractInt32 提取struct 切片为指定字段的切片
func SliceExtractInt32(slice interface{}, key string) []int32 {
	var result []int32

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return result
	}

	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)

		// 如果字段是指针类型，则获取其指向的类型
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		keyField := getKeyField(item, key)
		field := item.FieldByName(keyField)
		if reflect.ValueOf(field).IsZero() {
			continue
		}

		fieldValue := field.Interface()

		result = append(result, fieldValue.(int32))
	}

	return result
}

// SliceExtractString 提取struct 切片为指定字段的切片
func SliceExtractString(slice interface{}, key string) []string {
	var result []string

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return result
	}

	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)

		// 如果字段是指针类型，则获取其指向的类型
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		keyField := getKeyField(item, key)
		field := item.FieldByName(keyField)
		if reflect.ValueOf(field).IsZero() {
			continue
		}

		fieldValue := field.Interface()

		result = append(result, fieldValue.(string))
	}

	return result
}
