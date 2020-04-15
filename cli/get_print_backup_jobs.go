package main

import(
	"encoding/json"
	"fmt"
	"log"
	"strings"

	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	pangolinv1alpha1 "github.com/Netzlink/pangolin/pangolin-operator/pkg/apis/pangolin/v1alpha1"
)

func PrintBackupJobs(cr *unstructured.UnstructuredList) {
	fmt.Printf("Number\tName\t\tSchedule\tType\tDatabase\tBucket\n")
	for number, backupjob := range cr.Items {
		var conv []byte
		var bjobj pangolinv1alpha1.BackupJob

		conv, primErr := backupjob.MarshalJSON()
		if primErr != nil {
			log.Println(primErr)
		}
		secErr := json.Unmarshal(conv, &bjobj)
		if secErr != nil {
			log.Println(secErr)
		}

		backupJobType, typeErr := GetBackupJobDatabaseType(bjobj)
		if typeErr != nil {
			log.Println(typeErr)
		}

		fmt.Printf(
			"%d\t%s\t%s\t%s\t%s\t%s\n",
			number,
			backupjob.GetName(),
			bjobj.Spec.Schedule,
			backupJobType,
			strings.Split(bjobj.Spec.DatabaseConfig.Endpoint, ".")[0],
			bjobj.Spec.S3Config.Bucket,
		)
	}
}