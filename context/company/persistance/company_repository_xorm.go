package persistance

import (
	"bettertomorrow/common/dbxorm"
	"bettertomorrow/context/company/domain"
	"fmt"
	"sync"

	"xorm.io/xorm"
)

type CompanyRepositoryImplXorm struct {
	db *xorm.Engine
}

/* -- singleton for DI -- */

var companyRepositoryInstanceXorm *CompanyRepositoryImplXorm
var onceForCompanyRepositoryXorm sync.Once

func ProvideCompanyRepositoryImplXorm() *CompanyRepositoryImplXorm {
	onceForCompanyRepositoryXorm.Do(func() {
		dbConnection := dbxorm.DB()
		companyRepositoryInstanceXorm = &CompanyRepositoryImplXorm{dbConnection}
	})
	return companyRepositoryInstanceXorm
}

/* ---- */

func (crImpl *CompanyRepositoryImplXorm) Insert(company *domain.Company) error {
	company.ID = 0 // setting to 0 will trigger auto increment
	_, err := crImpl.db.Insert(&domain.Company{})
	return err
}

//xorm does not have insert or update
func (crImpl *CompanyRepositoryImplXorm) InsertOrUpdate(company *domain.Company) error {
	_, err := crImpl.db.ID(company.ID).Update(&domain.Company{})
	return err
}

func (crImpl *CompanyRepositoryImplXorm) Delete(id uint) error {
	_, err := crImpl.db.ID(id).Delete(&domain.Company{})
	return err
}

func (crImpl *CompanyRepositoryImplXorm) FindAll() ([]domain.Company, error) {
	var companies []domain.Company
	err := crImpl.db.Find(&companies)
	return companies, err
}

//TODO: remove companyName param from method and from interface
func (crImpl *CompanyRepositoryImplXorm) FindWithProducts(companyName string) ([]domain.CompanyWithProducts, error) {
	var companyWithProductArr []domain.CompanyWithProduct

	err := crImpl.db.Join("INNER", "products", "products.company_id = companies.id").Find(&companyWithProductArr)
	//TODO: test plain sql query

	companiesWIthProducts := mapJoin(companyWithProductArr)
	fmt.Printf("AA %v\n", companyWithProductArr)
	fmt.Printf("BB %v\n", companiesWIthProducts)

	return companiesWIthProducts, err
}

func mapJoin(companyWithProductArr []domain.CompanyWithProduct) []domain.CompanyWithProducts {
	// var productsOfCompany map[domain.Company][]domain.Product --> can't have a struct as key where at least one memeber of a struct is not comaprable - slice is not comaprable so it will not compile
	emptyCompany := domain.Company{} // zero value struct which is returned from map when it does nt have a given key
	
	companies := make(map[uint]domain.Company)
	productsOfCompany := make(map[uint][]domain.Product)
	var companiesWithProducts []domain.CompanyWithProducts

	for _, elem := range companyWithProductArr {
		companyId := elem.Company.ID
		company := companies[companyId]
		if  emptyCompany.Equals(&company) {   // if we dont have company in the map add it
			company = elem.Company
			companies[companyId] = company 
		}
		productsOfCompany[companyId] = append(productsOfCompany[companyId], elem.Product)	
	}

	for _, elem := range companies {
		newcompanyWithProducts := domain.CompanyWithProducts{elem, productsOfCompany[elem.ID]}
		companiesWithProducts = append(companiesWithProducts, newcompanyWithProducts)
	}

	return companiesWithProducts
}
