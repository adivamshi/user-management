package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-management-api/domain"
	"user-management-api/repository"
	"user-management-api/routes/users"
)

func MakeUserHandler(endpoint domain.Endpoint, request interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Parse User Request Body
		if request != nil {
			err := ctx.ShouldBind(request)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		response, err := endpoint(ctx, request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func SetupRoutes(r *gin.Engine) {
	dbConn, err := repository.ConnectToDB()
	if err != nil {
		panic(err)
	}
	userRepo := repository.NewUserRepository(dbConn)
	userService := users.NewService(userRepo)

	api := r.Group("/api/v1/users")
	{
		api.POST("/", MakeUserHandler(users.CreateUserHandler(userService), &domain.User{}))
		api.GET("/", MakeUserHandler(users.GetUsersHandler(userService), nil))
		api.GET("/:id", MakeUserHandler(users.GetUserHandler(userService), nil))
		api.PUT("/:id", MakeUserHandler(users.UpdateUserHandler(userService), &domain.User{}))
		api.DELETE("/:id", MakeUserHandler(users.DeleteUserHandler(userService), nil))
		api.GET("/search", MakeUserHandler(users.SearchUsersHandler(userService), nil))
	}

}
