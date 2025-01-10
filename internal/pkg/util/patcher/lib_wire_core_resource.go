package patcher

import (
	lib_wire_core_resource_permission "github.com/rodericusifo/employee-management-api/lib/injector/core/resource/permission"
	lib_wire_core_resource_role "github.com/rodericusifo/employee-management-api/lib/injector/core/resource/role"
	lib_wire_core_resource_role_permission "github.com/rodericusifo/employee-management-api/lib/injector/core/resource/role_permission"
	lib_wire_core_resource_user "github.com/rodericusifo/employee-management-api/lib/injector/core/resource/user"
)

var (
	UserResource           = lib_wire_core_resource_user.UserResource
	RoleResource           = lib_wire_core_resource_role.RoleResource
	PermissionResource     = lib_wire_core_resource_permission.PermissionResource
	RolePermissionResource = lib_wire_core_resource_role_permission.RolePermissionResource
)
