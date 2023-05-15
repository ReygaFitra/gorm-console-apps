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
	ShowProductsCli()
	EditProductCli()
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
2. Show Product
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
		d.ShowProductsCli()
	case "3":
		utils.ClearBash()
		d.EditProductCli()
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

func (d *productDelivery) ShowProductsCli() {
	scanner := bufio.NewScanner(os.Stdin)
	menu := `
=====================
Show Products Menu:

1. Show All Products
2. Show Product by Product Code
3. Back to Product Menu
4. Exit

Choose your menu: `

	fmt.Print(menu)
	scanner.Scan()
	userCommand := scanner.Text()
	switch userCommand {
	case "1":
		utils.ClearBash()
		products, err := d.productUsecase.ShowAllProducts()
		if err != nil {
			log.Fatal("Data is Empty!")
		}
		for _, product := range products {
			fmt.Println(product.ProductCode, product.ProductName, product.Stock, product.Price)
		}
		d.ProductMenu()
	case "2":
		utils.ClearBash()
		fmt.Printf("Input Product Code : ")
		scanner.Scan()
		productCode := scanner.Text()
		product, err := d.productUsecase.ShowProductByPcode(productCode)
		if err != nil {
			log.Fatal("Data is Empty!")
		}
		fmt.Println(product)
		d.ProductMenu()
	case "3":
		utils.ClearBash()
		d.ProductMenu()
	case "4":
		utils.ClearBash()
		fmt.Println("Thank You")
		os.Exit(1)
	default:
		utils.ClearBash()
		d.ProductMenu()
	}
}

func (d *productDelivery) EditProductCli() {
	fmt.Println("Product yang tersedia: ")
	products, err := d.productUsecase.ShowAllProducts()
	if err != nil {
			log.Fatal("Data is Empty!")
		}
	for _, product := range products {
			fmt.Println(product.ProductCode, product.ProductName, product.Stock, product.Price)
		}

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

	oldProduct := entity.Products{
		ProductCode: productCode,
		ProductName: productName,
		Stock: productStock,
		Price: productPrice,
	}
	err = d.productUsecase.EditProduct(productCode, &oldProduct)
	if err != nil {
		log.Fatal("Failed Update Product!")
	}
	fmt.Println("Product Update Successfully")
	d.ProductMenu()
}

func NewProductDelivery(productUsecase usecase.ProductUsecase) ProductDelivery {
	return &productDelivery{
		productUsecase: productUsecase,
	}
}
