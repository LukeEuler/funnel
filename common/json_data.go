package common

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/tidwall/gjson"

	"github.com/LukeEuler/funnel/model"
)

// implement model.EventData
type JSONData struct {
	Content   json.RawMessage
	checkList map[string]bool

	GetTimeFunc func(json.RawMessage) int64
}

func NewJSONData(conetnt json.RawMessage) *JSONData {
	return &JSONData{
		Content:     conetnt,
		checkList:   make(map[string]bool),
		GetTimeFunc: DefaultGetTime,
	}
}

func (d *JSONData) KeyExist(key string) bool {
	value := gjson.GetBytes(d.Content, key)
	return value.Exists()
}

func (d *JSONData) GetValueString(key string) (string, bool) {
	if !d.KeyExist(key) {
		return "", false
	}

	value := gjson.GetBytes(d.Content, key)
	if value.IsObject() || value.IsArray() {
		return "", false
	}
	return value.String(), true
}

func (d *JSONData) ValueEqual(key string, value string) bool {
	jValue, ok := d.GetValueString(key)
	if !ok {
		return false
	}
	return value == jValue
}

func (d *JSONData) ValueContains(key string, value string) bool {
	jValue, ok := d.GetValueString(key)
	if !ok {
		return false
	}
	return strings.Contains(jValue, value)
}

func (d *JSONData) SetRule(rule model.Rule, match bool) {
	d.checkList[rule.GetUniqueID()] = match
}

func (d *JSONData) Match(rule model.Rule) bool {
	match, checked := d.checkList[rule.GetUniqueID()]
	if !checked {
		return false
	}
	return match
}

func (d *JSONData) GetTime() int64 {
	return d.GetTimeFunc(d.Content)
}

func (d *JSONData) String() string {
	return string(d.Content)
}

func DefaultGetTime(content json.RawMessage) int64 {
	key := "time"
	value := gjson.GetBytes(content, key)
	if !value.Exists() {
		return 0
	}

	if value.IsObject() || value.IsArray() {
		return 0
	}

	t, err := time.Parse(time.RFC3339, value.String())
	if err != nil {
		return 0
	}
	return t.Unix()
}
