package model

type Data interface {
	KeyExist(key string) bool
	GetValueString(key string) (string, bool)
	ValueEqual(key string, value string) bool
	ValueContains(key string, value string) bool
	SetRule(Rule, bool)
	Match(Rule) bool
}

type Rule interface {
	GetUniqueID() string
	GetName() string
	GetContent() string
}

type EventData interface {
	Data
	GetTime() int64
}

/*
基于时间的 匹配规则的接口

- level+mutex 定义了过滤规则，level 较小的优先过滤 Data。如果 d1 命中了 r1，且 r1.IsMutex()。那么，d1 将不会命中任何 大于 r1.GetLevel() 的 规则。
- 每条规则都有生效时间 [start, end]
- duration+times 规定了命中阈值
*/
type EventRule interface {
	Rule

	// 规则匹配之间的忽略规则
	GetLevel() int
	IsMutex() bool

	// 定义规则生效的时间范围(timestamps), 0 值代表不考虑范围
	GetStart() int64
	GetEnd() int64

	// 定义了报警阈值: 在检测的时间范围内(Duration), 如果匹配的事件数不大于 阈值(times), 则可以忽略
	GetDuration() int64 // > 0
	GetTimes() int      // >= 0
}

type Event interface {
	EventData
	GetData() EventData
	GetHitRule() map[string]EventRule
	GetMissRule() map[string]EventRule
	Valid() bool
}
