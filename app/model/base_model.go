package model

import "gorm.io/gorm"

type Model struct {
	Id        int64
	CreatedAt int
	UpdateAt  int
	VenderId  int64
}

/**
获取round_rows总条数
*/
func GetFoundRows(db *gorm.DB, total *int) error {
	err := db.Raw("select FOUND_ROWS();").Find(total).Error
	return err
}
