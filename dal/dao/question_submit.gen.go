// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"back/dal/model"
)

func newQuestionSubmit(db *gorm.DB, opts ...gen.DOOption) questionSubmit {
	_questionSubmit := questionSubmit{}

	_questionSubmit.questionSubmitDo.UseDB(db, opts...)
	_questionSubmit.questionSubmitDo.UseModel(&model.QuestionSubmit{})

	tableName := _questionSubmit.questionSubmitDo.TableName()
	_questionSubmit.ALL = field.NewAsterisk(tableName)
	_questionSubmit.ID = field.NewInt64(tableName, "id")
	_questionSubmit.CodeLanguage = field.NewString(tableName, "code_language")
	_questionSubmit.Code = field.NewString(tableName, "code")
	_questionSubmit.JudgeInfo = field.NewString(tableName, "judge_info")
	_questionSubmit.Status = field.NewInt32(tableName, "status")
	_questionSubmit.QuestionID = field.NewInt64(tableName, "question_id")
	_questionSubmit.UserID = field.NewString(tableName, "user_id")
	_questionSubmit.CreateTime = field.NewTime(tableName, "create_time")
	_questionSubmit.UpdateTime = field.NewTime(tableName, "update_time")
	_questionSubmit.IsDelete = field.NewInt32(tableName, "is_delete")

	_questionSubmit.fillFieldMap()

	return _questionSubmit
}

// questionSubmit 题目提交
type questionSubmit struct {
	questionSubmitDo questionSubmitDo

	ALL          field.Asterisk
	ID           field.Int64  // id
	CodeLanguage field.String // 编程语言
	Code         field.String // 用户代码
	JudgeInfo    field.String // 判题信息（json 对象）
	Status       field.Int32  // 判题状态（0 - 待判题、1 - 判题中、2 - 成功、3 - 失败）
	QuestionID   field.Int64  // 题目 id
	UserID       field.String
	CreateTime   field.Time  // 创建时间
	UpdateTime   field.Time  // 更新时间
	IsDelete     field.Int32 // 是否删除

	fieldMap map[string]field.Expr
}

func (q questionSubmit) Table(newTableName string) *questionSubmit {
	q.questionSubmitDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q questionSubmit) As(alias string) *questionSubmit {
	q.questionSubmitDo.DO = *(q.questionSubmitDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *questionSubmit) updateTableName(table string) *questionSubmit {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewInt64(table, "id")
	q.CodeLanguage = field.NewString(table, "code_language")
	q.Code = field.NewString(table, "code")
	q.JudgeInfo = field.NewString(table, "judge_info")
	q.Status = field.NewInt32(table, "status")
	q.QuestionID = field.NewInt64(table, "question_id")
	q.UserID = field.NewString(table, "user_id")
	q.CreateTime = field.NewTime(table, "create_time")
	q.UpdateTime = field.NewTime(table, "update_time")
	q.IsDelete = field.NewInt32(table, "is_delete")

	q.fillFieldMap()

	return q
}

func (q *questionSubmit) WithContext(ctx context.Context) *questionSubmitDo {
	return q.questionSubmitDo.WithContext(ctx)
}

func (q questionSubmit) TableName() string { return q.questionSubmitDo.TableName() }

func (q questionSubmit) Alias() string { return q.questionSubmitDo.Alias() }

func (q questionSubmit) Columns(cols ...field.Expr) gen.Columns {
	return q.questionSubmitDo.Columns(cols...)
}

func (q *questionSubmit) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *questionSubmit) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 10)
	q.fieldMap["id"] = q.ID
	q.fieldMap["code_language"] = q.CodeLanguage
	q.fieldMap["code"] = q.Code
	q.fieldMap["judge_info"] = q.JudgeInfo
	q.fieldMap["status"] = q.Status
	q.fieldMap["question_id"] = q.QuestionID
	q.fieldMap["user_id"] = q.UserID
	q.fieldMap["create_time"] = q.CreateTime
	q.fieldMap["update_time"] = q.UpdateTime
	q.fieldMap["is_delete"] = q.IsDelete
}

func (q questionSubmit) clone(db *gorm.DB) questionSubmit {
	q.questionSubmitDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q questionSubmit) replaceDB(db *gorm.DB) questionSubmit {
	q.questionSubmitDo.ReplaceDB(db)
	return q
}

