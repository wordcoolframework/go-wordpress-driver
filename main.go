package main

import (
	"fmt"
	wordpressdriver "github.com/wordcoolframework/go-wordpress-driver/wordpress_driver"
	"log"
)

func main() {
	err := wordpressdriver.DBConnection()
	if err != nil {
		log.Fatal("Failed Connection!")
	}

	wp := wordpressdriver.WP

	posts, _ := wp.Post().GetPosts("wp")
	fmt.Println(posts)
}
