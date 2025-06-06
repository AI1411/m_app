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

func newPrefecture(db *gorm.DB, opts ...gen.DOOption) prefecture {
	_prefecture := prefecture{}

	_prefecture.prefectureDo.UseDB(db, opts...)
	_prefecture.prefectureDo.UseModel(&model.Prefecture{})

	tableName := _prefecture.prefectureDo.TableName()
	_prefecture.ALL = field.NewAsterisk(tableName)
	_prefecture.ID = field.NewInt32(tableName, "id")
	_prefecture.Code = field.NewInt16(tableName, "code")
	_prefecture.Name = field.NewString(tableName, "name")
	_prefecture.NameEn = field.NewString(tableName, "name_en")
	_prefecture.RegionID = field.NewInt32(tableName, "region_id")
	_prefecture.CreatedAt = field.NewTime(tableName, "created_at")
	_prefecture.UpdatedAt = field.NewTime(tableName, "updated_at")
	_prefecture.DeletedAt = field.NewField(tableName, "deleted_at")
	_prefecture.Region = prefectureBelongsToRegion{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Region", "model.Region"),
	}

	_prefecture.Users = prefectureHasManyUsers{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Users", "model.User"),
	}

	_prefecture.fillFieldMap()

	return _prefecture
}

type prefecture struct {
	prefectureDo

	ALL       field.Asterisk
	ID        field.Int32  // プライマリーキー
	Code      field.Int16  // 都道府県コード（JIS X 0401準拠）
	Name      field.String // 都道府県名
	NameEn    field.String // 都道府県名（英語）
	RegionID  field.Int32  // 地域区分ID
	CreatedAt field.Time   // 作成日時
	UpdatedAt field.Time   // 更新日時
	DeletedAt field.Field  // 削除日時（論理削除用）
	Region    prefectureBelongsToRegion

	Users prefectureHasManyUsers

	fieldMap map[string]field.Expr
}

func (p prefecture) Table(newTableName string) *prefecture {
	p.prefectureDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p prefecture) As(alias string) *prefecture {
	p.prefectureDo.DO = *(p.prefectureDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *prefecture) updateTableName(table string) *prefecture {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.Code = field.NewInt16(table, "code")
	p.Name = field.NewString(table, "name")
	p.NameEn = field.NewString(table, "name_en")
	p.RegionID = field.NewInt32(table, "region_id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *prefecture) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *prefecture) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 10)
	p.fieldMap["id"] = p.ID
	p.fieldMap["code"] = p.Code
	p.fieldMap["name"] = p.Name
	p.fieldMap["name_en"] = p.NameEn
	p.fieldMap["region_id"] = p.RegionID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt

}

func (p prefecture) clone(db *gorm.DB) prefecture {
	p.prefectureDo.ReplaceConnPool(db.Statement.ConnPool)
	p.Region.db = db.Session(&gorm.Session{Initialized: true})
	p.Region.db.Statement.ConnPool = db.Statement.ConnPool
	p.Users.db = db.Session(&gorm.Session{Initialized: true})
	p.Users.db.Statement.ConnPool = db.Statement.ConnPool
	return p
}

func (p prefecture) replaceDB(db *gorm.DB) prefecture {
	p.prefectureDo.ReplaceDB(db)
	p.Region.db = db.Session(&gorm.Session{})
	p.Users.db = db.Session(&gorm.Session{})
	return p
}

type prefectureBelongsToRegion struct {
	db *gorm.DB

	field.RelationField
}

func (a prefectureBelongsToRegion) Where(conds ...field.Expr) *prefectureBelongsToRegion {
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

func (a prefectureBelongsToRegion) WithContext(ctx context.Context) *prefectureBelongsToRegion {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a prefectureBelongsToRegion) Session(session *gorm.Session) *prefectureBelongsToRegion {
	a.db = a.db.Session(session)
	return &a
}

func (a prefectureBelongsToRegion) Model(m *model.Prefecture) *prefectureBelongsToRegionTx {
	return &prefectureBelongsToRegionTx{a.db.Model(m).Association(a.Name())}
}

func (a prefectureBelongsToRegion) Unscoped() *prefectureBelongsToRegion {
	a.db = a.db.Unscoped()
	return &a
}

type prefectureBelongsToRegionTx struct{ tx *gorm.Association }

func (a prefectureBelongsToRegionTx) Find() (result *model.Region, err error) {
	return result, a.tx.Find(&result)
}

func (a prefectureBelongsToRegionTx) Append(values ...*model.Region) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a prefectureBelongsToRegionTx) Replace(values ...*model.Region) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a prefectureBelongsToRegionTx) Delete(values ...*model.Region) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a prefectureBelongsToRegionTx) Clear() error {
	return a.tx.Clear()
}

