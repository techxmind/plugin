package plugin

type PluginProcessor struct {
	scope   *Scope
	name    string
	before  string
	after   string
	replace bool
	remove  bool
	plugin  *Plugin
}

func (r *PluginProcessor) Before(before string) *PluginProcessor {
	r.before = before
	return r
}

func (r *PluginProcessor) After(after string) *PluginProcessor {
	r.after = after
	return r
}

func (r *PluginProcessor) Replace(replace string, pl Plugin) *PluginProcessor {
	r.name = replace
	r.replace = true
	r.plugin = &pl
	r.scope.processors = append(r.scope.processors, r)
	r.scope.reorder()
	return r
}

func (r *PluginProcessor) Remove(remove string) *PluginProcessor {
	r.name = remove
	r.remove = true
	r.scope.processors = append(r.scope.processors, r)
	r.scope.reorder()
	return r
}

func (r *PluginProcessor) Register(name string, pl Plugin) {
	r.name = name
	r.plugin = &pl
	r.scope.processors = append(r.scope.processors, r)
	r.scope.reorder()
}
