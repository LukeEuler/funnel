package event

import (
	"sort"
	"time"

	"github.com/LukeEuler/funnel/common"
	"github.com/LukeEuler/funnel/core"
	"github.com/LukeEuler/funnel/model"
)

type event struct {
	model.EventData

	hitMap  map[string]model.EventRule
	missMap map[string]model.EventRule

	time int64

	hasMutex bool
	level    int
}

func newEvent(data model.EventData) *event {
	return &event{
		EventData: data,
		hitMap:    make(map[string]model.EventRule, 10),
		missMap:   make(map[string]model.EventRule, 10),
		time:      data.GetTime(),
	}
}

func (e *event) GetData() model.EventData {
	return e.EventData
}

func (e *event) GetHitRule() map[string]model.EventRule {
	temp := make(map[string]model.EventRule, len(e.hitMap))
	for k, v := range e.hitMap {
		temp[k] = v
	}
	return temp
}

func (e *event) GetMissRule() map[string]model.EventRule {
	temp := make(map[string]model.EventRule, len(e.missMap))
	for k, v := range e.missMap {
		temp[k] = v
	}
	return temp
}

/*
- 如果所有的规则都都未匹配, 那么这条数据应该参与报警. 此时发生的事情, 是预期之外的
- 如果存在部分规则与该数据匹配, 但所有规则都不命中. 那么, 这条数据可以在报警中忽略
- 如果存在部分规则与该数据匹配, 且这其中存在规则命中了. 那么, 这条数据应报警
*/
func (e *event) Valid() bool {
	if len(e.hitMap) > 0 {
		return true
	}
	return len(e.missMap) == 0
}

func (e *event) setMatchFalse(rule model.EventRule) {
	e.missMap[rule.GetUniqueID()] = rule
	delete(e.hitMap, rule.GetUniqueID())
	if !rule.IsMutex() {
		return
	}
	if !e.hasMutex {
		e.hasMutex = true
		e.level = rule.GetLevel()
		return
	}

	/*
		程序不应该运行到此处，所有rule都应该按照 level 由小到到到方式去做过滤
		如果真的不了解原则而误用了，此处也只能尽力处理了，但是结果可能会有诡异情况
		本程序不做负责
	*/
	if rule.GetLevel() < e.level {
		e.level = rule.GetLevel()
	}
}

func (e *event) hit(rule model.EventRule) bool {
	rm, ok := e.hitMap[rule.GetUniqueID()]
	return ok && rm != nil
}

func (e *event) shouldIgnoreForMutex(rule model.EventRule) bool {
	if !e.hasMutex {
		return false
	}
	return rule.GetLevel() > e.level
}

func Draw(datas []model.EventData, rules []model.EventRule) ([]model.Event, error) {
	common.TimeConsume(time.Now())

	// 初始化
	sort.SliceStable(rules, func(i, j int) bool {
		return rules[i].GetLevel() < rules[j].GetLevel()
	})
	sort.SliceStable(datas, func(i, j int) bool {
		return datas[i].GetTime() < datas[j].GetTime()
	})

	events := make([]*event, 0, len(datas))
	for _, data := range datas {
		events = append(events, newEvent(data))
	}

	for _, rule := range rules {
		// 先按照 rule 内容着色，暂不考虑时间
		for _, event := range events {
			if !timeMatch(event.EventData, rule) {
				continue
			}
			// shouldIgnoreForMutex 与 SetMatchFalse 联合，做 mutex 过滤
			if event.shouldIgnoreForMutex(rule) {
				continue
			}
			match, err := core.Match(rule, event.EventData)
			if err != nil {
				return nil, err
			}
			if match {
				event.hitMap[rule.GetUniqueID()] = rule
			}
		}

		// 然后根据 rule 的时间特性进行着色清洗
		hr := checkDurationLimit(events, rule)
		for index, event := range events {
			if event.hit(rule) {
				if !hr.hit(index) {
					event.setMatchFalse(rule)
				}
			}
		}
	}

	result := make([]model.Event, 0, len(events))
	for _, item := range events {
		result = append(result, item)
	}
	return result, nil
}

func timeMatch(data model.EventData, rule model.EventRule) bool {
	// invalid time setted, we should not ignore the rule.
	if rule.GetEnd() < rule.GetStart() {
		return true
	}
	if rule.GetStart() > 0 && rule.GetStart() > data.GetTime() {
		return false
	}
	if rule.GetEnd() > 0 && rule.GetEnd() < data.GetTime() {
		return false
	}
	return true
}

type hitRecord struct {
	content []struct {
		begin int
		end   int
	}
}

func newHitRecord() *hitRecord {
	return &hitRecord{content: make([]struct {
		begin int
		end   int
	}, 0, 2)}
}

func (h *hitRecord) hit(index int) bool {
	for _, item := range h.content {
		if item.begin <= index && index <= item.end {
			return true
		}
	}
	return false
}

func (h *hitRecord) add(begin, end int) {
	h.content = append(h.content, struct {
		begin int
		end   int
	}{begin: begin, end: end})
}

// 要求 events 是按照时间顺序做过排序的
func checkDurationLimit(events []*event, rule model.EventRule) *hitRecord {
	result := newHitRecord()
	lenght := len(events)
	if lenght <= rule.GetTimes() {
		return result
	}

	hit := false
	begin, end, count := 0, 0, 0
	var endTime int64
	for index := lenght - 1; index >= 0; index-- {
		item := events[index]
		if !item.hit(rule) {
			continue
		}

		if !hit {
			hit = true
			end = index
			endTime = item.time
		}

		// TODO
		if endTime-item.time > rule.GetDuration() {
			if count > rule.GetTimes() {
				result.add(begin, end)
				begin = index
				end = index
				endTime = item.time
				count = 1
			} else {
				count++
				for i := end - 1; i >= index; i-- {
					if !events[i].hit(rule) {
						continue
					}
					count--
					if events[i].time-item.time > rule.GetDuration() {
						continue
					}
					end = i
					endTime = events[i].time
					break
				}
			}
			continue
		}
		// endTime-item.time <= rule.GetDuration()
		begin = index
		count++
	}
	if count > rule.GetTimes() {
		result.add(begin, end)
	}
	return result
}