func (a prefectureBelongsToRegionTx) Count() int64 {
	return a.tx.Count()
}

func (a prefectureBelongsToRegionTx) Unscoped() *prefectureBelongsToRegionTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type prefectureHasManyUsers struct {
	db *gorm.DB

	field.RelationField
}

func (a prefectureHasManyUsers) Where(conds ...field.Expr) *prefectureHasManyUsers {
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

func (a prefectureHasManyUsers) WithContext(ctx context.Context) *prefectureHasManyUsers {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a prefectureHasManyUsers) Session(session *gorm.Session) *prefectureHasManyUsers {
	a.db = a.db.Session(session)
	return &a
}

func (a prefectureHasManyUsers) Model(m *model.Prefecture) *prefectureHasManyUsersTx {
	return &prefectureHasManyUsersTx{a.db.Model(m).Association(a.Name())}
}

func (a prefectureHasManyUsers) Unscoped() *prefectureHasManyUsers {
	a.db = a.db.Unscoped()
	return &a
}

type prefectureHasManyUsersTx struct{ tx *gorm.Association }

func (a prefectureHasManyUsersTx) Find() (result []*model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a prefectureHasManyUsersTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a prefectureHasManyUsersTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a prefectureHasManyUsersTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a prefectureHasManyUsersTx) Clear() error {
	return a.tx.Clear()
}

func (a prefectureHasManyUsersTx) Count() int64 {
	return a.tx.Count()
}

func (a prefectureHasManyUsersTx) Unscoped() *prefectureHasManyUsersTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type prefectureDo struct{ gen.DO }

type IPrefectureDo interface {
	gen.SubQuery
	Debug() IPrefectureDo
	WithContext(ctx context.Context) IPrefectureDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPrefectureDo
	WriteDB() IPrefectureDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPrefectureDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPrefectureDo
	Not(conds ...gen.Condition) IPrefectureDo
	Or(conds ...gen.Condition) IPrefectureDo
	Select(conds ...field.Expr) IPrefectureDo
	Where(conds ...gen.Condition) IPrefectureDo
	Order(conds ...field.Expr) IPrefectureDo
	Distinct(cols ...field.Expr) IPrefectureDo
	Omit(cols ...field.Expr) IPrefectureDo
	Join(table schema.Tabler, on ...field.Expr) IPrefectureDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPrefectureDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPrefectureDo
	Group(cols ...field.Expr) IPrefectureDo
	Having(conds ...gen.Condition) IPrefectureDo
	Limit(limit int) IPrefectureDo
	Offset(offset int) IPrefectureDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPrefectureDo
	Unscoped() IPrefectureDo
	Create(values ...*model.Prefecture) error
	CreateInBatches(values []*model.Prefecture, batchSize int) error
	Save(values ...*model.Prefecture) error
	First() (*model.Prefecture, error)
	Take() (*model.Prefecture, error)
	Last() (*model.Prefecture, error)
	Find() ([]*model.Prefecture, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Prefecture, err error)
	FindInBatches(result *[]*model.Prefecture, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Prefecture) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPrefectureDo
	Assign(attrs ...field.AssignExpr) IPrefectureDo
	Joins(fields ...field.RelationField) IPrefectureDo
	Preload(fields ...field.RelationField) IPrefectureDo
	FirstOrInit() (*model.Prefecture, error)
	FirstOrCreate() (*model.Prefecture, error)
	FindByPage(offset int, limit int) (result []*model.Prefecture, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPrefectureDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p prefectureDo) Debug() IPrefectureDo {
	return p.withDO(p.DO.Debug())
}

func (p prefectureDo) WithContext(ctx context.Context) IPrefectureDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p prefectureDo) ReadDB() IPrefectureDo {
	return p.Clauses(dbresolver.Read)
}

func (p prefectureDo) WriteDB() IPrefectureDo {
	return p.Clauses(dbresolver.Write)
}

func (p prefectureDo) Session(config *gorm.Session) IPrefectureDo {
	return p.withDO(p.DO.Session(config))
}

func (p prefectureDo) Clauses(conds ...clause.Expression) IPrefectureDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p prefectureDo) Returning(value interface{}, columns ...string) IPrefectureDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p prefectureDo) Not(conds ...gen.Condition) IPrefectureDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p prefectureDo) Or(conds ...gen.Condition) IPrefectureDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p prefectureDo) Select(conds ...field.Expr) IPrefectureDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p prefectureDo) Where(conds ...gen.Condition) IPrefectureDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p prefectureDo) Order(conds ...field.Expr) IPrefectureDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p prefectureDo) Distinct(cols ...field.Expr) IPrefectureDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p prefectureDo) Omit(cols ...field.Expr) IPrefectureDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p prefectureDo) Join(table schema.Tabler, on ...field.Expr) IPrefectureDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p prefectureDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPrefectureDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p prefectureDo) RightJoin(table schema.Tabler, on ...field.Expr) IPrefectureDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p prefectureDo) Group(cols ...field.Expr) IPrefectureDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p prefectureDo) Having(conds ...gen.Condition) IPrefectureDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p prefectureDo) Limit(limit int) IPrefectureDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p prefectureDo) Offset(offset int) IPrefectureDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p prefectureDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPrefectureDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p prefectureDo) Unscoped() IPrefectureDo {
	return p.withDO(p.DO.Unscoped())
}

