package main

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	backupjobGVR = schema.GroupVersionResource{
		Group:    "pangolin.netzlink.com",
		Version:  "v1alpha1",
		Resource: "backupjobs",
	}
)

func main() {
	verb := os.Args[1]
	switch verb {
		case "get":
			GetBackupJobs()
		case "create":
			CreateBackupJob()
		case "delete": 
			fmt.Println("Not implemented yet")
		default:
			fmt.Printf("Wrong verb: %s", verb)
	}
	
}