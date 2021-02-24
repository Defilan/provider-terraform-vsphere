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
func (in *DistributedPortGroup) DeepCopyInto(out *DistributedPortGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedPortGroup.
func (in *DistributedPortGroup) DeepCopy() *DistributedPortGroup {
	if in == nil {
		return nil
	}
	out := new(DistributedPortGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DistributedPortGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedPortGroupList) DeepCopyInto(out *DistributedPortGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DistributedPortGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedPortGroupList.
func (in *DistributedPortGroupList) DeepCopy() *DistributedPortGroupList {
	if in == nil {
		return nil
	}
	out := new(DistributedPortGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DistributedPortGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedPortGroupObservation) DeepCopyInto(out *DistributedPortGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedPortGroupObservation.
func (in *DistributedPortGroupObservation) DeepCopy() *DistributedPortGroupObservation {
	if in == nil {
		return nil
	}
	out := new(DistributedPortGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedPortGroupParameters) DeepCopyInto(out *DistributedPortGroupParameters) {
	*out = *in
	if in.ActiveUplinks != nil {
		in, out := &in.ActiveUplinks, &out.ActiveUplinks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StandbyUplinks != nil {
		in, out := &in.StandbyUplinks, &out.StandbyUplinks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.VlanRange = in.VlanRange
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedPortGroupParameters.
func (in *DistributedPortGroupParameters) DeepCopy() *DistributedPortGroupParameters {
	if in == nil {
		return nil
	}
	out := new(DistributedPortGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedPortGroupSpec) DeepCopyInto(out *DistributedPortGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedPortGroupSpec.
func (in *DistributedPortGroupSpec) DeepCopy() *DistributedPortGroupSpec {
	if in == nil {
		return nil
	}
	out := new(DistributedPortGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedPortGroupStatus) DeepCopyInto(out *DistributedPortGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedPortGroupStatus.
func (in *DistributedPortGroupStatus) DeepCopy() *DistributedPortGroupStatus {
	if in == nil {
		return nil
	}
	out := new(DistributedPortGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanRange) DeepCopyInto(out *VlanRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanRange.
func (in *VlanRange) DeepCopy() *VlanRange {
	if in == nil {
		return nil
	}
	out := new(VlanRange)
	in.DeepCopyInto(out)
	return out
}