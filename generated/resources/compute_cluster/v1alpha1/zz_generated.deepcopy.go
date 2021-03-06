// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
func (in *ComputeCluster) DeepCopyInto(out *ComputeCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeCluster.
func (in *ComputeCluster) DeepCopy() *ComputeCluster {
	if in == nil {
		return nil
	}
	out := new(ComputeCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeClusterList) DeepCopyInto(out *ComputeClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputeCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeClusterList.
func (in *ComputeClusterList) DeepCopy() *ComputeClusterList {
	if in == nil {
		return nil
	}
	out := new(ComputeClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeClusterObservation) DeepCopyInto(out *ComputeClusterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeClusterObservation.
func (in *ComputeClusterObservation) DeepCopy() *ComputeClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ComputeClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeClusterParameters) DeepCopyInto(out *ComputeClusterParameters) {
	*out = *in
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DrsAdvancedOptions != nil {
		in, out := &in.DrsAdvancedOptions, &out.DrsAdvancedOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HaAdmissionControlFailoverHostSystemIds != nil {
		in, out := &in.HaAdmissionControlFailoverHostSystemIds, &out.HaAdmissionControlFailoverHostSystemIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HaAdvancedOptions != nil {
		in, out := &in.HaAdvancedOptions, &out.HaAdvancedOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HaHeartbeatDatastoreIds != nil {
		in, out := &in.HaHeartbeatDatastoreIds, &out.HaHeartbeatDatastoreIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostSystemIds != nil {
		in, out := &in.HostSystemIds, &out.HostSystemIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProactiveHaProviderIds != nil {
		in, out := &in.ProactiveHaProviderIds, &out.ProactiveHaProviderIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.VsanDiskGroup.DeepCopyInto(&out.VsanDiskGroup)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeClusterParameters.
func (in *ComputeClusterParameters) DeepCopy() *ComputeClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ComputeClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeClusterSpec) DeepCopyInto(out *ComputeClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeClusterSpec.
func (in *ComputeClusterSpec) DeepCopy() *ComputeClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ComputeClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeClusterStatus) DeepCopyInto(out *ComputeClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeClusterStatus.
func (in *ComputeClusterStatus) DeepCopy() *ComputeClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ComputeClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsanDiskGroup) DeepCopyInto(out *VsanDiskGroup) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsanDiskGroup.
func (in *VsanDiskGroup) DeepCopy() *VsanDiskGroup {
	if in == nil {
		return nil
	}
	out := new(VsanDiskGroup)
	in.DeepCopyInto(out)
	return out
}
