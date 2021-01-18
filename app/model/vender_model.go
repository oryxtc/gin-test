package model

func (VenderMain) TableName() string {
	return "vender_main"
}

type VenderMain struct {
	Model
	Vid               int64
	FinCode           string
	Name              string
	ContractCode      string
	ContractDate      int
	Type              int
	IsPrepayer        int8
	IsPreStandbyer    int
	InvoiceOfferType  int8
	InvoiceOfferDate  int8
	SettleType        int8
	SettleDays        int
	SettlePeriod      int8
	ZoneId            string
	CompanyId         string
	Payee             string
	BankName          string
	BankAccount       string
	BankUnitCode      string
	Status            int
	ServiceChargeRate float64
	VenderType        int8
	OwnerType         int8
}
