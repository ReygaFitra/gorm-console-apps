package commandline

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ReygaFitra/gorm-console-apps/config/entity"
	"github.com/ReygaFitra/gorm-console-apps/usecase"
	"github.com/ReygaFitra/gorm-console-apps/utils"
)

type ProductDelivery interface {
	ProductMenu()
	AddProductCli()
}

type productDelivery struct {
	productUsecase usecase.ProductUsecase
}

func (d *productDelivery) ProductMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	menu := `
=====================
Products Menu:

1. Add Product
2. View Product
3. Edit Product
4. Remove Product
5. Back to Main Menu
6. Exit

Choose your menu: `

	fmt.Print(menu)
	scanner.Scan()
	userCommand := scanner.Text()
	switch userCommand {
	case "1":
		utils.ClearBash()
		d.AddProductCli()
	case "2":
		utils.ClearBash()

	case "3":
		utils.ClearBash()

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
		d.ProductMenu()
	}
}

func (d *productDelivery) AddProductCli() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input Product Code : ")
	scanner.Scan()
	productCode := scanner.Text()

	fmt.Printf("Input Product Name : ")
	scanner.Scan()
	productName := scanner.Text()

	fmt.Printf("Input Stock : ")
	scanner.Scan()
	productStock, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Invalid Format Input!")
	}

	fmt.Printf("Input Price : ")
	scanner.Scan()
	productPrice, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Invalid Format Input!")
	}

	products := entity.Products{
		ProductCode: productCode,
		ProductName: productName,
		Stock:       productStock,
		Price:       productPrice,
	}

	err = d.productUsecase.AddProduct(&products)
	if err != nil {
		log.Fatal("Failed Insert Product")
	}

	fmt.Println("Product added successfully!")
	d.ProductMenu()
}

func NewProductDelivery(productUsecase usecase.ProductUsecase) ProductDelivery {
	return &productDelivery{
		productUsecase: productUsecase,
	}
}
