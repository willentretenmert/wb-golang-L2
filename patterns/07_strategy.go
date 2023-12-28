package main

// Этот паттерн предоставляет возможность выбирать конкретное поведение (в данном случае способ оплаты)
// во время выполнения программы. стратегии могут быть изменены независимо от контекста, в котором они используются.

type Payment interface {
	Pay() error
}

type cardPayment struct {
	cardNumber, cvv string
}

func (p *cardPayment) Pay() error {
	// implementation
	return nil
}

type paypalPayment struct {
	account string
}

func (p *paypalPayment) Pay() error {
	// implementation
	return nil
}

type qiwiPayment struct {
	account string
}

func (q *qiwiPayment) Pay() error {
	// implementation
	return nil
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func NewPayPalPayment(account string) Payment {
	return &paypalPayment{account: account}
}

func NewQiwiPayment(account string) Payment {
	return &qiwiPayment{account: account}
}

func processOrder(orderNumber string, p Payment) {
	// implementation
	p.Pay()
}
