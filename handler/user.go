package handler

import (
    "net/http"
    "strconv"

    "crud/dto"
    "crud/service"

    "github.com/labstack/echo/v4"
)

// CreateUser handler
func CreateUser(c echo.Context) error {
    var req dto.CreateUserRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
    }

    // for now createdBy is 0 (system) — can be replaced with auth user ID later
    user, err := service.CreateUser(req.Username, req.Password, req.Email, "GO_USR_", 0)
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

    user, err := service.GetUserByID(uint(id64))
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

    user, err := service.UpdateUser(uint(id64), req.Username, req.Password, req.Email, req.Status, 0)
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

    if err := service.DeleteUser(uint(id64)); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.NoContent(http.StatusNoContent)
}

// ListUsers handler
func ListUsers(c echo.Context) error {
    users, err := service.ListUsers()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    res := dto.ListUsersResponse{Users: make([]dto.UserResponse, 0, len(users))}
    for _, u := range users {
        res.Users = append(res.Users, dto.UserResponse{
            ID:        u.ID,
            Code:      u.Code,
            Username:  u.Username,
            Email:     u.Email,
            Status:    u.Status,
            CreatedAt: u.CreatedAt,
            UpdatedAt: u.UpdatedAt,
        })
    }

    return c.JSON(http.StatusOK, res)
}
