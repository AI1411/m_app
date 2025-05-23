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

func newMatch(db *gorm.DB, opts ...gen.DOOption) match {
	_match := match{}

	_match.matchDo.UseDB(db, opts...)
	_match.matchDo.UseModel(&model.Match{})

	tableName := _match.matchDo.TableName()
	_match.ALL = field.NewAsterisk(tableName)
	_match.ID = field.NewString(tableName, "id")
	_match.UserId1 = field.NewString(tableName, "user_id_1")
	_match.UserId2 = field.NewString(tableName, "user_id_2")
	_match.MatchedAt = field.NewTime(tableName, "matched_at")
	_match.CreatedAt = field.NewTime(tableName, "created_at")
	_match.UpdatedAt = field.NewTime(tableName, "updated_at")
	_match.DeletedAt = field.NewField(tableName, "deleted_at")
	_match.User = matchBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "model.User"),
	}

	_match.MatchedUser = matchBelongsToMatchedUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("MatchedUser", "model.User"),
	}

	_match.fillFieldMap()

	return _match
}

type match struct {
	matchDo

	ALL       field.Asterisk
	ID        field.String // マッチングの一意識別子（UUIDv4）
	UserId1   field.String // マッチングしたユーザー1のID
	UserId2   field.String // マッチングしたユーザー2のID
	MatchedAt field.Time   // マッチング成立日時
	CreatedAt field.Time   // レコード作成日時
	UpdatedAt field.Time   // レコード更新日時
	DeletedAt field.Field  // 論理削除日時（NULLは有効なレコードを示す）
	User      matchBelongsToUser

	MatchedUser matchBelongsToMatchedUser

	fieldMap map[string]field.Expr
}

func (m match) Table(newTableName string) *match {
	m.matchDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m match) As(alias string) *match {
	m.matchDo.DO = *(m.matchDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *match) updateTableName(table string) *match {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewString(table, "id")
	m.UserId1 = field.NewString(table, "user_id_1")
	m.UserId2 = field.NewString(table, "user_id_2")
	m.MatchedAt = field.NewTime(table, "matched_at")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")

	m.fillFieldMap()

	return m
}

func (m *match) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *match) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 9)
	m.fieldMap["id"] = m.ID
	m.fieldMap["user_id_1"] = m.UserId1
	m.fieldMap["user_id_2"] = m.UserId2
	m.fieldMap["matched_at"] = m.MatchedAt
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt

}

func (m match) clone(db *gorm.DB) match {
	m.matchDo.ReplaceConnPool(db.Statement.ConnPool)
	m.User.db = db.Session(&gorm.Session{Initialized: true})
	m.User.db.Statement.ConnPool = db.Statement.ConnPool
	m.MatchedUser.db = db.Session(&gorm.Session{Initialized: true})
	m.MatchedUser.db.Statement.ConnPool = db.Statement.ConnPool
	return m
}

func (m match) replaceDB(db *gorm.DB) match {
	m.matchDo.ReplaceDB(db)
	m.User.db = db.Session(&gorm.Session{})
	m.MatchedUser.db = db.Session(&gorm.Session{})
	return m
}

type matchBelongsToUser struct {
	db *gorm.DB

	field.RelationField
}

