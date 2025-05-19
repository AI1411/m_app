package main

import (
	"go.uber.org/fx"

	"github.com/AI1411/m_app/internal/di"
)

func main() {
	app := fx.New(
		di.Module,
	)

	app.Run()
}
