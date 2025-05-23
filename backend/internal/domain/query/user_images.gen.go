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

func newUserImage(db *gorm.DB, opts ...gen.DOOption) userImage {
	_userImage := userImage{}

	_userImage.userImageDo.UseDB(db, opts...)
	_userImage.userImageDo.UseModel(&model.UserImage{})

	tableName := _userImage.userImageDo.TableName()
	_userImage.ALL = field.NewAsterisk(tableName)
	_userImage.ID = field.NewString(tableName, "id")
	_userImage.UserID = field.NewString(tableName, "user_id")
	_userImage.ImageURL = field.NewString(tableName, "image_url")
	_userImage.Description = field.NewString(tableName, "description")
	_userImage.IsPrimary = field.NewBool(tableName, "is_primary")
	_userImage.CreatedAt = field.NewTime(tableName, "created_at")
	_userImage.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userImage.DeletedAt = field.NewField(tableName, "deleted_at")
	_userImage.User = userImageBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "model.User"),
	}

	_userImage.fillFieldMap()

	return _userImage
}

type userImage struct {
	userImageDo

	ALL         field.Asterisk
	ID          field.String // 画像の一意識別子（UUIDv4）
	UserID      field.String // 画像の所有者であるユーザーのID
	ImageURL    field.String // 画像のURL
	Description field.String // 画像の説明
	IsPrimary   field.Bool   // プロフィールのメイン画像かどうか（trueの場合はメイン画像）
	CreatedAt   field.Time   // レコード作成日時
	UpdatedAt   field.Time   // レコード更新日時
	DeletedAt   field.Field  // 論理削除日時（NULLは有効なレコードを示す）
	User        userImageBelongsToUser

	fieldMap map[string]field.Expr
}

func (u userImage) Table(newTableName string) *userImage {
	u.userImageDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userImage) As(alias string) *userImage {
	u.userImageDo.DO = *(u.userImageDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userImage) updateTableName(table string) *userImage {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewString(table, "id")
	u.UserID = field.NewString(table, "user_id")
	u.ImageURL = field.NewString(table, "image_url")
	u.Description = field.NewString(table, "description")
	u.IsPrimary = field.NewBool(table, "is_primary")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userImage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userImage) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 9)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["image_url"] = u.ImageURL
	u.fieldMap["description"] = u.Description
	u.fieldMap["is_primary"] = u.IsPrimary
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt

}

func (u userImage) clone(db *gorm.DB) userImage {
	u.userImageDo.ReplaceConnPool(db.Statement.ConnPool)
	u.User.db = db.Session(&gorm.Session{Initialized: true})
	u.User.db.Statement.ConnPool = db.Statement.ConnPool
	return u
}

func (u userImage) replaceDB(db *gorm.DB) userImage {
	u.userImageDo.ReplaceDB(db)
	u.User.db = db.Session(&gorm.Session{})
	return u
}

type userImageBelongsToUser struct {
	db *gorm.DB

	field.RelationField
}

