package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	pangolinv1alpha1 "github.com/Netzlink/pangolin/pangolin-operator/pkg/apis/pangolin/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func CreateBackupJob() {
	unObj, err := ParseBackupToUnstructured()
	if err != nil {
		log.Fatalln(err)
	}

	dynClient := GetDynamicClientOrFail()
	cr, errCrd := dynClient.Resource(backupjobGVR).Create(unObj, metav1.CreateOptions{})

	if errCrd != nil {
		log.Println(errCrd)
	}
	fmt.Print(cr)
}

func ParseBackupToUnstructured() (*unstructured.Unstructured, error) {
	var backupjob pangolinv1alpha1.BackupJob
	var unObj *unstructured.Unstructured

	backupjobName := flag.String("name", "backup", "the job-name")
	backupjobSpecSchedule := flag.String("schedule", "0 */12 * * *", "the backup schedule")
	backupjobSpecExtras := flag.String("extras", "", "database-specific commands while exporting")
	backupjobSpecDatabaseConfigDatabase := flag.String("database", "default", "the database-name")
	backupjobSpecDatabaseConfigEndpoint := flag.String("endpoint", "localhost", "the database-endpoint")
	backupjobSpecDatabaseConfigUser := flag.String("user", "root", "the database-username")
	backupjobSpecDatabaseConfigPasswordSecret := flag.String("database-password", "my-secret", "the k8s-secret-name including the database password")
	backupjobSpecTypeMariadb := flag.Bool("mariadb", true, "mariadb database")
	backupjobSpecTypeMysql := flag.Bool("mysql", false, "mysql database")
	backupjobSpecTypeMssql := flag.Bool("sql-server", false, "sql-server database")
	backupjobSpecTypeMongodb := flag.Bool("mongodb", false, "mongodb database")
	backupjobSpecTypeCustomEnabled := flag.Bool("custom", false, "custom database")
	backupjobSpecTypeCustomImage := flag.String("custom-image", "", "database container image with client tools")
	backupjobSpecTypeCustomCommandTemplate := flag.String("custom-command-template", "", "A backup command with the ENVVARS dumping to /backup")
	backupjobSpecS3ConfigBucket := flag.String("bucket", "default", "S3 Bucket")
	backupjobSpecS3ConfigEndpoint := flag.String("s3-endpoint", "http://rook-ceph-rgw-ocp4-object-store.rook-ceph.svc.cluster.local", "S3 Endpoint")
	backupjobSpecS3ConfigSecret := flag.String("s3-secret", "my-secret", "S3 kubernetes secret")
	flag.Parse()
	backupjob.Name = *backupjobName
	backupjob.Spec.Schedule = *backupjobSpecSchedule
	backupjob.Spec.Extras = *backupjobSpecExtras
	backupjob.Spec.DatabaseConfig.Database = *backupjobSpecDatabaseConfigDatabase
	backupjob.Spec.DatabaseConfig.Endpoint = *backupjobSpecDatabaseConfigEndpoint
	backupjob.Spec.DatabaseConfig.User = *backupjobSpecDatabaseConfigUser
	backupjob.Spec.DatabaseConfig.PasswordSecret = *backupjobSpecDatabaseConfigPasswordSecret
	backupjob.Spec.Type.Mariadb = *backupjobSpecTypeMariadb
	backupjob.Spec.Type.Mysql = *backupjobSpecTypeMysql
	backupjob.Spec.Type.Mssql = *backupjobSpecTypeMssql
	backupjob.Spec.Type.Mongodb = *backupjobSpecTypeMongodb
	backupjob.Spec.Type.Custom.Enabled = *backupjobSpecTypeCustomEnabled
	backupjob.Spec.Type.Custom.Image = *backupjobSpecTypeCustomImage
	backupjob.Spec.Type.Custom.CommandTemplate = *backupjobSpecTypeCustomCommandTemplate
	backupjob.Spec.S3Config.Bucket = *backupjobSpecS3ConfigBucket
	backupjob.Spec.S3Config.Endpoint = *backupjobSpecS3ConfigEndpoint
	backupjob.Spec.S3Config.Secret = *backupjobSpecS3ConfigSecret

	strucJson, secErr := json.Marshal(backupjob)
	if secErr != nil {
		return nil, secErr
	}
	primErr := unObj.UnmarshalJSON(strucJson)
	if primErr != nil {
		return nil, primErr
	}
	return unObj, nil
}