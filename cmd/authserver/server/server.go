package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
	Shutdown()
}

type AuthServer struct {
	router *gin.Engine
}

func (auth *AuthServer) Run() error {

	auth.router = gin.New()

	api_v1 := auth.router.Group("api/v1")

	api_v1.POST("/signup", signupHandler)

	return auth.router.Run(":8080")
}

func (auth *AuthServer) Shutdown() {

}

type user struct {
	Email string `json:"email"`
}

// func logWrapper(handler func(c *gin.Context))

func signupHandler(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)

}
