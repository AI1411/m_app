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

func newCommunityMember(db *gorm.DB, opts ...gen.DOOption) communityMember {
	_communityMember := communityMember{}

	_communityMember.communityMemberDo.UseDB(db, opts...)
	_communityMember.communityMemberDo.UseModel(&model.CommunityMember{})

	tableName := _communityMember.communityMemberDo.TableName()
	_communityMember.ALL = field.NewAsterisk(tableName)
	_communityMember.CommunityID = field.NewString(tableName, "community_id")
	_communityMember.UserID = field.NewString(tableName, "user_id")
	_communityMember.Role = field.NewString(tableName, "role")
	_communityMember.IsApproved = field.NewBool(tableName, "is_approved")
	_communityMember.JoinedAt = field.NewTime(tableName, "joined_at")
	_communityMember.UpdatedAt = field.NewTime(tableName, "updated_at")
	_communityMember.Community = communityMemberBelongsToCommunity{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Community", "model.Community"),
	}

	_communityMember.User = communityMemberBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "model.User"),
	}

	_communityMember.fillFieldMap()

	return _communityMember
}

type communityMember struct {
	communityMemberDo

	ALL         field.Asterisk
	CommunityID field.String // コミュニティID
	UserID      field.String // ユーザーID
	Role        field.String // メンバーの役割（member, moderator, admin）
	IsApproved  field.Bool   // 承認状態（プライベートコミュニティの場合）
	JoinedAt    field.Time   // コミュニティ参加日時
	UpdatedAt   field.Time   // レコード更新日時
	Community   communityMemberBelongsToCommunity

	User communityMemberBelongsToUser

	fieldMap map[string]field.Expr
}

func (c communityMember) Table(newTableName string) *communityMember {
	c.communityMemberDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c communityMember) As(alias string) *communityMember {
	c.communityMemberDo.DO = *(c.communityMemberDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *communityMember) updateTableName(table string) *communityMember {
	c.ALL = field.NewAsterisk(table)
	c.CommunityID = field.NewString(table, "community_id")
	c.UserID = field.NewString(table, "user_id")
	c.Role = field.NewString(table, "role")
	c.IsApproved = field.NewBool(table, "is_approved")
	c.JoinedAt = field.NewTime(table, "joined_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *communityMember) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *communityMember) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 8)
	c.fieldMap["community_id"] = c.CommunityID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["role"] = c.Role
	c.fieldMap["is_approved"] = c.IsApproved
	c.fieldMap["joined_at"] = c.JoinedAt
	c.fieldMap["updated_at"] = c.UpdatedAt

}

func (c communityMember) clone(db *gorm.DB) communityMember {
	c.communityMemberDo.ReplaceConnPool(db.Statement.ConnPool)
	c.Community.db = db.Session(&gorm.Session{Initialized: true})
	c.Community.db.Statement.ConnPool = db.Statement.ConnPool
	c.User.db = db.Session(&gorm.Session{Initialized: true})
	c.User.db.Statement.ConnPool = db.Statement.ConnPool
	return c
}

func (c communityMember) replaceDB(db *gorm.DB) communityMember {
	c.communityMemberDo.ReplaceDB(db)
	c.Community.db = db.Session(&gorm.Session{})
	c.User.db = db.Session(&gorm.Session{})
	return c
}

type communityMemberBelongsToCommunity struct {
	db *gorm.DB

	field.RelationField
}

func (a communityMemberBelongsToCommunity) Where(conds ...field.Expr) *communityMemberBelongsToCommunity {
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

func (a communityMemberBelongsToCommunity) WithContext(ctx context.Context) *communityMemberBelongsToCommunity {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a communityMemberBelongsToCommunity) Session(session *gorm.Session) *communityMemberBelongsToCommunity {
	a.db = a.db.Session(session)
	return &a
}

func (a communityMemberBelongsToCommunity) Model(m *model.CommunityMember) *communityMemberBelongsToCommunityTx {
	return &communityMemberBelongsToCommunityTx{a.db.Model(m).Association(a.Name())}
}

func (a communityMemberBelongsToCommunity) Unscoped() *communityMemberBelongsToCommunity {
	a.db = a.db.Unscoped()
	return &a
}

type communityMemberBelongsToCommunityTx struct{ tx *gorm.Association }

func (a communityMemberBelongsToCommunityTx) Find() (result *model.Community, err error) {
	return result, a.tx.Find(&result)
}

func (a communityMemberBelongsToCommunityTx) Append(values ...*model.Community) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a communityMemberBelongsToCommunityTx) Replace(values ...*model.Community) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a communityMemberBelongsToCommunityTx) Delete(values ...*model.Community) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a communityMemberBelongsToCommunityTx) Clear() error {
	return a.tx.Clear()
}

