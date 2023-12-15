package books

import (
	d "implement_middleware/config"
	"implement_middleware/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// create new book
func CreateBookController(c echo.Context) error {
	book := model.Books{}
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := d.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// get all books
func GetBooksController(c echo.Context) error {
	books := []model.Books{}

	if err := d.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book := model.Books{}
	if err := d.DB.First(&book, idInt).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"book":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book := model.Books{}
	res := d.DB.Delete(&book, idInt)

	if err := res.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed delete data",
		})
	}

	d.DB.Unscoped().First(&book, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
		"book":    book,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	book := model.Books{}
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res := d.DB.Model(&book).Where("id = ?", idInt).
		Updates(model.Books{
			Title:     book.Title,
			Author:    book.Author,
			Publisher: book.Publisher,
		})

	if err := res.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed update data",
		})
	}

	d.DB.First(&book, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}
