package dto

import "time"

type CreateUserRequest struct {
    Username string `json:"username" form:"username"` 
    Password string `json:"password" form:"password"`
    Email    string `json:"email" form:"email"`
}

type UpdateUserRequest struct {
    Username *string `json:"username,omitempty" form:"username"`
    Password *string `json:"password,omitempty" form:"password"`
    Email    *string `json:"email,omitempty" form:"email"`
    Status   *string `json:"status,omitempty" form:"status"`
}

type UserResponse struct {
    ID        uint      `json:"id"`
    Code      string    `json:"code"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type ListUsersResponse struct {
    Users []UserResponse `json:"users"`
}
