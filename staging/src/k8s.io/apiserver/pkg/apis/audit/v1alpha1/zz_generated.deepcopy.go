// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/client-go/pkg/api/v1"
	v1 "k8s.io/client-go/pkg/apis/authentication/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Event, InType: reflect.TypeOf(&Event{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_EventList, InType: reflect.TypeOf(&EventList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_GroupKinds, InType: reflect.TypeOf(&GroupKinds{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Policy, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PolicyRule, InType: reflect.TypeOf(&PolicyRule{})},
	)
}

func DeepCopy_v1alpha1_Event(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Event)
		out := out.(*Event)
		*out = *in
		out.Timestamp = in.Timestamp.DeepCopy()
		if newVal, err := c.DeepCopy(&in.User); err != nil {
			return err
		} else {
			out.User = *newVal.(*v1.UserInfo)
		}
		if in.Impersonate != nil {
			in, out := &in.Impersonate, &out.Impersonate
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.UserInfo)
			}
		}
		if in.ObjectRef != nil {
			in, out := &in.ObjectRef, &out.ObjectRef
			*out = new(api_v1.ObjectReference)
			**out = **in
		}
		if in.ResponseStatus != nil {
			in, out := &in.ResponseStatus, &out.ResponseStatus
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*meta_v1.Status)
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_EventList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*EventList)
		out := out.(*EventList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Event, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Event)
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_GroupKinds(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupKinds)
		out := out.(*GroupKinds)
		*out = *in
		if in.Kinds != nil {
			in, out := &in.Kinds, &out.Kinds
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1alpha1_Policy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Policy)
		out := out.(*Policy)
		*out = *in
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*PolicyRule)
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_PolicyRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyRule)
		out := out.(*PolicyRule)
		*out = *in
		if in.Users != nil {
			in, out := &in.Users, &out.Users
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.UserGroups != nil {
			in, out := &in.UserGroups, &out.UserGroups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Verbs != nil {
			in, out := &in.Verbs, &out.Verbs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.ResourceKinds != nil {
			in, out := &in.ResourceKinds, &out.ResourceKinds
			*out = make([]GroupKinds, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*GroupKinds)
				}
			}
		}
		if in.Namespaces != nil {
			in, out := &in.Namespaces, &out.Namespaces
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.NonResourceURLs != nil {
			in, out := &in.NonResourceURLs, &out.NonResourceURLs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}