func (a userImageBelongsToUser) Where(conds ...field.Expr) *userImageBelongsToUser {
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

func (a userImageBelongsToUser) WithContext(ctx context.Context) *userImageBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userImageBelongsToUser) Session(session *gorm.Session) *userImageBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a userImageBelongsToUser) Model(m *model.UserImage) *userImageBelongsToUserTx {
	return &userImageBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

func (a userImageBelongsToUser) Unscoped() *userImageBelongsToUser {
	a.db = a.db.Unscoped()
	return &a
}

type userImageBelongsToUserTx struct{ tx *gorm.Association }

func (a userImageBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a userImageBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userImageBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userImageBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userImageBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a userImageBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

func (a userImageBelongsToUserTx) Unscoped() *userImageBelongsToUserTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type userImageDo struct{ gen.DO }

type IUserImageDo interface {
	gen.SubQuery
	Debug() IUserImageDo
	WithContext(ctx context.Context) IUserImageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserImageDo
	WriteDB() IUserImageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserImageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserImageDo
	Not(conds ...gen.Condition) IUserImageDo
	Or(conds ...gen.Condition) IUserImageDo
	Select(conds ...field.Expr) IUserImageDo
	Where(conds ...gen.Condition) IUserImageDo
	Order(conds ...field.Expr) IUserImageDo
	Distinct(cols ...field.Expr) IUserImageDo
	Omit(cols ...field.Expr) IUserImageDo
	Join(table schema.Tabler, on ...field.Expr) IUserImageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserImageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserImageDo
	Group(cols ...field.Expr) IUserImageDo
	Having(conds ...gen.Condition) IUserImageDo
	Limit(limit int) IUserImageDo
	Offset(offset int) IUserImageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserImageDo
	Unscoped() IUserImageDo
	Create(values ...*model.UserImage) error
	CreateInBatches(values []*model.UserImage, batchSize int) error
	Save(values ...*model.UserImage) error
	First() (*model.UserImage, error)
	Take() (*model.UserImage, error)
	Last() (*model.UserImage, error)
	Find() ([]*model.UserImage, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserImage, err error)
	FindInBatches(result *[]*model.UserImage, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserImage) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserImageDo
	Assign(attrs ...field.AssignExpr) IUserImageDo
	Joins(fields ...field.RelationField) IUserImageDo
	Preload(fields ...field.RelationField) IUserImageDo
	FirstOrInit() (*model.UserImage, error)
	FirstOrCreate() (*model.UserImage, error)
	FindByPage(offset int, limit int) (result []*model.UserImage, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserImageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userImageDo) Debug() IUserImageDo {
	return u.withDO(u.DO.Debug())
}

func (u userImageDo) WithContext(ctx context.Context) IUserImageDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userImageDo) ReadDB() IUserImageDo {
	return u.Clauses(dbresolver.Read)
}

func (u userImageDo) WriteDB() IUserImageDo {
	return u.Clauses(dbresolver.Write)
}

func (u userImageDo) Session(config *gorm.Session) IUserImageDo {
	return u.withDO(u.DO.Session(config))
}

func (u userImageDo) Clauses(conds ...clause.Expression) IUserImageDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userImageDo) Returning(value interface{}, columns ...string) IUserImageDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userImageDo) Not(conds ...gen.Condition) IUserImageDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userImageDo) Or(conds ...gen.Condition) IUserImageDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userImageDo) Select(conds ...field.Expr) IUserImageDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userImageDo) Where(conds ...gen.Condition) IUserImageDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userImageDo) Order(conds ...field.Expr) IUserImageDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userImageDo) Distinct(cols ...field.Expr) IUserImageDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userImageDo) Omit(cols ...field.Expr) IUserImageDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userImageDo) Join(table schema.Tabler, on ...field.Expr) IUserImageDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userImageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserImageDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userImageDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserImageDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userImageDo) Group(cols ...field.Expr) IUserImageDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userImageDo) Having(conds ...gen.Condition) IUserImageDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userImageDo) Limit(limit int) IUserImageDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userImageDo) Offset(offset int) IUserImageDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userImageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserImageDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userImageDo) Unscoped() IUserImageDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userImageDo) Create(values ...*model.UserImage) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userImageDo) CreateInBatches(values []*model.UserImage, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userImageDo) Save(values ...*model.UserImage) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userImageDo) First() (*model.UserImage, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserImage), nil
	}
}

func (u userImageDo) Take() (*model.UserImage, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserImage), nil
	}
}

func (u userImageDo) Last() (*model.UserImage, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserImage), nil
	}
}

func (u userImageDo) Find() ([]*model.UserImage, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserImage), err
}

func (u userImageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserImage, err error) {
	buf := make([]*model.UserImage, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userImageDo) FindInBatches(result *[]*model.UserImage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userImageDo) Attrs(attrs ...field.AssignExpr) IUserImageDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userImageDo) Assign(attrs ...field.AssignExpr) IUserImageDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userImageDo) Joins(fields ...field.RelationField) IUserImageDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userImageDo) Preload(fields ...field.RelationField) IUserImageDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userImageDo) FirstOrInit() (*model.UserImage, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserImage), nil
	}
}

func (u userImageDo) FirstOrCreate() (*model.UserImage, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserImage), nil
	}
}

func (u userImageDo) FindByPage(offset int, limit int) (result []*model.UserImage, count int64, err error) {
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

func (u userImageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userImageDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userImageDo) Delete(models ...*model.UserImage) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userImageDo) withDO(do gen.Dao) *userImageDo {
	u.DO = *do.(*gen.DO)
	return u
}
