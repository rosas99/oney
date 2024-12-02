//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	apps "github.com/rosas/onex/pkg/apis/apps"
	errors "github.com/rosas/onex/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Chain)(nil), (*apps.Chain)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Chain_To_apps_Chain(a.(*Chain), b.(*apps.Chain), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.Chain)(nil), (*Chain)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_Chain_To_v1beta1_Chain(a.(*apps.Chain), b.(*Chain), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChainList)(nil), (*apps.ChainList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ChainList_To_apps_ChainList(a.(*ChainList), b.(*apps.ChainList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.ChainList)(nil), (*ChainList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_ChainList_To_v1beta1_ChainList(a.(*apps.ChainList), b.(*ChainList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChainSpec)(nil), (*apps.ChainSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ChainSpec_To_apps_ChainSpec(a.(*ChainSpec), b.(*apps.ChainSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.ChainSpec)(nil), (*ChainSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_ChainSpec_To_v1beta1_ChainSpec(a.(*apps.ChainSpec), b.(*ChainSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChainStatus)(nil), (*apps.ChainStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ChainStatus_To_apps_ChainStatus(a.(*ChainStatus), b.(*apps.ChainStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.ChainStatus)(nil), (*ChainStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_ChainStatus_To_v1beta1_ChainStatus(a.(*apps.ChainStatus), b.(*ChainStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Condition)(nil), (*apps.Condition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Condition_To_apps_Condition(a.(*Condition), b.(*apps.Condition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.Condition)(nil), (*Condition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_Condition_To_v1beta1_Condition(a.(*apps.Condition), b.(*Condition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LocalObjectReference)(nil), (*apps.LocalObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_LocalObjectReference_To_apps_LocalObjectReference(a.(*LocalObjectReference), b.(*apps.LocalObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.LocalObjectReference)(nil), (*LocalObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_LocalObjectReference_To_v1beta1_LocalObjectReference(a.(*apps.LocalObjectReference), b.(*LocalObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Miner)(nil), (*apps.Miner)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Miner_To_apps_Miner(a.(*Miner), b.(*apps.Miner), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.Miner)(nil), (*Miner)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_Miner_To_v1beta1_Miner(a.(*apps.Miner), b.(*Miner), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerAddress)(nil), (*apps.MinerAddress)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerAddress_To_apps_MinerAddress(a.(*MinerAddress), b.(*apps.MinerAddress), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.MinerAddress)(nil), (*MinerAddress)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_MinerAddress_To_v1beta1_MinerAddress(a.(*apps.MinerAddress), b.(*MinerAddress), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerList)(nil), (*apps.MinerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerList_To_apps_MinerList(a.(*MinerList), b.(*apps.MinerList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.MinerList)(nil), (*MinerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_MinerList_To_v1beta1_MinerList(a.(*apps.MinerList), b.(*MinerList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerSet)(nil), (*apps.MinerSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerSet_To_apps_MinerSet(a.(*MinerSet), b.(*apps.MinerSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.MinerSet)(nil), (*MinerSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_MinerSet_To_v1beta1_MinerSet(a.(*apps.MinerSet), b.(*MinerSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerSetList)(nil), (*apps.MinerSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerSetList_To_apps_MinerSetList(a.(*MinerSetList), b.(*apps.MinerSetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.MinerSetList)(nil), (*MinerSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_MinerSetList_To_v1beta1_MinerSetList(a.(*apps.MinerSetList), b.(*MinerSetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerSetSpec)(nil), (*apps.MinerSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerSetSpec_To_apps_MinerSetSpec(a.(*MinerSetSpec), b.(*apps.MinerSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.MinerSetSpec)(nil), (*MinerSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_MinerSetSpec_To_v1beta1_MinerSetSpec(a.(*apps.MinerSetSpec), b.(*MinerSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerSetStatus)(nil), (*apps.MinerSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerSetStatus_To_apps_MinerSetStatus(a.(*MinerSetStatus), b.(*apps.MinerSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.MinerSetStatus)(nil), (*MinerSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_MinerSetStatus_To_v1beta1_MinerSetStatus(a.(*apps.MinerSetStatus), b.(*MinerSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerSpec)(nil), (*apps.MinerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerSpec_To_apps_MinerSpec(a.(*MinerSpec), b.(*apps.MinerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.MinerSpec)(nil), (*MinerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_MinerSpec_To_v1beta1_MinerSpec(a.(*apps.MinerSpec), b.(*MinerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerStatus)(nil), (*apps.MinerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerStatus_To_apps_MinerStatus(a.(*MinerStatus), b.(*apps.MinerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.MinerStatus)(nil), (*MinerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_MinerStatus_To_v1beta1_MinerStatus(a.(*apps.MinerStatus), b.(*MinerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerTemplateSpec)(nil), (*apps.MinerTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerTemplateSpec_To_apps_MinerTemplateSpec(a.(*MinerTemplateSpec), b.(*apps.MinerTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.MinerTemplateSpec)(nil), (*MinerTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_MinerTemplateSpec_To_v1beta1_MinerTemplateSpec(a.(*apps.MinerTemplateSpec), b.(*MinerTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ObjectMeta)(nil), (*apps.ObjectMeta)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ObjectMeta_To_apps_ObjectMeta(a.(*ObjectMeta), b.(*apps.ObjectMeta), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.ObjectMeta)(nil), (*ObjectMeta)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_ObjectMeta_To_v1beta1_ObjectMeta(a.(*apps.ObjectMeta), b.(*ObjectMeta), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodInfo)(nil), (*apps.PodInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PodInfo_To_apps_PodInfo(a.(*PodInfo), b.(*apps.PodInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apps.PodInfo)(nil), (*PodInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_PodInfo_To_v1beta1_PodInfo(a.(*apps.PodInfo), b.(*PodInfo), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_Chain_To_apps_Chain(in *Chain, out *apps.Chain, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ChainSpec_To_apps_ChainSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ChainStatus_To_apps_ChainStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Chain_To_apps_Chain is an autogenerated conversion function.
