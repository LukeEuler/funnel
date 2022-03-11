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
