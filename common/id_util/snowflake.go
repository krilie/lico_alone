package id_util

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/krilie/lico_alone/common/log"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		log.Fatal("init snowflake err: ", node)
	}
}

func NextSnowflakeId() snowflake.ID {
	return node.Generate()
}
