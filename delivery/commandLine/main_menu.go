package commandline

import (
	"bufio"
	"fmt"
	"gorm-basic/repository"
	"gorm-basic/usecase"
	"gorm-basic/utils"
	"os"

	"gorm.io/gorm"
)

type MainMenu struct {
	ProductDelivery ProductDelivery
	OrderDelivery OrderDelivery
}

func (m *MainMenu) DisplayMainMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	menu := `
======================
Main Menu:

1. Product Menu
2. Order Menu
3. Exit

Menu: `

	var isExit = true
	for isExit {
	fmt.Print(menu)
	scanner.Scan()
	switch scanner.Text() {
	case "1":
		utils.ClearBash()
		m.ProductDelivery.ProductMenu()
	case "2":
		utils.ClearBash()
		m.OrderDelivery.OrderMenu()
	case "3":
		isExit = false
		utils.ClearBash()
		fmt.Println("Thank You")
	default:
		utils.ClearBash()
	}	
	}
}

func (m *MainMenu) InitService() {
	productUsecase := usecase.NewProductUsecase(
		repository.NewProductRepository(&gorm.DB{}),
	)
	orderUsecase := usecase.NewOrderUsecase(
		repository.NewOrderRepository(&gorm.DB{}),
	)
	m.ProductDelivery = NewProductDelivery(productUsecase)
	m.OrderDelivery = NewOrderDelivery(orderUsecase)
}