package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rgafarov/xts/cmd/authserver/app"
)

type AuthContext struct {
	*gin.Context
	App *app.App
}

type AuthServer struct {
	router *gin.Engine
}

func (auth *AuthServer) Run(app *app.App) error {

	auth.router = gin.New()

	api_v1 := auth.router.Group("api/v1")

	a := appWrapper(app)
	api_v1.POST("/signup", a(signupHandler))

	return auth.router.Run(":8080")
}

func (auth *AuthServer) Shutdown() {

}

type user struct {
	Email string `json:"email"`
}

func appWrapper(app *app.App) func(func(*AuthContext)) gin.HandlerFunc {
	return func(authHandler func(*AuthContext)) gin.HandlerFunc {
		return func(ginContext *gin.Context) {
			actx := &AuthContext{
				Context: ginContext,
				App:     app,
			}
			authHandler(actx)
		}
	}
}

func signupHandler(c *AuthContext) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)

}
