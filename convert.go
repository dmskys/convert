package convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//any int to string
func IntToStr(value interface{}) string {
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int))
	case uint:
		return strconv.Itoa(int(value.(uint)))
	case int8:
		return strconv.Itoa(int(value.(int8)))
	case uint8:
		return strconv.Itoa(int(value.(uint8)))
	case int16:
		return strconv.Itoa(int(value.(int16)))
	case uint16:
		return strconv.Itoa(int(value.(uint16)))
	case int32:
		return strconv.Itoa(int(value.(int32)))
	case uint32:
		return strconv.Itoa(int(value.(uint32)))
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	case uint64:
		return strconv.FormatUint(value.(uint64), 10)
	default:
		return ""
	}
}

//Decimal to binary | 十进制转二进制
func DecToBin(value interface{}) string {
	switch value.(type) {
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 2)
	case int64:
		return strconv.FormatInt(value.(int64), 2)
	case int:
		return strconv.FormatInt(int64(value.(int)), 2)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 2)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 2)
	case uint:
		return strconv.FormatUint(uint64(value.(uint)), 2)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 2)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 2)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 2)
	default:
		return ""
	}
}

//Decimal to hexadecimal | 十进制转16进制
func DecToHex(value interface{}) string {
	switch value.(type) {
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 16)
	case int64:
		return strconv.FormatInt(value.(int64), 16)
	case int:
		return strconv.FormatInt(int64(value.(int)), 16)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 16)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 16)
	case uint:
		return strconv.FormatUint(uint64(value.(uint)), 16)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 16)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 16)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 16)
	default:
		return ""
	}
}

//Decimal to octal | 十进制转8进制
func DecToOct(value interface{}) string {
	switch value.(type) {
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 8)
	case int64:
		return strconv.FormatInt(value.(int64), 8)
	case int:
		return strconv.FormatInt(int64(value.(int)), 8)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 8)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 8)
	case uint:
		return strconv.FormatUint(uint64(value.(uint)), 8)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 8)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 8)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 8)
	default:
		return ""
	}
}

// FormatFloat 将浮点数 f 转换为字符串值
// f：要转换的浮点数
// fmt：格式标记（b、e、E、f、g、G）
// prec：精度（数字部分的长度，不包括指数部分）
// bitSize：指定浮点类型（32:float32、64:float64）
//// 格式标记：
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)
//// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
//Float64ToString
func Float64ToStr(f float64, decimalNum int) string {
	//decimal_num 小数点数量
	return strconv.FormatFloat(f, 'f', decimalNum, 64)
}

func Float32ToStr(f float32, decimalNum int) string {
	//decimal_num 小数点数量
	return strconv.FormatFloat(float64(f), 'f', decimalNum, 32)
}

//string到int64
func StrToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i = 0
	}
	return i
}

//string到int
func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = 0
	}
	return i
}

//string到int32
func StrToInt32(s string) int32 {
	var i32 int32 = 0
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		i32 = 0
	} else {
		i32 = int32(i)
	}
	return i32
}

//string到int16
func StrToInt16(s string) int16 {
	var i16 int16 = 0
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		i16 = 0
	} else {
		i16 = int16(i)
	}
	return i16
}

//string到int16
func StrToInt8(s string) int8 {
	var i8 int8 = 0
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		i8 = 0
	} else {
		i8 = int8(i)
	}
	return i8
}

//string到uint64
func StrToUInt64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		i = 0
	}
	return i
}

//string到uint32
func StrToUInt32(s string) uint32 {
	var ui32 uint32 = 0
	if i, err := strconv.ParseUint(s, 10, 64); err != nil {
		ui32 = 0
	} else {
		ui32 = uint32(i)
	}
	return ui32
}

//string到uint32
func StrToUInt16(s string) uint16 {
	var ui16 uint16 = 0
	if i, err := strconv.ParseUint(s, 10, 64); err != nil {
		ui16 = 0
	} else {
		ui16 = uint16(i)
	}
	return ui16
}

//string到uint32
func StrToUInt8(s string) uint8 {
	var ui8 uint8 = 0
	if i, err := strconv.ParseUint(s, 10, 64); err != nil {
		ui8 = 0
	} else {
		ui8 = uint8(i)
	}
	return ui8
}

//string到uint
func StrToUInt(s string) uint {
	var ui uint = 0
	if i, err := strconv.ParseUint(s, 10, 64); err != nil {
		ui = 0
	} else {
		ui = uint(i)
	}
	return ui
}

//string到float64
func StrToFloat64(s string) float64 {
	var unixTime float64
	_, _ = fmt.Sscanf(s, "%f", &unixTime)
	return unixTime
}

//二进制转十进制
func BinToDec(s []byte) int64 {
	i, err := strconv.ParseInt(string(s), 2, 64)
	if err != nil {
		return 0
	}
	return i
}

//十六进制转十进制
func HexToDec(s string) int {
	var i16 = 0
	if i, err := strconv.ParseInt(s, 16, 64); err != nil {
		i16 = 0
	} else {
		i16 = int(i)
	}
	return i16
}

//八进制转十进制
func OctToDec(s string) int {
	var i8 = 0
	if i, err := strconv.ParseInt(s, 8, 64); err != nil {
		i8 = 0
	} else {
		i8 = int(i)
	}
	return i8
}

