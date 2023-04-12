package main
import (
routes "Event_API/routes"
"github.com/gin-gonic/gin"
)
func main() {
router := gin.Default()

router.POST("/", routes.CreatePost)

// called as localhost:3000/getOne/{id}
router.GET("getOne/:eventId", routes.ReadOnePost)

// called as localhost:3000/update/{id}
// router.PUT("/update/:eventId", routes.UpdatePost)

// called as localhost:3000/delete/{id}
// router.DELETE("/delete/:eventId", routes.DeletePost)

router.Run("localhost: 3000")
}