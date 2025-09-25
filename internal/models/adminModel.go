package models

// UpdateUserRoleRequest defines the structure for a request to update a user's role.
// @Description Update user role request
type UpdateUserRoleRequest struct {
	Role string `json:"role" example:"admin" binding:"required,oneof=user admin"`
}
