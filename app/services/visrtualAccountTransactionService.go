package services

import (
	"errors"
	"mirae-va/app/helpers"
	"mirae-va/app/repositories"
	resources "mirae-va/app/resources"
	"time"
)

type VirtualAccountTransactionService struct {
	virtualAccountRepo repositories.VirtualAccountTransactionRepository
}

func NewVirtualAccountTransactionService() *VirtualAccountTransactionService {

	return &VirtualAccountTransactionService{
		virtualAccountRepo: repositories.NewVirtualAccountTransactionRepository(),
	}

}

func (h *VirtualAccountTransactionService) BillGenerate(dataResource resources.BillGenerate) (map[string]interface{}, error) {

	vaGenerated := helpers.GenerateVA()

	getCustomer, errGetCustomer := h.virtualAccountRepo.GetCustomer(dataResource.CustomerID)

	if errGetCustomer != nil {
		return nil, errGetCustomer
	}

	vaPayload := map[string]interface{}{
		"number":      vaGenerated,
		"customer_id": dataResource.CustomerID,
		"balance":     10000000,
	}

	_, errCreateVa := h.virtualAccountRepo.CreateVA(vaPayload)

	if errCreateVa != nil {
		return nil, errCreateVa
	}

	resGetLastVa, errGetLastVa := h.virtualAccountRepo.GetLastVaID()

	if errGetLastVa != nil {
		return nil, errGetLastVa
	}

	now := time.Now()

	dueDate := now.Add(3 * time.Hour)

	billPaylaod := map[string]interface{}{
		"virtual_account_id": resGetLastVa.ID,
		"amount":             dataResource.BillAmount,
		"due_date":           dueDate,
		"status":             "UNPAID",
	}

	errCreateBill := h.virtualAccountRepo.CreateBill(billPaylaod)

	if errCreateBill != nil {
		return nil, errCreateBill
	}

	responsePayload := map[string]interface{}{
		"customer_name":   getCustomer.Name,
		"amount":          dataResource.BillAmount,
		"virtual_account": vaGenerated,
	}

	return responsePayload, nil

}

func (h *VirtualAccountTransactionService) Inquiry(dataResource resources.Inquiry) (map[string]interface{}, error) {

	inquiry, errInquiry := h.virtualAccountRepo.Inquiry(dataResource)

	if errInquiry != nil {
		return nil, errors.New("transaction not found")
	}

	if inquiry.Bills.ID == 0 {
		return nil, errors.New("transaction not found")
	}

	responsePayload := map[string]interface{}{
		"customer_name": inquiry.Customer.Name,
		"amount":        inquiry.Bills.Amount,
	}

	return responsePayload, nil

}

func (h *VirtualAccountTransactionService) Payment(dataResource resources.Payment) (map[string]interface{}, error) {

	var promoPercentage float64

	promo, errPromo := h.virtualAccountRepo.GetPromo(dataResource.PromoCode)

	if errPromo != nil {
		promoPercentage = 0
	} else {
		promoPercentage = promo.DiscountPercent
	}

	getInquiry, errGetInquiry := h.virtualAccountRepo.GetInquiry(dataResource.VirtualAccountNumber)

	if errGetInquiry != nil {
		return nil, errors.New("transaction not found")
	}

	if getInquiry.Bills.Status == "PAID" {
		return nil, errors.New("your bill is already paid")
	}

	finalAmount := promoPercentage / 100 * float64(getInquiry.Bills.Amount)

	paymentPayload := map[string]interface{}{
		"virtual_account_id": getInquiry.Bills.VirtualAccountID,
		"bill_id":            getInquiry.Bills.ID,
		"amount":             finalAmount,
		"payment_date":       time.Now(),
	}

	_, errPayment := h.virtualAccountRepo.Payment(paymentPayload)

	if errPayment != nil {
		return nil, errors.New("transaction failed")
	}

	_, errUpdateBill := h.virtualAccountRepo.UpdateBillStatus(getInquiry.Bills.ID)

	if errUpdateBill != nil {
		return nil, errors.New("transaction failed")
	}

	responsePayload := map[string]interface{}{
		"customer_name": getInquiry.Customer.Name,
		"amount":        finalAmount,
	}

	return responsePayload, nil

}

func (h *VirtualAccountTransactionService) Report(dataResource resources.Report) (map[string]interface{}, error) {

	getInquiry, errGetInquiry := h.virtualAccountRepo.GetInquiry(dataResource.VirtualAccountNumber)

	if errGetInquiry != nil {
		return nil, errors.New("transaction not found")
	}

	if getInquiry.Bills.ID == 0 {
		return nil, errors.New("transaction not found")
	}

	report, errReport := h.virtualAccountRepo.Report(getInquiry.Bills.ID)

	if errReport != nil {
		return nil, errors.New("transaction not found")
	}

	responsePayload := map[string]interface{}{
		"bill_id":         report.ID,
		"virtual_account": getInquiry.VirtualAccount.Number,
		"amount":          report.Amount,
		"due_date":        report.DueDate,
		"status":          report.Status,
		"created_at":      report.CreatedAt,
		"updated_at":      report.UpdatedAt,
		"deleted_at":      report.CreatedAt,
	}

	return responsePayload, nil

}
