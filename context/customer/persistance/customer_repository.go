package persistance

import (
	"bettertomorrow/common/dbxorm"
	"bettertomorrow/context/customer/domain"
	"github.com/pkg/errors"
	"sync"

	"xorm.io/xorm"
)

type CustomerRepositoryImpl struct {
	db *xorm.Engine
}

type CustomerRepository interface {
	Insert(customer *domain.Customer) error
	Update(customer *domain.Customer) error
	InsertOrUpdate(customer *domain.Customer) error
	Delete(id uint) error
	FindAll() ([]domain.Customer, error)
	FindWithWallets() ([]domain.CustomerWithWallets, error)
}

/* -- singleton for DI -- */

var customerRepositoryInstance *CustomerRepositoryImpl
var onceForCustomerRepository sync.Once

func ProvideCustomerRepositoryImpl() *CustomerRepositoryImpl {
	onceForCustomerRepository.Do(func() {
		dbHandle := dbxorm.DB()
		customerRepositoryInstance = &CustomerRepositoryImpl{dbHandle}
	})
	return customerRepositoryInstance
}

/* ---- */

func (impl *CustomerRepositoryImpl) Insert(customer *domain.Customer) error {
	customer.Id = 0 // setting to 0 will trigger auto increment
	// _, err := impl.db.Insert(customer)
	err := errors.New("test error add customer")
	return err
}

//xorm does not have insert or update
//TODO: test what happens when you try to update non exsiting row 
func (impl *CustomerRepositoryImpl) Update(customer *domain.Customer) error {
	_, err := impl.db.ID(customer.Id).Update(customer)
	return err
}

//TODO: how to do it with plain query than?
func (impl *CustomerRepositoryImpl) InsertOrUpdate(customer *domain.Customer) error {
	return nil
}

func (impl *CustomerRepositoryImpl) Delete(id uint) error {
	_, err := impl.db.ID(id).Delete(&domain.Customer{})
	return err
}

func (impl *CustomerRepositoryImpl) FindAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	err := impl.db.Find(&customers)
	return customers, err
}

func (impl *CustomerRepositoryImpl) FindWithWallets() ([]domain.CustomerWithWallets, error) {
	var customerWithWalletArr []domain.CustomerWithWallet

	err := impl.db.Join("INNER", "wallets", "wallets.customer_id = customers.id").Find(&customerWithWalletArr)
	//TODO: test plain sql query

	customersWithWallets := mapJoin(customerWithWalletArr)

	return customersWithWallets, err
}

func mapJoin(customerWithWalletArr []domain.CustomerWithWallet) []domain.CustomerWithWallets {
	customersWithWalletsMap := make(map[domain.Customer][]domain.Wallet)
	var customersWithWallets []domain.CustomerWithWallets

	for _, elem := range customerWithWalletArr {
		// no need to check if key exists coz appending element to nil slice creates a new slice with this elemnent	=> append(nil, elem.Wallet) == []domain.Wallet{elem.Wallet}
		customersWithWalletsMap[elem.Customer] = append(customersWithWalletsMap[elem.Customer], elem.Wallet)
	}

	for customer, walletes := range customersWithWalletsMap {
		newcustomerWithWallets := domain.CustomerWithWallets{customer, walletes}
		customersWithWallets = append(customersWithWallets, newcustomerWithWallets)
	}

	return customersWithWallets
}
