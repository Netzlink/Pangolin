package main

import(
	"fmt"
	pangolinv1alpha1 "github.com/Netzlink/pangolin/pangolin-operator/pkg/apis/pangolin/v1alpha1"
)

func GetBackupJobDatabaseType(backJob pangolinv1alpha1.BackupJob) (string, error) {
	if backJob.Spec.Type.Mssql {
		return "SQL-Server", nil
	}
	if backJob.Spec.Type.Mariadb {
		return "MariaDB", nil
	}
	if backJob.Spec.Type.Mysql {
		return "MySQL", nil
	}
	if backJob.Spec.Type.Mongodb {
		return "MongoDB", nil
	}
	if backJob.Spec.Type.Custom.Enabled {
		return backJob.Spec.Type.Custom.Image, nil
	}
	return "", fmt.Errorf("Couldnt find type: %v", backJob.Spec.Type)
}