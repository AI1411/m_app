package main

import (
	"gorm.io/gen"

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
	all := g.GenerateAllTable() // database to table model.

	g.ApplyBasic(all...)

	// Generate the code
	g.Execute()
}
