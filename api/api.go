package api

import (
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest"
)

func PrepareAll() {
	apiRest := rest.NewAPIRest()
	ioc.Set(func() rest.APIRest {
		return *apiRest
	})

	apiMessaging := messaging.NewAPIMessaging()
	ioc.Set(func() messaging.APIMessaging {
		return *apiMessaging
	})

	apiGrpc := grpc.NewAPIGrpc()
	ioc.Set(func() grpc.APIGrpc {
		return *apiGrpc
	})
}
