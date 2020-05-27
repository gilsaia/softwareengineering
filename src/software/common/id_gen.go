package common

import "github.com/sony/sonyflake"

var sf *sonyflake.Sonyflake

func init() {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenId() int64 {
	id, _ := sf.NextID()
	return int64(id)
}
