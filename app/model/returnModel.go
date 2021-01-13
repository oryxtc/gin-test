package model

type ReturnHead struct {
	Model
	ReturnId      string
	AllocateId    int64
	Vid           int64
	ZoneId        string
	ItAmount      float64
	NtAmount      float64
	SettleId      int64
	CompanyId     string
	GenType       int
	Status        int8
	Date          int
	OrderDate     int
	IsWriteOff    int
	WarehouseCode string
}
