// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apis

type DomainizedResourceListInput struct {
	// swagger:ignore
	// Is an admin call? equivalent to scope=system
	// Deprecated
	Admin *bool `json:"admin"`

	// 指定查询的权限范围，可能值为project, domain or system
	Scope string `json:"scope"`

	// 查询指定的域（ID或名称）拥有的资源
	ProjectDomain string `json:"project_domain"`
	// swagger:ignore
	// Deprecated
	// Project domain Id filter, alias for project_domain
	ProjectDomainId string `json:"project_domain_id" deprecated-by:"project_domain"`
	// swagger:ignore
	// Deprecated
	// Domain Id filter, alias for project_domain
	DomainId string `json:"domain_id" deprecated-by:"project_domain"`
}

type DomainizedResourceCreateInput struct {
	// description: the owner domain name or id
	// required: false
	Domain string `json:"project_domain"`

	// description: the owner domain name or id, alias field of domain
	// required: false
	DomainId string `json:"domain_id"`
}

type ProjectizedResourceListInput struct {
	DomainizedResourceListInput

	// 查询指定的项目（ID或名称）拥有的资源
	Project string `json:"project"`
	// swagger:ignore
	// Deprecated
	// Filter by project_id, alias for project
	ProjectId string `json:"project_id" deprecated-by:"project"`
	// swagger:ignore
	// Deprecated
	// Filter by tenant ID or Name, alias for project
	Tenant string `json:"tenant" deprecated-by:"project"`
	// swagger:ignore
	// Deprecated
	// Filter by tenant_id, alias for project
	TenantId string `json:"tenant_id" deprecated-by:"project"`
}

type ProjectizedResourceCreateInput struct {
	DomainizedResourceCreateInput

	// description: the owner project name or id
	// required: false
	Project string `json:"project"`

	// description: the owner project name or id, alias field of project
	// required: false
	ProjectId string `json:"project_id"`
}

type UserResourceListInput struct {
	// 查询指定的用户（ID或名称）拥有的资源
	User string `json:"user"`
	// swagger:ignore
	// Deprecated
	// Filter by userId
	UserId string `json:"user_id" deprecated-by:"user"`
}

func (input UserResourceListInput) UserStr() string {
	if len(input.User) > 0 {
		return input.User
	}
	if len(input.UserId) > 0 {
		return input.UserId
	}
	return ""
}

type SharableVirtualResourceCreateInput struct {
	VirtualResourceCreateInput

	// description: indicate the resource is a public resource
	// required: false
	IsPublic *bool `json:"is_public"`

	// description: indicate the shared scope for a public resource, which can be domain or system or none
	// required: false
	PublicScope string `json:"public_scope"`
}

type VirtualResourceCreateInput struct {
	StatusStandaloneResourceCreateInput
	ProjectizedResourceCreateInput

	// description: indicate the resource is a system resource, which is not visible to user
	// required: false
	IsSystem *bool `json:"is_system"`
}

type EnabledStatusStandaloneResourceCreateInput struct {
	StatusStandaloneResourceCreateInput

	// description: indicate the resource is enabled/disabled by administrator
	// required: false
	Enabled *bool `json:"enabled"`
}

type StatusStandaloneResourceCreateInput struct {
	StandaloneResourceCreateInput

	// description: the status of the resource
	// required: false
	Status string `json:"status"`
}

type StandaloneResourceCreateInput struct {
	ResourceBaseCreateInput

	// description: resource name, required if generated_name is not given
	// unique: true
	// required: true
	// example: test-network
	Name string `json:"name"`

	// description: generated resource name, given a pattern to generate name, required if name is not given
	// unique: false
	// required: false
	// example: test###
	GenerateName string `json:"generate_name"`

	// description: resource description
	// required: false
	// example: test create network
	Description string `json:"description"`

	// description: the resource is an emulated resource
	// required: false
	IsEmulated *bool `json:"is_emulated"`

	// 标签列表,最多支持20个
	// example: { "user:rd": "op" }
	Metadata map[string]string `json:"__meta__"`
}

type JoinResourceBaseCreateInput struct {
	ResourceBaseCreateInput
}

type ResourceBaseCreateInput struct {
	ModelBaseCreateInput
}

type ModelBaseCreateInput struct {
	Meta
}
