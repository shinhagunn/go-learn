package bad

// SOLID
// O - Open-Closed Principle
// Objects or entities should be open for extension but closed for modification

// Good
type Payment struct {
	partnerCode   string
	paymentMethod string
	paymentNow    PaymentNowInterface
}

type PaymentInterface interface {
	PaymentNow() bool
}

func (p *Payment) PaymentNow() bool {
	return p.paymentNow.GetNow()
}

func NewPayment(partnerCode string) PaymentInterface {
	switch partnerCode {
	case "EX_1":
		{
			return &Payment{
				partnerCode: partnerCode,
			}
		}
	case "EX_2":
		{
			return &Payment{
				partnerCode: partnerCode,
			}
		}

	default:
		panic("partnerCode not valid")
	}

}

/////////////////////////////////////////////////////////////////////////////

type PaymentNowInterface interface {
	GetNow() bool
}

type PaypalPayNow struct {
	paymentMethod string
}

func (pp *PaypalPayNow) GetNow() bool {
	// todo action paynow paypal
	return true
}

type VisaPayNow struct {
	paymentMethod string
}

func (pp *VisaPayNow) GetNow() bool {
	// todo action paynow paypal
	return true
}

type MasterCardPayNow struct {
	paymentMethod string
}

func (pp *MasterCardPayNow) GetNow() bool {
	// todo action paynow paypal
	return true
}