func (p prefectureDo) Create(values ...*model.Prefecture) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p prefectureDo) CreateInBatches(values []*model.Prefecture, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p prefectureDo) Save(values ...*model.Prefecture) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p prefectureDo) First() (*model.Prefecture, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Prefecture), nil
	}
}

func (p prefectureDo) Take() (*model.Prefecture, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Prefecture), nil
	}
}

func (p prefectureDo) Last() (*model.Prefecture, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Prefecture), nil
	}
}

func (p prefectureDo) Find() ([]*model.Prefecture, error) {
	result, err := p.DO.Find()
	return result.([]*model.Prefecture), err
}

func (p prefectureDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Prefecture, err error) {
	buf := make([]*model.Prefecture, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p prefectureDo) FindInBatches(result *[]*model.Prefecture, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p prefectureDo) Attrs(attrs ...field.AssignExpr) IPrefectureDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p prefectureDo) Assign(attrs ...field.AssignExpr) IPrefectureDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p prefectureDo) Joins(fields ...field.RelationField) IPrefectureDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p prefectureDo) Preload(fields ...field.RelationField) IPrefectureDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p prefectureDo) FirstOrInit() (*model.Prefecture, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Prefecture), nil
	}
}

func (p prefectureDo) FirstOrCreate() (*model.Prefecture, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Prefecture), nil
	}
}

func (p prefectureDo) FindByPage(offset int, limit int) (result []*model.Prefecture, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p prefectureDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p prefectureDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p prefectureDo) Delete(models ...*model.Prefecture) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *prefectureDo) withDO(do gen.Dao) *prefectureDo {
	p.DO = *do.(*gen.DO)
	return p
}
