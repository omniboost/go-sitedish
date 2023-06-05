package sitedish

import "time"

type Orders []Order

type Order struct {
	ID            string     `json:"id"`
	OrderNumber   string     `json:"orderNumber"`
	Platform      string     `json:"platform"`
	RestaurantID  string     `json:"restaurantId"`
	ReceiptNumber int        `json:"receiptNumber"`
	OrderStatus   int        `json:"orderStatus"`
	OrderDate     time.Time  `json:"orderDate"`
	OrderType     string     `json:"orderType"`
	Courier       string     `json:"courier"`
	DeliveryCosts float64    `json:"deliveryCosts"`
	Tip           int        `json:"tip"`
	TotalPrice    float64    `json:"totalPrice"`
	TotalFee      float64    `json:"totalFee"`
	TotalDiscount float64    `json:"totalDiscount"`
	TotalDeposit  float64    `json:"totalDeposit"`
	PaymentMethod string     `json:"paymentMethod"`
	PaysExact     bool       `json:"paysExact"`
	PaysWith      float64    `json:"paysWith"`
	Customer      Customer   `json:"customer"`
	Products      Products   `json:"products"`
	Discounts     Discounts  `json:"discounts"`
	Remark        string     `json:"remark"`
	Location      Location   `json:"location"`
	Canceled      bool       `json:"canceled"`
	Deleted       bool       `json:"deleted"`
	Refunded      bool       `json:"refunded"`
	IsPaid        bool       `json:"isPaid"`
	RequestedAsap bool       `json:"requestedAsap"`
	ConfirmedTime DateTime   `json:"confirmedTime"`
	Version       string     `json:"version"`
	Accounting    Accounting `json:"accounting"`
	Meta          Meta       `json:"meta"`
}

type Customer struct {
	Name             string `json:"name"`
	CompanyName      string `json:"companyName"`
	PhoneNumber      string `json:"phoneNumber"`
	Street           string `json:"street"`
	StreetNumber     string `json:"streetNumber"`
	PostalCode       string `json:"postalCode"`
	City             string `json:"city"`
	ExtraAddressInfo string `json:"extraAddressInfo"`
	NewCustomer      bool   `json:"newCustomer"`
}

type Products []Product

type Product struct {
	ID           string  `json:"id"`
	Posid        string  `json:"posid"`
	Number       string  `json:"number"`
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	Instructions string  `json:"instructions"`
	Count        int     `json:"count"`
	UnitPrice    float64 `json:"unitPrice"`
	Price        float64 `json:"price"`
	UnitDiscount float64 `json:"unitDiscount"`
	Subtotal     float64 `json:"subtotal"`
	VAT          float64 `json:"vat"`
	Deposit      float64 `json:"deposit"`
	SideDishes   []struct {
		ID      string      `json:"id"`
		Posid   string      `json:"posid"`
		Number  interface{} `json:"number"`
		Name    string      `json:"name"`
		Count   int         `json:"count"`
		Price   float64     `json:"price"`
		VAT     float64     `json:"vat"`
		Deposit int         `json:"deposit"`
	} `json:"sideDishes"`
}

type Discounts []Discount

type Discount struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

type Location struct {
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Distance int     `json:"distance"`
	Duration int     `json:"duration"`
}

type Accounting struct {
	Revenue float64 `json:"revenue"`
	VAT     struct {
		VatLowPercentage  float64 `json:"vat_low_percentage"`
		VatLowAmount      float64 `json:"vat_low_amount"`
		VatLow            float64 `json:"vat_low"`
		VatHighPercentage float64 `json:"vat_high_percentage"`
		VatHighAmount     float64 `json:"vat_high_amount"`
		VatHigh           float64 `json:"vat_high"`
		VatNoneAmount     float64 `json:"vat_none_amount"`
	} `json:"vat"`
}

type Meta struct {
	Version     int         `json:"version"`
	LastUpdated interface{} `json:"lastUpdated"`
}
