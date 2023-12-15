package users

import (
	d "implement_middleware/config"
	"implement_middleware/lib"
	"implement_middleware/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Login controller
func LoginUsersController(c echo.Context) error {
	auth := model.Users{}
	c.Bind(&auth)

	user, err := lib.LoginUsers(&auth)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// create new user
func RegisterController(c echo.Context) error {
	user := model.Users{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := d.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// get all users
func GetUsersController(c echo.Context) error {
	users := []model.Users{}

	if err := d.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := model.Users{}
	if err := d.DB.First(&user, idInt).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"user":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := model.Users{}
	res := d.DB.Delete(&user, idInt)

	if err := res.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed delete data",
		})
	}

	d.DB.Unscoped().First(&user, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := model.Users{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res := d.DB.Model(&user).Where("id = ?", idInt).
		Updates(model.Users{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})

	if err := res.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed update data",
		})
	}

	d.DB.First(&user, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}
