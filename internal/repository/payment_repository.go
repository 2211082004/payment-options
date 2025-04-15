package repository

import "payment-options/internal/models"

type PaymentRepository interface {
	CallLINKAja() models.PaymentMethod
	CallJeniusPay() models.PaymentMethod
	CallBluePay() models.PaymentMethod
	CallOCTO() models.PaymentMethod
	CallQRIS() models.PaymentMethod
	CallBCAKlikPay() models.PaymentMethod
}
