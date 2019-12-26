package utils

import "encoding/json"

//interface{}可以储存任意类型的数值
func JsonEncode(code int, data interface{}, msg string) ([]byte, error) {
	result := struct {
		Code int
		Data interface{}
		Msg  string
	}{code, data, msg}
	return json.Marshal(result)
}
