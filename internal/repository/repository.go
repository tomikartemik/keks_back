package repository

import (
	"gorm.io/gorm"
	"tg_shop/internal/model"
)

type Repository struct {
	User
	Ad
	Category
	Invoice
	Payout
	Earning
	Premium
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:     NewUserRepository(db),
		Ad:       NewAdRepository(db),
		Category: NewCategoryRepository(db),
		Invoice:  NewInvoiceRepository(db),
		Payout:   NewPayoutRequestRepository(db),
		Earning:  NewEarningRepository(db),
		Premium:  NewPremiumRepository(db),
	}
}

type User interface {
	CreateUser(user model.User) (model.User, error)
	GetUserById(id int) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUserByUsername(username string) (model.User, error)
	SearchUsers(query string) ([]model.User, error)
	AddPurchase(userID, adID int) error
	ChangeBalance(userID int, newBlance float64) error
	ChangeHoldBalance(userID int, newBalance float64) error
	IncrementSalesAmount(userID int) error
}

type Ad interface {
	CreateAd(ad model.Ad) (model.Ad, error)
	GetAdListByCategory(categoryID int) ([]model.Ad, error)
	GetAllAds() ([]model.Ad, error)
	GetAdBySellerId(id int) (model.Ad, error)
	GetAdsByUserID(userID int) ([]model.Ad, error)
	GetAdById(id int) (model.Ad, error)
	UpdateAd(ad model.Ad) (model.Ad, error)
	DeleteAd(adID int) error // Добавляем метод удаления
	ChangeStock(adID, newStock int) error
	UpdateAdStatus(adID int, status string) error
	GetAdByIDTg(id int) (model.Ad, error)
	DisableExcessAds(userID int) error
	EnableAllDisabledAds(userID int) error
}

type Category interface {
	GetCategoryList() ([]model.Category, error)
	GetCategoryById(categoryID int) (model.Category, error)
}

type Invoice interface {
	CreateInvoice(TelegramID int, amount float64) (int, error)
	ChangeStatus(id int, status string) error
	GetInvoiceByID(id int) (model.Invoice, error)
}

type Payout interface {
	CreatePayoutRequest(telegramID int, amount float64) (int, error)
	UpdatePayoutStatus(requestID int, status string) error
	GetPayoutByID(telegramID int) (model.PayoutRequest, error)
}

type Earning interface {
	GetUnprocessedEarnings() ([]model.Earning, error)
	MarkAsProcessed(earning *model.Earning) error
	CreateEarning(newEarning model.Earning) error
	CountEarningsById(telegramID int) (int, error)
}

type Premium interface {
	GetExpiredPremiums() ([]model.User, []model.User, error)
	ResetPremiums(users []model.User) error
}
