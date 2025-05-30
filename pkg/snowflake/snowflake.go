package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(startTime string, machineId int64) (err error) {
	st, err := time.Parse("2006-07-01", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1_000_000
	node, err = snowflake.NewNode(machineId)
	return
}

func GenId() int64 {
	return node.Generate().Int64()
}
