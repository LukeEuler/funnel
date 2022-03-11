package model

type Data interface {
	KeyExist(key string) bool
	ValueEqual(key string, value string) bool
	ValueContains(key string, value string) bool
	SetRule(Rule, bool)
	// GetRule(Rule) (match bool, checked bool)
	Match(Rule) bool
}

type Rule interface {
	GetUniqueID() string
	GetName() string
	GetContent() string
}
