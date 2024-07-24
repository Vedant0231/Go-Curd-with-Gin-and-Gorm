package src

import (
	"encoding/json"
	"fmt"
	userstruct "go_gin/model"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:netcore@tcp(localhost:3306)/netcore_db?charset=utf8mb4&parseTime=True&loc=Local"

func Initalmigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&userstruct.User{})
	log.Println("Database migration complete")
}

func Createuser(c *gin.Context) {
	fmt.Println(DB)
	var user userstruct.User
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		log.Fatal("error decoding")
	}
	fmt.Println(user)
	DB.Create(&user)
	json.NewEncoder(c.Writer).Encode(user)
}

func Getusers(c *gin.Context) {
	var users []userstruct.User
	DB.Find(&users)
	json.NewEncoder(c.Writer).Encode(users)

}

func Getuser(c *gin.Context) {
	var users userstruct.User
	DB.Find(&users, c.Param("id"))
	json.NewEncoder(c.Writer).Encode(users)

}

func Updateuser(c *gin.Context) {
	var user userstruct.User
	DB.First(&user, c.Param("id"))
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		log.Fatal("error decoding")
	}
	DB.Save(&user)
	json.NewEncoder(c.Writer).Encode(user)

}

func Deleteuser(c *gin.Context) {
	var users userstruct.User
	DB.Delete(&users, c.Param("id"))
	json.NewEncoder(c.Writer).Encode("user is deleted")

}