func Convert_v1beta1_Chain_To_apps_Chain(in *Chain, out *apps.Chain, s conversion.Scope) error {
	return autoConvert_v1beta1_Chain_To_apps_Chain(in, out, s)
}

func autoConvert_apps_Chain_To_v1beta1_Chain(in *apps.Chain, out *Chain, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apps_ChainSpec_To_v1beta1_ChainSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apps_ChainStatus_To_v1beta1_ChainStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_apps_Chain_To_v1beta1_Chain is an autogenerated conversion function.
func Convert_apps_Chain_To_v1beta1_Chain(in *apps.Chain, out *Chain, s conversion.Scope) error {
	return autoConvert_apps_Chain_To_v1beta1_Chain(in, out, s)
}

func autoConvert_v1beta1_ChainList_To_apps_ChainList(in *ChainList, out *apps.ChainList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]apps.Chain)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_ChainList_To_apps_ChainList is an autogenerated conversion function.
func Convert_v1beta1_ChainList_To_apps_ChainList(in *ChainList, out *apps.ChainList, s conversion.Scope) error {
	return autoConvert_v1beta1_ChainList_To_apps_ChainList(in, out, s)
}

func autoConvert_apps_ChainList_To_v1beta1_ChainList(in *apps.ChainList, out *ChainList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Chain)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_apps_ChainList_To_v1beta1_ChainList is an autogenerated conversion function.
func Convert_apps_ChainList_To_v1beta1_ChainList(in *apps.ChainList, out *ChainList, s conversion.Scope) error {
	return autoConvert_apps_ChainList_To_v1beta1_ChainList(in, out, s)
}

func autoConvert_v1beta1_ChainSpec_To_apps_ChainSpec(in *ChainSpec, out *apps.ChainSpec, s conversion.Scope) error {
	out.DisplayName = in.DisplayName
	out.MinerType = in.MinerType
	out.Image = in.Image
	out.MinMineIntervalSeconds = in.MinMineIntervalSeconds
	out.BootstrapAccount = (*string)(unsafe.Pointer(in.BootstrapAccount))
	return nil
}

