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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1beta1 "github.com/crossplane-contrib/provider-aws/apis/ec2/v1beta1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	if mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo != nil {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.ClientSubnets),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.ClientSubnetRefs,
			Selector:      mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.ClientSubnetSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.ClientSubnets")
		}
		mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.ClientSubnets = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.ClientSubnetRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo != nil {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.SecurityGroupRefs,
			Selector:      mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.SecurityGroupSelector,
			To: reference.To{
				List:    &v1beta1.SecurityGroupList{},
				Managed: &v1beta1.SecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.SecurityGroups")
		}
		mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.CustomClusterParameters.CustomBrokerNodeGroupInfo.SecurityGroupRefs = mrsp.ResolvedReferences

	}

	return nil
}
