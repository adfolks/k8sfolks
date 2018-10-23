// An Interactive cli to start, stop and delete AKS Cluster
// At ADfolks we required this to reduce the AKS billing costs
// We use this to stop the Dev clusters when not in use
// Or delete and start again with from scratch in a single command

package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"log"
	"strings"
)

func main() {
	fmt.Println("Start setiing up your k8s Cluster")
	fmt.Println("---------------------------------")
	readerName := bufio.NewReader(os.Stdin)
	fmt.Println("Enter k8s cluster Name:")
	clusterName, _ := readerName.ReadString('\n')
	clusterName = strings.TrimSuffix(clusterName,"\n")
	readerRegion := bufio.NewReader(os.Stdin)
	fmt.Println("Enter k8s cloud region:")
	clusterRegion, _ := readerRegion.ReadString('\n')
	clusterRegion = strings.TrimSuffix(clusterRegion, "\n")
	//Create Azure Resource Group
	resourceGroupName := clusterName
	cmd := exec.Command("az", "group", "create", "-l", clusterRegion, "-n", resourceGroupName)
	out, err := cmd.CombinedOutput()
	if err != nil {
	log.Fatalf("cmd.Run() failed with %s\n", err)
    }
	fmt.Printf("Output:\n%s\n", string(out))
	//Create AKS Cluster
	cmd = exec.Command("az", "aks", "create", "--name", clusterName, "--resource-group", resourceGroupName, "--node-count", "6", "--kubernetes-version", "1.11.3")
	out, err = cmd.CombinedOutput()
	if err != nil {
	log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
}
