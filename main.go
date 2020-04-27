package main

import "github.com/codingXiang/gecko/cmd"

//go:generate go run main.go general model -s ./example -f user.go -d ./output/model
//go:generate go run main.go general repo -s ./output/model -f user.go -d ./output/module -p user
//go:generate go run main.go general svc -s ./output/model -f user.go -d ./output/module -p user
//go:generate go run main.go general delivery http -s ./output/module/user -f service.go -d ./output/module -p user
func main() {
	cmd.Execute()
}
