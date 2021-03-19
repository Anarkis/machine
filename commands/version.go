package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/rancher/machine/libmachine"
	"github.com/rancher/machine/libmachine/mcndockerclient"
)

func cmdVersion(c CommandLine, api libmachine.API) error {
	return printVersion(c, api, os.Stdout)
}

func printVersion(c CommandLine, api libmachine.API, out io.Writer) error {
	if len(c.Args()) == 0 {
		c.ShowVersion()
		return nil
	}

	if len(c.Args()) != 1 {
		return ErrExpectedOneMachine
	}

	host, err := api.Load(c.Args().First())
	if err != nil {
		return err
	}

	if host.HostOptions.AuthOptions != nil {
		version, err := mcndockerclient.DockerVersion(host)
		if err != nil {
			return err
		}

		fmt.Fprintln(out, version)
	} else {
		fmt.Fprintln(out, "Docker was not installed on machine")
	}

	return nil
}
