package repository

import (
	"payment-options/internal/models"
	// "time"
)

type paymentRepo struct{}

func NewPaymentRepo() PaymentRepository {
	return &paymentRepo{}
}

// CallLINKAja mengembalikan informasi metode pembayaran LINKAja.
func (r *paymentRepo) CallLINKAja() models.PaymentMethod {
	// time.Sleep(5 * time.Second) // Simulate network delay
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "12500",
		Icon:    "https://sampleurl.com/linkaja.jpg",
	}
}

func (r *paymentRepo) CallJeniusPay() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "15000",
		Icon:    "https://sampleurl.com/jenius.jpg",
	}
}

func (r *paymentRepo) CallBluePay() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "8000",
		Icon:    "https://sampleurl.com/bluepay.jpg",
	}
}

func (r *paymentRepo) CallOCTO() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "20000",
		Icon:    "https://sampleurl.com/octo.jpg",
	}
}

func (r *paymentRepo) CallQRIS() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "ID10xx",
		Status:  "Active",
		Balance: "0",
		Icon:    "https://sampleurl.com/qris.jpg",
	}
}

func (r *paymentRepo) CallBCAKlikPay() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "25000",
		Icon:    "https://sampleurl.com/bcaklikpay.jpg",
	}
}