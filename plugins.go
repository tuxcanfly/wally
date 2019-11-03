package main

import "errors"

var errunregistered = errors.New("plugin not registered")

type wallyPlugin interface {
	URL() (string, error)
}

type pluginManager map[string]wallyPlugin

func (p pluginManager) register(name string, plugin wallyPlugin) {
	p[name] = plugin
}

func (p pluginManager) get(name string) (wallyPlugin, error) {
	plugin, ok := p[name]
	if !ok {
		return nil, errunregistered
	}
	return plugin, nil
}

func (p pluginManager) list() []string {
	var plugins []string
	for k := range p {
		plugins = append(plugins, k)
	}
	return plugins
}
