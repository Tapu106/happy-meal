package main

import (
	"context"
	"os"

	"github.com/Tapu106/happy-meal/internal/routes"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda
var server *gin.Engine

func init() {
	server = gin.Default()
	routes.RegisterRoutes(server)
	ginLambda = ginadapter.New(server)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	if os.Getenv("IS_OFFLINE") == "true" || os.Getenv("LOCAL") == "true" {
		server.Run(":3000")
	} else {
		lambda.Start(Handler)
	}
}
