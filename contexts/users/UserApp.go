package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcosperez/remember-go/contexts/shared/model"
)

type UserApp struct {
	name string

	model.App
}

func (app *UserApp) Start() (bool, error) {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()

	if err != nil {
		return false, err
	}

	return true, nil
}

func NewUserApp() *UserApp {
	userApp := UserApp{name: "UserApp"}

	return &userApp
}
