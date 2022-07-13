package Services

import (
	"github.com/Mayank-CES/golang_summer_bootcamp/Day04/eCommerce/Models"
	"github.com/Mayank-CES/golang_summer_bootcamp/Day04/eCommerce/Repository"
)

type CustomerServices interface {
	/*---Customer---*/

	CreatCustomerAccount(customer *Models.Customer) (err error)
	// BuyProduct(product *Models.Product)	(err error)
	CheckOrderByID(transaction *Models.Transaction, id string) (err error)

	/*---Product---*/
	GetAllProducts(products *[]Models.Product) (err error)
	GetProductByID(product *Models.Product, id string) (err error)

	/*---Retailer---*/
	CreatRetailerAccount(retailer *Models.Retailer) (err error)
	AddProduct(product *Models.Product) (err error)
	PatchProduct(updatedProduct *Models.PatchProd, id string) (err error)
	GetRHistoryByID(transaction *Models.Transaction, id string) (err error)
}

type CustomerService struct {
	cust Repository.Repo
}

func NewCustomerService(repo Repository.Repo) CustomerServices {
	return &CustomerService{
		cust: repo,
	}
}

func (c *CustomerService) CreatCustomerAccount(customer *Models.Customer) (err error) {
	return c.cust.CreatCustomerAccount(customer)
}

/*
func (c *CustomerService) BuyProduct(product *Models.Product,quantity)	(err error){
	return c.cust.BuyProduct(product)
}
*/

func (c *CustomerService) CheckOrderByID(transaction *Models.Transaction, id string) (err error) {
	return c.cust.CheckOrderByID(transaction, id)
}

func (c *CustomerService) GetAllProducts(products *[]Models.Product) (err error) {
	return c.cust.GetAllProducts(products)
}

func (c *CustomerService) GetProductByID(product *Models.Product, id string) (err error) {
	return c.cust.GetProductByID(product, id)
}

/*----Retailer----*/

func (c *CustomerService) CreatRetailerAccount(retailer *Models.Retailer) (err error) {
	return c.cust.CreatRetailerAccount(retailer)
}

func (c *CustomerService) AddProduct(product *Models.Product) (err error) {
	return c.cust.AddProduct(product)
}

func (c *CustomerService) PatchProduct(updatedProduct *Models.PatchProd, id string) (err error) {
	return c.cust.PatchProduct(updatedProduct, id)
}

func (c *CustomerService) GetRHistoryByID(transaction *Models.Transaction, id string) (err error) {
	return c.cust.GetRHistoryByID(transaction, id)
}
