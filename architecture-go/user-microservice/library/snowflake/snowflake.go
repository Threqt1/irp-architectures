package snowflake

import (
	"strconv"

	"github.com/bwmarrin/snowflake"
)

type SnowflakeProvider struct {
	node *snowflake.Node
}

func (sp *SnowflakeProvider) Generate() string {
	return strconv.FormatInt(sp.node.Generate().Int64(), 10)
}

func CreateSnowflakeProvider() (SnowflakeProvider, error) {
	snowflakeProvider := SnowflakeProvider{}
	node, err := snowflake.NewNode(1)

	if err != nil {
		return snowflakeProvider, err
	}

	snowflakeProvider.node = node
	return snowflakeProvider, nil
}
