package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

//Person struktur
type Posts struct {
	ID           uint   `json:"id"`
	Title	     string `json:"title"`
	Content      string `json:"content"`
	Category	 string `json:"category"`
	Status       string `json:"status"`
	// Created_date date 	`json:"created_date"`
	// Updated_date date 	`json:"updated_date"`


}

func main() {
	db, _ = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/article")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&Posts{})
	r := gin.Default()
	r.GET("/article/", Getposts)
	r.GET("/article/:id", GetPosts)
	r.GET("/article/:limit/:offset", GetLimit)
	r.POST("/article", CreatePosts)
	r.PUT("/article/:id", UpdatePosts)
	r.DELETE("/article/:id", DeletePosts)
	r.Run(":8000")
}

//DeletePerson function for deleting person
func DeletePosts(c *gin.Context) {
	id := c.Params.ByName("id")
	var posts Posts
	d := db.Where("id = ?", id).Delete(&posts)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

//UpdatePerson function for updating person
func UpdatePosts(c *gin.Context) {
	var posts Posts
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&posts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	posts.Title = c.PostForm("title");
	posts.Content = c.PostForm("content");
	posts.Category = c.PostForm("category");
	posts.Status = c.PostForm("status");
	// c.BindJSON(&posts)
	db.Save(&posts)
	c.JSON(200, posts)
}

//CreatePerson function for creating person
func CreatePosts(c *gin.Context) {
	var posts Posts
	log.Println(posts, "ajsnaskjdnsakjndkjsandan")
	// posts.ID = c.PostForm(id);
	posts.Title = c.PostForm("title");
	posts.Content = c.PostForm("content");
	posts.Category = c.PostForm("category");
	posts.Status = c.PostForm("status");
	// c.BindJSON(&posts)
	db.Create(&posts)
	c.JSON(200, posts)
	fmt.Println(posts)
	
}

//GetPerson function for get one person
func GetPosts(c *gin.Context) {
	id := c.Params.ByName("id")
	var posts Posts
	if err := db.Where("id = ?", id).First(&posts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, posts)
	}
}

//GetPerson function for get one person
func GetLimit(c *gin.Context) {
	var posts []Posts
	if err := db.Limit(2).Offset(1).Find(&posts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, posts)
	}
}

//Getperson function for getting list
func Getposts(c *gin.Context) {
	var posts []Posts
	if err := db.Find(&posts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, posts)
	}
}