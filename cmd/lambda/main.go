package main

import "github.com/aws/aws-lambda-go/lambda"
import "github.com/fconnect/storylens-banglapaper/internal/lib"

func main() {
	lambda.Start(lib.RequestHandler)
}
