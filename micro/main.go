package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/pijalu/go.hands.one/micro/handler"

	holiday "github.com/pijalu/go.hands.one/micro/proto/holiday"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("be.ordina.micro.srv.micro"),
		micro.Version("latest"),
	)

	// Register Handler
	holiday.RegisterHolidayHandler(service.Server(), new(handler.Holiday))

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
