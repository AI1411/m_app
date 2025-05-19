package main

import (
	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/AI1411/m_app/internal/domain/model"
	"github.com/AI1411/m_app/internal/infra/db"
	applogger "github.com/AI1411/m_app/internal/infra/logger"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/domain/query", // 出力パス
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})

	sqlHandler, err := db.NewSqlHandler(
		"myuser",
		"mypassword",
		"localhost",
		"5432",
		"m_app",
		applogger.New(applogger.DefaultConfig()),
	)
	if err != nil {
		panic(err)
	}

	g.UseDB(sqlHandler.Conn)

	// Generate the code
	g.Execute()

	// 生成したModelにRelation情報を手動追加（これだけは手動対応が必要）
	allModels := []interface{}{
		g.GenerateModel(
			model.TableNameUser,
			gen.FieldRelateModel(field.HasMany, "UserInterests", model.UserInterest{}, nil),
		),
		g.GenerateModel(
			model.TableNameInterest,
			gen.FieldRelateModel(field.HasMany, "UserInterests", model.UserInterest{}, nil),
		),
		g.GenerateModel(
			model.TableNameUserInterest,
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "Interest", model.Interest{}, nil),
		),
	}

	g.ApplyBasic(allModels...)

	// Generate the code
	g.Execute()
}