type questionSubmitDo struct{ gen.DO }

func (q questionSubmitDo) Debug() *questionSubmitDo {
	return q.withDO(q.DO.Debug())
}

func (q questionSubmitDo) WithContext(ctx context.Context) *questionSubmitDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q questionSubmitDo) ReadDB() *questionSubmitDo {
	return q.Clauses(dbresolver.Read)
}

func (q questionSubmitDo) WriteDB() *questionSubmitDo {
	return q.Clauses(dbresolver.Write)
}

func (q questionSubmitDo) Session(config *gorm.Session) *questionSubmitDo {
	return q.withDO(q.DO.Session(config))
}

func (q questionSubmitDo) Clauses(conds ...clause.Expression) *questionSubmitDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q questionSubmitDo) Returning(value interface{}, columns ...string) *questionSubmitDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q questionSubmitDo) Not(conds ...gen.Condition) *questionSubmitDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q questionSubmitDo) Or(conds ...gen.Condition) *questionSubmitDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q questionSubmitDo) Select(conds ...field.Expr) *questionSubmitDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q questionSubmitDo) Where(conds ...gen.Condition) *questionSubmitDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q questionSubmitDo) Order(conds ...field.Expr) *questionSubmitDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q questionSubmitDo) Distinct(cols ...field.Expr) *questionSubmitDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q questionSubmitDo) Omit(cols ...field.Expr) *questionSubmitDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q questionSubmitDo) Join(table schema.Tabler, on ...field.Expr) *questionSubmitDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q questionSubmitDo) LeftJoin(table schema.Tabler, on ...field.Expr) *questionSubmitDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q questionSubmitDo) RightJoin(table schema.Tabler, on ...field.Expr) *questionSubmitDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q questionSubmitDo) Group(cols ...field.Expr) *questionSubmitDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q questionSubmitDo) Having(conds ...gen.Condition) *questionSubmitDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q questionSubmitDo) Limit(limit int) *questionSubmitDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q questionSubmitDo) Offset(offset int) *questionSubmitDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q questionSubmitDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *questionSubmitDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q questionSubmitDo) Unscoped() *questionSubmitDo {
	return q.withDO(q.DO.Unscoped())
}

func (q questionSubmitDo) Create(values ...*model.QuestionSubmit) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q questionSubmitDo) CreateInBatches(values []*model.QuestionSubmit, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q questionSubmitDo) Save(values ...*model.QuestionSubmit) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q questionSubmitDo) First() (*model.QuestionSubmit, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionSubmit), nil
	}
}

func (q questionSubmitDo) Take() (*model.QuestionSubmit, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionSubmit), nil
	}
}

func (q questionSubmitDo) Last() (*model.QuestionSubmit, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionSubmit), nil
	}
}

func (q questionSubmitDo) Find() ([]*model.QuestionSubmit, error) {
	result, err := q.DO.Find()
	return result.([]*model.QuestionSubmit), err
}

func (q questionSubmitDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QuestionSubmit, err error) {
	buf := make([]*model.QuestionSubmit, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q questionSubmitDo) FindInBatches(result *[]*model.QuestionSubmit, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q questionSubmitDo) Attrs(attrs ...field.AssignExpr) *questionSubmitDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q questionSubmitDo) Assign(attrs ...field.AssignExpr) *questionSubmitDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q questionSubmitDo) Joins(fields ...field.RelationField) *questionSubmitDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q questionSubmitDo) Preload(fields ...field.RelationField) *questionSubmitDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q questionSubmitDo) FirstOrInit() (*model.QuestionSubmit, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionSubmit), nil
	}
}

func (q questionSubmitDo) FirstOrCreate() (*model.QuestionSubmit, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionSubmit), nil
	}
}

func (q questionSubmitDo) FindByPage(offset int, limit int) (result []*model.QuestionSubmit, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q questionSubmitDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q questionSubmitDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q questionSubmitDo) Delete(models ...*model.QuestionSubmit) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *questionSubmitDo) withDO(do gen.Dao) *questionSubmitDo {
	q.DO = *do.(*gen.DO)
	return q
}
