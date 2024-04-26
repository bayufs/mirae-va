package repositories

import (
	"fmt"
	"mirae-va/app/models"
	"mirae-va/app/resources"

	"mirae-va/config"

	"gorm.io/gorm"
)

type InquiryResult struct {
	VirtualAccount models.VirtualAccount
	Customer       models.Customer
	Bills          models.Bill
}

type PaymentResult struct {
	VirtualAccount models.VirtualAccount
	Customer       models.Customer
	Bills          models.Bill
}

type virtualAccountTransactionDBConnection struct {
	connection *gorm.DB
}

type VirtualAccountTransactionRepository interface {
	GetCustomer(customerId string) (*models.Customer, error)
	GetLastVaID() (*models.VirtualAccount, error)
	CreateVA(va map[string]interface{}) (*models.VirtualAccount, error)
	CreateBill(billing map[string]interface{}) error
	Inquiry(dataResource resources.Inquiry) (*InquiryResult, error)
	GetPromo(promoCode string) (*models.PromoCode, error)
	GetInquiry(virtualAccountNumber string) (*PaymentResult, error)
	Payment(payment map[string]interface{}) (*models.Payment, error)
	UpdateBillStatus(billID int) (*models.Bill, error)
	Report(billID int) (*models.Bill, error)
}

func NewVirtualAccountTransactionRepository() VirtualAccountTransactionRepository {
	return &virtualAccountTransactionDBConnection{
		connection: config.ConnectDB(),
	}
}

func (db *virtualAccountTransactionDBConnection) CreateVA(va map[string]interface{}) (*models.VirtualAccount, error) {

	var virtualAccount models.VirtualAccount

	if err := db.connection.Debug().Model(&virtualAccount).Create(&va).Error; err != nil {
		return nil, err
	}

	return &virtualAccount, nil
}

func (db *virtualAccountTransactionDBConnection) CreateBill(billing map[string]interface{}) error {

	var bill models.Bill

	if err := db.connection.Debug().Model(&bill).Create(&billing).Error; err != nil {
		return err
	}

	return nil
}

func (db *virtualAccountTransactionDBConnection) GetLastVaID() (*models.VirtualAccount, error) {
	var virtualAccount models.VirtualAccount

	if err := db.connection.Debug().Model(&models.VirtualAccount{}).Last(&virtualAccount).Error; err != nil {
		return nil, err
	}

	return &virtualAccount, nil
}

func (db *virtualAccountTransactionDBConnection) GetCustomer(customerId string) (*models.Customer, error) {

	var Customer models.Customer

	if err := db.connection.Debug().Model(&Customer).Where("id", customerId).Find(&Customer).Error; err != nil {
		return nil, err
	}

	return &Customer, nil
}

func (db *virtualAccountTransactionDBConnection) Inquiry(dataResource resources.Inquiry) (*InquiryResult, error) {
	var inquiryResult InquiryResult

	if err := db.connection.Debug().
		Model(&inquiryResult.VirtualAccount).
		Where("virtual_accounts.number = ?", dataResource.VirtualAccountNumber).
		Joins("JOIN bills ON virtual_accounts.id = bills.virtual_account_id").
		Joins("JOIN customers ON virtual_accounts.customer_id = customers.id").
		Find(&inquiryResult.VirtualAccount).Error; err != nil {
		return nil, err
	}

	if err := db.connection.Debug().
		Model(&inquiryResult.Customer).
		Where("id = ?", inquiryResult.VirtualAccount.CustomerID).
		Find(&inquiryResult.Customer).Error; err != nil {
		return nil, err
	}

	if err := db.connection.Debug().
		Where("virtual_account_id = ?", inquiryResult.VirtualAccount.ID).
		Find(&inquiryResult.Bills).Error; err != nil {
		return nil, err
	}

	return &inquiryResult, nil
}

func (db *virtualAccountTransactionDBConnection) GetPromo(promoCode string) (*models.PromoCode, error) {
	var promoCodeModel models.PromoCode

	if err := db.connection.Debug().Model(&promoCodeModel).Where("code", promoCode).Find(&promoCodeModel).Error; err != nil {
		return nil, err
	}

	if promoCodeModel.ID == 0 {
		return nil, fmt.Errorf("promo code not found")
	}

	return &promoCodeModel, nil
}

func (db *virtualAccountTransactionDBConnection) GetInquiry(virtualAccountNumber string) (*PaymentResult, error) {
	var paymentResult PaymentResult

	if err := db.connection.Debug().
		Model(&paymentResult.VirtualAccount).
		Where("virtual_accounts.number = ?", virtualAccountNumber).
		Joins("JOIN bills ON virtual_accounts.id = bills.virtual_account_id").
		Joins("JOIN customers ON virtual_accounts.customer_id = customers.id").
		Find(&paymentResult.VirtualAccount).Error; err != nil {
		return nil, err
	}

	if err := db.connection.Debug().
		Model(&paymentResult.Customer).
		Where("id = ?", paymentResult.VirtualAccount.CustomerID).
		Find(&paymentResult.Customer).Error; err != nil {
		return nil, err
	}

	if err := db.connection.Debug().
		Where("virtual_account_id = ?", paymentResult.VirtualAccount.ID).
		Find(&paymentResult.Bills).Error; err != nil {
		return nil, err
	}

	return &paymentResult, nil
}

func (db *virtualAccountTransactionDBConnection) Payment(paymentPayload map[string]interface{}) (*models.Payment, error) {

	var payment models.Payment

	if err := db.connection.Debug().Model(&payment).Create(&paymentPayload).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

func (db *virtualAccountTransactionDBConnection) UpdateBillStatus(billID int) (*models.Bill, error) {

	var bill models.Bill

	if err := db.connection.Debug().Model(&bill).Where("id", billID).Update("status", "PAID").Error; err != nil {
		return nil, err
	}

	return &bill, nil
}

func (db *virtualAccountTransactionDBConnection) Report(billID int) (*models.Bill, error) {

	var bill models.Bill

	if err := db.connection.Debug().Model(&bill).Where("id", billID).Find(&bill).Error; err != nil {
		return nil, err
	}

	return &bill, nil
}
