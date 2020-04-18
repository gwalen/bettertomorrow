package persistance

import (
	"bettertomorrow/common/dbsqlx"
	"bettertomorrow/context/customer/domain"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryImplSqlx struct {
	db *sqlx.DB
}

/* -- singleton for DI -- */

var customerRepositoryInstanceSqlx *CustomerRepositoryImplSqlx
var onceForCustomerRepositorySqlx sync.Once

func ProvideCustomerRepositoryImplSqlx() *CustomerRepositoryImplSqlx {
	onceForCustomerRepositorySqlx.Do(func() {
		dbHandle := dbsqlx.DB()
		customerRepositoryInstanceSqlx = &CustomerRepositoryImplSqlx{dbHandle}
	})
	return customerRepositoryInstanceSqlx
}

/* ---- */

func (impl *CustomerRepositoryImplSqlx) Insert(customer *domain.Customer) error {
	customer.Id = 0 // setting to 0 will trigger auto increment
	queryStr := "INSERT INTO customers(id, first_name, last_name, created_at) VALUES (0, ?, ?, ?);"
	impl.db.MustExec(queryStr, customer.FirstName, customer.LastName, time.Now) //MustExec will panic on error (Exec will not)
	return nil // no error returned it will panic on error
}

//this will only work on mysql (its hwo you do insertOnUpdate in Mysql with plan query)
func (impl *CustomerRepositoryImplSqlx) InsertOrUpdate(customer *domain.Customer) error {
	queryStr := "INSERT INTO customers(id, first_name, last_name, created_at) VALUES (0, ?, ?, ?) ON DUPLICATE KEY UPDATE first_name = ?, last_name = ?"
	impl.db.MustExec(queryStr, customer.FirstName, customer.LastName, customer.FirstName, customer.LastName)
	return nil // no error returned it will panic on error
}

func (impl *CustomerRepositoryImplSqlx) Update(customer *domain.Customer) error {
	queryStr := "UPDATE customers SET first_name = ?, last_name = ? where id = ?"
	impl.db.MustExec(queryStr, customer.FirstName, customer.LastName)
	return nil
}

func (impl *CustomerRepositoryImplSqlx) Delete(id uint) error {
	queryStr := "DELETE FROM customers WHERE id = ?"
	impl.db.MustExec(queryStr, id)
	return nil
}

//Select featches all rows to memory and populats the slice, for large data sets when paging is needed Queryx
func (impl *CustomerRepositoryImplSqlx) FindAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	err := impl.db.Select(&customers, "SELECT customers.* FROM customers")
	return customers, err
}


//TODO: add to notes super annoying (lot of boiler plate) when doing join and you must alias each field from joined tables
func (impl *CustomerRepositoryImplSqlx) FindWithWallets() ([]domain.CustomerWithWallets, error) {
	var customerWithWalletArr []domain.CustomerWithWallet
	
	query := `
		SELECT 
			customers.id "customers.id",
			customers.first_name "customers.first_name",
			customers.last_name "customers.last_name",
			customers.created_at "customers.created_at",
			wallets.id "wallets.id",
			wallets.currency "wallets.currency",
			wallets.amount "wallets.amount",
			wallets.customer_id "wallets.customer_id",
			wallets.created_at "wallets.created_at"
		FROM customers JOIN wallets ON wallets.customer_id = customers.id;
	`
	rows, err := impl.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		var cs domain.CustomerWithWallet
		err :=  rows.StructScan(&cs)
		if err != nil {
			return nil, err //most porbably some extral logging would be necessary here
		}
		customerWithWalletArr = append(customerWithWalletArr, cs)
	}

	customersWithWallets := mapJoinSqlx(customerWithWalletArr)
	return customersWithWallets, err
}

func mapJoinSqlx(customerWithWalletArr []domain.CustomerWithWallet) []domain.CustomerWithWallets {
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
