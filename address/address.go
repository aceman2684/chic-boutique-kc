package address

// Address represents an address in the United States
type Address struct {
	Id      int     `json:"id"`      // Id represents the unique identifier of the address
	LineOne string  `json:"lineOne"` // LineOne represents the first line of the address
	LineTwo *string `json:"lineTwo"` // LineTwo represents the second line of the address
	City    string  `json:"city"`    // City represents the city the address resides in
	State   string  `json:"state"`   // State represents the state the address resides in
	ZipCode string  `json:"zipCode"` // ZipCode represents the ZIP code the address resides in
}
