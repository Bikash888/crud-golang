package Routes

import (
	Controller "first/src/first-api/Controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()
	u1 := r.Group("/admin")
	{
		u1.POST("save", Controller.CreateUser)
		u1.GET("all", Controller.GetAllUsers)
		u1.GET("user/:id", Controller.GetUserById)
		u1.PUT("update/:id", Controller.UpdateUserById)
		u1.DELETE("delete/:id", Controller.DeleteUserById)

	}
	return r

}
