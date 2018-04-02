package qcommon

import (
	"github.com/go-mgo/mgo"
	"sync"
)
///mongodb 查询数据实体类
type PostData struct {
	Method string//get,count,update,exists,add,delete,func
	DBName string
	Table     string
	Data      interface{}
	OrderBy   string
	Limit     int
	Skip      int
	Select    map[string]int
	UpdateAll    bool        //插入时判断是否更新，更新时判断是否批量更新
	Distinct  string      //是否去重，返回一个[]string
	Condition interface{} //update的时候查询条件
}
var mgomutex sync.Mutex

func InitMongo(url string) (*mgo.Session, error) {
	mgomutex.Lock()
	defer mgomutex.Unlock()
	mongo := url
	session, err := mgo.Dial(mongo) //127.0.0.1:27017
	if err != nil {
		return nil, err
	} else {
		session.SetCursorTimeout(0)
		session.SetMode(mgo.Monotonic, true)
		return session, nil
	}
}
