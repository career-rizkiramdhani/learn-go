package handler

import (
	"net/http"
	"net/mail"
	"strconv"

	userapp "crud/application/user"
	"crud/dto"

	"github.com/labstack/echo/v4"
)

// CreateUser handler
func CreateUser(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// basic validation
	if req.Username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "username is required"})
	}
	if len(req.Password) < 6 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "password must be at least 6 characters"})
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid email address"})
	}

	// for now createdBy is 0 (system) — can be replaced with auth user ID later
	user, err := userapp.CreateUser(req.Username, req.Password, req.Email, "GO_USR_", 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res := dto.UserResponse{
		ID:        user.ID,
		Code:      user.Code,
		Username:  user.Username,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, res)
}

// GetUser handler
func GetUser(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	user, err := userapp.GetUserByID(uint(id64))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	res := dto.UserResponse{
		ID:        user.ID,
		Code:      user.Code,
		Username:  user.Username,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, res)
}

// UpdateUser handler
func UpdateUser(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// validate provided fields
	if req.Password != nil && len(*req.Password) > 0 && len(*req.Password) < 6 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "password must be at least 6 characters"})
	}
	if req.Email != nil {
		if _, err := mail.ParseAddress(*req.Email); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid email address"})
		}
	}

	user, err := userapp.UpdateUser(uint(id64), req.Username, req.Password, req.Email, req.Status, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res := dto.UserResponse{
		ID:        user.ID,
		Code:      user.Code,
		Username:  user.Username,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, res)
}

// DeleteUser handler
func DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	if err := userapp.DeleteUser(uint(id64)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// ListUsers handler
func ListUsers(c echo.Context) error {
	// parse pagination query params
	page := 1
	size := 10
	if p := c.QueryParam("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if s := c.QueryParam("size"); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v > 0 {
			size = v
		}
	}

	users, total, err := userapp.ListUsers(page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	resp := dto.ListUsersResponse{Users: make([]dto.UserResponse, 0, len(users))}
	for _, u := range users {
		resp.Users = append(resp.Users, dto.UserResponse{
			ID:        u.ID,
			Code:      u.Code,
			Username:  u.Username,
			Email:     u.Email,
			Status:    u.Status,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	resp.Meta = dto.ListMeta{Page: page, Size: size, Total: total}

	return c.JSON(http.StatusOK, resp)
}
