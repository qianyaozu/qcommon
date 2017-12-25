package qcommon

import (
	"sync"
	"time"
)

var mu sync.Mutex

type UUID struct {
	LastTimeStamp int64
	Sequence      int64
	WorkId        int64
}

//workerId为机器码，生成的uuid格式为时间戳精确到毫秒+3位序列+3位机器码
func UUIDDefault(workerId int64) *UUID {
	return &UUID{LastTimeStamp: 0, Sequence: -1, WorkId: workerId}
}
func (uuid *UUID) GetUUID() int64 {
	mu.Lock()
	defer mu.Unlock()
	var t = time.Now().UnixNano() / 1e6
	if t == uuid.LastTimeStamp {
		uuid.Sequence = (uuid.Sequence + 1)
	} else {
		uuid.Sequence = 0
	}
	uuid.LastTimeStamp = t
	return t*1000000 + uuid.Sequence*1000 + uuid.WorkId
}
