package common

import (
	"encoding/json"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/LukeEuler/funnel/model"
)

// implement model.Data
type JSONData struct {
	Content   json.RawMessage
	checkList map[string]bool
}

func NewJSONData(conetnt json.RawMessage) *JSONData {
	return &JSONData{
		Content:   conetnt,
		checkList: make(map[string]bool),
	}
}

func (d *JSONData) KeyExist(key string) bool {
	value := gjson.GetBytes(d.Content, key)
	return value.Exists()
}

func (d *JSONData) getValueString(key string) (string, bool) {
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
	jValue, ok := d.getValueString(key)
	if !ok {
		return false
	}
	return value == jValue
}

func (d *JSONData) ValueContains(key string, value string) bool {
	jValue, ok := d.getValueString(key)
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
