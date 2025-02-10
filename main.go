package main

import (
	"fmt"
	wordpressdriver "go-wordpress-driver/wordpress_driver"
	"log"
)

func main() {

	err := wordpressdriver.DBConnection()

	if err != nil {
		log.Fatal("Faild Connection!")
	}

	post := wordpressdriver.Post{}

	category := wordpressdriver.Category{}

	GetById, err := post.GetPostByID(22, "wp")

	GetAllCategory, err := category.GetAllCategories("wp")

	if err != nil {
		log.Fatal("error", err)
	}

	//fmt.Println(GetById.ID)

	fmt.Println(GetAllCategory)
}
