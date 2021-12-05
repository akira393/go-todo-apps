package main

import (
	"fmt"

	"udemy.com/golang-webgosql/app/controllers"
	"udemy.com/golang-webgosql/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()

}
