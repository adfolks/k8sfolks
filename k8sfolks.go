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
	//Get Kubectl credentials copied to ~/.kube/config
	getKubectlConfig := exec.Command("az", "aks", "get-credentials", ""--resource-group", resourceGroupName, "--name",clusterName)
	out, err = getKubectlConfig.CombinedOutput()
	if err != nil {
	log.Fatalf("getKubectlConfig.Run() Failed with %s\n", err)
	}
	fmt.Printf("Get kubectl config Output: \n%s\n", string(out))
    //Initialize Helm
	fmt.Println("Initialising Package manager Helm")
	cmd = exec.Command("helm", "init")
    out, err = cmd.CombinedOutput()
    if err != nil {
    log.Fatalf("cmd.Run() Failed with:\n%s\n", err)
	}
	fmt.Printf("Helm Init Output:\n%s\n", string(out))
	//Install ElasticSearch
	fmt.Println("Installing ElasticSearch with Helm")
	helmIncubatorRepo := "http://storage.googleapis.com/kubernetes-charts-incubator"
	addRepo := exec.Command("helm", "repo", "add", "incubator", helmIncubatorRepo)
    out, err = addRepo.CombineOutput()
    if err != nil {
        log.Fatalf("Add Helm Repo failed with: \n%s\n" err)
    }
    fmt.Println("Helm Add repo output:\n%s\n", string(out))
	installPackage := exec.Command("helm", "install", "--name", "elasticsearch", "incubator/elasticsearch")
    //Install Kafka
    installPackage = exec.Command("helm", "install", "--name", "kafka", "incubator/kafka")
    out, err = installPackage.CombinedOutput()
    if err != nil {
    log.Fatalf("cmd.run() Failed with:\n%s\n", err)
    }
    fmt.Println("Package install Output:\n%s\n", string(out))
}
