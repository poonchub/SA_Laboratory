package genders


import (

   "net/http"


   "example.com/sa-67-example/config"

//    "example.com/sa-67-example/models"

   "example.com/sa-67-example/entity"

   "github.com/gin-gonic/gin"

)


func GetAll(c *gin.Context) {


   db := config.DB()


   var genders []models.Genders

   db.Find(&genders)


   c.JSON(http.StatusOK, &genders)


}