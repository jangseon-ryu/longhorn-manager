/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "k8s.io/api/networking/v1alpha1"
	networkingv1alpha1 "k8s.io/client-go/applyconfigurations/networking/v1alpha1"
	gentype "k8s.io/client-go/gentype"
	typednetworkingv1alpha1 "k8s.io/client-go/kubernetes/typed/networking/v1alpha1"
)

// fakeServiceCIDRs implements ServiceCIDRInterface
type fakeServiceCIDRs struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.ServiceCIDR, *v1alpha1.ServiceCIDRList, *networkingv1alpha1.ServiceCIDRApplyConfiguration]
	Fake *FakeNetworkingV1alpha1
}

func newFakeServiceCIDRs(fake *FakeNetworkingV1alpha1) typednetworkingv1alpha1.ServiceCIDRInterface {
	return &fakeServiceCIDRs{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.ServiceCIDR, *v1alpha1.ServiceCIDRList, *networkingv1alpha1.ServiceCIDRApplyConfiguration](
			fake.Fake,
			"",
			v1alpha1.SchemeGroupVersion.WithResource("servicecidrs"),
			v1alpha1.SchemeGroupVersion.WithKind("ServiceCIDR"),
			func() *v1alpha1.ServiceCIDR { return &v1alpha1.ServiceCIDR{} },
			func() *v1alpha1.ServiceCIDRList { return &v1alpha1.ServiceCIDRList{} },
			func(dst, src *v1alpha1.ServiceCIDRList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.ServiceCIDRList) []*v1alpha1.ServiceCIDR {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.ServiceCIDRList, items []*v1alpha1.ServiceCIDR) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
