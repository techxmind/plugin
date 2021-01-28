package plugin

import (
	"context"
)

type Scope struct {
	name       string
	plugins    []*Plugin
	processors []*PluginProcessor
}

func NewScope(name string) *Scope {
	return &Scope{
		name:       name,
		plugins:    make([]*Plugin, 0),
		processors: make([]*PluginProcessor, 0),
	}
}

func (s *Scope) Plugin() *PluginProcessor {
	return &PluginProcessor{
		scope: s,
	}
}

func (s *Scope) Get(name string) Plugin {
	for _, p := range s.processors {
		if p.name == name {
			return *p.plugin
		}
	}

	return nil
}

func (s *Scope) Execute(ctx context.Context, data interface{}) {
	message := messagePool.Get().(*Message)
	message.Data = data
	for _, plugin := range s.plugins {
		(*plugin)(ctx, message)
	}
	messagePool.Put(message)
}

// getRIndex get right index from string slice
func getRIndex(strs []string, str string) int {
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] == str {
			return i
		}
	}
	return -1
}

func (s *Scope) reorder() {
	var (
		pls                   = s.processors
		allNames, sortedNames []string
		sortPluginProcessor   func(p *PluginProcessor)
	)

	for _, pl := range pls {
		if index := getRIndex(allNames, pl.name); index > -1 && !pl.replace && !pl.remove {
			logger.Warnf("[plugin-scope][%s] duplicated plugin `%v`", s.name, pl.name)
		}
		allNames = append(allNames, pl.name)
	}

	sortPluginProcessor = func(p *PluginProcessor) {
		if getRIndex(sortedNames, p.name) != -1 { // if sorted
			return
		}
		if p.before != "" { // if defined before plugin
			if index := getRIndex(sortedNames, p.before); index != -1 {
				// if before plugin already sorted, append current plugin just after it
				sortedNames = append(sortedNames[:index], append([]string{p.name}, sortedNames[index:]...)...)
			} else if index := getRIndex(allNames, p.before); index != -1 {
				// if before plugin exists but haven't sorted, append current plugin to last
				sortedNames = append(sortedNames, p.name)
				sortPluginProcessor(pls[index])
			}
		}

		if p.after != "" { // if defined after plugin
			if index := getRIndex(sortedNames, p.after); index != -1 {
				// if after plugin already sorted, append current plugin just before it
				sortedNames = append(sortedNames[:index+1], append([]string{p.name}, sortedNames[index+1:]...)...)
			} else if index := getRIndex(allNames, p.after); index != -1 {
				// if after plugin exists but haven't sorted
				pl := pls[index]
				// set after plugin's before callback to current callback
				if pl.before == "" {
					pl.before = p.name
				}
				sortPluginProcessor(pl)
			}
		}

		// if current plugin haven't been sorted, append it to last
		if getRIndex(sortedNames, p.name) == -1 {
			sortedNames = append(sortedNames, p.name)
		}
	}

	for _, pl := range pls {
		sortPluginProcessor(pl)
	}

	var sortedPlugins []*Plugin
	for _, name := range sortedNames {
		if index := getRIndex(allNames, name); !pls[index].remove {
			sortedPlugins = append(sortedPlugins, pls[index].plugin)
		}
	}

	s.plugins = sortedPlugins
}
