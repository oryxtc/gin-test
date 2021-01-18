package service

import (
	"strings"
	"vss/app"
)

//校验请求数据
type GetReceiptAndReturnInput struct {
	SettleId string `form:"settle_id,omitempty" json:"settle_id,omitempty" binding:"required" `
	FinCode  string `form:"fin_code,omitempty"`
	Vid      string `json:"vid,omitempty"`
	Pn       int    `form:"pn,omitempty" json:",omitempty"`
	Rn       int    `form:"rn,omitempty"`
}

/**
获取收货和退货数据
*/
func GetReceiptAndReturnData(p GetReceiptAndReturnInput) ([]map[string]interface{}, int) {
	var where []string
	var whereSql string
	binds := make(map[string]interface{}, 0)
	binds["venderId"] = 10005
	//var binds []interface{}
	db := app.GetDb()
	if p.SettleId != "" {
		where = append(where, "settle_id=@settleId")
		binds["settleId"] = p.SettleId
	}
	if p.FinCode != "" {
		where = append(where, "fin_code=@finCode")
		binds["finCode"] = p.FinCode
	}
	if len(where) > 0 {
		whereSql += " WHERE " + strings.Join(where, " AND ")
	}
	sql := "SELECT SQL_CALC_FOUND_ROWS * FROM (" +
		"SELECT 1 order_type,b.vid,b.fin_code,b.`name` AS vender_name,b.type AS vender_type,b.settle_type,a.venderId,a.zone_id,a.receipt_id AS rr_num,a.allocate_id,a.order_date,a.order_id,a.order_name,a.reference,a.type,a.enter_date AS rr_date,a.created_at,a.it_amount,a.nt_amount,a.`status`,a.settle_id,a.is_write_off,a.warehouse_code FROM receipt_head AS a LEFT JOIN vender_main AS b ON a.vid=b.vid AND a.venderId=b.venderId WHERE a.venderId=@venderId " +
		"UNION ALL " +
		"SELECT 2 order_type,b.vid,b.fin_code, b.name AS vender_name,b.type AS vendr_type,b.settle_type,a.venderId,a.zone_id,a.return_id AS order_num,a.allocate_id,a.order_date,0 order_id,'' order_name,'' reference,0 type,a.date AS rr_date,a.created_at,a.it_amount,a.nt_amount,a.`status`,a.settle_id,a.is_write_off,a.warehouse_code FROM return_head AS a LEFT JOIN vender_main AS b ON a.vid=b.vid AND a.venderId=b.venderId WHERE a.venderId=@venderId " +
		") a"
	sql += whereSql
	//拼接limit
	if p.Rn > 0 {
		sql += " Limit @rn "
		binds["rn"] = p.Rn
	}
	if p.Pn > 0 {
		sql += " Offset @pn "
		binds["pn"] = p.Pn
	}
	//获取数据
	list := make([]map[string]interface{}, 0)
	if err := db.Raw(sql, binds).Find(&list).Error; err != nil {
		panic(err.Error())
	}
	//获取总条数
	var total int
	if err := db.Raw("select FOUND_ROWS();", binds).Find(&total).Error; err != nil {
		panic(err)
	}
	return list, total
}
