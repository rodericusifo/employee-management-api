package patcher

import (
	registry_resource_permission "github.com/rodericusifo/employee-management-api/registry/resource/permission"
	registry_resource_role "github.com/rodericusifo/employee-management-api/registry/resource/role"
	registry_resource_role_permission "github.com/rodericusifo/employee-management-api/registry/resource/role_permission"
	registry_resource_user "github.com/rodericusifo/employee-management-api/registry/resource/user"
)

var (
	UserResource           = registry_resource_user.UserResource
	RoleResource           = registry_resource_role.RoleResource
	PermissionResource     = registry_resource_permission.PermissionResource
	RolePermissionResource = registry_resource_role_permission.RolePermissionResource
)
