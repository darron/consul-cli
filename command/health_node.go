package command

import (
	"strings"
)

type HealthNodeCommand struct {
	Meta
}

func (c *HealthNodeCommand) Help() string {
	helpText := `
Usage: consul-cli health-node [options] nodeName

  Get the health info for a node

Options: 

` + c.ConsulHelp()

	return strings.TrimSpace(helpText)
}

func (c *HealthNodeCommand) Run(args []string) int {
	c.AddDataCenter()
	flags := c.Meta.FlagSet()
	flags.Usage = func() { c.UI.Output(c.Help()) }

	if err := flags.Parse(args); err != nil {
		return 1
	}

	extra := flags.Args()
	if len(extra) < 1 {
		c.UI.Error("Node name must be specified")
		c.UI.Error("")
		c.UI.Error(c.Help())
		return 1
	}
	node := extra[0]

	client, err := c.Client()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	queryOpts := c.QueryOptions()
	healthClient := client.Health()

	h, _, err := healthClient.Node(node, queryOpts)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.OutputJSON(h, true)

	return 0
}

func (c *HealthNodeCommand) Synopsis() string {
	return "Get the health info for a node"
}
