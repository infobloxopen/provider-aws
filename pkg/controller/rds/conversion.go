package rds

import (
	svcsdk "github.com/aws/aws-sdk-go/service/rds"

	databasev1beta1 "github.com/crossplane-contrib/provider-aws/apis/database/v1beta1"
	rdsv1alpha1 "github.com/crossplane-contrib/provider-aws/apis/rds/v1alpha1"
	"github.com/crossplane-contrib/provider-aws/pkg/generics"
)

type Tag interface {
	*rdsv1alpha1.Tag |
		*databasev1beta1.Tag
}

type WithTagFuncs interface {
	GetKey() string
	GetValue() string
	SetKey(string)
	SetValue(string)
}

type databasev1beta1Tag struct {
	*databasev1beta1.Tag
}

func (p *databasev1beta1Tag) GetKey() string {
	return p.Key
}

func (p *databasev1beta1Tag) GetValue() string {
	return p.Value
}

func (p *databasev1beta1Tag) SetKey(v string) {
	p.Key = v
}

func (p *databasev1beta1Tag) SetValue(v string) {
	p.Value = v
}

type rdsv1alpha1Tag struct {
	*rdsv1alpha1.Tag
}

func (p *rdsv1alpha1Tag) GetKey() string {
	return *p.Key
}

func (p *rdsv1alpha1Tag) GetValue() string {
	return *p.Value
}

func (p *rdsv1alpha1Tag) SetKey(v string) {
	p.Key = &v
}

func (p *rdsv1alpha1Tag) SetValue(v string) {
	p.Value = &v
}

func ToAWSSDKTag[INPUT Tag](i INPUT) *svcsdk.Tag {
	switch any(i).(type) {
	case *rdsv1alpha1.Tag:
		t := any(i).(*rdsv1alpha1.Tag)
		return &svcsdk.Tag{Key: t.Key, Value: t.Value}
	case *databasev1beta1.Tag:
		t := any(i).(*databasev1beta1.Tag)
		k, v := t.Key, t.Value
		return &svcsdk.Tag{Key: &k, Value: &v}
	default:
		return nil
	}
}

var _ generics.ItemConversionFunc[*rdsv1alpha1.Tag, *svcsdk.Tag] = ToAWSSDKTag[*rdsv1alpha1.Tag]

func FromAWSSDKTag[OUTPUT Tag](i *svcsdk.Tag, o OUTPUT) OUTPUT {
	v := any(o).(WithTagFuncs)
	v.SetKey(*i.Key)
	v.SetValue(*i.Value)
	return o
}

type ProcessorFeature interface {
	*databasev1beta1.ProcessorFeature |
		*rdsv1alpha1.ProcessorFeature
}

func ToAWSProcessorFeature[INPUT ProcessorFeature](i INPUT) *svcsdk.ProcessorFeature {
	var v WithProcessorFeatureFuncs
	switch any(i).(type) {
	case *databasev1beta1.ProcessorFeature:
		v = &databasev1beta1ProcessorFeature{any(i).(*databasev1beta1.ProcessorFeature)}
	case *rdsv1alpha1.ProcessorFeature:
		v = &rdsv1alpha1ProcessorFeature{any(i).(*rdsv1alpha1.ProcessorFeature)}
	default:
		return nil
	}
	name, value := v.GetName(), v.GetValue()
	return &svcsdk.ProcessorFeature{Name: &name, Value: &value}
}

func FromAWSProcessorFeature[OUTPUT ProcessorFeature](i *svcsdk.ProcessorFeature, o OUTPUT) OUTPUT {
	v := any(o).(WithProcessorFeatureFuncs)
	v.SetName(*i.Name)
	v.SetValue(*i.Value)
	return o
}

type WithProcessorFeatureFuncs interface {
	GetName() string
	GetValue() string
	SetName(string)
	SetValue(string)
}

type databasev1beta1ProcessorFeature struct {
	*databasev1beta1.ProcessorFeature
}

func (p *databasev1beta1ProcessorFeature) GetName() string {
	return p.Name
}

func (p *databasev1beta1ProcessorFeature) GetValue() string {
	return p.Value
}

func (p *databasev1beta1ProcessorFeature) SetName(v string) {
	p.Name = v
}

func (p *databasev1beta1ProcessorFeature) SetValue(v string) {
	p.Value = v
}

type rdsv1alpha1ProcessorFeature struct {
	*rdsv1alpha1.ProcessorFeature
}

func (p *rdsv1alpha1ProcessorFeature) GetName() string {
	return *p.Name
}

func (p *rdsv1alpha1ProcessorFeature) GetValue() string {
	return *p.Value
}

func (p *rdsv1alpha1ProcessorFeature) SetName(v string) {
	p.Name = &v
}

func (p *rdsv1alpha1ProcessorFeature) SetValue(v string) {
	p.Value = &v
}
