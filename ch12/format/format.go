package format

import (
	"reflect"
	"strconv"
)

// Any把任何值格式化为一个字符串
func Any(value interface{}) string {
	return FormatAtom(reflect.ValueOf(value))
}

// formatAtom格式化一个值，且不分析他的内部结构
func FormatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid: // 无效值
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64: // 整数类型
		return strconv.FormatInt(v.Int(), 10) // 获取该值的整数表示，按照十进制输出
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr: // 无符号整型
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool: //布尔类型
		return strconv.FormatBool(v.Bool())
	case reflect.String: // 字符串类型
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map: // 指针，通道，函数，切片，映射类型
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16) // 返回类型名和指针地址（通过v.Pointer()获取)。这个地址用十六进制表示
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
