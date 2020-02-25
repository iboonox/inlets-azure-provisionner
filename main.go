package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/iboonox/inlets-azure-provisionner/pkg/provision"
)

func main() {

	var userDataFile string
	var userdata string
	var hostname string

	flag.StringVar(&userDataFile, "userdata-file", "", "Apply user-data from a file to configure the host")
	flag.StringVar(&hostname, "hostname", "provision-example", "Name for the host")
	flag.Parse()

	// Init Azure provisionner
	provisioner, err := provision.NewAzureProvisioner()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	if len(userDataFile) > 0 {
		res, err := ioutil.ReadFile(userDataFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		userdata = string(res)
	}

	// Provision inlets exit node
	err = provisioner.Provision(hostname, userdata)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

}
