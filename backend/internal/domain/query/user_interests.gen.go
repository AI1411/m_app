// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/AI1411/m_app/internal/domain/model"
)

func newUserInterest(db *gorm.DB, opts ...gen.DOOption) userInterest {
	_userInterest := userInterest{}

	_userInterest.userInterestDo.UseDB(db, opts...)
	_userInterest.userInterestDo.UseModel(&model.UserInterest{})

	tableName := _userInterest.userInterestDo.TableName()
	_userInterest.ALL = field.NewAsterisk(tableName)
	_userInterest.UserID = field.NewString(tableName, "user_id")
	_userInterest.InterestID = field.NewInt32(tableName, "interest_id")
	_userInterest.CreatedAt = field.NewTime(tableName, "created_at")
	_userInterest.User = userInterestBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "model.User"),
	}

	_userInterest.Interest = userInterestBelongsToInterest{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Interest", "model.Interest"),
	}

	_userInterest.fillFieldMap()

	return _userInterest
}

type userInterest struct {
	userInterestDo

	ALL        field.Asterisk
	UserID     field.String // ユーザーID
	InterestID field.Int32  // 趣味ID
	CreatedAt  field.Time   // 関連付け作成日時
	User       userInterestBelongsToUser

	Interest userInterestBelongsToInterest

	fieldMap map[string]field.Expr
}

func (u userInterest) Table(newTableName string) *userInterest {
	u.userInterestDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userInterest) As(alias string) *userInterest {
	u.userInterestDo.DO = *(u.userInterestDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userInterest) updateTableName(table string) *userInterest {
	u.ALL = field.NewAsterisk(table)
	u.UserID = field.NewString(table, "user_id")
	u.InterestID = field.NewInt32(table, "interest_id")
	u.CreatedAt = field.NewTime(table, "created_at")

	u.fillFieldMap()

	return u
}

func (u *userInterest) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userInterest) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 5)
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["interest_id"] = u.InterestID
	u.fieldMap["created_at"] = u.CreatedAt

}

func (u userInterest) clone(db *gorm.DB) userInterest {
	u.userInterestDo.ReplaceConnPool(db.Statement.ConnPool)
	u.User.db = db.Session(&gorm.Session{Initialized: true})
	u.User.db.Statement.ConnPool = db.Statement.ConnPool
	u.Interest.db = db.Session(&gorm.Session{Initialized: true})
	u.Interest.db.Statement.ConnPool = db.Statement.ConnPool
	return u
}

func (u userInterest) replaceDB(db *gorm.DB) userInterest {
	u.userInterestDo.ReplaceDB(db)
	u.User.db = db.Session(&gorm.Session{})
	u.Interest.db = db.Session(&gorm.Session{})
	return u
}

type userInterestBelongsToUser struct {
	db *gorm.DB

	field.RelationField
}

func (a userInterestBelongsToUser) Where(conds ...field.Expr) *userInterestBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userInterestBelongsToUser) WithContext(ctx context.Context) *userInterestBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userInterestBelongsToUser) Session(session *gorm.Session) *userInterestBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a userInterestBelongsToUser) Model(m *model.UserInterest) *userInterestBelongsToUserTx {
	return &userInterestBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

func (a userInterestBelongsToUser) Unscoped() *userInterestBelongsToUser {
	a.db = a.db.Unscoped()
	return &a
}

type userInterestBelongsToUserTx struct{ tx *gorm.Association }

func (a userInterestBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a userInterestBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userInterestBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userInterestBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userInterestBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a userInterestBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

func (a userInterestBelongsToUserTx) Unscoped() *userInterestBelongsToUserTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type userInterestBelongsToInterest struct {
	db *gorm.DB

	field.RelationField
}

func (a userInterestBelongsToInterest) Where(conds ...field.Expr) *userInterestBelongsToInterest {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userInterestBelongsToInterest) WithContext(ctx context.Context) *userInterestBelongsToInterest {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userInterestBelongsToInterest) Session(session *gorm.Session) *userInterestBelongsToInterest {
	a.db = a.db.Session(session)
	return &a
}

func (a userInterestBelongsToInterest) Model(m *model.UserInterest) *userInterestBelongsToInterestTx {
	return &userInterestBelongsToInterestTx{a.db.Model(m).Association(a.Name())}
}

func (a userInterestBelongsToInterest) Unscoped() *userInterestBelongsToInterest {
	a.db = a.db.Unscoped()
	return &a
}

type userInterestBelongsToInterestTx struct{ tx *gorm.Association }

func (a userInterestBelongsToInterestTx) Find() (result *model.Interest, err error) {
	return result, a.tx.Find(&result)
}

func (a userInterestBelongsToInterestTx) Append(values ...*model.Interest) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userInterestBelongsToInterestTx) Replace(values ...*model.Interest) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userInterestBelongsToInterestTx) Delete(values ...*model.Interest) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userInterestBelongsToInterestTx) Clear() error {
	return a.tx.Clear()
}

func (a userInterestBelongsToInterestTx) Count() int64 {
	return a.tx.Count()
}

func (a userInterestBelongsToInterestTx) Unscoped() *userInterestBelongsToInterestTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type userInterestDo struct{ gen.DO }

