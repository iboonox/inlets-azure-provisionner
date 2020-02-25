package provision

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

const (
	resourceGroupName     = "InletsRG"
	resourceGroupLocation = "westeurope"
	deploymentName        = "InletsExitNode"
	templateFile          = "templates/vm-inlets-exitnode-template.json"
	parametersFile        = "templates/vm-inlets-exitnode-params.json"
)

var (
	ctx = context.Background()
)

// AzureProvisioner provision a VM on azure.com
type AzureProvisioner struct {
	client autorest.Authorizer
}

// NewAzureProvisioner with an environement variables
func NewAzureProvisioner() (*AzureProvisioner, error) {

	client, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}

	return &AzureProvisioner{
		client: client,
	}, nil
}

// Provision creates an exit node
func (p *AzureProvisioner) Provision(hostname string, userdata string) error {

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	// Create a resource group
	log.Printf("Starting creating resource group: %s", resourceGroupName)
	rgClient := resources.NewGroupsClient(subscriptionID)
	rgClient.Authorizer = p.client
	_, err := rgClient.CreateOrUpdate(context.Background(), resourceGroupName, resources.Group{
		Location: to.StringPtr(resourceGroupLocation),
	})
	if err != nil {
		return err
	}
	fmt.Printf("created resource group %s\n", resourceGroupName)

	// Create deployment
	template, err := readJSON(templateFile)
	if err != nil {
		return err
	}
	params, err := readJSON(parametersFile)
	if err != nil {
		return err
	}

	// Override default params
	(*params)["vm_password"] = map[string]string{
		"value": os.Getenv("AZURE_CLIENT_SECRET"),
	}
	(*params)["virtualMachines_InletsVM_name"] = map[string]string{
		"value": hostname,
	}
	(*params)["customData"] = map[string]string{
		"value": userdata,
	}

	deploymentsClient := resources.NewDeploymentsClient(subscriptionID)
	deploymentsClient.Authorizer = p.client

	deploymentFuture, err := deploymentsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		deploymentName,
		resources.Deployment{
			Properties: &resources.DeploymentProperties{
				Template:   template,
				Parameters: params,
				Mode:       resources.Incremental,
			},
		},
	)
	if err != nil {
		return err
	}
	err = deploymentFuture.Future.WaitForCompletionRef(ctx, deploymentsClient.BaseClient.Client)
	if err != nil {
		return err
	}
	return nil

}

func readJSON(path string) (*map[string]interface{}, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	contents := make(map[string]interface{})
	json.Unmarshal(data, &contents)
	return &contents, nil
}
