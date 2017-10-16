package qcommon

import (
	"testing"
	"fmt"
	"time"
)

func TestQconf(t *testing.T) {
	uuid := UUIDDefault(1)
	var i = 0

	fmt.Println(uuid.GetUUID())
	fmt.Println(uuid.GetUUID())
	fmt.Println(uuid.GetUUID())
	fmt.Println(uuid.GetUUID())
	fmt.Println(uuid.GetUUID())
	fmt.Println(uuid.GetUUID())
	fmt.Println(uuid.GetUUID())
	fmt.Println(uuid.GetUUID())
	fmt.Println(uuid.GetUUID())
	var t1 = time.Now().UnixNano() / 1e6
	for {
		uuid.GetUUID()
		i++
		if (time.Now().UnixNano()/1e6)-1000 > t1 {
			break
		}
	}
	fmt.Println(uuid.GetUUID())
	fmt.Println(i)
}
