package model

type ReceiptHead struct {
	Model
	ReceiptId     string
	AllocateId    int64
	OrderId       int64
	OrderName     string
	Vid           int64
	ZoneId        string
	Reference     string
	NtAmount      float64
	ItAmount      float64
	SettleId      int64
	CompanyId     string
	Type          int
	GenType       int
	EnterDate     int
	OrderDate     int
	Status        int8
	IsWriteOff    int
	WarehouseCode string
}
