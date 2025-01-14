// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Train Travel API
 *
 * API for finding and booking train trips across Europe.  ## Run in Postman  Experiment with this API in Postman, using our Postman Collection.  [![Run In Postman](https://run.pstmn.io/button.svg =128pxx32px)](https://app.getpostman.com/run-collection/9265903-7a75a0d0-b108-4436-ba54-c6139698dc08?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D9265903-7a75a0d0-b108-4436-ba54-c6139698dc08%26entityType%3Dcollection%26workspaceId%3Df507f69d-9564-419c-89a2-cb8e4c8c7b8f)
 *
 * API version: 1.2.0
 * Contact: support@example.com
 */

package oas


// BookingPaymentSource - The payment source to take the payment from. This can be a card or a bank account. Some of these properties will be hidden on read to protect PII leaking.
type BookingPaymentSource struct {
	Object string `json:"object,omitempty"`

	Name string `json:"name"`

	// The account number for the bank account, in string form. Must be a current account.
	Number string `json:"number"`

	// Card security code, 3 or 4 digits usually found on the back of the card.
	Cvc string `json:"cvc"`

	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`

	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`

	AddressLine1 string `json:"address_line1,omitempty"`

	AddressLine2 string `json:"address_line2,omitempty"`

	AddressCity string `json:"address_city,omitempty"`

	AddressCountry string `json:"address_country"`

	AddressPostCode string `json:"address_post_code,omitempty"`

	// The sort code for the bank account, in string form. Must be a six-digit number.
	SortCode string `json:"sort_code,omitempty"`

	// The type of entity that holds the account. This can be either `individual` or `company`.
	AccountType string `json:"account_type"`

	// The name of the bank associated with the routing number.
	BankName string `json:"bank_name"`

	// Two-letter country code (ISO 3166-1 alpha-2).
	Country string `json:"country"`
}

// AssertBookingPaymentSourceRequired checks if the required fields are not zero-ed
func AssertBookingPaymentSourceRequired(obj BookingPaymentSource) error {
	elements := map[string]interface{}{
		"name":            obj.Name,
		"number":          obj.Number,
		"cvc":             obj.Cvc,
		"exp_month":       obj.ExpMonth,
		"exp_year":        obj.ExpYear,
		"address_country": obj.AddressCountry,
		"account_type":    obj.AccountType,
		"bank_name":       obj.BankName,
		"country":         obj.Country,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertBookingPaymentSourceConstraints checks if the values respects the defined constraints
func AssertBookingPaymentSourceConstraints(obj BookingPaymentSource) error {
	return nil
}
