//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 Percona, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Storages != nil {
		in, out := &in.Storages, &out.Storages
		*out = make(map[string]*BackupStorageSpec, len(*in))
		for key, val := range *in {
			var outVal *BackupStorageSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(BackupStorageSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageAzureSpec) DeepCopyInto(out *BackupStorageAzureSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageAzureSpec.
func (in *BackupStorageAzureSpec) DeepCopy() *BackupStorageAzureSpec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageAzureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageGCSSpec) DeepCopyInto(out *BackupStorageGCSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageGCSSpec.
func (in *BackupStorageGCSSpec) DeepCopy() *BackupStorageGCSSpec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageGCSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageS3Spec) DeepCopyInto(out *BackupStorageS3Spec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageS3Spec.
func (in *BackupStorageS3Spec) DeepCopy() *BackupStorageS3Spec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageS3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageSpec) DeepCopyInto(out *BackupStorageSpec) {
	*out = *in
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(VolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(BackupStorageS3Spec)
		**out = **in
	}
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		*out = new(BackupStorageGCSSpec)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(BackupStorageAzureSpec)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	if in.VerifyTLS != nil {
		in, out := &in.VerifyTLS, &out.VerifyTLS
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageSpec.
func (in *BackupStorageSpec) DeepCopy() *BackupStorageSpec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSpec) DeepCopyInto(out *ContainerSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.StartupProbe.DeepCopyInto(&out.StartupProbe)
	in.ReadinessProbe.DeepCopyInto(&out.ReadinessProbe)
	in.LivenessProbe.DeepCopyInto(&out.LivenessProbe)
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSpec.
func (in *ContainerSpec) DeepCopy() *ContainerSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HAProxySpec) DeepCopyInto(out *HAProxySpec) {
	*out = *in
	in.Expose.DeepCopyInto(&out.Expose)
	in.PodSpec.DeepCopyInto(&out.PodSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HAProxySpec.
func (in *HAProxySpec) DeepCopy() *HAProxySpec {
	if in == nil {
		return nil
	}
	out := new(HAProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLRouterSpec) DeepCopyInto(out *MySQLRouterSpec) {
	*out = *in
	in.Expose.DeepCopyInto(&out.Expose)
	in.PodSpec.DeepCopyInto(&out.PodSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLRouterSpec.
func (in *MySQLRouterSpec) DeepCopy() *MySQLRouterSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLRouterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLSpec) DeepCopyInto(out *MySQLSpec) {
	*out = *in
	out.SizeSemiSync = in.SizeSemiSync
	in.Expose.DeepCopyInto(&out.Expose)
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SidecarVolumes != nil {
		in, out := &in.SidecarVolumes, &out.SidecarVolumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SidecarPVCs != nil {
		in, out := &in.SidecarPVCs, &out.SidecarPVCs
		*out = make([]SidecarPVC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.PodSpec.DeepCopyInto(&out.PodSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLSpec.
func (in *MySQLSpec) DeepCopy() *MySQLSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestratorSpec) DeepCopyInto(out *OrchestratorSpec) {
	*out = *in
	in.Expose.DeepCopyInto(&out.Expose)
	in.PodSpec.DeepCopyInto(&out.PodSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestratorSpec.
func (in *OrchestratorSpec) DeepCopy() *OrchestratorSpec {
	if in == nil {
		return nil
	}
	out := new(OrchestratorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PMMSpec) DeepCopyInto(out *PMMSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PMMSpec.
func (in *PMMSpec) DeepCopy() *PMMSpec {
	if in == nil {
		return nil
	}
	out := new(PMMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQL) DeepCopyInto(out *PerconaServerMySQL) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQL.
func (in *PerconaServerMySQL) DeepCopy() *PerconaServerMySQL {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMySQL) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLBackup) DeepCopyInto(out *PerconaServerMySQLBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLBackup.
func (in *PerconaServerMySQLBackup) DeepCopy() *PerconaServerMySQLBackup {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMySQLBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLBackupList) DeepCopyInto(out *PerconaServerMySQLBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaServerMySQLBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLBackupList.
func (in *PerconaServerMySQLBackupList) DeepCopy() *PerconaServerMySQLBackupList {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMySQLBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLBackupSpec) DeepCopyInto(out *PerconaServerMySQLBackupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLBackupSpec.
func (in *PerconaServerMySQLBackupSpec) DeepCopy() *PerconaServerMySQLBackupSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLBackupStatus) DeepCopyInto(out *PerconaServerMySQLBackupStatus) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(BackupStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLBackupStatus.
func (in *PerconaServerMySQLBackupStatus) DeepCopy() *PerconaServerMySQLBackupStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLList) DeepCopyInto(out *PerconaServerMySQLList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaServerMySQL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLList.
func (in *PerconaServerMySQLList) DeepCopy() *PerconaServerMySQLList {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMySQLList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLRestore) DeepCopyInto(out *PerconaServerMySQLRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLRestore.
func (in *PerconaServerMySQLRestore) DeepCopy() *PerconaServerMySQLRestore {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMySQLRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLRestoreList) DeepCopyInto(out *PerconaServerMySQLRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaServerMySQLRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLRestoreList.
func (in *PerconaServerMySQLRestoreList) DeepCopy() *PerconaServerMySQLRestoreList {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMySQLRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLRestoreSpec) DeepCopyInto(out *PerconaServerMySQLRestoreSpec) {
	*out = *in
	if in.BackupSource != nil {
		in, out := &in.BackupSource, &out.BackupSource
		*out = new(PerconaServerMySQLBackupStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLRestoreSpec.
func (in *PerconaServerMySQLRestoreSpec) DeepCopy() *PerconaServerMySQLRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLRestoreStatus) DeepCopyInto(out *PerconaServerMySQLRestoreStatus) {
	*out = *in
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLRestoreStatus.
func (in *PerconaServerMySQLRestoreStatus) DeepCopy() *PerconaServerMySQLRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLSpec) DeepCopyInto(out *PerconaServerMySQLSpec) {
	*out = *in
	if in.IgnoreAnnotations != nil {
		in, out := &in.IgnoreAnnotations, &out.IgnoreAnnotations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IgnoreLabels != nil {
		in, out := &in.IgnoreLabels, &out.IgnoreLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.MySQL.DeepCopyInto(&out.MySQL)
	in.Orchestrator.DeepCopyInto(&out.Orchestrator)
	if in.PMM != nil {
		in, out := &in.PMM, &out.PMM
		*out = new(PMMSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(BackupSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Proxy.DeepCopyInto(&out.Proxy)
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Toolkit != nil {
		in, out := &in.Toolkit, &out.Toolkit
		*out = new(ToolkitSpec)
		(*in).DeepCopyInto(*out)
	}
	out.UpgradeOptions = in.UpgradeOptions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLSpec.
func (in *PerconaServerMySQLSpec) DeepCopy() *PerconaServerMySQLSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMySQLStatus) DeepCopyInto(out *PerconaServerMySQLStatus) {
	*out = *in
	out.MySQL = in.MySQL
	out.Orchestrator = in.Orchestrator
	out.HAProxy = in.HAProxy
	out.Router = in.Router
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMySQLStatus.
func (in *PerconaServerMySQLStatus) DeepCopy() *PerconaServerMySQLStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMySQLStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodAffinity) DeepCopyInto(out *PodAffinity) {
	*out = *in
	if in.TopologyKey != nil {
		in, out := &in.TopologyKey, &out.TopologyKey
		*out = new(string)
		**out = **in
	}
	if in.Advanced != nil {
		in, out := &in.Advanced, &out.Advanced
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodAffinity.
func (in *PodAffinity) DeepCopy() *PodAffinity {
	if in == nil {
		return nil
	}
	out := new(PodAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSpec) DeepCopyInto(out *PodSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VolumeSpec != nil {
		in, out := &in.VolumeSpec, &out.VolumeSpec
		*out = new(VolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(PodAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSpec.
func (in *PodSpec) DeepCopy() *PodSpec {
	if in == nil {
		return nil
	}
	out := new(PodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySpec) DeepCopyInto(out *ProxySpec) {
	*out = *in
	if in.Router != nil {
		in, out := &in.Router, &out.Router
		*out = new(MySQLRouterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.HAProxy != nil {
		in, out := &in.HAProxy, &out.HAProxy
		*out = new(HAProxySpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySpec.
func (in *ProxySpec) DeepCopy() *ProxySpec {
	if in == nil {
		return nil
	}
	out := new(ProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExpose) DeepCopyInto(out *ServiceExpose) {
	*out = *in
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.InternalTrafficPolicy != nil {
		in, out := &in.InternalTrafficPolicy, &out.InternalTrafficPolicy
		*out = new(corev1.ServiceInternalTrafficPolicyType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExpose.
func (in *ServiceExpose) DeepCopy() *ServiceExpose {
	if in == nil {
		return nil
	}
	out := new(ServiceExpose)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExposeTogglable) DeepCopyInto(out *ServiceExposeTogglable) {
	*out = *in
	in.ServiceExpose.DeepCopyInto(&out.ServiceExpose)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExposeTogglable.
func (in *ServiceExposeTogglable) DeepCopy() *ServiceExposeTogglable {
	if in == nil {
		return nil
	}
	out := new(ServiceExposeTogglable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarPVC) DeepCopyInto(out *SidecarPVC) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarPVC.
func (in *SidecarPVC) DeepCopy() *SidecarPVC {
	if in == nil {
		return nil
	}
	out := new(SidecarPVC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulAppStatus) DeepCopyInto(out *StatefulAppStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulAppStatus.
func (in *StatefulAppStatus) DeepCopy() *StatefulAppStatus {
	if in == nil {
		return nil
	}
	out := new(StatefulAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSpec) DeepCopyInto(out *TLSSpec) {
	*out = *in
	if in.SANs != nil {
		in, out := &in.SANs, &out.SANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IssuerConf != nil {
		in, out := &in.IssuerConf, &out.IssuerConf
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSpec.
func (in *TLSSpec) DeepCopy() *TLSSpec {
	if in == nil {
		return nil
	}
	out := new(TLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolkitSpec) DeepCopyInto(out *ToolkitSpec) {
	*out = *in
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolkitSpec.
func (in *ToolkitSpec) DeepCopy() *ToolkitSpec {
	if in == nil {
		return nil
	}
	out := new(ToolkitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeOptions) DeepCopyInto(out *UpgradeOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeOptions.
func (in *UpgradeOptions) DeepCopy() *UpgradeOptions {
	if in == nil {
		return nil
	}
	out := new(UpgradeOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(corev1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(corev1.HostPathVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(corev1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