// Convert_v1beta1_ChainSpec_To_apps_ChainSpec is an autogenerated conversion function.
func Convert_v1beta1_ChainSpec_To_apps_ChainSpec(in *ChainSpec, out *apps.ChainSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_ChainSpec_To_apps_ChainSpec(in, out, s)
}

func autoConvert_apps_ChainSpec_To_v1beta1_ChainSpec(in *apps.ChainSpec, out *ChainSpec, s conversion.Scope) error {
	out.DisplayName = in.DisplayName
	out.MinerType = in.MinerType
	out.Image = in.Image
	out.MinMineIntervalSeconds = in.MinMineIntervalSeconds
	out.BootstrapAccount = (*string)(unsafe.Pointer(in.BootstrapAccount))
	return nil
}

// Convert_apps_ChainSpec_To_v1beta1_ChainSpec is an autogenerated conversion function.
func Convert_apps_ChainSpec_To_v1beta1_ChainSpec(in *apps.ChainSpec, out *ChainSpec, s conversion.Scope) error {
	return autoConvert_apps_ChainSpec_To_v1beta1_ChainSpec(in, out, s)
}

func autoConvert_v1beta1_ChainStatus_To_apps_ChainStatus(in *ChainStatus, out *apps.ChainStatus, s conversion.Scope) error {
	out.ConfigMapRef = (*apps.LocalObjectReference)(unsafe.Pointer(in.ConfigMapRef))
	out.MinerRef = (*apps.LocalObjectReference)(unsafe.Pointer(in.MinerRef))
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*apps.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1beta1_ChainStatus_To_apps_ChainStatus is an autogenerated conversion function.
func Convert_v1beta1_ChainStatus_To_apps_ChainStatus(in *ChainStatus, out *apps.ChainStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_ChainStatus_To_apps_ChainStatus(in, out, s)
}

func autoConvert_apps_ChainStatus_To_v1beta1_ChainStatus(in *apps.ChainStatus, out *ChainStatus, s conversion.Scope) error {
	out.ConfigMapRef = (*LocalObjectReference)(unsafe.Pointer(in.ConfigMapRef))
	out.MinerRef = (*LocalObjectReference)(unsafe.Pointer(in.MinerRef))
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_apps_ChainStatus_To_v1beta1_ChainStatus is an autogenerated conversion function.
func Convert_apps_ChainStatus_To_v1beta1_ChainStatus(in *apps.ChainStatus, out *ChainStatus, s conversion.Scope) error {
	return autoConvert_apps_ChainStatus_To_v1beta1_ChainStatus(in, out, s)
}

func autoConvert_v1beta1_Condition_To_apps_Condition(in *Condition, out *apps.Condition, s conversion.Scope) error {
	out.Type = apps.ConditionType(in.Type)
	out.Status = core.ConditionStatus(in.Status)
	out.Severity = apps.ConditionSeverity(in.Severity)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1beta1_Condition_To_apps_Condition is an autogenerated conversion function.
func Convert_v1beta1_Condition_To_apps_Condition(in *Condition, out *apps.Condition, s conversion.Scope) error {
	return autoConvert_v1beta1_Condition_To_apps_Condition(in, out, s)
}

func autoConvert_apps_Condition_To_v1beta1_Condition(in *apps.Condition, out *Condition, s conversion.Scope) error {
	out.Type = ConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.Severity = ConditionSeverity(in.Severity)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_apps_Condition_To_v1beta1_Condition is an autogenerated conversion function.
func Convert_apps_Condition_To_v1beta1_Condition(in *apps.Condition, out *Condition, s conversion.Scope) error {
	return autoConvert_apps_Condition_To_v1beta1_Condition(in, out, s)
}

func autoConvert_v1beta1_LocalObjectReference_To_apps_LocalObjectReference(in *LocalObjectReference, out *apps.LocalObjectReference, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_v1beta1_LocalObjectReference_To_apps_LocalObjectReference is an autogenerated conversion function.
func Convert_v1beta1_LocalObjectReference_To_apps_LocalObjectReference(in *LocalObjectReference, out *apps.LocalObjectReference, s conversion.Scope) error {
	return autoConvert_v1beta1_LocalObjectReference_To_apps_LocalObjectReference(in, out, s)
}

func autoConvert_apps_LocalObjectReference_To_v1beta1_LocalObjectReference(in *apps.LocalObjectReference, out *LocalObjectReference, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_apps_LocalObjectReference_To_v1beta1_LocalObjectReference is an autogenerated conversion function.
func Convert_apps_LocalObjectReference_To_v1beta1_LocalObjectReference(in *apps.LocalObjectReference, out *LocalObjectReference, s conversion.Scope) error {
	return autoConvert_apps_LocalObjectReference_To_v1beta1_LocalObjectReference(in, out, s)
}

func autoConvert_v1beta1_Miner_To_apps_Miner(in *Miner, out *apps.Miner, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_MinerSpec_To_apps_MinerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_MinerStatus_To_apps_MinerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Miner_To_apps_Miner is an autogenerated conversion function.
func Convert_v1beta1_Miner_To_apps_Miner(in *Miner, out *apps.Miner, s conversion.Scope) error {
	return autoConvert_v1beta1_Miner_To_apps_Miner(in, out, s)
}

func autoConvert_apps_Miner_To_v1beta1_Miner(in *apps.Miner, out *Miner, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apps_MinerSpec_To_v1beta1_MinerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apps_MinerStatus_To_v1beta1_MinerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_apps_Miner_To_v1beta1_Miner is an autogenerated conversion function.
func Convert_apps_Miner_To_v1beta1_Miner(in *apps.Miner, out *Miner, s conversion.Scope) error {
	return autoConvert_apps_Miner_To_v1beta1_Miner(in, out, s)
}

func autoConvert_v1beta1_MinerAddress_To_apps_MinerAddress(in *MinerAddress, out *apps.MinerAddress, s conversion.Scope) error {
	out.Type = apps.MinerAddressType(in.Type)
	out.Address = in.Address
	return nil
}

// Convert_v1beta1_MinerAddress_To_apps_MinerAddress is an autogenerated conversion function.
func Convert_v1beta1_MinerAddress_To_apps_MinerAddress(in *MinerAddress, out *apps.MinerAddress, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerAddress_To_apps_MinerAddress(in, out, s)
}

func autoConvert_apps_MinerAddress_To_v1beta1_MinerAddress(in *apps.MinerAddress, out *MinerAddress, s conversion.Scope) error {
	out.Type = MinerAddressType(in.Type)
	out.Address = in.Address
	return nil
}

// Convert_apps_MinerAddress_To_v1beta1_MinerAddress is an autogenerated conversion function.
func Convert_apps_MinerAddress_To_v1beta1_MinerAddress(in *apps.MinerAddress, out *MinerAddress, s conversion.Scope) error {
	return autoConvert_apps_MinerAddress_To_v1beta1_MinerAddress(in, out, s)
}

func autoConvert_v1beta1_MinerList_To_apps_MinerList(in *MinerList, out *apps.MinerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]apps.Miner)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_MinerList_To_apps_MinerList is an autogenerated conversion function.
func Convert_v1beta1_MinerList_To_apps_MinerList(in *MinerList, out *apps.MinerList, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerList_To_apps_MinerList(in, out, s)
}

func autoConvert_apps_MinerList_To_v1beta1_MinerList(in *apps.MinerList, out *MinerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Miner)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_apps_MinerList_To_v1beta1_MinerList is an autogenerated conversion function.
func Convert_apps_MinerList_To_v1beta1_MinerList(in *apps.MinerList, out *MinerList, s conversion.Scope) error {
	return autoConvert_apps_MinerList_To_v1beta1_MinerList(in, out, s)
}

func autoConvert_v1beta1_MinerSet_To_apps_MinerSet(in *MinerSet, out *apps.MinerSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_MinerSetSpec_To_apps_MinerSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_MinerSetStatus_To_apps_MinerSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_MinerSet_To_apps_MinerSet is an autogenerated conversion function.
func Convert_v1beta1_MinerSet_To_apps_MinerSet(in *MinerSet, out *apps.MinerSet, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerSet_To_apps_MinerSet(in, out, s)
}

func autoConvert_apps_MinerSet_To_v1beta1_MinerSet(in *apps.MinerSet, out *MinerSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apps_MinerSetSpec_To_v1beta1_MinerSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apps_MinerSetStatus_To_v1beta1_MinerSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_apps_MinerSet_To_v1beta1_MinerSet is an autogenerated conversion function.
func Convert_apps_MinerSet_To_v1beta1_MinerSet(in *apps.MinerSet, out *MinerSet, s conversion.Scope) error {
	return autoConvert_apps_MinerSet_To_v1beta1_MinerSet(in, out, s)
}

func autoConvert_v1beta1_MinerSetList_To_apps_MinerSetList(in *MinerSetList, out *apps.MinerSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]apps.MinerSet)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_MinerSetList_To_apps_MinerSetList is an autogenerated conversion function.
func Convert_v1beta1_MinerSetList_To_apps_MinerSetList(in *MinerSetList, out *apps.MinerSetList, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerSetList_To_apps_MinerSetList(in, out, s)
}

func autoConvert_apps_MinerSetList_To_v1beta1_MinerSetList(in *apps.MinerSetList, out *MinerSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]MinerSet)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_apps_MinerSetList_To_v1beta1_MinerSetList is an autogenerated conversion function.
func Convert_apps_MinerSetList_To_v1beta1_MinerSetList(in *apps.MinerSetList, out *MinerSetList, s conversion.Scope) error {
	return autoConvert_apps_MinerSetList_To_v1beta1_MinerSetList(in, out, s)
}

func autoConvert_v1beta1_MinerSetSpec_To_apps_MinerSetSpec(in *MinerSetSpec, out *apps.MinerSetSpec, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Selector = in.Selector
	if err := Convert_v1beta1_MinerTemplateSpec_To_apps_MinerTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.DisplayName = in.DisplayName
	out.DeletePolicy = in.DeletePolicy
	out.MinReadySeconds = in.MinReadySeconds
	out.ProgressDeadlineSeconds = (*int32)(unsafe.Pointer(in.ProgressDeadlineSeconds))
	return nil
}

// Convert_v1beta1_MinerSetSpec_To_apps_MinerSetSpec is an autogenerated conversion function.
func Convert_v1beta1_MinerSetSpec_To_apps_MinerSetSpec(in *MinerSetSpec, out *apps.MinerSetSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerSetSpec_To_apps_MinerSetSpec(in, out, s)
}

func autoConvert_apps_MinerSetSpec_To_v1beta1_MinerSetSpec(in *apps.MinerSetSpec, out *MinerSetSpec, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Selector = in.Selector
	if err := Convert_apps_MinerTemplateSpec_To_v1beta1_MinerTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.DisplayName = in.DisplayName
	out.DeletePolicy = in.DeletePolicy
	out.MinReadySeconds = in.MinReadySeconds
	out.ProgressDeadlineSeconds = (*int32)(unsafe.Pointer(in.ProgressDeadlineSeconds))
	return nil
}

// Convert_apps_MinerSetSpec_To_v1beta1_MinerSetSpec is an autogenerated conversion function.
func Convert_apps_MinerSetSpec_To_v1beta1_MinerSetSpec(in *apps.MinerSetSpec, out *MinerSetSpec, s conversion.Scope) error {
	return autoConvert_apps_MinerSetSpec_To_v1beta1_MinerSetSpec(in, out, s)
}

func autoConvert_v1beta1_MinerSetStatus_To_apps_MinerSetStatus(in *MinerSetStatus, out *apps.MinerSetStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	out.FullyLabeledReplicas = in.FullyLabeledReplicas
	out.ReadyReplicas = in.ReadyReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.ObservedGeneration = in.ObservedGeneration
	out.FailureReason = (*errors.MinerSetStatusError)(unsafe.Pointer(in.FailureReason))
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.Conditions = *(*apps.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1beta1_MinerSetStatus_To_apps_MinerSetStatus is an autogenerated conversion function.
func Convert_v1beta1_MinerSetStatus_To_apps_MinerSetStatus(in *MinerSetStatus, out *apps.MinerSetStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerSetStatus_To_apps_MinerSetStatus(in, out, s)
}

func autoConvert_apps_MinerSetStatus_To_v1beta1_MinerSetStatus(in *apps.MinerSetStatus, out *MinerSetStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	out.FullyLabeledReplicas = in.FullyLabeledReplicas
	out.ReadyReplicas = in.ReadyReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.ObservedGeneration = in.ObservedGeneration
	out.FailureReason = (*errors.MinerSetStatusError)(unsafe.Pointer(in.FailureReason))
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.Conditions = *(*Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_apps_MinerSetStatus_To_v1beta1_MinerSetStatus is an autogenerated conversion function.
func Convert_apps_MinerSetStatus_To_v1beta1_MinerSetStatus(in *apps.MinerSetStatus, out *MinerSetStatus, s conversion.Scope) error {
	return autoConvert_apps_MinerSetStatus_To_v1beta1_MinerSetStatus(in, out, s)
}

func autoConvert_v1beta1_MinerSpec_To_apps_MinerSpec(in *MinerSpec, out *apps.MinerSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_ObjectMeta_To_apps_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.DisplayName = in.DisplayName
	out.MinerType = in.MinerType
	out.ChainName = in.ChainName
	out.RestartPolicy = core.RestartPolicy(in.RestartPolicy)
	out.PodDeletionTimeout = (*metav1.Duration)(unsafe.Pointer(in.PodDeletionTimeout))
	return nil
}

// Convert_v1beta1_MinerSpec_To_apps_MinerSpec is an autogenerated conversion function.
func Convert_v1beta1_MinerSpec_To_apps_MinerSpec(in *MinerSpec, out *apps.MinerSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerSpec_To_apps_MinerSpec(in, out, s)
}

func autoConvert_apps_MinerSpec_To_v1beta1_MinerSpec(in *apps.MinerSpec, out *MinerSpec, s conversion.Scope) error {
	if err := Convert_apps_ObjectMeta_To_v1beta1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.DisplayName = in.DisplayName
	out.MinerType = in.MinerType
	out.ChainName = in.ChainName
	out.RestartPolicy = v1.RestartPolicy(in.RestartPolicy)
	out.PodDeletionTimeout = (*metav1.Duration)(unsafe.Pointer(in.PodDeletionTimeout))
	return nil
}

// Convert_apps_MinerSpec_To_v1beta1_MinerSpec is an autogenerated conversion function.
func Convert_apps_MinerSpec_To_v1beta1_MinerSpec(in *apps.MinerSpec, out *MinerSpec, s conversion.Scope) error {
	return autoConvert_apps_MinerSpec_To_v1beta1_MinerSpec(in, out, s)
}

func autoConvert_v1beta1_MinerStatus_To_apps_MinerStatus(in *MinerStatus, out *apps.MinerStatus, s conversion.Scope) error {
	out.PodRef = (*core.ObjectReference)(unsafe.Pointer(in.PodRef))
	out.LastUpdated = (*metav1.Time)(unsafe.Pointer(in.LastUpdated))
	out.FailureReason = (*errors.MinerStatusError)(unsafe.Pointer(in.FailureReason))
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.Addresses = *(*apps.MinerAddresses)(unsafe.Pointer(&in.Addresses))
	out.Phase = in.Phase
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*apps.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1beta1_MinerStatus_To_apps_MinerStatus is an autogenerated conversion function.
func Convert_v1beta1_MinerStatus_To_apps_MinerStatus(in *MinerStatus, out *apps.MinerStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerStatus_To_apps_MinerStatus(in, out, s)
}

func autoConvert_apps_MinerStatus_To_v1beta1_MinerStatus(in *apps.MinerStatus, out *MinerStatus, s conversion.Scope) error {
	out.PodRef = (*v1.ObjectReference)(unsafe.Pointer(in.PodRef))
	out.LastUpdated = (*metav1.Time)(unsafe.Pointer(in.LastUpdated))
	out.FailureReason = (*errors.MinerStatusError)(unsafe.Pointer(in.FailureReason))
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.Addresses = *(*MinerAddresses)(unsafe.Pointer(&in.Addresses))
	out.Phase = in.Phase
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_apps_MinerStatus_To_v1beta1_MinerStatus is an autogenerated conversion function.
func Convert_apps_MinerStatus_To_v1beta1_MinerStatus(in *apps.MinerStatus, out *MinerStatus, s conversion.Scope) error {
	return autoConvert_apps_MinerStatus_To_v1beta1_MinerStatus(in, out, s)
}

func autoConvert_v1beta1_MinerTemplateSpec_To_apps_MinerTemplateSpec(in *MinerTemplateSpec, out *apps.MinerTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_ObjectMeta_To_apps_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_MinerSpec_To_apps_MinerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_MinerTemplateSpec_To_apps_MinerTemplateSpec is an autogenerated conversion function.
func Convert_v1beta1_MinerTemplateSpec_To_apps_MinerTemplateSpec(in *MinerTemplateSpec, out *apps.MinerTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerTemplateSpec_To_apps_MinerTemplateSpec(in, out, s)
}

func autoConvert_apps_MinerTemplateSpec_To_v1beta1_MinerTemplateSpec(in *apps.MinerTemplateSpec, out *MinerTemplateSpec, s conversion.Scope) error {
	if err := Convert_apps_ObjectMeta_To_v1beta1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_apps_MinerSpec_To_v1beta1_MinerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_apps_MinerTemplateSpec_To_v1beta1_MinerTemplateSpec is an autogenerated conversion function.
func Convert_apps_MinerTemplateSpec_To_v1beta1_MinerTemplateSpec(in *apps.MinerTemplateSpec, out *MinerTemplateSpec, s conversion.Scope) error {
	return autoConvert_apps_MinerTemplateSpec_To_v1beta1_MinerTemplateSpec(in, out, s)
}

func autoConvert_v1beta1_ObjectMeta_To_apps_ObjectMeta(in *ObjectMeta, out *apps.ObjectMeta, s conversion.Scope) error {
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

// Convert_v1beta1_ObjectMeta_To_apps_ObjectMeta is an autogenerated conversion function.
func Convert_v1beta1_ObjectMeta_To_apps_ObjectMeta(in *ObjectMeta, out *apps.ObjectMeta, s conversion.Scope) error {
	return autoConvert_v1beta1_ObjectMeta_To_apps_ObjectMeta(in, out, s)
}

func autoConvert_apps_ObjectMeta_To_v1beta1_ObjectMeta(in *apps.ObjectMeta, out *ObjectMeta, s conversion.Scope) error {
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

// Convert_apps_ObjectMeta_To_v1beta1_ObjectMeta is an autogenerated conversion function.
func Convert_apps_ObjectMeta_To_v1beta1_ObjectMeta(in *apps.ObjectMeta, out *ObjectMeta, s conversion.Scope) error {
	return autoConvert_apps_ObjectMeta_To_v1beta1_ObjectMeta(in, out, s)
}

func autoConvert_v1beta1_PodInfo_To_apps_PodInfo(in *PodInfo, out *apps.PodInfo, s conversion.Scope) error {
	out.OperatingSystem = in.OperatingSystem
	out.Architecture = in.Architecture
	return nil
}

// Convert_v1beta1_PodInfo_To_apps_PodInfo is an autogenerated conversion function.
func Convert_v1beta1_PodInfo_To_apps_PodInfo(in *PodInfo, out *apps.PodInfo, s conversion.Scope) error {
	return autoConvert_v1beta1_PodInfo_To_apps_PodInfo(in, out, s)
}

func autoConvert_apps_PodInfo_To_v1beta1_PodInfo(in *apps.PodInfo, out *PodInfo, s conversion.Scope) error {
	out.OperatingSystem = in.OperatingSystem
	out.Architecture = in.Architecture
	return nil
}

// Convert_apps_PodInfo_To_v1beta1_PodInfo is an autogenerated conversion function.
func Convert_apps_PodInfo_To_v1beta1_PodInfo(in *apps.PodInfo, out *PodInfo, s conversion.Scope) error {
	return autoConvert_apps_PodInfo_To_v1beta1_PodInfo(in, out, s)
}
