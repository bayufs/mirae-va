package resources

type BillGenerate struct {
	CustomerID string `json:"customer_id" binding:"required"`
	BillAmount string `json:"bill_amount" binding:"required"`
}

type Inquiry struct {
	VirtualAccountNumber string `json:"virtual_account_number" binding:"required"`
}

type Payment struct {
	VirtualAccountNumber string `json:"virtual_account_number" binding:"required"`
	PromoCode            string `json:"promo_code"`
}

type Report struct {
	VirtualAccountNumber string `json:"virtual_account_number" binding:"required"`
}
