package core

import (
	"encoding/json"
	"testing"

	"github.com/LukeEuler/funnel/common"
	"github.com/LukeEuler/funnel/model"
)

func TestMatch(t *testing.T) {
	rule1 := &common.RuleInfo{
		ID:      "9527",
		Name:    "xxx",
		Content: "region = 'ru' & message > '404'",
	}
	rule2 := &common.RuleInfo{
		ID:      "110",
		Name:    "xxxxcsdcx",
		Content: "region = 'jp' & region = 'jp' & message > '404'",
	}
	rule3 := &common.RuleInfo{
		ID:      "999",
		Name:    "dnakcsadc",
		Content: "msg+ & message > '404'",
	}
	rule4 := &common.RuleInfo{
		ID:      "110",
		Name:    "dsakvs",
		Content: "(!(region = 'us' | region = 'eu')) & (message > '403' | message = '110') ",
	}
	rule5 := &common.RuleInfo{
		ID:      "114",
		Name:    "kcsdlcmasc",
		Content: "a.bb+",
	}
	rule6 := &common.RuleInfo{
		ID:      "9527",
		Name:    "xxx",
		Content: "region = 'ru' & message in '404'",
	}

	raw1 := `{"region":"ru","message": "xxxx 404 xxxxx"}`
	raw2 := `{"region":"jp","message": "xxxx 404 xxxxx"}`
	raw3 := `{"region":"jp","message": "110"}`
	raw4 := `{"a":{"bb":{"ccc":123}}}`

	tests := []struct {
		name    string
		rule    model.Rule
		data    model.Data
		want    bool
		wantErr bool
	}{
		{
			"test 1",
			rule1,
			common.NewJSONData(json.RawMessage(raw1)),
			true,
			false,
		},
		{
			"test 2",
			rule1,
			common.NewJSONData(json.RawMessage(raw2)),
			false,
			false,
		},
		{
			"test 3",
			rule2,
			common.NewJSONData(json.RawMessage(raw2)),
			true,
			false,
		},
		{
			"test 4",
			rule3,
			common.NewJSONData(json.RawMessage(raw2)),
			false,
			false,
		},
		{
			"test 5",
			rule4,
			common.NewJSONData(json.RawMessage(raw3)),
			true,
			false,
		},
		{
			"test 6",
			rule5,
			common.NewJSONData(json.RawMessage(raw4)),
			true,
			false,
		},
		{
			"test 7",
			rule6,
			common.NewJSONData(json.RawMessage(raw1)),
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Match(tt.rule, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Match() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
