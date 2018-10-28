/* An Interactive cli to start, stop and delete AKS Cluster
   At ADfolks we required this to reduce the AKS billing costs
   We use this to stop the Dev clusters when not in use
*/ Or delete and start again with from scratch in a single command

package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"log"
	"strings"
)

func getClusterName() string {
	readerClusterName := bufio.NewReader(os.Stdin)
	fmt.Println("Enter k8s cluster Name:")
	clusterName, _ := readerClusterName.ReadString('\n')
	clusterName = strings.TrimSuffix(clusterName,"\n")
	return clusterName
}

func getClusterRegion() string {
	readerClusterRegion := bufio.NewReader(os.Stdin)
	fmt.Println("Enter k8s cloud region:")
	clusterRegion, _ := readerClusterRegion.ReadString('\n')
	clusterRegion = strings.TrimSuffix(clusterRegion, "\n")
	retun clusterRegion
}
func createResourceGroup(resourceGroupName string, clusterRegion string) {
	cmd := exec.Command("az", "group", "create", "-l", clusterRegion, "-n",
		 resourceGroupName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
}

func createCluster( clusterName string) sting, string {
	fmt.Println("Start setiing up your k8s Cluster")
	fmt.Println("---------------------------------")
	//Create AKS Cluster
	cmd = exec.Command("az", "aks", "create", "--name", clusterName,
		"--resource-group", resourceGroupName, "--node-count",
		 "6", "--kubernetes-version", "1.11.3")
	out, err = cmd.CombinedOutput()
	if err != nil {
	log.Fatalf("cmd.Run() failed with %s\n", err)
	}
    fmt.Printf("Output:\n%s\n", string(out))
}

func getKubectlConfig(clusterName string) {
	//Get Kubectl credentials copied to ~/.kube/config
	cmd := exec.Command("az", "aks", "get-credentials",
		 "--resource-group", resourceGroupName, "--name",clusterName)
	out, err = cmd.CombinedOutput()
	if err != nil {
	log.Fatalf("cmd.Run() Failed with %s\n", err)
	}
	fmt.Printf("Get kubectl config Output: \n%s\n", string(out))
}

func intializeHelm() string {
	//Initialize Helm
	fmt.Println("Initialising Package manager Helm")
	cmd = exec.Command("helm", "init")
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() Failed with:\n%s\n", err)
	}
	fmt.Printf("Helm Init Output:\n%s\n", string(out))
}

func addHelmRepo(repoName string, repoURL string) string {
	fmt.Println("Adding Helm repo: %s", repoName)
	helmIncubatorRepo := "http://storage.googleapis.com/kubernetes-charts-incubator"
	cmd := exec.Command("helm", "repo", "add", "incubator", helmIncubatorRepo)
	out, err = cmd.CombinedOutput()
	if err != nil {
				log.Fatalf("Add Helm Repo failed with: \n%s\n", err)
		}
	fmt.Println("Helm Add repo output:\n%s\n", string(out))

}
func addPackage(packageName) string {
	//Get Kubectl credentials copied to ~/.kube/config
	cmd := exec.Command("az", "aks", "get-credentials",
		 "--resource-group", resourceGroupName, "--name",clusterName)
	out, err = cmd.CombinedOutput()
	if err != nil {
	log.Fatalf("cmd.Run() Failed with %s\n", err)
	}
	fmt.Printf("Get kubectl config Output: \n%s\n", string(out))
}
func createNamespace(nameSpace string) {
	cmd := exec.Command("kubectl", "create", "namespace", "elasticsearch")
	out, err = cmd.CombinedOutput()
  if err != nil {
        log.Fatalf("Kubectl create namespace failed with: \n%s\n", err)
    }
  fmt.Println("kubectl create namespace output:\n%s\n", string(out))

}
func main() {	 clusterRegion := getClusterRegion()
   clusterName := getClusterName()
	 //assign Azure Resource Group same name as cluster name
	 resourceGroupName := clusterName


	//Install ElasticSearch
  //Install Kafka
	createNamespace = exec.Command("kubectl", "create", "namespace", "kafka")
	out, err = createNamespace.CombinedOutput()
  if err != nil {
        log.Fatalf("Kubectl create namespace Kafka failed with: \n%s\n", err)
			}
  fmt.Println("kubectl create namespace output:\n%s\n", string(out))
  installPackage = exec.Command("helm", "install", "--name", "kafka", "--namespace", "kafka", "incubator/kafka")
  out, err = installPackage.CombinedOutput()
  if err != nil {
		log.Fatalf("cmd.run() Failed with:\n%s\n", err)
	}
  fmt.Println("Package install Output:\n%s\n", string(out))
}
