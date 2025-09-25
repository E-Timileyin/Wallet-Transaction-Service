package models

// UpdateUserRoleRequest defines the structure for a request to update a user's role.
type UpdateUserRoleRequest struct {
	Role string `json:"role" binding:"required,oneof=user admin"`
}
