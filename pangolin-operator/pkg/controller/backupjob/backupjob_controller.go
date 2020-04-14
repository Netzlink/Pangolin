package backupjob

import (
	"context"
	"fmt"

	pangolinv1alpha1 "github.com/Netzlink/pangolin/pangolin-operator/pkg/apis/pangolin/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_backupjob")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new BackupJob Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileBackupJob{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("backupjob-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource BackupJob
	err = c.Watch(&source.Kind{Type: &pangolinv1alpha1.BackupJob{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner BackupJob
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &pangolinv1alpha1.BackupJob{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileBackupJob implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileBackupJob{}

// ReconcileBackupJob reconciles a BackupJob object
type ReconcileBackupJob struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a BackupJob object and makes changes based on the state read
// and what is in the BackupJob.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileBackupJob) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling BackupJob")

	// Fetch the BackupJob instance
	instance := &pangolinv1alpha1.BackupJob{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Pod object
	job := newJobForCR(instance)

	// Set BackupJob instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, job, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Job already exists
	found := &batchv1beta1.CronJob{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: job.Name, Namespace: job.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new job", "Job.Namespace", job.Namespace, "Job.Name", job.Name)
		err = r.client.Create(context.TODO(), job)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: backup-job already exists", "Job.Namespace", found.Namespace, "Job.Name", found.Name)
	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newJobForCR(cr *pangolinv1alpha1.BackupJob) *batchv1beta1.CronJob {
	backupImageName, backupCommand := getBackupImageNameAndCommand(cr)
	labels := map[string]string{
		"app": cr.Name,
	}
	return &batchv1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pangolin-backupjob",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: batchv1beta1.CronJobSpec{
			Schedule: cr.Spec.Schedule,
			JobTemplate: batchv1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Volumes: []v1.Volume{
								v1.Volume{
									Name: cr.Name + "-pangolin-volume",
									VolumeSource: v1.VolumeSource{
										EmptyDir: &v1.EmptyDirVolumeSource{},
									},
								},
							},
							InitContainers: []v1.Container{
								v1.Container{
									Image:   backupImageName,
									Command: []string{backupCommand},
									VolumeMounts: []v1.VolumeMount{
										v1.VolumeMount{
											Name:      cr.Name + "-pangolin-volume",
											MountPath: "/backup",
										},
									},
									Env: []v1.EnvVar{
										v1.EnvVar{
											Name:  "PANGOLIN_HOST",
											Value: cr.Spec.DatabaseConfig.Endpoint,
										},
										v1.EnvVar{
											Name:  "PANGOLIN_DATABASE",
											Value: cr.Spec.DatabaseConfig.Database,
										},
										v1.EnvVar{
											Name:  "PANGOLIN_USER",
											Value: cr.Spec.DatabaseConfig.User,
										},
										v1.EnvVar{
											Name: "PANGOLIN_HOST",
											ValueFrom: &v1.EnvVarSource{
												SecretKeyRef: &v1.SecretKeySelector{
													LocalObjectReference: v1.LocalObjectReference{
														Name: cr.Spec.DatabaseConfig.PasswordSecret, 
													},
													Key: "password",
												},
											},
										},
									},
								},
							},
							Containers: []v1.Container{
								v1.Container{
									Image:   "minio/mc:RELEASE.2020-04-04T05-28-55Z",
									Command: []string{
										"mc config host add s3 $PANGULIN_ENDPOINT $PANGULIN_ACCESS_KEY $PANGULIN_SECRET_KEY ;",
										"mc cp -r /backup s3/$PANGULIN_BUCKET",
									},
									VolumeMounts: []v1.VolumeMount{
										v1.VolumeMount{
											Name:      cr.Name + "-pangolin-volume",
											MountPath: "/backup",
										},
									},
									Env: []v1.EnvVar{
										v1.EnvVar{
											Name:  "PANGOLIN_ENDPOINT",
											Value: cr.Spec.DatabaseConfig.Endpoint,
										},
										v1.EnvVar{
											Name:  "PANGOLIN_BUCKET",
											Value: cr.Spec.DatabaseConfig.Database,
										},
										v1.EnvVar{
											Name:  "PANGOLIN_ACCESS_KEY",
											ValueFrom: &v1.EnvVarSource{
												SecretKeyRef: &v1.SecretKeySelector{
													LocalObjectReference: v1.LocalObjectReference{
														Name: cr.Spec.S3Config.Secret, 
													},
													Key: "access-key",
												},
											},
										},
										v1.EnvVar{
											Name: "PANGOLIN_SECRET_KEY",
											ValueFrom: &v1.EnvVarSource{
												SecretKeyRef: &v1.SecretKeySelector{
													LocalObjectReference: v1.LocalObjectReference{
														Name: cr.Spec.S3Config.Secret, 
													},
													Key: "secret-key",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func getBackupImageNameAndCommand(cr *pangolinv1alpha1.BackupJob) (string, string) {
	var imageName string
	var command string

	if cr.Spec.Type.Mariadb || cr.Spec.Type.Mysql {
		imageName = "mariadb:10.5.2"
		command = "mysqldump --host $PANGOLIN_HOST--databases $PANGOLIN_DATABASE --user $PANGOLIN_USER --password $PANGOLIN_PASSWORD $PANGULIN_EXTRAS > /backup/dump.sql"
		return imageName, command
	}
	if cr.Spec.Type.Mongodb {
		imageName = "mongo:3.6.17"
		command = "mongodump --host=$PANGOLIN_HOST --db $PANGOLIN_DATABASE  --username=$PANGOLIN_USER --password=$PANGOLIN_PASSWORD $PANGULIN_EXTRAS --out /backup"
		return imageName, command
	}
	if cr.Spec.Type.Postgres {
		imageName = "postgres:12.2"
		command = "pg_dump $PANGULIN_EXTRAS -h $PANGOLIN_HOST -U $PANGULIN_USER -W $PANGULIN_PASSWORD $PANGULIN_DATABASE > /backup/dump.sql"
		return imageName, command
	}
	if cr.Spec.Type.Custom.Enabled {
		imageName = cr.Spec.Type.Custom.Image
		command = cr.Spec.Type.Custom.CommandTemplate
		return imageName, command
	}
	fmt.Println("No Image specified in Job: " + cr.Name)
	return "", ""
}
