package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BackupJobSpec struct {
	Schedule string `json:"schedule,omitempty"`
	Type     struct {
		Mssql	 bool `json:"mssql"`
		Mariadb  bool `json:"mariadb"`
		Mysql    bool `json:"mysql"`
		Mongodb  bool `json:"mongodb"`
		Postgres bool `json:"postgres"`
		Custom   struct {
			Enabled         bool   `json:"enabled"`
			Image           string `json:"image"`
			CommandTemplate string `json:"commandTemplate"`
		} `json:"custom"`
	} `json:"type,omitemtpy"`
	Extras         string `json:"extras"`
	DatabaseConfig struct {
		Endpoint       string `json:"endpoint,omitempty"`
		Database       string `json:"database"` // if nil all databases will be included
		User           string `json:"user,omitempty"`
		PasswordSecret string `json:"passwordSecret,omitempty"`
	} `json:"databaseConfig,omitemtpy"`
	S3Config struct {
		Endpoint string `json:"endpoint,omitempty"`
		Bucket   string `json:"bucket,omitempty"`
		Secret   string `json:"secret,omitempty"`
	} `json:"s3Config,omitempty"`
}

type BackupJobStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BackupJob is the Schema for the backupjobs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=backupjobs,scope=Namespaced
type BackupJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupJobSpec   `json:"spec,omitempty"`
	Status BackupJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BackupJobList contains a list of BackupJob
type BackupJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupJob{}, &BackupJobList{})
}