type IUserInterestDo interface {
	gen.SubQuery
	Debug() IUserInterestDo
	WithContext(ctx context.Context) IUserInterestDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserInterestDo
	WriteDB() IUserInterestDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserInterestDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserInterestDo
	Not(conds ...gen.Condition) IUserInterestDo
	Or(conds ...gen.Condition) IUserInterestDo
	Select(conds ...field.Expr) IUserInterestDo
	Where(conds ...gen.Condition) IUserInterestDo
	Order(conds ...field.Expr) IUserInterestDo
	Distinct(cols ...field.Expr) IUserInterestDo
	Omit(cols ...field.Expr) IUserInterestDo
	Join(table schema.Tabler, on ...field.Expr) IUserInterestDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserInterestDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserInterestDo
	Group(cols ...field.Expr) IUserInterestDo
	Having(conds ...gen.Condition) IUserInterestDo
	Limit(limit int) IUserInterestDo
	Offset(offset int) IUserInterestDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserInterestDo
	Unscoped() IUserInterestDo
	Create(values ...*model.UserInterest) error
	CreateInBatches(values []*model.UserInterest, batchSize int) error
	Save(values ...*model.UserInterest) error
	First() (*model.UserInterest, error)
	Take() (*model.UserInterest, error)
	Last() (*model.UserInterest, error)
	Find() ([]*model.UserInterest, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserInterest, err error)
	FindInBatches(result *[]*model.UserInterest, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserInterest) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserInterestDo
	Assign(attrs ...field.AssignExpr) IUserInterestDo
	Joins(fields ...field.RelationField) IUserInterestDo
	Preload(fields ...field.RelationField) IUserInterestDo
	FirstOrInit() (*model.UserInterest, error)
	FirstOrCreate() (*model.UserInterest, error)
	FindByPage(offset int, limit int) (result []*model.UserInterest, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserInterestDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userInterestDo) Debug() IUserInterestDo {
	return u.withDO(u.DO.Debug())
}

func (u userInterestDo) WithContext(ctx context.Context) IUserInterestDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userInterestDo) ReadDB() IUserInterestDo {
	return u.Clauses(dbresolver.Read)
}

func (u userInterestDo) WriteDB() IUserInterestDo {
	return u.Clauses(dbresolver.Write)
}

func (u userInterestDo) Session(config *gorm.Session) IUserInterestDo {
	return u.withDO(u.DO.Session(config))
}

func (u userInterestDo) Clauses(conds ...clause.Expression) IUserInterestDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userInterestDo) Returning(value interface{}, columns ...string) IUserInterestDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userInterestDo) Not(conds ...gen.Condition) IUserInterestDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userInterestDo) Or(conds ...gen.Condition) IUserInterestDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userInterestDo) Select(conds ...field.Expr) IUserInterestDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userInterestDo) Where(conds ...gen.Condition) IUserInterestDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userInterestDo) Order(conds ...field.Expr) IUserInterestDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userInterestDo) Distinct(cols ...field.Expr) IUserInterestDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userInterestDo) Omit(cols ...field.Expr) IUserInterestDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userInterestDo) Join(table schema.Tabler, on ...field.Expr) IUserInterestDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userInterestDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserInterestDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userInterestDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserInterestDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userInterestDo) Group(cols ...field.Expr) IUserInterestDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userInterestDo) Having(conds ...gen.Condition) IUserInterestDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userInterestDo) Limit(limit int) IUserInterestDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userInterestDo) Offset(offset int) IUserInterestDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userInterestDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserInterestDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userInterestDo) Unscoped() IUserInterestDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userInterestDo) Create(values ...*model.UserInterest) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userInterestDo) CreateInBatches(values []*model.UserInterest, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userInterestDo) Save(values ...*model.UserInterest) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userInterestDo) First() (*model.UserInterest, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserInterest), nil
	}
}

func (u userInterestDo) Take() (*model.UserInterest, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserInterest), nil
	}
}

func (u userInterestDo) Last() (*model.UserInterest, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserInterest), nil
	}
}

func (u userInterestDo) Find() ([]*model.UserInterest, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserInterest), err
}

func (u userInterestDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserInterest, err error) {
	buf := make([]*model.UserInterest, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userInterestDo) FindInBatches(result *[]*model.UserInterest, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userInterestDo) Attrs(attrs ...field.AssignExpr) IUserInterestDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userInterestDo) Assign(attrs ...field.AssignExpr) IUserInterestDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userInterestDo) Joins(fields ...field.RelationField) IUserInterestDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userInterestDo) Preload(fields ...field.RelationField) IUserInterestDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userInterestDo) FirstOrInit() (*model.UserInterest, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserInterest), nil
	}
}

func (u userInterestDo) FirstOrCreate() (*model.UserInterest, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserInterest), nil
	}
}

func (u userInterestDo) FindByPage(offset int, limit int) (result []*model.UserInterest, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userInterestDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userInterestDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userInterestDo) Delete(models ...*model.UserInterest) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userInterestDo) withDO(do gen.Dao) *userInterestDo {
	u.DO = *do.(*gen.DO)
	return u
}
