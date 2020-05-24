// package test_application_customer_wallet_service
package application

import (
	"bettertomorrow/common/logger"
	"bettertomorrow/context/customer/domain"
	"bettertomorrow/context/customer/persistance"
	"github.com/google/go-cmp/cmp"
	// "fmt"
	"testing"
	"time"
)

func init() {
	// regular logger is initalised using config file which path when run from test is not reachable, so we need to override it
	logCws = logger.ProvideDefaultLogger()
}

/*** Stubs ***/

type CustomerRepositoryStub struct {
	// persistance.CustomerRepository
	InsertMock          func(customer *domain.Customer) error
	UpdateMock          func(customer *domain.Customer) error
	InsertOrUpdateMock  func(customer *domain.Customer) error
	DeleteMock          func(id uint) error
	FindAllMock         func() ([]domain.Customer, error)
	FindWithWalletsMock func() ([]domain.CustomerWithWallets, error)
}

func (cr *CustomerRepositoryStub) Insert(customer *domain.Customer) error {
	return cr.InsertMock(customer)
}

func (cr *CustomerRepositoryStub) Update(customer *domain.Customer) error {
	return cr.UpdateMock(customer)
}

func (cr *CustomerRepositoryStub) InsertOrUpdate(customer *domain.Customer) error {
	return cr.InsertOrUpdateMock(customer)
}

func (cr *CustomerRepositoryStub) Delete(id uint) error {
	return cr.DeleteMock(id)
}

func (cr *CustomerRepositoryStub) FindAll() ([]domain.Customer, error) {
	return cr.FindAllMock()
}

func (cr *CustomerRepositoryStub) FindWithWallets() ([]domain.CustomerWithWallets, error) {
	return cr.FindWithWalletsMock()
}

// we dont use any methods from WalletRepository so empty stub is enough
type WalletRepositoryStub struct {
	persistance.WalletRepository
}

/*** Tests ***/

func TestCustomerWalletsServiceAggregateCustomerSavings(t *testing.T) {
	customersRepositoryStub := CustomerRepositoryStub{
		FindWithWalletsMock: func() ([]domain.CustomerWithWallets, error) {
			return createCustomerWithWalletsSmallSet(), nil
		},
	}

	walletRepositoryStub := WalletRepositoryStub{}

	customerWalletsService := &CustomerWalletsServiceImpl{&customersRepositoryStub, &walletRepositoryStub}

	result, err := customerWalletsService.AggregateCustomerSavings()
	expectedResult := createAggregatedWalletResultForSmallSet()

	if err != nil {
		t.Errorf("Failed to aggreggate customer savings, err: %v\n", err)
	} else {
		if !cmp.Equal(result, expectedResult) {
			t.Errorf("Wring result, got : %v, expected: %v \n", result, expectedResult)
		} else {
			t.Logf("Correct result for AggregateWallets small set")
		}
	}
}

func createCustomerWithWalletsSmallSet() []domain.CustomerWithWallets {
	return []domain.CustomerWithWallets{
		domain.CustomerWithWallets{
			Customer: domain.Customer{101, "john", "doe", time.Unix(0, 0)},
			Wallet: []domain.Wallet{
				domain.Wallet{1, 100, "usd", time.Unix(0, 0), 101},
				domain.Wallet{2, 100, "eur", time.Unix(0, 0), 101},
				domain.Wallet{3, 100, "idr", time.Unix(0, 0), 101},
				domain.Wallet{4, 100, "aud", time.Unix(0, 0), 101},
			},
		},
		domain.CustomerWithWallets{
			Customer: domain.Customer{202, "alice", "wanderland", time.Now()},
			Wallet: []domain.Wallet{
				domain.Wallet{1, 200, "pln", time.Unix(0, 0), 202},
				domain.Wallet{2, 300, "nzd", time.Unix(0, 0), 202},
				domain.Wallet{3, 100, "usd", time.Unix(0, 0), 202},
				domain.Wallet{4, 50, "chf", time.Unix(0, 0), 202},
			},
		},
	}
}

func createAggregatedWalletResultForSmallSet() []domain.AggregatedWallet {
	return []domain.AggregatedWallet{
		domain.AggregatedWallet{101, []string{"idr", "eur", "aud", "usd"}, 400},
		domain.AggregatedWallet{202, []string{"usd", "pln", "nzd", "chf"}, 650},
	}
}
