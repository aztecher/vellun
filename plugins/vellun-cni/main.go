package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	cniVersion "github.com/containernetworking/cni/pkg/version"

	"github.com/aztecher/vellun/pkg/version"
	"github.com/aztecher/vellun/plugins/vellun-cni/cmd"
)

func main() {
	cmd := cmd.NewCmd()
	skel.PluginMainFuncs(
		cmd.CNIFuncs(),
		cniVersion.PluginSupports(
			"0.1.0",
			"0.2.0",
			"0.3.0",
			"0.3.1",
			"0.4.0",
			"1.0.0",
			"1.1.0",
		),
		"Vellun CNI plugin "+version.Version,
	)
}
