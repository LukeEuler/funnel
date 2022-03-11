package event

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/LukeEuler/funnel/common"
	"github.com/LukeEuler/funnel/model"
)

func TestDraw(t *testing.T) {
	raw1 := `{"time":"2022-03-17T13:45:12+08:00","region":"ru","level":"info","message":"... timeout ..."}`
	raw2 := `{"time":"2022-03-17T13:45:13+08:00","region":"jp","level":"error","message":"invalid type, ..."}`
	raw3 := `{"time":"2022-03-17T13:45:14+08:00","region":"ru","level":"error","message":"... 404 ..."}`
	raw4 := `{"time":"2022-03-17T13:45:15+08:00","region":"test","level":"error","message":"... 404 ..."}`
	raw5 := `{"time":"2022-03-17T13:45:16+08:00","region":"us","level":"error","message":"... 404 ..."}`
	raw6 := `{"time":"2022-03-17T13:45:17+08:00","region":"ru","level":"error","message":"... 500 ..."}`
	raw7 := `{"time":"2022-03-17T13:45:18+08:00","region":"ru","level":"debu","message":"cost: 13ms, func: abc.efg, job done"}`

	events, err := Draw([]model.EventData{
		common.NewJSONData(json.RawMessage(raw1)),
		common.NewJSONData(json.RawMessage(raw2)),
		common.NewJSONData(json.RawMessage(raw3)),
		common.NewJSONData(json.RawMessage(raw4)),
		common.NewJSONData(json.RawMessage(raw5)),
		common.NewJSONData(json.RawMessage(raw6)),
		common.NewJSONData(json.RawMessage(raw7)),
	}, []model.EventRule{
		&common.EventRuleInfo{
			RuleInfo: &common.RuleInfo{
				ID:      "0_1",
				Name:    "test",
				Content: "region = 'test'",
			},
			Level:    0,
			Mutex:    true,
			Start:    0,
			End:      0,
			Duration: 60,
			Times:    100,
		},
		&common.EventRuleInfo{
			RuleInfo: &common.RuleInfo{
				ID:      "3_1",
				Name:    "404",
				Content: "message > '404'",
			},
			Level:    3,
			Mutex:    false,
			Start:    0,
			End:      0,
			Duration: 60,
			Times:    1,
		},
		&common.EventRuleInfo{
			RuleInfo: &common.RuleInfo{
				ID:      "2_1",
				Name:    "error",
				Content: "level = 'error'",
			},
			Level:    2,
			Mutex:    false,
			Start:    0,
			End:      0,
			Duration: 3,
			Times:    2,
		},
	})
	assert.NoError(t, err)
	assert.Len(t, events, 7)

	expectList := []map[string]bool{
		{},
		{},
		{
			"2_1": true,
			"3_1": true,
		},
		{},
		{
			"2_1": true,
			"3_1": true,
		},
		{
			"2_1": true,
		},
		{},
	}

	drawTest(t, expectList, events)

}

func TestDraw_axx(t *testing.T) {
	raw1 := `{"time":"2022-03-17T13:45:12+08:00","region":"ru","level":"info","message":"... 404 1 ..."}`
	raw2 := `{"time":"2022-03-17T13:45:13+08:00","region":"jp","level":"error","message":"404 type, ..."}`
	raw3 := `{"time":"2022-03-17T13:45:14+08:00","region":"ru","level":"error","message":"... 404 ..."}`
	raw4 := `{"time":"2022-03-17T13:55:14+08:00","region":"test","level":"error","message":"... 999 ..."}`
	raw5 := `{"time":"2022-03-17T13:55:15+08:00","region":"test","level":"error","message":"... 404 ..."}`
	raw6 := `{"time":"2022-03-17T13:55:16+08:00","region":"test","level":"error","message":"... 404 ..."}`

	events, err := Draw([]model.EventData{
		common.NewJSONData(json.RawMessage(raw1)),
		common.NewJSONData(json.RawMessage(raw2)),
		common.NewJSONData(json.RawMessage(raw3)),
		common.NewJSONData(json.RawMessage(raw4)),
		common.NewJSONData(json.RawMessage(raw5)),
		common.NewJSONData(json.RawMessage(raw6)),
	}, []model.EventRule{
		&common.EventRuleInfo{
			RuleInfo: &common.RuleInfo{
				ID:      "3_1",
				Name:    "404",
				Content: "message > '404'",
			},
			Level:    3,
			Mutex:    false,
			Start:    0,
			End:      0,
			Duration: 60,
			Times:    2,
		},
	})
	assert.NoError(t, err)
	assert.Len(t, events, 6)

	expectList := []map[string]bool{
		{
			"3_1": true,
		},
		{
			"3_1": true,
		},
		{
			"3_1": true,
		},
		{},
		{},
		{},
	}

	drawTest(t, expectList, events)
}

