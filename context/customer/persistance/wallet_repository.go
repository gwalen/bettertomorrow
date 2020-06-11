package persistance

import (
	"bettertomorrow/common/dbxorm"
	"bettertomorrow/context/customer/domain"
	"sync"

	"xorm.io/xorm"
)

type WalletRepository interface {
	Insert(wallet *domain.Wallet) error
	InsertOrUpdate(wallet *domain.Wallet) error
	Delete(id uint) error
	FindAll() ([]domain.Wallet, error)
}

type WalletRepositoryImpl struct {
	db *xorm.Engine
}

/* -- singleton for DI -- */

var walletRepositoryInstance *WalletRepositoryImpl
var onceForWalletRepository sync.Once

func ProvideWalletRepositoryImpl() *WalletRepositoryImpl {
	onceForWalletRepository.Do(func() {
		dbHandle := dbxorm.DB()
		walletRepositoryInstance = &WalletRepositoryImpl{dbHandle}
	})
	return walletRepositoryInstance
}

/* ---- */

func (impl *WalletRepositoryImpl) Insert(wallet *domain.Wallet) error {
	_, err := impl.db.Insert(wallet)
	return err
}

//TODO: do it with plain query
func (impl *WalletRepositoryImpl) InsertOrUpdate(wallet *domain.Wallet) error {
	return nil
}

func (impl *WalletRepositoryImpl) Delete(id uint) error {
	_, err := impl.db.ID(id).Delete(&domain.Wallet{})
	return err
}

func (impl *WalletRepositoryImpl) FindAll() ([]domain.Wallet, error) {
	var wallets []domain.Wallet
	err := impl.db.Find(&wallets)
	return wallets, err
}
