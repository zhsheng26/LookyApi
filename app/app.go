package app

import (
	"api.welooky/conf"
	"fmt"
	"google.golang.org/grpc"
)

type App struct {
}

func Initialize(conf conf.Setting) {
	_, err := grpc.Dial("", grpc.WithInsecure())
	fmt.Println(err)
}