func (a communityMemberBelongsToCommunityTx) Count() int64 {
	return a.tx.Count()
}

func (a communityMemberBelongsToCommunityTx) Unscoped() *communityMemberBelongsToCommunityTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type communityMemberBelongsToUser struct {
	db *gorm.DB

	field.RelationField
}

func (a communityMemberBelongsToUser) Where(conds ...field.Expr) *communityMemberBelongsToUser {
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

func (a communityMemberBelongsToUser) WithContext(ctx context.Context) *communityMemberBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a communityMemberBelongsToUser) Session(session *gorm.Session) *communityMemberBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a communityMemberBelongsToUser) Model(m *model.CommunityMember) *communityMemberBelongsToUserTx {
	return &communityMemberBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

func (a communityMemberBelongsToUser) Unscoped() *communityMemberBelongsToUser {
	a.db = a.db.Unscoped()
	return &a
}

type communityMemberBelongsToUserTx struct{ tx *gorm.Association }

func (a communityMemberBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a communityMemberBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a communityMemberBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a communityMemberBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a communityMemberBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a communityMemberBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

func (a communityMemberBelongsToUserTx) Unscoped() *communityMemberBelongsToUserTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type communityMemberDo struct{ gen.DO }

type ICommunityMemberDo interface {
	gen.SubQuery
	Debug() ICommunityMemberDo
	WithContext(ctx context.Context) ICommunityMemberDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommunityMemberDo
	WriteDB() ICommunityMemberDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommunityMemberDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommunityMemberDo
	Not(conds ...gen.Condition) ICommunityMemberDo
	Or(conds ...gen.Condition) ICommunityMemberDo
	Select(conds ...field.Expr) ICommunityMemberDo
	Where(conds ...gen.Condition) ICommunityMemberDo
	Order(conds ...field.Expr) ICommunityMemberDo
	Distinct(cols ...field.Expr) ICommunityMemberDo
	Omit(cols ...field.Expr) ICommunityMemberDo
	Join(table schema.Tabler, on ...field.Expr) ICommunityMemberDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommunityMemberDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommunityMemberDo
	Group(cols ...field.Expr) ICommunityMemberDo
	Having(conds ...gen.Condition) ICommunityMemberDo
	Limit(limit int) ICommunityMemberDo
	Offset(offset int) ICommunityMemberDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommunityMemberDo
	Unscoped() ICommunityMemberDo
	Create(values ...*model.CommunityMember) error
	CreateInBatches(values []*model.CommunityMember, batchSize int) error
	Save(values ...*model.CommunityMember) error
	First() (*model.CommunityMember, error)
	Take() (*model.CommunityMember, error)
	Last() (*model.CommunityMember, error)
	Find() ([]*model.CommunityMember, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommunityMember, err error)
	FindInBatches(result *[]*model.CommunityMember, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CommunityMember) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommunityMemberDo
	Assign(attrs ...field.AssignExpr) ICommunityMemberDo
	Joins(fields ...field.RelationField) ICommunityMemberDo
	Preload(fields ...field.RelationField) ICommunityMemberDo
	FirstOrInit() (*model.CommunityMember, error)
	FirstOrCreate() (*model.CommunityMember, error)
	FindByPage(offset int, limit int) (result []*model.CommunityMember, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommunityMemberDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c communityMemberDo) Debug() ICommunityMemberDo {
	return c.withDO(c.DO.Debug())
}

func (c communityMemberDo) WithContext(ctx context.Context) ICommunityMemberDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c communityMemberDo) ReadDB() ICommunityMemberDo {
	return c.Clauses(dbresolver.Read)
}

func (c communityMemberDo) WriteDB() ICommunityMemberDo {
	return c.Clauses(dbresolver.Write)
}

func (c communityMemberDo) Session(config *gorm.Session) ICommunityMemberDo {
	return c.withDO(c.DO.Session(config))
}

func (c communityMemberDo) Clauses(conds ...clause.Expression) ICommunityMemberDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c communityMemberDo) Returning(value interface{}, columns ...string) ICommunityMemberDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c communityMemberDo) Not(conds ...gen.Condition) ICommunityMemberDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c communityMemberDo) Or(conds ...gen.Condition) ICommunityMemberDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c communityMemberDo) Select(conds ...field.Expr) ICommunityMemberDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c communityMemberDo) Where(conds ...gen.Condition) ICommunityMemberDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c communityMemberDo) Order(conds ...field.Expr) ICommunityMemberDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c communityMemberDo) Distinct(cols ...field.Expr) ICommunityMemberDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c communityMemberDo) Omit(cols ...field.Expr) ICommunityMemberDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c communityMemberDo) Join(table schema.Tabler, on ...field.Expr) ICommunityMemberDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c communityMemberDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommunityMemberDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c communityMemberDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommunityMemberDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c communityMemberDo) Group(cols ...field.Expr) ICommunityMemberDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c communityMemberDo) Having(conds ...gen.Condition) ICommunityMemberDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c communityMemberDo) Limit(limit int) ICommunityMemberDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c communityMemberDo) Offset(offset int) ICommunityMemberDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c communityMemberDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommunityMemberDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c communityMemberDo) Unscoped() ICommunityMemberDo {
	return c.withDO(c.DO.Unscoped())
}

