/*
mkdir custom
cd custom
go mod init github.com/weebnetsu/custom
touch math.go

Inside math.go we create a few functions

cd ..
mkdir calculate
cd calculate
go mod init github.com/weebnetsu/calculate
touch main.go

we use the custom package in here!
We can't do this inside this folder structure, but do it for the tutorial!

Below code will only work if above is followed!
*/

package main

import (
	"fmt"

	"github.com/weebnetsu/custom"
)

func main() {
	custom.SayHello("Nick")
	num := custom.Add(1, 8)
	fmt.Println(num)
}
