package main

import (
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetBackupJobs() {
	dynClient := GetDynamicClientOrFail()
	cr, errCrd := dynClient.Resource(backupjobGVR).List(metav1.ListOptions{})
	//Get(nil, )

	if errCrd != nil {
		log.Println(errCrd)
	}

	PrintBackupJobs(cr)
}