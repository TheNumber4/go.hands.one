package handler

import (
	"fmt"
	"log"
	"time"

	types "github.com/golang/protobuf/ptypes"
	h "github.com/pijalu/go.hands.one/micro/proto/holiday"

	"golang.org/x/net/context"
)

type Holiday struct{}

func isHoliday(t *time.Time) bool {
	day := t.Weekday()
	return day == time.Saturday ||
		day == time.Sunday
}

func (h *Holiday) GetNextHoliday(ctx context.Context, in *h.HolidayRequest, out *h.HolidayReply) error {
	ts, err := types.Timestamp(in.RequestDate)
	if err != nil {
		return fmt.Errorf("Failed to convert timestamp: %v: %v", in.RequestDate, err)
	}
	inTs := ts

	var inc time.Duration
	cnt := in.OpenDays
	if cnt < 0 {
		inc = -1
		cnt = -cnt
	} else {
		inc = 1
	}

	for cnt > 0 {
		if !isHoliday(&ts) {
			cnt -= 1
		}
		ts = ts.Add(time.Hour * 24 * inc)
	}
	tspt, err := types.TimestampProto(ts)
	if err != nil {
		return fmt.Errorf("Failed to convert timestamp to proto: %v: %v", ts, err)
	}
	out.ReplyDate = tspt

	log.Printf("Request %s + %d => Reply %s", inTs.String(), in.OpenDays, ts.String())
	return nil
}
