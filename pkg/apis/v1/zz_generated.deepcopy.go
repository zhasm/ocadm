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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	net "net"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfiguration) DeepCopyInto(out *ClusterConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.MysqlConnection = in.MysqlConnection
	in.Keystone.DeepCopyInto(&out.Keystone)
	in.RegionServer.DeepCopyInto(&out.RegionServer)
	in.Glance.DeepCopyInto(&out.Glance)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfiguration.
func (in *ClusterConfiguration) DeepCopy() *ClusterConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClusterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBInfo) DeepCopyInto(out *DBInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBInfo.
func (in *DBInfo) DeepCopy() *DBInfo {
	if in == nil {
		return nil
	}
	out := new(DBInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Glance) DeepCopyInto(out *Glance) {
	*out = *in
	in.ServiceCommonOptions.DeepCopyInto(&out.ServiceCommonOptions)
	out.ServiceDBOptions = in.ServiceDBOptions
	if in.TargetImageFormats != nil {
		in, out := &in.TargetImageFormats, &out.TargetImageFormats
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Glance.
func (in *Glance) DeepCopy() *Glance {
	if in == nil {
		return nil
	}
	out := new(Glance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostLocalInfo) DeepCopyInto(out *HostLocalInfo) {
	*out = *in
	in.ManagementNetInterface.DeepCopyInto(&out.ManagementNetInterface)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostLocalInfo.
func (in *HostLocalInfo) DeepCopy() *HostLocalInfo {
	if in == nil {
		return nil
	}
	out := new(HostLocalInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitConfiguration) DeepCopyInto(out *InitConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.InitConfiguration.DeepCopyInto(&out.InitConfiguration)
	in.ClusterConfiguration.DeepCopyInto(&out.ClusterConfiguration)
	in.HostLocalInfo.DeepCopyInto(&out.HostLocalInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitConfiguration.
func (in *InitConfiguration) DeepCopy() *InitConfiguration {
	if in == nil {
		return nil
	}
	out := new(InitConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InitConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Keystone) DeepCopyInto(out *Keystone) {
	*out = *in
	in.ServiceBaseOptions.DeepCopyInto(&out.ServiceBaseOptions)
	out.ServiceDBOptions = in.ServiceDBOptions
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Keystone.
func (in *Keystone) DeepCopy() *Keystone {
	if in == nil {
		return nil
	}
	out := new(Keystone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlConnection) DeepCopyInto(out *MysqlConnection) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlConnection.
func (in *MysqlConnection) DeepCopy() *MysqlConnection {
	if in == nil {
		return nil
	}
	out := new(MysqlConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetInterface) DeepCopyInto(out *NetInterface) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetInterface.
func (in *NetInterface) DeepCopy() *NetInterface {
	if in == nil {
		return nil
	}
	out := new(NetInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionServer) DeepCopyInto(out *RegionServer) {
	*out = *in
	in.ServiceCommonOptions.DeepCopyInto(&out.ServiceCommonOptions)
	out.ServiceDBOptions = in.ServiceDBOptions
	if in.DNSResolvers != nil {
		in, out := &in.DNSResolvers, &out.DNSResolvers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionServer.
func (in *RegionServer) DeepCopy() *RegionServer {
	if in == nil {
		return nil
	}
	out := new(RegionServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBaseOptions) DeepCopyInto(out *ServiceBaseOptions) {
	*out = *in
	if in.CorsHosts != nil {
		in, out := &in.CorsHosts, &out.CorsHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotifyAdminUsers != nil {
		in, out := &in.NotifyAdminUsers, &out.NotifyAdminUsers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotifyAdminGroups != nil {
		in, out := &in.NotifyAdminGroups, &out.NotifyAdminGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBaseOptions.
func (in *ServiceBaseOptions) DeepCopy() *ServiceBaseOptions {
	if in == nil {
		return nil
	}
	out := new(ServiceBaseOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceCommonOptions) DeepCopyInto(out *ServiceCommonOptions) {
	*out = *in
	in.ServiceBaseOptions.DeepCopyInto(&out.ServiceBaseOptions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceCommonOptions.
func (in *ServiceCommonOptions) DeepCopy() *ServiceCommonOptions {
	if in == nil {
		return nil
	}
	out := new(ServiceCommonOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDBOptions) DeepCopyInto(out *ServiceDBOptions) {
	*out = *in
	out.DBInfo = in.DBInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDBOptions.
func (in *ServiceDBOptions) DeepCopy() *ServiceDBOptions {
	if in == nil {
		return nil
	}
	out := new(ServiceDBOptions)
	in.DeepCopyInto(out)
	return out
}
