package delivery

import (
	"gorm-basic/config"
	commandline "gorm-basic/delivery/commandLine"
	"gorm-basic/repository"
	"gorm-basic/usecase"
)

func RunServer() {
	db, err := config.LoadDatabase()
	if err != nil {
		panic(err)
	}

	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	productUsecase := usecase.NewProductUsecase(productRepo)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)

	productCli := commandline.NewProductDelivery(productUsecase)
	orderCli := commandline.NewOrderDelivery(orderUsecase)

	mainMenu := commandline.MainMenu{
		ProductDelivery: productCli,
		OrderDelivery:   orderCli,
	}
	mainMenu.DisplayMainMenu()
	// var mainMenu commandline.MainMenu
	// mainMenu.InitService()
	// mainMenu.DisplayMainMenu()
	
}