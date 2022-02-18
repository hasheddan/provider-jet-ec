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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoffeeObservation) DeepCopyInto(out *CoffeeObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Price != nil {
		in, out := &in.Price, &out.Price
		*out = new(int64)
		**out = **in
	}
	if in.Teaser != nil {
		in, out := &in.Teaser, &out.Teaser
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoffeeObservation.
func (in *CoffeeObservation) DeepCopy() *CoffeeObservation {
	if in == nil {
		return nil
	}
	out := new(CoffeeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoffeeParameters) DeepCopyInto(out *CoffeeParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoffeeParameters.
func (in *CoffeeParameters) DeepCopy() *CoffeeParameters {
	if in == nil {
		return nil
	}
	out := new(CoffeeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ItemsObservation) DeepCopyInto(out *ItemsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ItemsObservation.
func (in *ItemsObservation) DeepCopy() *ItemsObservation {
	if in == nil {
		return nil
	}
	out := new(ItemsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ItemsParameters) DeepCopyInto(out *ItemsParameters) {
	*out = *in
	if in.Coffee != nil {
		in, out := &in.Coffee, &out.Coffee
		*out = make([]CoffeeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Quantity != nil {
		in, out := &in.Quantity, &out.Quantity
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ItemsParameters.
func (in *ItemsParameters) DeepCopy() *ItemsParameters {
	if in == nil {
		return nil
	}
	out := new(ItemsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Order) DeepCopyInto(out *Order) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Order.
func (in *Order) DeepCopy() *Order {
	if in == nil {
		return nil
	}
	out := new(Order)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Order) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderList) DeepCopyInto(out *OrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Order, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderList.
func (in *OrderList) DeepCopy() *OrderList {
	if in == nil {
		return nil
	}
	out := new(OrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderObservation) DeepCopyInto(out *OrderObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderObservation.
func (in *OrderObservation) DeepCopy() *OrderObservation {
	if in == nil {
		return nil
	}
	out := new(OrderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderParameters) DeepCopyInto(out *OrderParameters) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ItemsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderParameters.
func (in *OrderParameters) DeepCopy() *OrderParameters {
	if in == nil {
		return nil
	}
	out := new(OrderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderSpec) DeepCopyInto(out *OrderSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderSpec.
func (in *OrderSpec) DeepCopy() *OrderSpec {
	if in == nil {
		return nil
	}
	out := new(OrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderStatus) DeepCopyInto(out *OrderStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderStatus.
func (in *OrderStatus) DeepCopy() *OrderStatus {
	if in == nil {
		return nil
	}
	out := new(OrderStatus)
	in.DeepCopyInto(out)
	return out
}
