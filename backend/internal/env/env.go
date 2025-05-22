package env

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Values struct {
	DB
	Env        string `default:"local" split_words:"true"`
	ServerPort string `required:"true" split_words:"true"`
}

type DB struct {
	DatabaseHost          string        `required:"true" split_words:"true"`
	DatabaseUsername      string        `required:"true" split_words:"true"`
	DatabasePassword      string        `required:"true" split_words:"true"`
	DatabaseName          string        `required:"true" split_words:"true"`
	DatabasePort          string        `default:"3306" split_words:"true"`
	ConnectionMaxOpen     int           `default:"10" split_words:"true"`
	ConnectionMaxIdle     int           `default:"2" split_words:"true"`
	ConnectionMaxLifetime time.Duration `default:"300s" split_words:"true"`
}

func NewValues() (*Values, error) {
	var v Values
	err := envconfig.Process("", &v)
	if err != nil {
		s := fmt.Sprintf("need to set all env values %+v", v)
		return nil, errors.Wrap(err, s)
	}
	return &v, nil
}

func (v Values) IsLocal() bool {
	return v.Env == "local" || v.Env == "test"
}
