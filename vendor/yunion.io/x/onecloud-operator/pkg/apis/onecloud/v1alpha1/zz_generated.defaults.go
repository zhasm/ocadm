// +build !ignore_autogenerated

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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&OnecloudCluster{}, func(obj interface{}) { SetObjectDefaults_OnecloudCluster(obj.(*OnecloudCluster)) })
	scheme.AddTypeDefaultingFunc(&OnecloudClusterConfig{}, func(obj interface{}) { SetObjectDefaults_OnecloudClusterConfig(obj.(*OnecloudClusterConfig)) })
	scheme.AddTypeDefaultingFunc(&OnecloudClusterList{}, func(obj interface{}) { SetObjectDefaults_OnecloudClusterList(obj.(*OnecloudClusterList)) })
	return nil
}

func SetObjectDefaults_OnecloudCluster(in *OnecloudCluster) {
	SetDefaults_OnecloudCluster(in)
	SetDefaults_Mysql(&in.Spec.Mysql)
}

func SetObjectDefaults_OnecloudClusterConfig(in *OnecloudClusterConfig) {
	SetDefaults_OnecloudClusterConfig(in)
	SetDefaults_KeystoneConfig(&in.Keystone)
}

func SetObjectDefaults_OnecloudClusterList(in *OnecloudClusterList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_OnecloudCluster(a)
	}
}
