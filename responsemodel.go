package qcommon

import (
	"encoding/json"
	"fmt"
)



type ResponseModel struct {
	State   int
	Message string
	Data    interface{}
}


func ResponseJson(js interface{}, e error) ResponseModel {
	var data ResponseModel
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
	var data ResponseModel
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

func Recovery() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}