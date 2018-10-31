/* An Interactive cli to start, stop and delete AKS Cluster
   At ADfolks we required this to reduce the AKS billing costs
   We use this to stop the Dev clusters when not in use
 Or delete and start again with from scratch in a single command
*/

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
	readerClusterName := bufio.NewReader(os.Stdin)
	fmt.Println("Enter k8s cluster Region:")
	clusterRegion, _ := readerClusterName.ReadString('\n')
	clusterRegion = strings.TrimSuffix(clusterRegion,"\n")
	return clusterRegion
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

func createCluster( clusterName string, resourceGroupName string) {
	fmt.Println("Starting to set up your k8s Cluster")
	fmt.Println("This would take a few minutes...")
	fmt.Println("---------------------------------")
	//Create AKS Cluster
	cmd := exec.Command("az", "aks", "create", "--name", clusterName,
		"--resource-group", resourceGroupName, "--node-count",
		 "6", "--kubernetes-version", "1.11.3")
	out, err := cmd.CombinedOutput()
	if err != nil {
	log.Fatalf("cmd.Run() failed with %s\n", err)
	}
    fmt.Printf("Output:\n%s\n", string(out))
}

func getKubectlConfig(clusterName string, resourceGroupName string) {
	//Get Kubectl credentials copied to ~/.kube/config
	cmd := exec.Command("az", "aks", "get-credentials",
		 "--resource-group", resourceGroupName, "--name",clusterName)
	out, err := cmd.CombinedOutput()
	if err != nil {
	log.Fatalf("cmd.Run() Failed with %s\n", err)
	}
	fmt.Printf("Get kubectl config Output: \n%s\n", string(out))
}

func intializeHelm() {
	//Initialize Helm
	fmt.Println("Initialising Package manager Helm")
	cmd := exec.Command("helm", "init")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() Failed with:\n%s\n", err)
	}
	fmt.Printf("Helm Init Output:\n%s\n", string(out))
}

func addHelmRepo(repoName string, repoURL string) {
	fmt.Println("Adding Helm repo:", repoName)
	cmd := exec.Command("helm", "repo", "add", repoName, repoURL)
	out, err := cmd.CombinedOutput()
	if err != nil {
				log.Fatalf("Add Helm Repo failed with: \n%s\n", err)
		}
	fmt.Println("Helm Add repo output:\n%s\n", string(out))

}

func createNamespace(nameSpace string) {
	cmd := exec.Command("kubectl", "create", "namespace", nameSpace)
	out, err := cmd.CombinedOutput()
  if err != nil {
        log.Fatalf("Kubectl create namespace %s failed with: \n%s\n",
					nameSpace, err)
    }
  fmt.Println("kubectl create namespace %s output:\n%s\n", nameSpace,
		string(out))

}

func installPackage(packageName string, repoName string) {
	//Create Namespace in same name as packageName
	createNamespace(packageName)
	repowithPackage := repoName + "/" + packageName
  cmd := exec.Command("helm", "install", "--name", packageName,
		 "--namespace", packageName, repowithPackage)
  out, err := cmd.CombinedOutput()
  if err != nil {
		log.Fatalf("cmd.run() Failed with:\n%s\n", err)
	}
  fmt.Println("Package install %s Output:\n%s\n", packageName, string(out))
}

func main() {
   clusterName := getClusterName()
	 clusterRegion := getClusterRegion()
   //assign Azure Resource Group same name as cluster name
   resourceGroupName := clusterName
   //Create Azure resource group
   createResourceGroup(clusterName, clusterRegion)
   //Create AKS cluster
   createCluster(clusterName, resourceGroupName)
   //Get kubectl config copied to ~/.kube/getKubectlConfig
   getKubectlConfig(clusterName, resourceGroupName)
   //intialize helm
   intializeHelm()
   //Add helm repo inucubator
   repoName := "incubator"
   repoURL := "http://storage.googleapis.com/kubernetes-charts-incubator"
   addHelmRepo(repoName, repoURL)
   //Install ElasticSearch
   installPackage("elasticsearch", repoName)
   //Install Kafka
   installPackage("kafka", repoName)
}
