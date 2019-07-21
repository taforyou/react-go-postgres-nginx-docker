package level11api

import (
	"fmt"
	"level11infrastructure"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// ต้อง Refactor
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//----------
// Handlers
//----------

func CreateUser(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}

	params := make(map[interface{}]interface{})
	params["sql"] = fmt.Sprintf("INSERT INTO users (name,email) VALUES ('%v','%v')", u.Name, u.Email)
	level11infrastructure.Execute(params)
	return c.JSON(http.StatusOK, u)
}

func GetUsers(c echo.Context) (err error) {

	params := make(map[interface{}]interface{})
	params["sql"] = fmt.Sprintf("SELECT id,name,email FROM users ")
	u := level11infrastructure.Query(params)
	return c.JSON(http.StatusOK, u)
}

func TestFetch(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, "TestFetch test"+os.Getenv("POSTGRES_PORT")+os.Getenv("POSTGRES_USER")+os.Getenv("POSTGRES_PASSWORD")+os.Getenv("POSTGRES_DB"))
}
