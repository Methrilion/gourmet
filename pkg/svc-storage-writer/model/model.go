package model

// Currency struct
type Currency struct {
	ID   uint32
	Name string
	Code string
}

// TableName set Currency's table name to be `currency`
func (Currency) TableName() string {
	return "currency"
}

// RateOfExchange struct
type RateOfExchange struct {
	ID     uint32
	FromID uint32
	ToID   uint32
	Price  float32
}

// Location struct
type Location struct {
	ID          uint32
	Name        string
	Description string
	CurrencyID  uint32
	Status      bool
}

// Product struct
type Product struct {
	ID          uint32
	Name        string
	Description string
}

// Price struct
type Price struct {
	ID         uint32
	ProductID  uint32
	LocationID uint32
	Price      float32
}

// Position struct
type Position struct {
	ID          uint32
	Name        string
	Description string
}

// Employee struct
type Employee struct {
	ID         uint32
	FirstName  string
	LastName   string
	PositionID uint32
	LocationID uint32
}

// Method struct
type Method struct {
	ID   uint32
	Name string
}

// Receipt struct
// TODO: change datetime type
type Receipt struct {
	ID         uint32
	EmployeeID uint32
	LocationID uint32
	MethodID   uint32
	Datetime   string
}

// Purchase struct
type Purchase struct {
	ID        uint32
	ReceiptID uint32
	PriceID   uint32
	Amount    float32
	Price     float32
	Result    float32
}
