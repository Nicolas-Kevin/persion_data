///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package s_log

import (
	"fmt"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *SLog {
	return new(SLog)
}

func NewQueryBuilder() *sLogQueryBuilder {
	return new(sLogQueryBuilder)
}

func (t *SLog) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type sLogQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *sLogQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *sLogQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&SLog{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *sLogQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&SLog{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *sLogQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&SLog{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *sLogQueryBuilder) First(db *gorm.DB) (*SLog, error) {
	ret := &SLog{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *sLogQueryBuilder) QueryOne(db *gorm.DB) (*SLog, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *sLogQueryBuilder) QueryAll(db *gorm.DB) ([]*SLog, error) {
	var ret []*SLog
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *sLogQueryBuilder) Limit(limit int) *sLogQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *sLogQueryBuilder) Offset(offset int) *sLogQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *sLogQueryBuilder) WhereId(p mysql.Predicate, value int32) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) WhereIdIn(value []int32) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) WhereIdNotIn(value []int32) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) OrderById(asc bool) *sLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *sLogQueryBuilder) WhereOperaUser(p mysql.Predicate, value int32) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "opera_user", p),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) WhereOperaUserIn(value []int32) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "opera_user", "IN"),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) WhereOperaUserNotIn(value []int32) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "opera_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) OrderByOperaUser(asc bool) *sLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "opera_user "+order)
	return qb
}

func (qb *sLogQueryBuilder) WhereOperaTime(p mysql.Predicate, value string) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "opera_time", p),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) WhereOperaTimeIn(value []string) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "opera_time", "IN"),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) WhereOperaTimeNotIn(value []string) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "opera_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) OrderByOperaTime(asc bool) *sLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "opera_time "+order)
	return qb
}

func (qb *sLogQueryBuilder) WhereOperaId(p mysql.Predicate, value string) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "opera_id", p),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) WhereOperaIdIn(value []string) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "opera_id", "IN"),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) WhereOperaIdNotIn(value []string) *sLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "opera_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sLogQueryBuilder) OrderByOperaId(asc bool) *sLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "opera_id "+order)
	return qb
}
