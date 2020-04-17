package persistance

import (
	"bettertomorrow/common/dbgorm"
	"bettertomorrow/context/customer/domain"
	"sync"

	"github.com/jinzhu/gorm"
)

type WalletRepository interface {
	Insert(wallet *domain.Wallet) error
	InsertOrUpdate(wallet *domain.Wallet) error
	Delete(id uint) error
	FindAll() ([]domain.Wallet, error)
}

type WalletRepositoryImpl struct {
	db *gorm.DB
}

/* -- singleton for DI -- */

var walletRepositoryInstance *WalletRepositoryImpl
var onceForWalletRepository sync.Once

func ProvideWalletRepositoryImpl() *WalletRepositoryImpl {
	onceForWalletRepository.Do(func() {
		dbHandle := dbgorm.DB()
		walletRepositoryInstance = &WalletRepositoryImpl{dbHandle}
	})
	return walletRepositoryInstance
}

/* ---- */

func (impl *WalletRepositoryImpl) Insert(wallet *domain.Wallet) error {
	return impl.db.Create(wallet).Error
}

func (impl *WalletRepositoryImpl) InsertOrUpdate(wallet *domain.Wallet) error {
	return impl.db.Save(wallet).Error
}

func (impl *WalletRepositoryImpl) Delete(id uint) error {
	return impl.db.Where("id = ?", id).Delete(domain.Wallet{}).Error
}

func (impl *WalletRepositoryImpl) FindAll() ([]domain.Wallet, error) {
	var wallets []domain.Wallet
	error := impl.db.Find(&wallets).Error
	return wallets, error
}
