package cockroach

import (

	//sql import
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/zitadel/zitadel/internal/database/dialect"
)

func init() {
	config := &Config{}
	dialect.Register(config, config, true)
}
