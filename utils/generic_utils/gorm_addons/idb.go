package gorm_addons

import "gorm.io/gorm"

// IDB
// GORM is very hard to test, this interface is useless for testing GORM codes.
type IDB interface {
	Model(value interface{}) *gorm.DB
	Distinct(args ...interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Group(name string) *gorm.DB
	Having(query interface{}, args ...interface{}) *gorm.DB
	Order(value interface{}) *gorm.DB
	Limit(limit int) *gorm.DB
	Offset(offset int) *gorm.DB
	Preload(query string, args ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB
	Update(column string, value interface{}) *gorm.DB
	UpdateColumn(column string, value interface{}) *gorm.DB
	UpdateColumns(values interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Count(count *int64) *gorm.DB
}
