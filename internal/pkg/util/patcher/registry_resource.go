package patcher

import (
	internal_registry_resource_permission "github.com/rodericusifo/employee-management-api/internal/registry/resource/permission"
	internal_registry_resource_role "github.com/rodericusifo/employee-management-api/internal/registry/resource/role"
	internal_registry_resource_role_permission "github.com/rodericusifo/employee-management-api/internal/registry/resource/role_permission"
	internal_registry_resource_user "github.com/rodericusifo/employee-management-api/internal/registry/resource/user"
)

var (
	UserResource           = internal_registry_resource_user.UserResource
	RoleResource           = internal_registry_resource_role.RoleResource
	PermissionResource     = internal_registry_resource_permission.PermissionResource
	RolePermissionResource = internal_registry_resource_role_permission.RolePermissionResource
)
