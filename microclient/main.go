package main

import (
	"context"
	"log"
	"time"

	"github.com/micro/go-micro"

	holiday "github.com/pijalu/go.hands.one/micro/proto/holiday"

	types "github.com/golang/protobuf/ptypes"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("be.ordina.micro.client.micro"),
	)
	service.Init()
	holidayClient := holiday.NewHolidayClient("be.ordina.micro.srv.micro", service.Client())

	date := time.Date(2012, time.December, 21, 6, 6, 6, 0, time.UTC)
	ts, err := types.TimestampProto(date)
	if err != nil {
		panic(err)
	}

	var openDays int64 = 2
	resp, err := holidayClient.GetNextHoliday(context.TODO(), &holiday.HolidayRequest{
		RequestDate: ts,
		OpenDays:    openDays,
		IsoCountry:  "BE",
	})
	if err != nil {
		panic(err)
	}
	rTs, err := types.Timestamp(resp.ReplyDate)
	if err != nil {
		panic(err)
	}

	log.Printf("Requested: %s + %d => %s", date.String(), openDays, rTs.String())
}