func (a matchBelongsToUser) Where(conds ...field.Expr) *matchBelongsToUser {
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

func (a matchBelongsToUser) WithContext(ctx context.Context) *matchBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a matchBelongsToUser) Session(session *gorm.Session) *matchBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a matchBelongsToUser) Model(m *model.Match) *matchBelongsToUserTx {
	return &matchBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

func (a matchBelongsToUser) Unscoped() *matchBelongsToUser {
	a.db = a.db.Unscoped()
	return &a
}

type matchBelongsToUserTx struct{ tx *gorm.Association }

func (a matchBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a matchBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a matchBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a matchBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a matchBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a matchBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

func (a matchBelongsToUserTx) Unscoped() *matchBelongsToUserTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type matchBelongsToMatchedUser struct {
	db *gorm.DB

	field.RelationField
}

func (a matchBelongsToMatchedUser) Where(conds ...field.Expr) *matchBelongsToMatchedUser {
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

func (a matchBelongsToMatchedUser) WithContext(ctx context.Context) *matchBelongsToMatchedUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a matchBelongsToMatchedUser) Session(session *gorm.Session) *matchBelongsToMatchedUser {
	a.db = a.db.Session(session)
	return &a
}

func (a matchBelongsToMatchedUser) Model(m *model.Match) *matchBelongsToMatchedUserTx {
	return &matchBelongsToMatchedUserTx{a.db.Model(m).Association(a.Name())}
}

func (a matchBelongsToMatchedUser) Unscoped() *matchBelongsToMatchedUser {
	a.db = a.db.Unscoped()
	return &a
}

type matchBelongsToMatchedUserTx struct{ tx *gorm.Association }

func (a matchBelongsToMatchedUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a matchBelongsToMatchedUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a matchBelongsToMatchedUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a matchBelongsToMatchedUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a matchBelongsToMatchedUserTx) Clear() error {
	return a.tx.Clear()
}

func (a matchBelongsToMatchedUserTx) Count() int64 {
	return a.tx.Count()
}

func (a matchBelongsToMatchedUserTx) Unscoped() *matchBelongsToMatchedUserTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type matchDo struct{ gen.DO }

type IMatchDo interface {
	gen.SubQuery
	Debug() IMatchDo
	WithContext(ctx context.Context) IMatchDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMatchDo
	WriteDB() IMatchDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMatchDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMatchDo
	Not(conds ...gen.Condition) IMatchDo
	Or(conds ...gen.Condition) IMatchDo
	Select(conds ...field.Expr) IMatchDo
	Where(conds ...gen.Condition) IMatchDo
	Order(conds ...field.Expr) IMatchDo
	Distinct(cols ...field.Expr) IMatchDo
	Omit(cols ...field.Expr) IMatchDo
	Join(table schema.Tabler, on ...field.Expr) IMatchDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMatchDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMatchDo
	Group(cols ...field.Expr) IMatchDo
	Having(conds ...gen.Condition) IMatchDo
	Limit(limit int) IMatchDo
	Offset(offset int) IMatchDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMatchDo
	Unscoped() IMatchDo
	Create(values ...*model.Match) error
	CreateInBatches(values []*model.Match, batchSize int) error
	Save(values ...*model.Match) error
	First() (*model.Match, error)
	Take() (*model.Match, error)
	Last() (*model.Match, error)
	Find() ([]*model.Match, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Match, err error)
	FindInBatches(result *[]*model.Match, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Match) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMatchDo
	Assign(attrs ...field.AssignExpr) IMatchDo
	Joins(fields ...field.RelationField) IMatchDo
	Preload(fields ...field.RelationField) IMatchDo
	FirstOrInit() (*model.Match, error)
	FirstOrCreate() (*model.Match, error)
	FindByPage(offset int, limit int) (result []*model.Match, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMatchDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m matchDo) Debug() IMatchDo {
	return m.withDO(m.DO.Debug())
}

func (m matchDo) WithContext(ctx context.Context) IMatchDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m matchDo) ReadDB() IMatchDo {
	return m.Clauses(dbresolver.Read)
}

func (m matchDo) WriteDB() IMatchDo {
	return m.Clauses(dbresolver.Write)
}

func (m matchDo) Session(config *gorm.Session) IMatchDo {
	return m.withDO(m.DO.Session(config))
}

func (m matchDo) Clauses(conds ...clause.Expression) IMatchDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m matchDo) Returning(value interface{}, columns ...string) IMatchDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m matchDo) Not(conds ...gen.Condition) IMatchDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m matchDo) Or(conds ...gen.Condition) IMatchDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m matchDo) Select(conds ...field.Expr) IMatchDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m matchDo) Where(conds ...gen.Condition) IMatchDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m matchDo) Order(conds ...field.Expr) IMatchDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m matchDo) Distinct(cols ...field.Expr) IMatchDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m matchDo) Omit(cols ...field.Expr) IMatchDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m matchDo) Join(table schema.Tabler, on ...field.Expr) IMatchDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m matchDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMatchDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m matchDo) RightJoin(table schema.Tabler, on ...field.Expr) IMatchDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m matchDo) Group(cols ...field.Expr) IMatchDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m matchDo) Having(conds ...gen.Condition) IMatchDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m matchDo) Limit(limit int) IMatchDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m matchDo) Offset(offset int) IMatchDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m matchDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMatchDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m matchDo) Unscoped() IMatchDo {
	return m.withDO(m.DO.Unscoped())
}

func (m matchDo) Create(values ...*model.Match) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m matchDo) CreateInBatches(values []*model.Match, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m matchDo) Save(values ...*model.Match) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m matchDo) First() (*model.Match, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Match), nil
	}
}

func (m matchDo) Take() (*model.Match, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Match), nil
	}
}

func (m matchDo) Last() (*model.Match, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Match), nil
	}
}

func (m matchDo) Find() ([]*model.Match, error) {
	result, err := m.DO.Find()
	return result.([]*model.Match), err
}

func (m matchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Match, err error) {
	buf := make([]*model.Match, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m matchDo) FindInBatches(result *[]*model.Match, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m matchDo) Attrs(attrs ...field.AssignExpr) IMatchDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m matchDo) Assign(attrs ...field.AssignExpr) IMatchDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m matchDo) Joins(fields ...field.RelationField) IMatchDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m matchDo) Preload(fields ...field.RelationField) IMatchDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m matchDo) FirstOrInit() (*model.Match, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Match), nil
	}
}

func (m matchDo) FirstOrCreate() (*model.Match, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Match), nil
	}
}

func (m matchDo) FindByPage(offset int, limit int) (result []*model.Match, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m matchDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m matchDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m matchDo) Delete(models ...*model.Match) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *matchDo) withDO(do gen.Dao) *matchDo {
	m.DO = *do.(*gen.DO)
	return m
}
