package main

import (
	"fmt"

	_ "github.com/containers/ocicrypt"
	"github.com/google/uuid"
	_ "github.com/kardianos/service"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(uuid.New().String())
}
