![Go WordPress Driver Logo](logo_80.webp)
# Go WordPress Driver

If you don't have a complex business and don't want to develop a custom backend, you can use WordPress as your CMS.
This package allows you to directly connect to the WordPress database using Golang and build your custom API.

## 🎯 Use Cases
With this package, you can:
- Use WordPress for content management, e-commerce, or startup introduction.
- Connect directly to the database instead of relying on the WordPress REST API.
- Manage posts, categories, users, and other data.
- Build a custom API for WordPress using Golang.

## 🚀Installation & Setup
To install this package, run:
```sh
go get github.com/wordcoolframework/go-wordpress-driver

how use : 

package main

import (
    "fmt"
    "github.com/wordcoolframework/go-wordpress-driver/wordpressdriver"
)

func main() {
    posts, err := wordpressdriver.GetPosts("wp")
    if err != nil {
        fmt.Println("خطا در دریافت پست‌ها:", err)
        return
    }

    for _, post := range posts {
        fmt.Println("عنوان:", post.Title)
    }
}

categories, err := wordpressdriver.GetCategories("wp")
if err != nil {
    log.Fatal(err)
}

for _, category := range categories {
    fmt.Println("نام دسته‌بندی:", category.Name)
}