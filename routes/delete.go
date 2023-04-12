package routes

import (
 getcollection "Event_API/Collection"
 database "Event_API/databases"
 "context"
 "net/http"
 "time"
 "github.com/gin-gonic/gin"
 "go.mongodb.org/mongo-driver/bson"
 "go.mongodb.org/mongo-driver/bson/primitive"
)

func DeletePost(c *gin.Context) {
 ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
 var DB = database.ConnectDB()
 postId := c.Param("eventId")

 var postCollection = getcollection.GetCollection(DB, "Event") 
 defer cancel()
 objId, _ := primitive.ObjectIDFromHex(postId) 
 result, err := postCollection.DeleteOne(ctx, bson.M{"id": objId}) 
 res := map[string]interface{}{"data": result}

 if err != nil {
 c.JSON(http.StatusInternalServerError, gin.H{"message": err})
 return
 }

 if result.DeletedCount < 1 {
 c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
 return
 }

 c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully", "Data": res})
}