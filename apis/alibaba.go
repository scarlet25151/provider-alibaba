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

// Package apis contains Kubernetes API for the Alibaba Cloud provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	databasev1alpha1 "github.com/crossplane/provider-alibaba/apis/database/v1alpha1"
	nasv1alpha1 "github.com/crossplane/provider-alibaba/apis/nas/v1alpha1"
	ossv1alpha1 "github.com/crossplane/provider-alibaba/apis/oss/v1alpha1"
	redisv1alpha1 "github.com/crossplane/provider-alibaba/apis/redis/v1alpha1"
	slbv1alpha1 "github.com/crossplane/provider-alibaba/apis/slb/v1alpha1"
	slsv1alpha1 "github.com/crossplane/provider-alibaba/apis/sls/v1alpha1"
	"github.com/crossplane/provider-alibaba/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		databasev1alpha1.SchemeBuilder.AddToScheme,
		slsv1alpha1.AddToScheme,
		ossv1alpha1.SchemeBuilder.AddToScheme,
		nasv1alpha1.AddToScheme,
		slbv1alpha1.AddToScheme,
		redisv1alpha1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