func (c communityMemberDo) Create(values ...*model.CommunityMember) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c communityMemberDo) CreateInBatches(values []*model.CommunityMember, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c communityMemberDo) Save(values ...*model.CommunityMember) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c communityMemberDo) First() (*model.CommunityMember, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommunityMember), nil
	}
}

func (c communityMemberDo) Take() (*model.CommunityMember, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommunityMember), nil
	}
}

func (c communityMemberDo) Last() (*model.CommunityMember, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommunityMember), nil
	}
}

func (c communityMemberDo) Find() ([]*model.CommunityMember, error) {
	result, err := c.DO.Find()
	return result.([]*model.CommunityMember), err
}

func (c communityMemberDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommunityMember, err error) {
	buf := make([]*model.CommunityMember, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c communityMemberDo) FindInBatches(result *[]*model.CommunityMember, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c communityMemberDo) Attrs(attrs ...field.AssignExpr) ICommunityMemberDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c communityMemberDo) Assign(attrs ...field.AssignExpr) ICommunityMemberDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c communityMemberDo) Joins(fields ...field.RelationField) ICommunityMemberDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c communityMemberDo) Preload(fields ...field.RelationField) ICommunityMemberDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c communityMemberDo) FirstOrInit() (*model.CommunityMember, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommunityMember), nil
	}
}

func (c communityMemberDo) FirstOrCreate() (*model.CommunityMember, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommunityMember), nil
	}
}

func (c communityMemberDo) FindByPage(offset int, limit int) (result []*model.CommunityMember, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c communityMemberDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c communityMemberDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c communityMemberDo) Delete(models ...*model.CommunityMember) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *communityMemberDo) withDO(do gen.Dao) *communityMemberDo {
	c.DO = *do.(*gen.DO)
	return c
}
