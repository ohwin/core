package core

import (
	"github.com/ohwin/core/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type DBModel struct {
	gorm.Model
}

func (m DBModel) Id() uint {
	return m.Model.ID
}

type ModelI interface {
	Id() uint
}

type ParamI interface {
	Where() func(db *gorm.DB) *gorm.DB
	Order() clause.OrderByColumn
}
type BaseRepo struct {
}

func (repo *BaseRepo) Create(model ModelI) (uint, error) {
	if err := global.DB.Create(model).Error; err != nil {
		return 0, err
	}
	return model.Id(), nil
}

func (repo *BaseRepo) Del(model ModelI, param ParamI) error {
	if err := global.DB.Delete(model).Scopes(param.Where()).Error; err != nil {
		return err
	}
	return nil
}

func (repo *BaseRepo) Update(model ModelI) error {
	if err := global.DB.Updates(model).Error; err != nil {
		return err
	}
	return nil
}

func (repo *BaseRepo) Get(model ModelI, param ParamI) (interface{}, error) {
	err := global.DB.Model(model).Scopes(param.Where()).First(model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return model, nil
}

func (repo *BaseRepo) GetList(model ModelI, res interface{}, param ParamI, page, pageSize int) (int64, interface{}, error) {
	var total int64

	db := global.DB.Model(model).Scopes(param.Where()).Count(&total)
	if db.Error != nil {
		return 0, nil, db.Error
	}
	if err := db.Scopes(Paginate(page, pageSize)).Order(param.Order()).Find(res).Error; err != nil {
		return 0, nil, err
	}
	return total, res, nil
}

// Paginate 分页查询
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// TimeInterval 给定时间范围查询
func TimeInterval(times ...int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch len(times) {
		case 1:
			db = db.Where("created_at >= ?", time.Unix(times[0], 0))
		case 2:
			db = db.Where("created_at BETWEEN ? AND ?", time.Unix(times[0], 0), time.Unix(times[1], 0))
		}
		return db
	}
}

// NumInterval 给定数字范围查询
func NumInterval(db *gorm.DB, column string, nums ...int64) *gorm.DB {
	switch len(nums) {
	case 1:
		db = db.Where("? >= ?", column, nums[0])
	case 2:
		db = db.Where("? BETWEEN ? AND ?", column, nums[0], nums[1])
	}
	return db
}

// Id 根据ID查询
func Id(id uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

// Custom 自定义单个条件查询
func Custom(column string, value any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("? = ?", column, value)
	}
}
