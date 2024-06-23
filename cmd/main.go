package main

import (
	"fmt"

	"github.com/ProovGroup/worker-claim-declaration/pkg"
	// "github.com/aws/aws-lambda-go/lambda"
)

var dummyData = pkg.Payload{
	ProovCode: "GG4QXE",
	Params: pkg.DeclarationParams{
		Api: "VEOS",
	},
}

func main() {
	// lambda.Start(pkg.Handler)
	fmt.Println("Start")
	pkg.Handler(dummyData)
	fmt.Println("End")
}