func TestDraw_extra(t *testing.T) {
	raw1 := `{"time":"2022-03-17T13:45:12+08:00","region":"ru","level":"info","message":"... 404 1 ..."}`
	raw2 := `{"time":"2022-03-17T13:45:13+08:00","region":"jp","level":"error","message":"404 type, ..."}`
	raw3 := `{"time":"2022-03-17T13:45:14+08:00","region":"ru","level":"error","message":"... 404 ..."}`
	raw4 := `{"time":"2022-03-17T13:45:15+08:00","region":"test","level":"error","message":"... 404 ..."}`
	raw5 := `{"time":"2022-03-17T13:45:16+08:00","region":"us","level":"error","message":"... 404 ..."}`
	raw6 := `{"time":"2022-03-17T13:45:17+08:00","region":"ru","level":"error","message":"... 404 ..."}`
	raw7 := `{"time":"2022-03-17T13:45:18+08:00","region":"ru","level":"debu","message":"404"}`

	events, err := Draw([]model.EventData{
		common.NewJSONData(json.RawMessage(raw1)),
		common.NewJSONData(json.RawMessage(raw2)),
		common.NewJSONData(json.RawMessage(raw3)),
		common.NewJSONData(json.RawMessage(raw4)),
		common.NewJSONData(json.RawMessage(raw5)),
		common.NewJSONData(json.RawMessage(raw6)),
		common.NewJSONData(json.RawMessage(raw7)),
	}, []model.EventRule{
		&common.EventRuleInfo{
			RuleInfo: &common.RuleInfo{
				ID:      "3_1",
				Name:    "404",
				Content: "message > '404'",
			},
			Level:    3,
			Mutex:    false,
			Start:    0,
			End:      0,
			Duration: 2,
			Times:    2,
		},
	})
	assert.NoError(t, err)
	assert.Len(t, events, 7)

	expectList := []map[string]bool{
		{},
		{
			"3_1": true,
		},
		{
			"3_1": true,
		},
		{
			"3_1": true,
		},
		{
			"3_1": true,
		},
		{
			"3_1": true,
		},
		{
			"3_1": true,
		},
	}

	drawTest(t, expectList, events)
}

func TestDraw_timeMatch(t *testing.T) {
	raw1 := `{"time":"2022-03-17T13:45:12+08:00","region":"ru","level":"info","message":"... timeout ..."}`
	raw2 := `{"time":"2022-03-17T13:45:13+08:00","region":"jp","level":"error","message":"invalid type, ..."}`
	raw3 := `{"time":"2022-03-17T13:45:14+08:00","region":"ru","level":"error","message":"... 404 ..."}`
	raw4 := `{"time":"2022-03-17T13:45:15+08:00","region":"test","level":"error","message":"... 404 ..."}`
	raw5 := `{"time":"2022-03-17T13:45:16+08:00","region":"us","level":"error","message":"... 404 ..."}`
	raw6 := `{"time":"2022-03-17T13:45:17+08:00","region":"ru","level":"error","message":"... 500 ..."}`
	raw7 := `{"time":"2022-03-17T13:45:18+08:00","region":"ru","level":"debu","message":"cost: 13ms, func: abc.efg, job done"}`

	events, err := Draw([]model.EventData{
		common.NewJSONData(json.RawMessage(raw1)),
		common.NewJSONData(json.RawMessage(raw2)),
		common.NewJSONData(json.RawMessage(raw3)),
		common.NewJSONData(json.RawMessage(raw4)),
		common.NewJSONData(json.RawMessage(raw5)),
		common.NewJSONData(json.RawMessage(raw6)),
		common.NewJSONData(json.RawMessage(raw7)),
	}, []model.EventRule{
		&common.EventRuleInfo{
			RuleInfo: &common.RuleInfo{
				ID:      "2_1",
				Name:    "error",
				Content: "level = 'error'",
			},
			Level:    2,
			Mutex:    false,
			Start:    1,
			End:      2,
			Duration: 3000,
			Times:    0,
		},
		&common.EventRuleInfo{
			RuleInfo: &common.RuleInfo{
				ID:      "2_2",
				Name:    "error",
				Content: "level = 'error'",
			},
			Level:    2,
			Mutex:    false,
			Start:    1700000000,
			End:      1800000000,
			Duration: 3000,
			Times:    0,
		},
		&common.EventRuleInfo{
			RuleInfo: &common.RuleInfo{
				ID:      "2_3",
				Name:    "error",
				Content: "level = 'error'",
			},
			Level:    2,
			Mutex:    false,
			Start:    10,
			End:      5,
			Duration: 3000,
			Times:    0,
		},
	})
	assert.NoError(t, err)
	assert.Len(t, events, 7)

	expectList := []map[string]bool{
		{},
		{
			"2_3": true,
		},
		{
			"2_3": true,
		},
		{
			"2_3": true,
		},
		{
			"2_3": true,
		},
		{
			"2_3": true,
		},
		{},
	}

	drawTest(t, expectList, events)
}

func drawTest(t *testing.T, expectList []map[string]bool, events []model.Event) {
	for i, item := range expectList {
		match := events[i].GetHitRule()
		assert.Len(t, match, len(item), "%d. expect match %d rule, but got %d", i, len(item), len(match))
		for name := range item {
			assert.Contains(t, match, name, "%d. should not contains %s", i, name)
		}
	}
}
