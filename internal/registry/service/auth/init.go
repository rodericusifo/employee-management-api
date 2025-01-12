package auth

import (
	internal_app_core_auth_service "github.com/rodericusifo/employee-management-api/internal/app/core/auth/service"
	internal_registry_resource_role "github.com/rodericusifo/employee-management-api/internal/registry/resource/role"
	internal_registry_resource_user "github.com/rodericusifo/employee-management-api/internal/registry/resource/user"
)

func AuthService() internal_app_core_auth_service.IAuthService {
	iUserResource := internal_registry_resource_user.UserResource()
	iRoleResource := internal_registry_resource_role.RoleResource()
	iAuthService := internal_app_core_auth_service.InitAuthService(iUserResource, iRoleResource)
	return iAuthService
}
