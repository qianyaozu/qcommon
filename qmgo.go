package qcommon

import (
	"github.com/go-mgo/mgo"
	"sync"
)

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
