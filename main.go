package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	vmName := flag.String("vm", "", "-vm <name>")
	resourceGroupName := flag.String("rg", "", "-rg <RGName>")
	command := flag.String("c", "", "command  start or stop")

	flag.Parse()

	if *command == "" {
		fmt.Println("You must supply an command  start or stop (-c <cmd>")
		return
	}
	if *vmName == "" {
		fmt.Println("You must supply VM name  (-vm <vmName>")
		return
	}
	if *resourceGroupName == "" {
		fmt.Println("You must supply Resource Group Name (-rg <resGrp>")
		return
	}

	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a compute client
	computeClient := compute.NewVirtualMachinesClient(subscriptionID)
	computeClient.Authorizer = authorizer

	// Set the resource group name and virtual machine name

	ctx := context.Background()

	if *command == "stop" {
		// stop call
		fmt.Println("Stoping Virtual machine")

		stopFuture, err := computeClient.Deallocate(ctx, *resourceGroupName, *vmName)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = stopFuture.WaitForCompletionRef(ctx, computeClient.Client)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Virtual machine stopped.")
	}
	// stop end

	if *command == "start" {
		// start call
		fmt.Println("Starting Virtual machine")

		startFuture, err := computeClient.Start(ctx, *resourceGroupName, *vmName)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = startFuture.WaitForCompletionRef(ctx, computeClient.Client)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Virtual machine started.")

	}
	// start end

}
