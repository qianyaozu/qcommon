package qcommon

///mongodb 通用查询数据实体类
type PostData struct {
	Method    string //get,count,update,exists,add,delete,func
	DBName    string
	Table     string
	Data      interface{}
	OrderBy   string
	Limit     int
	Skip      int
	Select    map[string]int
	ObjectID  string
	KeyWords  []string
	UpdateAll bool        //插入时判断是否更新，更新时判断是否批量更新
	Distinct  string      //是否去重，返回一个[]string
	Condition interface{} //update的时候查询条件
}

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