func StructToStringMap(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		switch t.Field(i).Type.Name() {
		case "string":
			data[t.Field(i).Name] = v.Field(i).String()
		case "int":
			data[t.Field(i).Name] = strconv.FormatInt(v.Field(i).Int(), 10)
		case "int8":
			data[t.Field(i).Name] = strconv.FormatInt(v.Field(i).Int(), 10)
		case "int16":
			data[t.Field(i).Name] = strconv.FormatInt(v.Field(i).Int(), 10)
		case "int32":
			data[t.Field(i).Name] = strconv.FormatInt(v.Field(i).Int(), 10)
		case "int64":
			data[t.Field(i).Name] = strconv.FormatInt(v.Field(i).Int(), 10)
		case "uint":
			data[t.Field(i).Name] = strconv.FormatUint(v.Field(i).Uint(), 10)
		case "uint8":
			data[t.Field(i).Name] = strconv.FormatUint(v.Field(i).Uint(), 10)
		case "uint16":
			data[t.Field(i).Name] = strconv.FormatUint(v.Field(i).Uint(), 10)
		case "uint32":
			data[t.Field(i).Name] = strconv.FormatUint(v.Field(i).Uint(), 10)
		case "uint64":
			data[t.Field(i).Name] = strconv.FormatUint(v.Field(i).Uint(), 10)
		case "float32":
			data[t.Field(i).Name] = strconv.FormatFloat(v.Field(i).Float(), 'f', -1, 64)
		case "float64":
			data[t.Field(i).Name] = strconv.FormatFloat(v.Field(i).Float(), 'f', -1, 64)
		case "bool":
			data[t.Field(i).Name] = strconv.FormatBool(v.Field(i).Bool())
		}
	}
	return data
}

func StringMapToStruct(data map[string]string, result interface{}) error {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		if !val.IsValid() {
			fmt.Printf("StringMapToStruct No such field: %s in datamap \n", k)
			continue
		}
		if !val.CanSet() {
			fmt.Printf("StringMapToStruct Cannot set %s field value \n", k)
			continue
		}
		valueOf := reflect.ValueOf(v)
		switch val.Type().Name() {
		case "string":
			val.SetString(valueOf.String())
		case "int":
			pv, _ := strconv.ParseInt(valueOf.String(), 10, 64)
			val.SetInt(pv)
		case "int8":
			pv, _ := strconv.ParseInt(valueOf.String(), 10, 64)
			val.SetInt(pv)
		case "int16":
			pv, _ := strconv.ParseInt(valueOf.String(), 10, 64)
			val.SetInt(pv)
		case "int32":
			pv, _ := strconv.ParseInt(valueOf.String(), 10, 64)
			val.SetInt(pv)
		case "int64":
			pv, _ := strconv.ParseInt(valueOf.String(), 10, 64)
			val.SetInt(pv)
		case "uint":
			pv, _ := strconv.ParseUint(valueOf.String(), 10, 64)
			val.SetUint(pv)
		case "uint8":
			pv, _ := strconv.ParseUint(valueOf.String(), 10, 64)
			val.SetUint(pv)
		case "uint16":
			pv, _ := strconv.ParseUint(valueOf.String(), 10, 64)
			val.SetUint(pv)
		case "uint32":
			pv, _ := strconv.ParseUint(valueOf.String(), 10, 64)
			val.SetUint(pv)
		case "uint64":
			pv, _ := strconv.ParseUint(valueOf.String(), 10, 64)
			val.SetUint(pv)
		case "float32":
			pv, _ := strconv.ParseFloat(valueOf.String(), 64)
			val.SetFloat(pv)
		case "float64":
			pv, _ := strconv.ParseFloat(valueOf.String(), 64)
			val.SetFloat(pv)
		case "bool":
			pv, _ := strconv.ParseBool(valueOf.String())
			val.SetBool(pv)
		}
	}
	return nil
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StructToDbMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		// 获取tag中的内容
		data[t.Field(i).Tag.Get("db")] = v.Field(i).Interface()
	}
	return data
}

//myData := make(map[string]interface{})
//myData["Name"] = "Tony"
//myData["Age"] = 23
//result := &MyStruct{}
//MapToStruct(myData, result)
//fmt.Println(result.Name)
func MapToStruct(data map[string]interface{}, result interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
}

// ToStr 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func ToStr(value interface{}) string {
	var strVal string
	if value == nil {
		return strVal
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		strVal = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		strVal = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		strVal = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		strVal = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		strVal = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		strVal = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		strVal = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		strVal = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		strVal = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		strVal = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		strVal = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		strVal = strconv.FormatUint(it, 10)
	case string:
		strVal = value.(string)
	case []byte:
		strVal = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		strVal = string(newValue)
	}

	return strVal
}

/**
 * 添加单引号
 */
func AddSingleQuotes(data interface{}) string {
	return "'" + strings.Trim(ToStr(data), " ") + "'"
}

/**
 * 数组转字符串, 接受混合类型, 最终输出的是字符串类型
 */
func SliceJoin(data interface{}, glue string) string {
	var tmp []string
	for _, item := range data.([]interface{}) {
		tmp = append(tmp, ToStr(item))
	}
	return strings.Join(tmp, glue)
}

// model中类型提取其中的 idField(int 类型) 属性组成 slice 返回
func ModelsToIntSlice(models interface{}, idField string) []int {
	if models == nil {
		return []int{}
	}

	// 类型检查
	modelsValue := reflect.ValueOf(models)
	if modelsValue.Kind() != reflect.Slice {
		return []int{}
	}

	var modelValue reflect.Value

	length := modelsValue.Len()
	ids := make([]int, 0, length)

	for i := 0; i < length; i++ {
		modelValue = reflect.Indirect(modelsValue.Index(i))
		if modelValue.Kind() != reflect.Struct {
			continue
		}

		val := modelValue.FieldByName(idField)
		if val.Kind() != reflect.Int {
			continue
		}

		ids = append(ids, int(val.Int()))
	}

	return ids
}
