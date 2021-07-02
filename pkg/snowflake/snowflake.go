package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)
var node *snowflake.Node
func Init(startTime string, machineId int64)(err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if nil != err {
		return
	}
	// Epoch 默认是 2010-11-04 01:42:54 UTC，这个就是 "某个特定的时间"
	// 在本例子中由调用方来指定，即 startTime
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)
	return
}
func GetID() int64 {
	if nil == node {
		return 0
	}
	return node.Generate().Int64()
}