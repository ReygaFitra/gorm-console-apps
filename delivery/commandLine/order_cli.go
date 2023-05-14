package commandline

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ReygaFitra/gorm-console-apps/config/entity"
	"github.com/ReygaFitra/gorm-console-apps/usecase"
	"github.com/ReygaFitra/gorm-console-apps/utils"
)

type OrderDelivery interface {
	OrderMenu()
}

type ordeDelivery struct {
	orderUsecase usecase.OrderUsecase
}

func(d *ordeDelivery) OrderMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	menu := `
=====================
Orders Menu:

1. Add Order
2. Add Order Detail
3. Show All Order
4. Edit Order
5. Back to Main Menu
6. Exit

Choose your menu: `

	fmt.Print(menu)
	scanner.Scan()
	userCommand := scanner.Text()
	switch userCommand {
	case "1":
		utils.ClearBash()
		d.AddOrderCli()
	case "2":
		utils.ClearBash()
		d.AddOrderDetailCli()
	case "3":
		utils.ClearBash()
		orders, err := d.orderUsecase.ShowAllOrders()
		if err != nil {
			log.Fatal("Data is Empty")
		}
		for _, order := range orders {
			fmt.Println(order.IdOrderDetail, order.OrdersIdOrder, order.ProductsProductCode, order.Qty)
		}
	case "4":
		utils.ClearBash()
		
	case "5":
		utils.ClearBash()
	case "6":
		utils.ClearBash()
		fmt.Println("Thank You")
		os.Exit(1)
	default:
		utils.ClearBash()
		d.OrderMenu()
	}
}

func (d *ordeDelivery) AddOrderCli() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input Order Year : ")
	scanner.Scan()
	yearCli, _ := strconv.Atoi(scanner.Text())

	fmt.Printf("Input Order Month : ")
	scanner.Scan()
	monthCli, _ := strconv.Atoi(scanner.Text())

	fmt.Printf("Input Order Date : ")
	scanner.Scan()
	dateCli, _ := strconv.Atoi(scanner.Text())

	fmt.Printf("Input Payment Methode : ")
	scanner.Scan()
	paymentMethod := scanner.Text()

	fmt.Printf("Input Order Total : ")
	scanner.Scan()
	total, _ := strconv.Atoi(scanner.Text())

	order := entity.Orders{
		OrderDate:     time.Date(int(yearCli), time.Month(monthCli), int(dateCli), 0, 0, 0, 0, time.Local),
		PaymentMethod: paymentMethod,
		OrderTotal:    total,
	}

	newOrder, err := d.orderUsecase.AddOrder(&order)
	if err != nil {
		log.Fatal("Failed Insert Product")
	}

	fmt.Println("Order added successfully!")
	fmt.Println("Order ID: ", newOrder.IdOrders)
	d.OrderMenu()
}

func (d *ordeDelivery) AddOrderDetailCli() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input ID Order : ")
	scanner.Scan()
	idOrder, _ := strconv.Atoi(scanner.Text())


	fmt.Printf("Input Product Code : ")
	scanner.Scan()
	productCode := scanner.Text()

	fmt.Printf("Input Quantity : ")
	scanner.Scan()
	quantity, _ := strconv.Atoi(scanner.Text())
	
	detail := []entity.OrderDetails{
		{
			Orders: entity.Orders{
				IdOrders: idOrder,
			},
			Products: entity.Products{
				ProductCode: productCode,
			},
			Qty: quantity,
		},
	} 
	err := d.orderUsecase.AddOrderDetails(&detail)
	if err != nil {
		log.Fatal("Failed Insert Product")
	}

	fmt.Println("Product added successfully!")
	d.OrderMenu()
}

func NewOrderDelivery(orderUsecase usecase.OrderUsecase) OrderDelivery {
	return &ordeDelivery{
		orderUsecase: orderUsecase,
	}
}