// +build !ignore_autogenerated

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha2

import (
	unsafe "unsafe"

	v1alpha3 "github.com/talos-systems/cluster-api-provider-metal/api/v1alpha3"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*MetalCluster)(nil), (*v1alpha3.MetalCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalCluster_To_v1alpha3_MetalCluster(a.(*MetalCluster), b.(*v1alpha3.MetalCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.MetalCluster)(nil), (*MetalCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalCluster_To_v1alpha2_MetalCluster(a.(*v1alpha3.MetalCluster), b.(*MetalCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetalClusterList)(nil), (*v1alpha3.MetalClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalClusterList_To_v1alpha3_MetalClusterList(a.(*MetalClusterList), b.(*v1alpha3.MetalClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.MetalClusterList)(nil), (*MetalClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalClusterList_To_v1alpha2_MetalClusterList(a.(*v1alpha3.MetalClusterList), b.(*MetalClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetalMachine)(nil), (*v1alpha3.MetalMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalMachine_To_v1alpha3_MetalMachine(a.(*MetalMachine), b.(*v1alpha3.MetalMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.MetalMachine)(nil), (*MetalMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalMachine_To_v1alpha2_MetalMachine(a.(*v1alpha3.MetalMachine), b.(*MetalMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetalMachineList)(nil), (*v1alpha3.MetalMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalMachineList_To_v1alpha3_MetalMachineList(a.(*MetalMachineList), b.(*v1alpha3.MetalMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.MetalMachineList)(nil), (*MetalMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalMachineList_To_v1alpha2_MetalMachineList(a.(*v1alpha3.MetalMachineList), b.(*MetalMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetalMachineTemplate)(nil), (*v1alpha3.MetalMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalMachineTemplate_To_v1alpha3_MetalMachineTemplate(a.(*MetalMachineTemplate), b.(*v1alpha3.MetalMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.MetalMachineTemplate)(nil), (*MetalMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalMachineTemplate_To_v1alpha2_MetalMachineTemplate(a.(*v1alpha3.MetalMachineTemplate), b.(*MetalMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetalMachineTemplateList)(nil), (*v1alpha3.MetalMachineTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalMachineTemplateList_To_v1alpha3_MetalMachineTemplateList(a.(*MetalMachineTemplateList), b.(*v1alpha3.MetalMachineTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.MetalMachineTemplateList)(nil), (*MetalMachineTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalMachineTemplateList_To_v1alpha2_MetalMachineTemplateList(a.(*v1alpha3.MetalMachineTemplateList), b.(*MetalMachineTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetalMachineTemplateResource)(nil), (*v1alpha3.MetalMachineTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalMachineTemplateResource_To_v1alpha3_MetalMachineTemplateResource(a.(*MetalMachineTemplateResource), b.(*v1alpha3.MetalMachineTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.MetalMachineTemplateResource)(nil), (*MetalMachineTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalMachineTemplateResource_To_v1alpha2_MetalMachineTemplateResource(a.(*v1alpha3.MetalMachineTemplateResource), b.(*MetalMachineTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetalMachineTemplateSpec)(nil), (*v1alpha3.MetalMachineTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalMachineTemplateSpec_To_v1alpha3_MetalMachineTemplateSpec(a.(*MetalMachineTemplateSpec), b.(*v1alpha3.MetalMachineTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.MetalMachineTemplateSpec)(nil), (*MetalMachineTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalMachineTemplateSpec_To_v1alpha2_MetalMachineTemplateSpec(a.(*v1alpha3.MetalMachineTemplateSpec), b.(*MetalMachineTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetalMachineTemplateStatus)(nil), (*v1alpha3.MetalMachineTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalMachineTemplateStatus_To_v1alpha3_MetalMachineTemplateStatus(a.(*MetalMachineTemplateStatus), b.(*v1alpha3.MetalMachineTemplateStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.MetalMachineTemplateStatus)(nil), (*MetalMachineTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalMachineTemplateStatus_To_v1alpha2_MetalMachineTemplateStatus(a.(*v1alpha3.MetalMachineTemplateStatus), b.(*MetalMachineTemplateStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*MetalClusterSpec)(nil), (*v1alpha3.MetalClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalClusterSpec_To_v1alpha3_MetalClusterSpec(a.(*MetalClusterSpec), b.(*v1alpha3.MetalClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*MetalClusterStatus)(nil), (*v1alpha3.MetalClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalClusterStatus_To_v1alpha3_MetalClusterStatus(a.(*MetalClusterStatus), b.(*v1alpha3.MetalClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*MetalMachineSpec)(nil), (*v1alpha3.MetalMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalMachineSpec_To_v1alpha3_MetalMachineSpec(a.(*MetalMachineSpec), b.(*v1alpha3.MetalMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*MetalMachineStatus)(nil), (*v1alpha3.MetalMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetalMachineStatus_To_v1alpha3_MetalMachineStatus(a.(*MetalMachineStatus), b.(*v1alpha3.MetalMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha3.MetalClusterSpec)(nil), (*MetalClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalClusterSpec_To_v1alpha2_MetalClusterSpec(a.(*v1alpha3.MetalClusterSpec), b.(*MetalClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha3.MetalClusterStatus)(nil), (*MetalClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalClusterStatus_To_v1alpha2_MetalClusterStatus(a.(*v1alpha3.MetalClusterStatus), b.(*MetalClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha3.MetalMachineSpec)(nil), (*MetalMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalMachineSpec_To_v1alpha2_MetalMachineSpec(a.(*v1alpha3.MetalMachineSpec), b.(*MetalMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha3.MetalMachineStatus)(nil), (*MetalMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_MetalMachineStatus_To_v1alpha2_MetalMachineStatus(a.(*v1alpha3.MetalMachineStatus), b.(*MetalMachineStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_MetalCluster_To_v1alpha3_MetalCluster(in *MetalCluster, out *v1alpha3.MetalCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_MetalClusterSpec_To_v1alpha3_MetalClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_MetalClusterStatus_To_v1alpha3_MetalClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_MetalCluster_To_v1alpha3_MetalCluster is an autogenerated conversion function.
func Convert_v1alpha2_MetalCluster_To_v1alpha3_MetalCluster(in *MetalCluster, out *v1alpha3.MetalCluster, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetalCluster_To_v1alpha3_MetalCluster(in, out, s)
}

func autoConvert_v1alpha3_MetalCluster_To_v1alpha2_MetalCluster(in *v1alpha3.MetalCluster, out *MetalCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_MetalClusterSpec_To_v1alpha2_MetalClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_MetalClusterStatus_To_v1alpha2_MetalClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_MetalCluster_To_v1alpha2_MetalCluster is an autogenerated conversion function.
func Convert_v1alpha3_MetalCluster_To_v1alpha2_MetalCluster(in *v1alpha3.MetalCluster, out *MetalCluster, s conversion.Scope) error {
	return autoConvert_v1alpha3_MetalCluster_To_v1alpha2_MetalCluster(in, out, s)
}

func autoConvert_v1alpha2_MetalClusterList_To_v1alpha3_MetalClusterList(in *MetalClusterList, out *v1alpha3.MetalClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha3.MetalCluster, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_MetalCluster_To_v1alpha3_MetalCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_MetalClusterList_To_v1alpha3_MetalClusterList is an autogenerated conversion function.
func Convert_v1alpha2_MetalClusterList_To_v1alpha3_MetalClusterList(in *MetalClusterList, out *v1alpha3.MetalClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetalClusterList_To_v1alpha3_MetalClusterList(in, out, s)
}

func autoConvert_v1alpha3_MetalClusterList_To_v1alpha2_MetalClusterList(in *v1alpha3.MetalClusterList, out *MetalClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetalCluster, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_MetalCluster_To_v1alpha2_MetalCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_MetalClusterList_To_v1alpha2_MetalClusterList is an autogenerated conversion function.
func Convert_v1alpha3_MetalClusterList_To_v1alpha2_MetalClusterList(in *v1alpha3.MetalClusterList, out *MetalClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha3_MetalClusterList_To_v1alpha2_MetalClusterList(in, out, s)
}

func autoConvert_v1alpha2_MetalClusterSpec_To_v1alpha3_MetalClusterSpec(in *MetalClusterSpec, out *v1alpha3.MetalClusterSpec, s conversion.Scope) error {
	// WARNING: in.APIEndpoints requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha3_MetalClusterSpec_To_v1alpha2_MetalClusterSpec(in *v1alpha3.MetalClusterSpec, out *MetalClusterSpec, s conversion.Scope) error {
	// WARNING: in.ControlPlaneEndpoint requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_MetalClusterStatus_To_v1alpha3_MetalClusterStatus(in *MetalClusterStatus, out *v1alpha3.MetalClusterStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	// WARNING: in.APIEndpoints requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha3_MetalClusterStatus_To_v1alpha2_MetalClusterStatus(in *v1alpha3.MetalClusterStatus, out *MetalClusterStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	return nil
}

func autoConvert_v1alpha2_MetalMachine_To_v1alpha3_MetalMachine(in *MetalMachine, out *v1alpha3.MetalMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_MetalMachineSpec_To_v1alpha3_MetalMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_MetalMachineStatus_To_v1alpha3_MetalMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_MetalMachine_To_v1alpha3_MetalMachine is an autogenerated conversion function.
func Convert_v1alpha2_MetalMachine_To_v1alpha3_MetalMachine(in *MetalMachine, out *v1alpha3.MetalMachine, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetalMachine_To_v1alpha3_MetalMachine(in, out, s)
}

func autoConvert_v1alpha3_MetalMachine_To_v1alpha2_MetalMachine(in *v1alpha3.MetalMachine, out *MetalMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_MetalMachineSpec_To_v1alpha2_MetalMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_MetalMachineStatus_To_v1alpha2_MetalMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_MetalMachine_To_v1alpha2_MetalMachine is an autogenerated conversion function.
func Convert_v1alpha3_MetalMachine_To_v1alpha2_MetalMachine(in *v1alpha3.MetalMachine, out *MetalMachine, s conversion.Scope) error {
	return autoConvert_v1alpha3_MetalMachine_To_v1alpha2_MetalMachine(in, out, s)
}

func autoConvert_v1alpha2_MetalMachineList_To_v1alpha3_MetalMachineList(in *MetalMachineList, out *v1alpha3.MetalMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha3.MetalMachine, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_MetalMachine_To_v1alpha3_MetalMachine(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_MetalMachineList_To_v1alpha3_MetalMachineList is an autogenerated conversion function.
func Convert_v1alpha2_MetalMachineList_To_v1alpha3_MetalMachineList(in *MetalMachineList, out *v1alpha3.MetalMachineList, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetalMachineList_To_v1alpha3_MetalMachineList(in, out, s)
}

func autoConvert_v1alpha3_MetalMachineList_To_v1alpha2_MetalMachineList(in *v1alpha3.MetalMachineList, out *MetalMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetalMachine, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_MetalMachine_To_v1alpha2_MetalMachine(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_MetalMachineList_To_v1alpha2_MetalMachineList is an autogenerated conversion function.
func Convert_v1alpha3_MetalMachineList_To_v1alpha2_MetalMachineList(in *v1alpha3.MetalMachineList, out *MetalMachineList, s conversion.Scope) error {
	return autoConvert_v1alpha3_MetalMachineList_To_v1alpha2_MetalMachineList(in, out, s)
}

func autoConvert_v1alpha2_MetalMachineSpec_To_v1alpha3_MetalMachineSpec(in *MetalMachineSpec, out *v1alpha3.MetalMachineSpec, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.ServerRef = (*v1.ObjectReference)(unsafe.Pointer(in.ServerRef))
	return nil
}

func autoConvert_v1alpha3_MetalMachineSpec_To_v1alpha2_MetalMachineSpec(in *v1alpha3.MetalMachineSpec, out *MetalMachineSpec, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.ServerRef = (*v1.ObjectReference)(unsafe.Pointer(in.ServerRef))
	// WARNING: in.ServerClassRef requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_MetalMachineStatus_To_v1alpha3_MetalMachineStatus(in *MetalMachineStatus, out *v1alpha3.MetalMachineStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	// WARNING: in.ErrorReason requires manual conversion: does not exist in peer-type
	// WARNING: in.ErrorMessage requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha3_MetalMachineStatus_To_v1alpha2_MetalMachineStatus(in *v1alpha3.MetalMachineStatus, out *MetalMachineStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	// WARNING: in.FailureReason requires manual conversion: does not exist in peer-type
	// WARNING: in.FailureMessage requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_MetalMachineTemplate_To_v1alpha3_MetalMachineTemplate(in *MetalMachineTemplate, out *v1alpha3.MetalMachineTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_MetalMachineTemplateSpec_To_v1alpha3_MetalMachineTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_MetalMachineTemplateStatus_To_v1alpha3_MetalMachineTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_MetalMachineTemplate_To_v1alpha3_MetalMachineTemplate is an autogenerated conversion function.
func Convert_v1alpha2_MetalMachineTemplate_To_v1alpha3_MetalMachineTemplate(in *MetalMachineTemplate, out *v1alpha3.MetalMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetalMachineTemplate_To_v1alpha3_MetalMachineTemplate(in, out, s)
}

func autoConvert_v1alpha3_MetalMachineTemplate_To_v1alpha2_MetalMachineTemplate(in *v1alpha3.MetalMachineTemplate, out *MetalMachineTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_MetalMachineTemplateSpec_To_v1alpha2_MetalMachineTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_MetalMachineTemplateStatus_To_v1alpha2_MetalMachineTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_MetalMachineTemplate_To_v1alpha2_MetalMachineTemplate is an autogenerated conversion function.
func Convert_v1alpha3_MetalMachineTemplate_To_v1alpha2_MetalMachineTemplate(in *v1alpha3.MetalMachineTemplate, out *MetalMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha3_MetalMachineTemplate_To_v1alpha2_MetalMachineTemplate(in, out, s)
}

func autoConvert_v1alpha2_MetalMachineTemplateList_To_v1alpha3_MetalMachineTemplateList(in *MetalMachineTemplateList, out *v1alpha3.MetalMachineTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha3.MetalMachineTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_MetalMachineTemplate_To_v1alpha3_MetalMachineTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_MetalMachineTemplateList_To_v1alpha3_MetalMachineTemplateList is an autogenerated conversion function.
func Convert_v1alpha2_MetalMachineTemplateList_To_v1alpha3_MetalMachineTemplateList(in *MetalMachineTemplateList, out *v1alpha3.MetalMachineTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetalMachineTemplateList_To_v1alpha3_MetalMachineTemplateList(in, out, s)
}

func autoConvert_v1alpha3_MetalMachineTemplateList_To_v1alpha2_MetalMachineTemplateList(in *v1alpha3.MetalMachineTemplateList, out *MetalMachineTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetalMachineTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_MetalMachineTemplate_To_v1alpha2_MetalMachineTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_MetalMachineTemplateList_To_v1alpha2_MetalMachineTemplateList is an autogenerated conversion function.
func Convert_v1alpha3_MetalMachineTemplateList_To_v1alpha2_MetalMachineTemplateList(in *v1alpha3.MetalMachineTemplateList, out *MetalMachineTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha3_MetalMachineTemplateList_To_v1alpha2_MetalMachineTemplateList(in, out, s)
}

func autoConvert_v1alpha2_MetalMachineTemplateResource_To_v1alpha3_MetalMachineTemplateResource(in *MetalMachineTemplateResource, out *v1alpha3.MetalMachineTemplateResource, s conversion.Scope) error {
	if err := Convert_v1alpha2_MetalMachineSpec_To_v1alpha3_MetalMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_MetalMachineTemplateResource_To_v1alpha3_MetalMachineTemplateResource is an autogenerated conversion function.
func Convert_v1alpha2_MetalMachineTemplateResource_To_v1alpha3_MetalMachineTemplateResource(in *MetalMachineTemplateResource, out *v1alpha3.MetalMachineTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetalMachineTemplateResource_To_v1alpha3_MetalMachineTemplateResource(in, out, s)
}

func autoConvert_v1alpha3_MetalMachineTemplateResource_To_v1alpha2_MetalMachineTemplateResource(in *v1alpha3.MetalMachineTemplateResource, out *MetalMachineTemplateResource, s conversion.Scope) error {
	if err := Convert_v1alpha3_MetalMachineSpec_To_v1alpha2_MetalMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_MetalMachineTemplateResource_To_v1alpha2_MetalMachineTemplateResource is an autogenerated conversion function.
func Convert_v1alpha3_MetalMachineTemplateResource_To_v1alpha2_MetalMachineTemplateResource(in *v1alpha3.MetalMachineTemplateResource, out *MetalMachineTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha3_MetalMachineTemplateResource_To_v1alpha2_MetalMachineTemplateResource(in, out, s)
}

func autoConvert_v1alpha2_MetalMachineTemplateSpec_To_v1alpha3_MetalMachineTemplateSpec(in *MetalMachineTemplateSpec, out *v1alpha3.MetalMachineTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha2_MetalMachineTemplateResource_To_v1alpha3_MetalMachineTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_MetalMachineTemplateSpec_To_v1alpha3_MetalMachineTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha2_MetalMachineTemplateSpec_To_v1alpha3_MetalMachineTemplateSpec(in *MetalMachineTemplateSpec, out *v1alpha3.MetalMachineTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetalMachineTemplateSpec_To_v1alpha3_MetalMachineTemplateSpec(in, out, s)
}

func autoConvert_v1alpha3_MetalMachineTemplateSpec_To_v1alpha2_MetalMachineTemplateSpec(in *v1alpha3.MetalMachineTemplateSpec, out *MetalMachineTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha3_MetalMachineTemplateResource_To_v1alpha2_MetalMachineTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_MetalMachineTemplateSpec_To_v1alpha2_MetalMachineTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha3_MetalMachineTemplateSpec_To_v1alpha2_MetalMachineTemplateSpec(in *v1alpha3.MetalMachineTemplateSpec, out *MetalMachineTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_MetalMachineTemplateSpec_To_v1alpha2_MetalMachineTemplateSpec(in, out, s)
}

func autoConvert_v1alpha2_MetalMachineTemplateStatus_To_v1alpha3_MetalMachineTemplateStatus(in *MetalMachineTemplateStatus, out *v1alpha3.MetalMachineTemplateStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha2_MetalMachineTemplateStatus_To_v1alpha3_MetalMachineTemplateStatus is an autogenerated conversion function.
func Convert_v1alpha2_MetalMachineTemplateStatus_To_v1alpha3_MetalMachineTemplateStatus(in *MetalMachineTemplateStatus, out *v1alpha3.MetalMachineTemplateStatus, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetalMachineTemplateStatus_To_v1alpha3_MetalMachineTemplateStatus(in, out, s)
}

func autoConvert_v1alpha3_MetalMachineTemplateStatus_To_v1alpha2_MetalMachineTemplateStatus(in *v1alpha3.MetalMachineTemplateStatus, out *MetalMachineTemplateStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha3_MetalMachineTemplateStatus_To_v1alpha2_MetalMachineTemplateStatus is an autogenerated conversion function.
func Convert_v1alpha3_MetalMachineTemplateStatus_To_v1alpha2_MetalMachineTemplateStatus(in *v1alpha3.MetalMachineTemplateStatus, out *MetalMachineTemplateStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_MetalMachineTemplateStatus_To_v1alpha2_MetalMachineTemplateStatus(in, out, s)
}
