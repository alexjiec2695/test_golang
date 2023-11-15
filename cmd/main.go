package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"runtime"
	"strings"
	"test/internal/di"
)

func main() {
	environment()
	di.Start()
}

func environment() {
	path, _ := os.Getwd()

	var operativeSystem = runtime.GOOS
	switch operativeSystem {
	case "windows":
		path = strings.ReplaceAll(path, `\`, "/")
		break
	}

	godotenv.Load(fmt.Sprintf("%v/cmd/%v", path, ".env"))
}
