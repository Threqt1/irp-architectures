package main

import "github.com/Threqt1/architecture-go/api-gateway/api"

func main() {
	apiv1 := api.CreateAPIv1()
	apiv1.Start()
}
