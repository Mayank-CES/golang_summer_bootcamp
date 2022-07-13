package Repository

import (
	"github.com/Mayank-CES/golang_summer_bootcamp/Day04/eCommerce/Config"
	"github.com/Mayank-CES/golang_summer_bootcamp/Day04/eCommerce/Models"
)

type Repo interface {
	CreatCustomerAccount(customer *Models.Customer) (err error)
	CheckOrderByID(transaction *Models.Transaction, id string) (err error)

	GetAllProducts(products *[]Models.Product) (err error)
	GetProductByID(product *Models.Product, id string) (err error)

	CreatRetailerAccount(retailer *Models.Retailer) (err error)
	AddProduct(product *Models.Product) (err error)
	PatchProduct(updatedProduct *Models.PatchProd, id string) (err error)
	GetRHistoryByID(order *Models.Transaction, id string) (err error)
	// IsPresent(id string, Student *Models.Student) (*Models.Student, error)
	// DoCreate(Student *Models.Student)
	// DoFind(Student *[]Models.Student)
	// DoUpdate(Student *Models.Student, newStudent Models.UpdatedStudent)
	// DoDelete(Student *Models.Student) error
}

type repository struct{}

func NewRepo() Repo {
	return &repository{}
}

var err error

// CreatCustomerAccount  ... Insert New Customer data
func (r *repository) CreatCustomerAccount(customer *Models.Customer) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

// CheckOrderByID ... Fetch only one Order by Id
func (r *repository) CheckOrderByID(transaction *Models.Transaction, id string) (err error) {
	if err = Config.DB.Where("transaction_id = ?", id).First(transaction).Error; err != nil {
		return err
	}
	return nil
}

// GetAllProducts -----  Fetch all product data
func (r *repository) GetAllProducts(products *[]Models.Product) (err error) {
	if err = Config.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}

// GetProductByID ... Fetch only one product by Id
func (r *repository) GetProductByID(product *Models.Product, id string) (err error) {
	if err = Config.DB.Where("product_id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

// CreatRetailerAccount  ... Insert New Retailer data
func (r *repository) CreatRetailerAccount(retailer *Models.Retailer) (err error) {
	if err = Config.DB.Create(retailer).Error; err != nil {
		return err
	}
	return nil
}

// AddProduct ... Insert New Retailer data
func (r *repository) AddProduct(product *Models.Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// PatchProduct ... Insert New Retailer data
func (r *repository) PatchProduct(updatedProduct *Models.PatchProd, id string) (err error) {
	if err = Config.DB.Model(&Models.Product{}).Where("product_id = ?", id).Update("product_quantity", updatedProduct.Quantity).Error; err != nil {
		return err
	}
	return nil
}

// GetRHistoryByID ... Insert New Retailer data
func (r *repository) GetRHistoryByID(order *Models.Transaction, id string) (err error) {
	if err = Config.DB.Where("retailer_id = ?", id).Find(order).Error; err != nil {
		return err
	}
	return nil
}

/*

// UpdateInfo ... Update student
func UpdateInfo(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

//DeleteStudent ... Delete student
func DeleteStudent(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}


//GetAllMarks ... Fetch all Marks
func GetAllMarks(marks *[]Marks) (err error) {
	if err = Config.DB.Find(marks).Error; err != nil {
		return err
	}
	return nil
}

//GetMarksByID ... Fetch only one user by Id
func GetMarksByID(marks *Marks, id string) (err error) {
	if err = Config.DB.Where("student_id = ?", id).Find(marks).Error; err != nil {
		return err
	}
	return nil
}

//UpdateMarks ... Update marks
func UpdateMarks(marks *Marks, id string) (err error) {
	fmt.Println(marks)
	Config.DB.Save(marks)
	return nil
}

//DeleteMarks ... Delete marks
func DeleteMarks(marks *Marks, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(marks)
	return nil
}

//DeleteAllMarks ... Delete marks
func DeleteAllMarks(marks *[]Marks, id string) (err error) {
	Config.DB.Where("student_id = ?", id).Delete(marks)
	return nil
}

*/
