package structs

type SkuType int
type EntitlementType int

const (
	SkuTypeApplication = 1
	SkuTypeDLC         = 2
	SkuTypeConsumable  = 3
	SkuTypeBundle      = 4

	EntitlementTypePurchase            = 1
	EntitlementTypePremiumSubscription = 2
	EntitlementTypeDeveloerGift        = 3
	EntitlementTypeTestModePurchase    = 4
	EntitlementTypeFreePurchase        = 5
	EntitlementTypeUserGift            = 6
	EntitlementTypePremiumPurchase     = 7
)

type SKU struct {
	// Id the unique ID of the SKU
	Id int64 `json:"Id"`
	// Type what sort of SKU it is
	Type SkuType `json:"Type"`
	// Name the name of the SKU
	Name string `json:"Name"`
	// Price the price of the SKU
	Price SkuPrice `json:"Price"`
}

type SkuPrice struct {
	// Amount the amount of money the SKU costs
	Amount uint32 `json:"Amount"`
	// Currency the currency the amount is in
	Currency string `json:"Currency"`
}

type Entitlement struct {
	// Id the unique ID of the entitlement
	Id int64 `json:"Id"`
	// Type the kind of entitlement it is
	Type EntitlementType `json:"Type"`
	// SkuId the ID of the SKU to which the user is entitled
	SkuId int64 `json:"SkuId"`
}

type LimitedPaymentData struct {
	// Id unique ID of the payment
	Id string `json:"id"`
	// Currency the currency the payment was made in
	Currency string `json:"currency"`
	// Amount the amount paid
	Amount int `json:"amount"`
	// Tax the amount of tax
	Tax int `json:"tax"`
	// TaxInclusive whether the amount is tax-inclusive
	TaxInclusive bool `json:"tax_inclusive"`
}
