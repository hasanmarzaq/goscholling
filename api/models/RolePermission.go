package models

type RolePermission struct {
	RoleID       uint32 `json:"role_id"`
	PermissionID uint32 `json:"permission_id"`
}
type RolePermissions []RolePermission
