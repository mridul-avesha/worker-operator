//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 Copyright (c) 2022 Avesha, Inc. All rights reserved.   SPDX-License-Identifier: Apache-2.0
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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppPod) DeepCopyInto(out *AppPod) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppPod.
func (in *AppPod) DeepCopy() *AppPod {
	if in == nil {
		return nil
	}
	out := new(AppPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQuota) DeepCopyInto(out *ClusterQuota) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.NamespaceQuota != nil {
		in, out := &in.NamespaceQuota, &out.NamespaceQuota
		*out = make([]NamespaceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQuota.
func (in *ClusterQuota) DeepCopy() *ClusterQuota {
	if in == nil {
		return nil
	}
	out := new(ClusterQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaStatus) DeepCopyInto(out *ClusterResourceQuotaStatus) {
	*out = *in
	in.ResourcesUsage.DeepCopyInto(&out.ResourcesUsage)
	if in.NamespaceResourceQuotaStatus != nil {
		in, out := &in.NamespaceResourceQuotaStatus, &out.NamespaceResourceQuotaStatus
		*out = make([]NamespaceResourceQuotaStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaStatus.
func (in *ClusterResourceQuotaStatus) DeepCopy() *ClusterResourceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalGatewayConfig) DeepCopyInto(out *ExternalGatewayConfig) {
	*out = *in
	out.Ingress = in.Ingress
	out.Egress = in.Egress
	out.NsIngress = in.NsIngress
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalGatewayConfig.
func (in *ExternalGatewayConfig) DeepCopy() *ExternalGatewayConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalGatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalGatewayConfigOptions) DeepCopyInto(out *ExternalGatewayConfigOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalGatewayConfigOptions.
func (in *ExternalGatewayConfigOptions) DeepCopy() *ExternalGatewayConfigOptions {
	if in == nil {
		return nil
	}
	out := new(ExternalGatewayConfigOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayCredentials) DeepCopyInto(out *GatewayCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayCredentials.
func (in *GatewayCredentials) DeepCopy() *GatewayCredentials {
	if in == nil {
		return nil
	}
	out := new(GatewayCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceConfig) DeepCopyInto(out *NamespaceConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceConfig.
func (in *NamespaceConfig) DeepCopy() *NamespaceConfig {
	if in == nil {
		return nil
	}
	out := new(NamespaceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceIsolationProfile) DeepCopyInto(out *NamespaceIsolationProfile) {
	*out = *in
	if in.ApplicationNamespaces != nil {
		in, out := &in.ApplicationNamespaces, &out.ApplicationNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedNamespaces != nil {
		in, out := &in.AllowedNamespaces, &out.AllowedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceIsolationProfile.
func (in *NamespaceIsolationProfile) DeepCopy() *NamespaceIsolationProfile {
	if in == nil {
		return nil
	}
	out := new(NamespaceIsolationProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceQuota) DeepCopyInto(out *NamespaceQuota) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceQuota.
func (in *NamespaceQuota) DeepCopy() *NamespaceQuota {
	if in == nil {
		return nil
	}
	out := new(NamespaceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceResourceQuotaStatus) DeepCopyInto(out *NamespaceResourceQuotaStatus) {
	*out = *in
	in.ResourceUsage.DeepCopyInto(&out.ResourceUsage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceResourceQuotaStatus.
func (in *NamespaceResourceQuotaStatus) DeepCopy() *NamespaceResourceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceResourceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QOSProfile) DeepCopyInto(out *QOSProfile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QOSProfile.
func (in *QOSProfile) DeepCopy() *QOSProfile {
	if in == nil {
		return nil
	}
	out := new(QOSProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	out.Memory = in.Memory.DeepCopy()
	out.Cpu = in.Cpu.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaProfile) DeepCopyInto(out *ResourceQuotaProfile) {
	*out = *in
	in.SliceQuota.DeepCopyInto(&out.SliceQuota)
	in.ClusterQuota.DeepCopyInto(&out.ClusterQuota)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaProfile.
func (in *ResourceQuotaProfile) DeepCopy() *ResourceQuotaProfile {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryEndpoint) DeepCopyInto(out *ServiceDiscoveryEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryEndpoint.
func (in *ServiceDiscoveryEndpoint) DeepCopy() *ServiceDiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryPort) DeepCopyInto(out *ServiceDiscoveryPort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryPort.
func (in *ServiceDiscoveryPort) DeepCopy() *ServiceDiscoveryPort {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceGatewayConfig) DeepCopyInto(out *SliceGatewayConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceGatewayConfig.
func (in *SliceGatewayConfig) DeepCopy() *SliceGatewayConfig {
	if in == nil {
		return nil
	}
	out := new(SliceGatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceQuota) DeepCopyInto(out *SliceQuota) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceQuota.
func (in *SliceQuota) DeepCopy() *SliceQuota {
	if in == nil {
		return nil
	}
	out := new(SliceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerServiceImport) DeepCopyInto(out *WorkerServiceImport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerServiceImport.
func (in *WorkerServiceImport) DeepCopy() *WorkerServiceImport {
	if in == nil {
		return nil
	}
	out := new(WorkerServiceImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerServiceImport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerServiceImportList) DeepCopyInto(out *WorkerServiceImportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerServiceImport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerServiceImportList.
func (in *WorkerServiceImportList) DeepCopy() *WorkerServiceImportList {
	if in == nil {
		return nil
	}
	out := new(WorkerServiceImportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerServiceImportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerServiceImportSpec) DeepCopyInto(out *WorkerServiceImportSpec) {
	*out = *in
	if in.SourceClusters != nil {
		in, out := &in.SourceClusters, &out.SourceClusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceDiscoveryEndpoints != nil {
		in, out := &in.ServiceDiscoveryEndpoints, &out.ServiceDiscoveryEndpoints
		*out = make([]ServiceDiscoveryEndpoint, len(*in))
		copy(*out, *in)
	}
	if in.ServiceDiscoveryPorts != nil {
		in, out := &in.ServiceDiscoveryPorts, &out.ServiceDiscoveryPorts
		*out = make([]ServiceDiscoveryPort, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerServiceImportSpec.
func (in *WorkerServiceImportSpec) DeepCopy() *WorkerServiceImportSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerServiceImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerServiceImportStatus) DeepCopyInto(out *WorkerServiceImportStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerServiceImportStatus.
func (in *WorkerServiceImportStatus) DeepCopy() *WorkerServiceImportStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerServiceImportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceConfig) DeepCopyInto(out *WorkerSliceConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceConfig.
func (in *WorkerSliceConfig) DeepCopy() *WorkerSliceConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerSliceConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceConfigList) DeepCopyInto(out *WorkerSliceConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerSliceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceConfigList.
func (in *WorkerSliceConfigList) DeepCopy() *WorkerSliceConfigList {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerSliceConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceConfigSpec) DeepCopyInto(out *WorkerSliceConfigSpec) {
	*out = *in
	out.SliceGatewayProvider = in.SliceGatewayProvider
	out.QosProfileDetails = in.QosProfileDetails
	in.NamespaceIsolationProfile.DeepCopyInto(&out.NamespaceIsolationProfile)
	out.ExternalGatewayConfig = in.ExternalGatewayConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceConfigSpec.
func (in *WorkerSliceConfigSpec) DeepCopy() *WorkerSliceConfigSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceConfigStatus) DeepCopyInto(out *WorkerSliceConfigStatus) {
	*out = *in
	if in.ConnectedAppPods != nil {
		in, out := &in.ConnectedAppPods, &out.ConnectedAppPods
		*out = make([]AppPod, len(*in))
		copy(*out, *in)
	}
	if in.OnboardedAppNamespaces != nil {
		in, out := &in.OnboardedAppNamespaces, &out.OnboardedAppNamespaces
		*out = make([]NamespaceConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceConfigStatus.
func (in *WorkerSliceConfigStatus) DeepCopy() *WorkerSliceConfigStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceGateway) DeepCopyInto(out *WorkerSliceGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceGateway.
func (in *WorkerSliceGateway) DeepCopy() *WorkerSliceGateway {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerSliceGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceGatewayList) DeepCopyInto(out *WorkerSliceGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerSliceGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceGatewayList.
func (in *WorkerSliceGatewayList) DeepCopy() *WorkerSliceGatewayList {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerSliceGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceGatewayProvider) DeepCopyInto(out *WorkerSliceGatewayProvider) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceGatewayProvider.
func (in *WorkerSliceGatewayProvider) DeepCopy() *WorkerSliceGatewayProvider {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceGatewayProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceGatewaySpec) DeepCopyInto(out *WorkerSliceGatewaySpec) {
	*out = *in
	out.GatewayCredentials = in.GatewayCredentials
	out.LocalGatewayConfig = in.LocalGatewayConfig
	out.RemoteGatewayConfig = in.RemoteGatewayConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceGatewaySpec.
func (in *WorkerSliceGatewaySpec) DeepCopy() *WorkerSliceGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceGatewayStatus) DeepCopyInto(out *WorkerSliceGatewayStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceGatewayStatus.
func (in *WorkerSliceGatewayStatus) DeepCopy() *WorkerSliceGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceResourceQuota) DeepCopyInto(out *WorkerSliceResourceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceResourceQuota.
func (in *WorkerSliceResourceQuota) DeepCopy() *WorkerSliceResourceQuota {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerSliceResourceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceResourceQuotaList) DeepCopyInto(out *WorkerSliceResourceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerSliceResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceResourceQuotaList.
func (in *WorkerSliceResourceQuotaList) DeepCopy() *WorkerSliceResourceQuotaList {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerSliceResourceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceResourceQuotaSpec) DeepCopyInto(out *WorkerSliceResourceQuotaSpec) {
	*out = *in
	in.ResourceQuotaProfile.DeepCopyInto(&out.ResourceQuotaProfile)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceResourceQuotaSpec.
func (in *WorkerSliceResourceQuotaSpec) DeepCopy() *WorkerSliceResourceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceResourceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceResourceQuotaStatus) DeepCopyInto(out *WorkerSliceResourceQuotaStatus) {
	*out = *in
	in.ClusterResourceQuotaStatus.DeepCopyInto(&out.ClusterResourceQuotaStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceResourceQuotaStatus.
func (in *WorkerSliceResourceQuotaStatus) DeepCopy() *WorkerSliceResourceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceResourceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}
