package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *BackupJob) DeepCopyInto(out *BackupJob) {
        out.TypeMeta = in.TypeMeta
        out.ObjectMeta = in.ObjectMeta
        out.Spec = in.Spec
        out.Status = in.Status
}

// DeepCopyObject returns a generically typed copy of an object
func (in *BackupJob) DeepCopyObject() runtime.Object {
        out := BackupJob{}
        in.DeepCopyInto(&out)

        return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *BackupJobList) DeepCopyObject() runtime.Object {
        out := BackupJobList{}
        out.TypeMeta = in.TypeMeta
        out.ListMeta = in.ListMeta

        if in.Items != nil {
                out.Items = make([]BackupJob, len(in.Items))
                for i := range in.Items {
                        in.Items[i].DeepCopyInto(&out.Items[i])
                }
        }

        return &out
}