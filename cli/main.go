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
	logo = `
	     ___
	    /. '\_._._._.
   	   / _/_ _ _ _ _|\
   	  |/  /\____/__/\\
	     ///   |/|/  \_`
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(logo, "\n\nPangolin holds your back!\nVersion:\tv1alpha1")
		return
	}
	verb := os.Args[1]
	switch verb {
		case "get":
			GetBackupJobs()
		case "create":
			fmt.Println("Not implemented yet")
		case "delete": 
			fmt.Println("Not implemented yet")
		default:
			fmt.Printf("Wrong verb: %s", verb)
	}
	
}