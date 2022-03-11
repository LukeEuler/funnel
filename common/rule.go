package common

// implement model.Rule
type RuleInfo struct {
	ID      string
	Name    string
	Content string
}

func (r *RuleInfo) GetUniqueID() string {
	return r.ID
}

func (r *RuleInfo) GetName() string {
	return r.Name
}

func (r *RuleInfo) GetContent() string {
	return r.Content
}

// implement model.EventRule
type EventRuleInfo struct {
	*RuleInfo

	Level int
	Mutex bool

	Start int64
	End   int64

	Duration int64
	Times    int
}

func (r EventRuleInfo) GetLevel() int {
	return r.Level
}

func (r EventRuleInfo) IsMutex() bool {
	return r.Mutex
}

func (r EventRuleInfo) GetStart() int64 {
	return r.Start
}

func (r EventRuleInfo) GetEnd() int64 {
	return r.End
}

func (r EventRuleInfo) GetDuration() int64 {
	if r.Duration < 1 {
		return 1
	}
	return r.Duration
}

func (r EventRuleInfo) GetTimes() int {
	if r.Times < 0 {
		return 0
	}
	return r.Times
}
