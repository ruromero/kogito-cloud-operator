// +build !ignore_autogenerated

// Copyright 2020 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Artifact) DeepCopyInto(out *Artifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Artifact.
func (in *Artifact) DeepCopy() *Artifact {
	if in == nil {
		return nil
	}
	out := new(Artifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Builds) DeepCopyInto(out *Builds) {
	*out = *in
	if in.New != nil {
		in, out := &in.New, &out.New
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pending != nil {
		in, out := &in.Pending, &out.Pending
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Running != nil {
		in, out := &in.Running, &out.Running
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Complete != nil {
		in, out := &in.Complete, &out.Complete
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Cancelled != nil {
		in, out := &in.Cancelled, &out.Cancelled
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Builds.
func (in *Builds) DeepCopy() *Builds {
	if in == nil {
		return nil
	}
	out := new(Builds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsMeta) DeepCopyInto(out *ConditionsMeta) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsMeta.
func (in *ConditionsMeta) DeepCopy() *ConditionsMeta {
	if in == nil {
		return nil
	}
	out := new(ConditionsMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Deployments) DeepCopyInto(out *Deployments) {
	*out = *in
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Starting != nil {
		in, out := &in.Starting, &out.Starting
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Stopped != nil {
		in, out := &in.Stopped, &out.Stopped
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Deployments.
func (in *Deployments) DeepCopy() *Deployments {
	if in == nil {
		return nil
	}
	out := new(Deployments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSource) DeepCopyInto(out *GitSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSource.
func (in *GitSource) DeepCopy() *GitSource {
	if in == nil {
		return nil
	}
	out := new(GitSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanConnectionProperties) DeepCopyInto(out *InfinispanConnectionProperties) {
	*out = *in
	out.Credentials = in.Credentials
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanConnectionProperties.
func (in *InfinispanConnectionProperties) DeepCopy() *InfinispanConnectionProperties {
	if in == nil {
		return nil
	}
	out := new(InfinispanConnectionProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanInstallStatus) DeepCopyInto(out *InfinispanInstallStatus) {
	*out = *in
	in.InfraComponentInstallStatusType.DeepCopyInto(&out.InfraComponentInstallStatusType)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanInstallStatus.
func (in *InfinispanInstallStatus) DeepCopy() *InfinispanInstallStatus {
	if in == nil {
		return nil
	}
	out := new(InfinispanInstallStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanMeta) DeepCopyInto(out *InfinispanMeta) {
	*out = *in
	out.InfinispanProperties = in.InfinispanProperties
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanMeta.
func (in *InfinispanMeta) DeepCopy() *InfinispanMeta {
	if in == nil {
		return nil
	}
	out := new(InfinispanMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraComponentInstallStatusType) DeepCopyInto(out *InfraComponentInstallStatusType) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]InstallCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraComponentInstallStatusType.
func (in *InfraComponentInstallStatusType) DeepCopy() *InfraComponentInstallStatusType {
	if in == nil {
		return nil
	}
	out := new(InfraComponentInstallStatusType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallCondition) DeepCopyInto(out *InstallCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallCondition.
func (in *InstallCondition) DeepCopy() *InstallCondition {
	if in == nil {
		return nil
	}
	out := new(InstallCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectionProperties) DeepCopyInto(out *KafkaConnectionProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectionProperties.
func (in *KafkaConnectionProperties) DeepCopy() *KafkaConnectionProperties {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectionProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaMeta) DeepCopyInto(out *KafkaMeta) {
	*out = *in
	out.KafkaProperties = in.KafkaProperties
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaMeta.
func (in *KafkaMeta) DeepCopy() *KafkaMeta {
	if in == nil {
		return nil
	}
	out := new(KafkaMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoApp) DeepCopyInto(out *KogitoApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoApp.
func (in *KogitoApp) DeepCopy() *KogitoApp {
	if in == nil {
		return nil
	}
	out := new(KogitoApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoAppBuildObject) DeepCopyInto(out *KogitoAppBuildObject) {
	*out = *in
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.GitSource = in.GitSource
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]WebhookSecret, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	out.Artifact = in.Artifact
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoAppBuildObject.
func (in *KogitoAppBuildObject) DeepCopy() *KogitoAppBuildObject {
	if in == nil {
		return nil
	}
	out := new(KogitoAppBuildObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoAppList) DeepCopyInto(out *KogitoAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoAppList.
func (in *KogitoAppList) DeepCopy() *KogitoAppList {
	if in == nil {
		return nil
	}
	out := new(KogitoAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoAppServiceObject) DeepCopyInto(out *KogitoAppServiceObject) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoAppServiceObject.
func (in *KogitoAppServiceObject) DeepCopy() *KogitoAppServiceObject {
	if in == nil {
		return nil
	}
	out := new(KogitoAppServiceObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoAppSpec) DeepCopyInto(out *KogitoAppSpec) {
	*out = *in
	in.KogitoServiceSpec.DeepCopyInto(&out.KogitoServiceSpec)
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(KogitoAppBuildObject)
		(*in).DeepCopyInto(*out)
	}
	in.Service.DeepCopyInto(&out.Service)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoAppSpec.
func (in *KogitoAppSpec) DeepCopy() *KogitoAppSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoAppStatus) DeepCopyInto(out *KogitoAppStatus) {
	*out = *in
	in.ConditionsMeta.DeepCopyInto(&out.ConditionsMeta)
	in.Deployments.DeepCopyInto(&out.Deployments)
	in.Builds.DeepCopyInto(&out.Builds)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoAppStatus.
func (in *KogitoAppStatus) DeepCopy() *KogitoAppStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoBuild) DeepCopyInto(out *KogitoBuild) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoBuild.
func (in *KogitoBuild) DeepCopy() *KogitoBuild {
	if in == nil {
		return nil
	}
	out := new(KogitoBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoBuild) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoBuildConditions) DeepCopyInto(out *KogitoBuildConditions) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoBuildConditions.
func (in *KogitoBuildConditions) DeepCopy() *KogitoBuildConditions {
	if in == nil {
		return nil
	}
	out := new(KogitoBuildConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoBuildList) DeepCopyInto(out *KogitoBuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoBuild, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoBuildList.
func (in *KogitoBuildList) DeepCopy() *KogitoBuildList {
	if in == nil {
		return nil
	}
	out := new(KogitoBuildList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoBuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoBuildSpec) DeepCopyInto(out *KogitoBuildSpec) {
	*out = *in
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.GitSource = in.GitSource
	if in.WebHooks != nil {
		in, out := &in.WebHooks, &out.WebHooks
		*out = make([]WebhookSecret, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	out.BuildImage = in.BuildImage
	out.RuntimeImage = in.RuntimeImage
	out.Artifact = in.Artifact
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoBuildSpec.
func (in *KogitoBuildSpec) DeepCopy() *KogitoBuildSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoBuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoBuildStatus) DeepCopyInto(out *KogitoBuildStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KogitoBuildConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Builds.DeepCopyInto(&out.Builds)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoBuildStatus.
func (in *KogitoBuildStatus) DeepCopy() *KogitoBuildStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoBuildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoDataIndex) DeepCopyInto(out *KogitoDataIndex) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoDataIndex.
func (in *KogitoDataIndex) DeepCopy() *KogitoDataIndex {
	if in == nil {
		return nil
	}
	out := new(KogitoDataIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoDataIndex) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoDataIndexList) DeepCopyInto(out *KogitoDataIndexList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoDataIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoDataIndexList.
func (in *KogitoDataIndexList) DeepCopy() *KogitoDataIndexList {
	if in == nil {
		return nil
	}
	out := new(KogitoDataIndexList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoDataIndexList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoDataIndexSpec) DeepCopyInto(out *KogitoDataIndexSpec) {
	*out = *in
	out.InfinispanMeta = in.InfinispanMeta
	out.KafkaMeta = in.KafkaMeta
	in.KogitoServiceSpec.DeepCopyInto(&out.KogitoServiceSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoDataIndexSpec.
func (in *KogitoDataIndexSpec) DeepCopy() *KogitoDataIndexSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoDataIndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoDataIndexStatus) DeepCopyInto(out *KogitoDataIndexStatus) {
	*out = *in
	in.KogitoServiceStatus.DeepCopyInto(&out.KogitoServiceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoDataIndexStatus.
func (in *KogitoDataIndexStatus) DeepCopy() *KogitoDataIndexStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoDataIndexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoInfra) DeepCopyInto(out *KogitoInfra) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoInfra.
func (in *KogitoInfra) DeepCopy() *KogitoInfra {
	if in == nil {
		return nil
	}
	out := new(KogitoInfra)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoInfra) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoInfraCondition) DeepCopyInto(out *KogitoInfraCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoInfraCondition.
func (in *KogitoInfraCondition) DeepCopy() *KogitoInfraCondition {
	if in == nil {
		return nil
	}
	out := new(KogitoInfraCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoInfraList) DeepCopyInto(out *KogitoInfraList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoInfra, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoInfraList.
func (in *KogitoInfraList) DeepCopy() *KogitoInfraList {
	if in == nil {
		return nil
	}
	out := new(KogitoInfraList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoInfraList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoInfraSpec) DeepCopyInto(out *KogitoInfraSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoInfraSpec.
func (in *KogitoInfraSpec) DeepCopy() *KogitoInfraSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoInfraSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoInfraStatus) DeepCopyInto(out *KogitoInfraStatus) {
	*out = *in
	out.Condition = in.Condition
	in.Infinispan.DeepCopyInto(&out.Infinispan)
	in.Kafka.DeepCopyInto(&out.Kafka)
	in.Keycloak.DeepCopyInto(&out.Keycloak)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoInfraStatus.
func (in *KogitoInfraStatus) DeepCopy() *KogitoInfraStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoInfraStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoJobsService) DeepCopyInto(out *KogitoJobsService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoJobsService.
func (in *KogitoJobsService) DeepCopy() *KogitoJobsService {
	if in == nil {
		return nil
	}
	out := new(KogitoJobsService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoJobsService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoJobsServiceList) DeepCopyInto(out *KogitoJobsServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoJobsService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoJobsServiceList.
func (in *KogitoJobsServiceList) DeepCopy() *KogitoJobsServiceList {
	if in == nil {
		return nil
	}
	out := new(KogitoJobsServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoJobsServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoJobsServiceSpec) DeepCopyInto(out *KogitoJobsServiceSpec) {
	*out = *in
	out.InfinispanMeta = in.InfinispanMeta
	out.KafkaMeta = in.KafkaMeta
	in.KogitoServiceSpec.DeepCopyInto(&out.KogitoServiceSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoJobsServiceSpec.
func (in *KogitoJobsServiceSpec) DeepCopy() *KogitoJobsServiceSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoJobsServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoJobsServiceStatus) DeepCopyInto(out *KogitoJobsServiceStatus) {
	*out = *in
	in.KogitoServiceStatus.DeepCopyInto(&out.KogitoServiceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoJobsServiceStatus.
func (in *KogitoJobsServiceStatus) DeepCopy() *KogitoJobsServiceStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoJobsServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoMgmtConsole) DeepCopyInto(out *KogitoMgmtConsole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoMgmtConsole.
func (in *KogitoMgmtConsole) DeepCopy() *KogitoMgmtConsole {
	if in == nil {
		return nil
	}
	out := new(KogitoMgmtConsole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoMgmtConsole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoMgmtConsoleList) DeepCopyInto(out *KogitoMgmtConsoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoMgmtConsole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoMgmtConsoleList.
func (in *KogitoMgmtConsoleList) DeepCopy() *KogitoMgmtConsoleList {
	if in == nil {
		return nil
	}
	out := new(KogitoMgmtConsoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoMgmtConsoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoMgmtConsoleSpec) DeepCopyInto(out *KogitoMgmtConsoleSpec) {
	*out = *in
	in.KogitoServiceSpec.DeepCopyInto(&out.KogitoServiceSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoMgmtConsoleSpec.
func (in *KogitoMgmtConsoleSpec) DeepCopy() *KogitoMgmtConsoleSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoMgmtConsoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoMgmtConsoleStatus) DeepCopyInto(out *KogitoMgmtConsoleStatus) {
	*out = *in
	in.KogitoServiceStatus.DeepCopyInto(&out.KogitoServiceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoMgmtConsoleStatus.
func (in *KogitoMgmtConsoleStatus) DeepCopy() *KogitoMgmtConsoleStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoMgmtConsoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoRuntime) DeepCopyInto(out *KogitoRuntime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoRuntime.
func (in *KogitoRuntime) DeepCopy() *KogitoRuntime {
	if in == nil {
		return nil
	}
	out := new(KogitoRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoRuntime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoRuntimeList) DeepCopyInto(out *KogitoRuntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoRuntime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoRuntimeList.
func (in *KogitoRuntimeList) DeepCopy() *KogitoRuntimeList {
	if in == nil {
		return nil
	}
	out := new(KogitoRuntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoRuntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoRuntimeSpec) DeepCopyInto(out *KogitoRuntimeSpec) {
	*out = *in
	out.InfinispanMeta = in.InfinispanMeta
	out.KafkaMeta = in.KafkaMeta
	in.KogitoServiceSpec.DeepCopyInto(&out.KogitoServiceSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoRuntimeSpec.
func (in *KogitoRuntimeSpec) DeepCopy() *KogitoRuntimeSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoRuntimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoRuntimeStatus) DeepCopyInto(out *KogitoRuntimeStatus) {
	*out = *in
	in.KogitoServiceStatus.DeepCopyInto(&out.KogitoServiceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoRuntimeStatus.
func (in *KogitoRuntimeStatus) DeepCopy() *KogitoRuntimeStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoRuntimeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoServiceSpec) DeepCopyInto(out *KogitoServiceSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Image = in.Image
	in.Resources.DeepCopyInto(&out.Resources)
	if in.DeploymentLabels != nil {
		in, out := &in.DeploymentLabels, &out.DeploymentLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceLabels != nil {
		in, out := &in.ServiceLabels, &out.ServiceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoServiceSpec.
func (in *KogitoServiceSpec) DeepCopy() *KogitoServiceSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoServiceStatus) DeepCopyInto(out *KogitoServiceStatus) {
	*out = *in
	in.ConditionsMeta.DeepCopyInto(&out.ConditionsMeta)
	if in.DeploymentConditions != nil {
		in, out := &in.DeploymentConditions, &out.DeploymentConditions
		*out = make([]appsv1.DeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoServiceStatus.
func (in *KogitoServiceStatus) DeepCopy() *KogitoServiceStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretCredentialsType) DeepCopyInto(out *SecretCredentialsType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretCredentialsType.
func (in *SecretCredentialsType) DeepCopy() *SecretCredentialsType {
	if in == nil {
		return nil
	}
	out := new(SecretCredentialsType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSecret) DeepCopyInto(out *WebhookSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSecret.
func (in *WebhookSecret) DeepCopy() *WebhookSecret {
	if in == nil {
		return nil
	}
	out := new(WebhookSecret)
	in.DeepCopyInto(out)
	return out
}
