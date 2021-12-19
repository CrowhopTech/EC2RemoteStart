package main

import (
	"bufio"
	"fmt"
	"os"

	"crowhop.tech/gaming-aws-starter/pkg/gamingvm"
)

func main() {
	userName, err := gamingvm.GetCurrentUserName()
	if err != nil {
		fmt.Println("Failed to get current user. Do you have credentials at ~/.aws/credentials under the default name?")
		fmt.Println(err)
		os.Exit(1)
		return
	}

	instance, instanceName, err := gamingvm.GetGamingVMForUser(userName)
	if err != nil {
		fmt.Println("Failed to get gaming VM for current user.")
		fmt.Println(err)
		os.Exit(1)
		return
	}

	instanceNameStr := "<no name tag>"
	if instanceName != nil {
		instanceNameStr = *instanceName
	}

	err = gamingvm.StopVM(instance)
	if err != nil {
		fmt.Printf("Failed to stop VM '%s' (instance ID '%s')\n", instanceNameStr, *instance.InstanceId)
		fmt.Println(err)
		os.Exit(1)
		return
	}

	// TODO: add a spinner to wait until it's fully stopped

	fmt.Printf("VM '%s' stopped successfully.\n", instanceNameStr)

	waitForUserToContinue()

	os.Exit(0)
}

func waitForUserToContinue() {
	fmt.Print("Press enter to continue...")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
