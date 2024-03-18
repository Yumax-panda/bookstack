package gormutil

import "gorm.io/gorm"

// RecordExists 指定した条件のレコードが1行以上存在するかどうか
func RecordExists(db *gorm.DB, where interface{}, tableName ...string) (exists bool, err error) {
	if len(tableName) > 0 {
		db = db.Table(tableName[0])
	} else {
		db = db.Model(where)
	}
	return Exists(db.Where(where))
}

// Exists 行数が1行以上かどうかを返す
func Exists(db *gorm.DB) (exists bool, err error) {
	n, err := Count(db.Limit(1))
	return n > 0, err
}

// Count 行数を数える
func Count(db *gorm.DB) (n int64, err error) {
	return n, db.Count(&n).Error
}

// LimitAndOffset limit句とoffset句を指定する。値が0以下の場合は指定されない。
func LimitAndOffset(limit, offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if offset > 0 {
			db = db.Offset(offset)
		}
		if limit > 0 {
			db = db.Limit(limit)
		}
		return db
	}
}
