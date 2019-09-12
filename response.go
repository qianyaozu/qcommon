package qcommon

import "encoding/json"

type Response struct {
	State   int
	Message string
	Data    interface{}
}


func ResponseJson(js interface{}, e error) Response {
	var data Response
	if e != nil {
		data.State = 0
		data.Message = e.Error()
	} else {
		data.State = 1
		data.Message = "success"
	}
	if js != nil {
		data.Data = js
	}
	return data
}


func ResponseData(js interface{},e error) []byte {
	var data Response
	if e != nil {
		data.State = 0
		data.Message = e.Error()
	} else {
		data.State = 1
		data.Message = "success"
	}
	if js != nil {
		data.Data = js
	}
	d, _ := json.Marshal(data)
	return d
}