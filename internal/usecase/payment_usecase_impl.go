package usecase

import (
	"payment-options/internal/models"
	"payment-options/internal/repository"
	"sync"
)

type paymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(r repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo: r}
}

func (u *paymentUsecase) GetPaymentOptions() (map[string]models.PaymentMethod, error) {
		result := make(map[string]models.PaymentMethod)
		result["linkaja"] = u.repo.CallLINKAja()
		result["jeniuspay"] = u.repo.CallJeniusPay()
		result["bluepay"] =u.repo.CallBluePay()
		result["octo"] = u.repo.CallOCTO()
		result["qris"] = u.repo.CallQRIS()
		result["bcaklikpay"] = u.repo.CallBCAKlikPay()
		return result, nil
}

func (u *paymentUsecase) GetPaymentOptions2() (map[string]models.PaymentMethod, error) { //jika tanpa goroutine, bedakan nama pemanggilan get nya
	var wg sync.WaitGroup
	result := make(map[string]models.PaymentMethod)
	mu := sync.Mutex{}

	wg.Add(6)

	go func() {
		defer wg.Done()
		mu.Lock()
		result["linkaja"] = u.repo.CallLINKAja()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["jeniuspay"] = u.repo.CallJeniusPay()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["bluepay"] = u.repo.CallBluePay()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["octo"] =  u.repo.CallOCTO()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["qris"] = u.repo.CallQRIS()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["bcaklikpay"] = u.repo.CallBCAKlikPay()
		mu.Unlock()
	}()

	wg.Wait()
	return result, nil
}