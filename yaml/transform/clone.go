package transform

import "github.com/drone/drone-exec/yaml"

const clone = "clone"

// Clone transforms the Yaml to include a clone step.
func Clone(c *yaml.Config, plugin string) error {
	for _, p := range c.Pipeline {
		if p.Name == clone {
			return nil
		}
	}
	if plugin == "" {
		plugin = "git"
	}

	s := &yaml.Container{
		Image: plugin,
		Name:  clone,
	}

	c.Pipeline = append([]*yaml.Container{s}, c.Pipeline...)
	return nil
}
