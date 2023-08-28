package initializers

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

var SnowflakeNode *snowflake.Node

func InitSnowflakeNode(config *Configuration) {
	node, err := snowflake.NewNode(config.SnowflakeNode)
	if err != nil {
		log.Panicln("failed to init snowflake node", err)
	}
	SnowflakeNode = node
}
