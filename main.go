package main

import (
	"github.com/cryptogrampus/terraform-provider-grafana/grafana"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: grafana.Provider})
}
