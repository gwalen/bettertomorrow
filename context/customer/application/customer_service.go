package application

import (
	"bettertomorrow/context/customer/domain"
	"bettertomorrow/context/customer/persistance"
	"sync"
)

type CustomerService interface {
	CreateCustomer(customer *domain.Customer) error
}

type CustomerServiceImpl struct {
	customerRepository persistance.CustomerRepository
}

/* -- singleton for DI -- */

var customerServiceInstance *CustomerServiceImpl
var onceForCompantService sync.Once

func ProvideCustomerServiceImpl() *CustomerServiceImpl {
	onceForCompantService.Do(func() {
		customerRepository := persistance.ProvideCustomerRepositoryImpl()
		customerServiceInstance = &CustomerServiceImpl{customerRepository}
	})
	return customerServiceInstance
}

/* ---- */

func (impl *CustomerServiceImpl) CreateCustomer(customer *domain.Customer) error {
	return impl.customerRepository.Insert(customer)
}

func (impl *CustomerServiceImpl) CreateOrUpdateCustomer(customer *domain.Customer) error {
	return impl.customerRepository.InsertOrUpdate(customer)
}

func (impl *CustomerServiceImpl) UpdateCustomer(customer *domain.Customer) error {
	return impl.customerRepository.Update(customer)
}

func (impl *CustomerServiceImpl) DeleteCustomer(id uint) error {
	return impl.customerRepository.Delete(id)
}

func (impl *CustomerServiceImpl) FindAllCustomers() ([]domain.Customer, error) {
	customers, err := impl.customerRepository.FindAll()
	return customers, err
}
