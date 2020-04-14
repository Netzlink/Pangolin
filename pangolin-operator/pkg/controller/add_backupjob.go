package controller

import (
	"github.com/Netzlink/pangolin/pangolin-operator/pkg/controller/backupjob"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, backupjob.Add)
}
