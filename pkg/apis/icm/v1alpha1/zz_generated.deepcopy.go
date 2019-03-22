// +build !ignore_autogenerated

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/policy/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VCLStatus) DeepCopyInto(out *VCLStatus) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VCLStatus.
func (in *VCLStatus) DeepCopy() *VCLStatus {
	if in == nil {
		return nil
	}
	out := new(VCLStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarnishContainer) DeepCopyInto(out *VarnishContainer) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ImagePullSecret != nil {
		in, out := &in.ImagePullSecret, &out.ImagePullSecret
		*out = new(string)
		**out = **in
	}
	if in.VarnishArgs != nil {
		in, out := &in.VarnishArgs, &out.VarnishArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarnishContainer.
func (in *VarnishContainer) DeepCopy() *VarnishContainer {
	if in == nil {
		return nil
	}
	out := new(VarnishContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarnishDeployment) DeepCopyInto(out *VarnishDeployment) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Container.DeepCopyInto(&out.Container)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarnishDeployment.
func (in *VarnishDeployment) DeepCopy() *VarnishDeployment {
	if in == nil {
		return nil
	}
	out := new(VarnishDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarnishService) DeepCopyInto(out *VarnishService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarnishService.
func (in *VarnishService) DeepCopy() *VarnishService {
	if in == nil {
		return nil
	}
	out := new(VarnishService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VarnishService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarnishServiceList) DeepCopyInto(out *VarnishServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VarnishService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarnishServiceList.
func (in *VarnishServiceList) DeepCopy() *VarnishServiceList {
	if in == nil {
		return nil
	}
	out := new(VarnishServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VarnishServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarnishServiceService) DeepCopyInto(out *VarnishServiceService) {
	*out = *in
	in.ServiceSpec.DeepCopyInto(&out.ServiceSpec)
	out.VarnishPort = in.VarnishPort
	out.VarnishExporterPort = in.VarnishExporterPort
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarnishServiceService.
func (in *VarnishServiceService) DeepCopy() *VarnishServiceService {
	if in == nil {
		return nil
	}
	out := new(VarnishServiceService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarnishServiceServiceStatus) DeepCopyInto(out *VarnishServiceServiceStatus) {
	*out = *in
	in.ServiceStatus.DeepCopyInto(&out.ServiceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarnishServiceServiceStatus.
func (in *VarnishServiceServiceStatus) DeepCopy() *VarnishServiceServiceStatus {
	if in == nil {
		return nil
	}
	out := new(VarnishServiceServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarnishServiceSpec) DeepCopyInto(out *VarnishServiceSpec) {
	*out = *in
	out.VCLConfigMap = in.VCLConfigMap
	in.Deployment.DeepCopyInto(&out.Deployment)
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(v1beta1.PodDisruptionBudgetSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Service.DeepCopyInto(&out.Service)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarnishServiceSpec.
func (in *VarnishServiceSpec) DeepCopy() *VarnishServiceSpec {
	if in == nil {
		return nil
	}
	out := new(VarnishServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarnishServiceStatus) DeepCopyInto(out *VarnishServiceStatus) {
	*out = *in
	in.Deployment.DeepCopyInto(&out.Deployment)
	in.Service.DeepCopyInto(&out.Service)
	in.ServiceNoCache.DeepCopyInto(&out.ServiceNoCache)
	in.VCL.DeepCopyInto(&out.VCL)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarnishServiceStatus.
func (in *VarnishServiceStatus) DeepCopy() *VarnishServiceStatus {
	if in == nil {
		return nil
	}
	out := new(VarnishServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarnishVCLConfigMap) DeepCopyInto(out *VarnishVCLConfigMap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarnishVCLConfigMap.
func (in *VarnishVCLConfigMap) DeepCopy() *VarnishVCLConfigMap {
	if in == nil {
		return nil
	}
	out := new(VarnishVCLConfigMap)
	in.DeepCopyInto(out)
	return out
}
