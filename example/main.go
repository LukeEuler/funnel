package main

import (
	"encoding/json"
	"fmt"

	"github.com/LukeEuler/funnel/common"
	"github.com/LukeEuler/funnel/event"
	"github.com/LukeEuler/funnel/model"
)

func main() {
	rawData := make([]json.RawMessage, 0, 400)
	err := json.Unmarshal([]byte(raw), &rawData)
	if err != nil {
		panic(err)
	}

	datas := make([]model.EventData, 0, len(rawData))
	for _, item := range rawData {
		datas = append(datas, common.NewJSONData(item))
	}
	fmt.Printf("total %d pieces of data\n", len(datas))

	rules := []model.EventRule{
		&common.EventRuleInfo{
			RuleInfo: &common.RuleInfo{
				ID:      "0_1",
				Name:    "no test",
				Content: "platform > 'test'",
			},
			Level:    0,
			Mutex:    true,
			Start:    0,
			End:      0,
			Duration: 60,
			Times:    5,
		},
		&common.EventRuleInfo{
			RuleInfo: &common.RuleInfo{
				ID:      "1_1",
				Name:    "502",
				Content: "message > '502'",
			},
			Level:    1,
			Mutex:    false,
			Start:    0,
			End:      0,
			Duration: 5,
			Times:    1,
		},
	}
	events, _ := event.Draw(datas, rules)
	fmt.Printf("get %d events\n", len(events))

	for i, item := range events {
		fmt.Printf("%d, valid: %v. hit %d, miss %d \n", i+1, item.Valid(), len(item.GetHitRule()), len(item.GetMissRule()))
	}
}

const raw = `
[
    {
        "func": "log.ErrWithFields",
        "host": "a.bb.com",
        "time": "2022-04-20T05:51:01Z",
        "message": "error code: -100, data not found",
        "component": "client abc",
        "devlang": "golang",
        "platform": "test2",
        "level": "error"
    },
    {
        "platform": "dev",
        "func": "log.ErrWithFields",
        "level": "error",
        "host": "a.bb.com",
        "time": "2022-04-20T13:51:02+08:00",
        "message": "502 Bad Gateway",
        "component": "client abc"
    },
    {
        "message": "502 Bad Gateway",
        "func": "log.ErrWithFields",
        "level": "error",
        "time": "2022-04-20T13:51:04+08:00",
        "platform": "prd",
        "host": "c.dd.com",
        "component": "consumer xxx"
    },
    {
        "time": "2022-04-20T13:51:10+08:00",
        "component": "consumer xxx",
        "func": "log.ErrWithFields",
        "host": "e.ff.com",
        "level": "error",
        "platform": "test 1",
        "message": "404 Not Found: 404 page not found"
    },
    {
        "func": "log.ErrWithFields",
        "level": "error",
        "time": "2022-04-20T05:52:08Z",
        "devlang": "golang",
        "message": "error code: -100, data not found",
        "platform": "test2",
        "host": "h.ii.com",
        "component": "client abc"
    },
    {
        "func": "log.ErrWithFields",
        "message": "context deadline exceeded (Client.Timeout exceeded while awaiting headers)",
        "host": "a.bb.com",
        "time": "2022-04-20T13:52:09+08:00",
        "component": "client abc",
        "platform": "dev",
        "level": "error"
    },
    {
        "time": "2022-04-20T13:52:10+08:00",
        "component": "consumer xxx",
        "func": "log.ErrWithFields",
        "host": "c.dd.com",
        "level": "error",
        "platform": "prd",
        "message": "502 Bad Gateway"
    }
]`
