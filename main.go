package main

import (
	"github.com/kindai-csg/D-Chat/infrastructure"
)

func main() {
	infrastructure.Router.Run(":3000")
}
