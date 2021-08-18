package internal

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

type App struct {
	Delivery *delivery.Delivery
}

func NewIOC(conf config.Config) di.IOC {
	ioc := di.IOC{}
	ioc[di.DI_CONFIG] = conf
	ioc[di.DI_REPOSITORY] = repository.NewRepository(ioc)
	ioc[di.DI_SERVICE] = service.NewService(ioc)
	ioc[di.DI_APP] = NewApp(ioc)
	return ioc
}

func NewApp(ioc di.IOC) *App {
	return &App{
		Delivery: delivery.NewDelivery(ioc),
	}
}
