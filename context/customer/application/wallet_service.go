package application

import (
	"bettertomorrow/context/customer/domain"
	"bettertomorrow/context/customer/persistance"
	"sync"
)

type WalletService interface {
	CreateWallet(customer *domain.Wallet) error
	UpdateWallet(customer *domain.Wallet) error
	DeleteWallet(id uint) error
	FindAllWallets() ([]domain.Wallet, error)
}

type WalletServiceImpl struct {
	walletRepository persistance.WalletRepository
}

/* -- singleton for DI -- */

var walletServiceInstance *WalletServiceImpl
var onceForWalletService sync.Once

func ProvideWalletServiceImpl() *WalletServiceImpl {
	onceForWalletService.Do(func() {
		var walletRepository persistance.WalletRepository
		walletRepositoryImpl := persistance.ProvideWalletRepositoryImpl()
		walletRepository = walletRepositoryImpl
		walletServiceInstance = &WalletServiceImpl{walletRepository}
	})
	return walletServiceInstance
}

/* ---- */

func (impl *WalletServiceImpl) CreateWallet(customer *domain.Wallet) error {
	return impl.walletRepository.Insert(customer)
}

func (impl *WalletServiceImpl) UpdateWallet(customer *domain.Wallet) error {
	return impl.walletRepository.InsertOrUpdate(customer)
}

func (impl *WalletServiceImpl) DeleteWallet(id uint) error {
	return impl.walletRepository.Delete(id)
}

func (impl *WalletServiceImpl) FindAllWallets() ([]domain.Wallet, error) {
	wallets, err := impl.walletRepository.FindAll()
	return wallets, err
}
