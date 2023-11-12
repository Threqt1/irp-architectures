package main

import (
	"log"

	"github.com/Threqt1/architecture-go/user-microservice/api"
	"github.com/Threqt1/architecture-go/user-microservice/library/snowflake"
)

func main() {
	snowflake, err := snowflake.CreateSnowflakeProvider()
	if err != nil {
		log.Fatal(err)
	}

	apiv1 := api.CreateAPIv1(snowflake)
	apiv1.Start()
}
