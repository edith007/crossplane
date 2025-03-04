//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Combine) DeepCopyInto(out *Combine) {
	*out = *in
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make([]CombineVariable, len(*in))
		copy(*out, *in)
	}
	if in.String != nil {
		in, out := &in.String, &out.String
		*out = new(StringCombine)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Combine.
func (in *Combine) DeepCopy() *Combine {
	if in == nil {
		return nil
	}
	out := new(Combine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CombineVariable) DeepCopyInto(out *CombineVariable) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CombineVariable.
func (in *CombineVariable) DeepCopy() *CombineVariable {
	if in == nil {
		return nil
	}
	out := new(CombineVariable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComposedTemplate) DeepCopyInto(out *ComposedTemplate) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	in.Base.DeepCopyInto(&out.Base)
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]Patch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConnectionDetails != nil {
		in, out := &in.ConnectionDetails, &out.ConnectionDetails
		*out = make([]ConnectionDetail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReadinessChecks != nil {
		in, out := &in.ReadinessChecks, &out.ReadinessChecks
		*out = make([]ReadinessCheck, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComposedTemplate.
func (in *ComposedTemplate) DeepCopy() *ComposedTemplate {
	if in == nil {
		return nil
	}
	out := new(ComposedTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositionRevision) DeepCopyInto(out *CompositionRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositionRevision.
func (in *CompositionRevision) DeepCopy() *CompositionRevision {
	if in == nil {
		return nil
	}
	out := new(CompositionRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompositionRevision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositionRevisionList) DeepCopyInto(out *CompositionRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CompositionRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositionRevisionList.
func (in *CompositionRevisionList) DeepCopy() *CompositionRevisionList {
	if in == nil {
		return nil
	}
	out := new(CompositionRevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompositionRevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositionRevisionSpec) DeepCopyInto(out *CompositionRevisionSpec) {
	*out = *in
	out.CompositeTypeRef = in.CompositeTypeRef
	if in.PatchSets != nil {
		in, out := &in.PatchSets, &out.PatchSets
		*out = make([]PatchSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ComposedTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WriteConnectionSecretsToNamespace != nil {
		in, out := &in.WriteConnectionSecretsToNamespace, &out.WriteConnectionSecretsToNamespace
		*out = new(string)
		**out = **in
	}
	if in.PublishConnectionDetailsWithStoreConfigRef != nil {
		in, out := &in.PublishConnectionDetailsWithStoreConfigRef, &out.PublishConnectionDetailsWithStoreConfigRef
		*out = new(v1.Reference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositionRevisionSpec.
func (in *CompositionRevisionSpec) DeepCopy() *CompositionRevisionSpec {
	if in == nil {
		return nil
	}
	out := new(CompositionRevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositionRevisionStatus) DeepCopyInto(out *CompositionRevisionStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositionRevisionStatus.
func (in *CompositionRevisionStatus) DeepCopy() *CompositionRevisionStatus {
	if in == nil {
		return nil
	}
	out := new(CompositionRevisionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionDetail) DeepCopyInto(out *ConnectionDetail) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ConnectionDetailType)
		**out = **in
	}
	if in.FromConnectionSecretKey != nil {
		in, out := &in.FromConnectionSecretKey, &out.FromConnectionSecretKey
		*out = new(string)
		**out = **in
	}
	if in.FromFieldPath != nil {
		in, out := &in.FromFieldPath, &out.FromFieldPath
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionDetail.
func (in *ConnectionDetail) DeepCopy() *ConnectionDetail {
	if in == nil {
		return nil
	}
	out := new(ConnectionDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConvertTransform) DeepCopyInto(out *ConvertTransform) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConvertTransform.
func (in *ConvertTransform) DeepCopy() *ConvertTransform {
	if in == nil {
		return nil
	}
	out := new(ConvertTransform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MapTransform) DeepCopyInto(out *MapTransform) {
	*out = *in
	if in.Pairs != nil {
		in, out := &in.Pairs, &out.Pairs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MapTransform.
func (in *MapTransform) DeepCopy() *MapTransform {
	if in == nil {
		return nil
	}
	out := new(MapTransform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MathTransform) DeepCopyInto(out *MathTransform) {
	*out = *in
	if in.Multiply != nil {
		in, out := &in.Multiply, &out.Multiply
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MathTransform.
func (in *MathTransform) DeepCopy() *MathTransform {
	if in == nil {
		return nil
	}
	out := new(MathTransform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Patch) DeepCopyInto(out *Patch) {
	*out = *in
	if in.FromFieldPath != nil {
		in, out := &in.FromFieldPath, &out.FromFieldPath
		*out = new(string)
		**out = **in
	}
	if in.Combine != nil {
		in, out := &in.Combine, &out.Combine
		*out = new(Combine)
		(*in).DeepCopyInto(*out)
	}
	if in.ToFieldPath != nil {
		in, out := &in.ToFieldPath, &out.ToFieldPath
		*out = new(string)
		**out = **in
	}
	if in.PatchSetName != nil {
		in, out := &in.PatchSetName, &out.PatchSetName
		*out = new(string)
		**out = **in
	}
	if in.Transforms != nil {
		in, out := &in.Transforms, &out.Transforms
		*out = make([]Transform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(PatchPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Patch.
func (in *Patch) DeepCopy() *Patch {
	if in == nil {
		return nil
	}
	out := new(Patch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchPolicy) DeepCopyInto(out *PatchPolicy) {
	*out = *in
	if in.FromFieldPath != nil {
		in, out := &in.FromFieldPath, &out.FromFieldPath
		*out = new(FromFieldPathPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchPolicy.
func (in *PatchPolicy) DeepCopy() *PatchPolicy {
	if in == nil {
		return nil
	}
	out := new(PatchPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchSet) DeepCopyInto(out *PatchSet) {
	*out = *in
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]Patch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchSet.
func (in *PatchSet) DeepCopy() *PatchSet {
	if in == nil {
		return nil
	}
	out := new(PatchSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessCheck) DeepCopyInto(out *ReadinessCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessCheck.
func (in *ReadinessCheck) DeepCopy() *ReadinessCheck {
	if in == nil {
		return nil
	}
	out := new(ReadinessCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringCombine) DeepCopyInto(out *StringCombine) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringCombine.
func (in *StringCombine) DeepCopy() *StringCombine {
	if in == nil {
		return nil
	}
	out := new(StringCombine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringTransform) DeepCopyInto(out *StringTransform) {
	*out = *in
	if in.Convert != nil {
		in, out := &in.Convert, &out.Convert
		*out = new(StringConversionType)
		**out = **in
	}
	if in.Trim != nil {
		in, out := &in.Trim, &out.Trim
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringTransform.
func (in *StringTransform) DeepCopy() *StringTransform {
	if in == nil {
		return nil
	}
	out := new(StringTransform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transform) DeepCopyInto(out *Transform) {
	*out = *in
	if in.Math != nil {
		in, out := &in.Math, &out.Math
		*out = new(MathTransform)
		(*in).DeepCopyInto(*out)
	}
	if in.Map != nil {
		in, out := &in.Map, &out.Map
		*out = new(MapTransform)
		(*in).DeepCopyInto(*out)
	}
	if in.String != nil {
		in, out := &in.String, &out.String
		*out = new(StringTransform)
		(*in).DeepCopyInto(*out)
	}
	if in.Convert != nil {
		in, out := &in.Convert, &out.Convert
		*out = new(ConvertTransform)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transform.
func (in *Transform) DeepCopy() *Transform {
	if in == nil {
		return nil
	}
	out := new(Transform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TypeReference) DeepCopyInto(out *TypeReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeReference.
func (in *TypeReference) DeepCopy() *TypeReference {
	if in == nil {
		return nil
	}
	out := new(TypeReference)
	in.DeepCopyInto(out)
	return out
}
