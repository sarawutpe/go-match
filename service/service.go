package service

import (
	"fmt"
	"main/db"
	"main/helper"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveFile(c *gin.Context) {
	id := c.Param("id")

	if _, err := helper.RemoveFile(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"success": true, "data": id})

}

func Upload(c *gin.Context) {
	file, err := c.FormFile("image")

	upload, err := helper.SaveFile(c, file, err)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully" + upload})
}

func CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	collection, _ := db.Collection(db.UsersCollection)
	user := model.User{}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("erro")
		return
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error  df": err.Error()})
		return
	}

	fmt.Println("Inserted document with ID:", result.InsertedID)
}

// Retrieve the inserted document from MongoDB
// filter := bson.M{"_id": result.InsertedID}
// insertedDoc := &model.User{}
// err = collection.FindOne(ctx, filter).Decode(insertedDoc)

// if err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	return
// }

// c.JSON(http.StatusCreated, gin.H{"success": true, "data": insertedDoc})