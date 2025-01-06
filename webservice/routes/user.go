package routes

import (
	"net/http"
	"webservice/env"
	"webservice/pkg/auth"

	"github.com/gin-gonic/gin"
)

func LoadUserRoutes(route *gin.Engine) {
	userAuth := route.Group("/auth/user")
	{
		userAuth.POST("/create", createUser)
		userAuth.PATCH("/claims", updateUserClaims)
		userAuth.GET("", BearerTokenMiddleware, getUserDetails)
	}
}

type createUserRequestObj struct {
	Username     string        `json:"username"`
	PasswordHash string        `json:"passwordHash"`
	Role         auth.UserRole `json:"role"`
}

type createUserErrorResponse struct {
	Error string `json:"error"`
}

func createUser(c *gin.Context) {
	authRepo := c.MustGet(env.AuthRepo).(auth.UserDataStorage)
	var userDetails createUserRequestObj

	if err := c.BindJSON(&userDetails); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	_, createUserError := auth.CreateUser(authRepo, userDetails.Username, userDetails.PasswordHash, userDetails.Role)
	if createUserError != nil {
		c.IndentedJSON(http.StatusInternalServerError, createUserErrorResponse{
			Error: createUserError.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
	return
}

type updateUserClaimsRequestObj struct {
	Username string                    `json:"username"`
	Claims   map[auth.Claim]auth.Claim `json:"claims"`
}

func updateUserClaims(c *gin.Context) {
	authRepo := c.MustGet(env.AuthRepo).(auth.UserDataStorage)
	var newUserClaims updateUserClaimsRequestObj

	if err := c.BindJSON(&newUserClaims); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	auth.UpdateUserClaims(authRepo, newUserClaims.Username, newUserClaims.Claims)
	c.Status(http.StatusOK)

}

type getUserDetailsResponse struct {
	Username string `json:"username"`
}

func getUserDetails(c *gin.Context) {
	currentUser := c.MustGet(env.User).(auth.User)

	response := getUserDetailsResponse{
		Username: currentUser.Username,
	}
	c.IndentedJSON(200, response)

}
