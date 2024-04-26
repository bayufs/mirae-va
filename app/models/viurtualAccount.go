package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID        int       `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email" json:"email"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type VirtualAccount struct {
	gorm.Model
	ID         int       `gorm:"column:id;primaryKey" json:"id"`
	Number     string    `gorm:"column:number;" json:"number"`
	CustomerID int       `gorm:"column:customer_id" json:"customer_id"`
	Balance    float64   `gorm:"column:balance" json:"balance"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type Bill struct {
	gorm.Model
	ID               int       `gorm:"column:id;primaryKey" json:"id"`
	VirtualAccountID int       `gorm:"column:virtual_account_id" json:"virtual_account_id"`
	Amount           float64   `gorm:"column:amount" json:"amount"`
	DueDate          time.Time `gorm:"column:due_date" json:"due_date"`
	Status           string    `gorm:"column:status;" json:"status"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type Payment struct {
	gorm.Model
	ID               int       `gorm:"column:id;primaryKey" json:"id"`
	VirtualAccountID int       `gorm:"column:virtual_account_id" json:"virtual_account_id"`
	BillID           int       `gorm:"column:bill_id" json:"bill_id"`
	Amount           float64   `gorm:"column:amount" json:"amount"`
	PaymentDate      time.Time `gorm:"column:payment_date" json:"payment_date"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type PromoCode struct {
	gorm.Model
	ID              int       `gorm:"column:id;primaryKey" json:"id"`
	Code            string    `gorm:"column:code;unique" json:"code"`
	DiscountPercent float64   `gorm:"column:discount_percent" json:"discount_percent"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
