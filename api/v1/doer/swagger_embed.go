package doer

import (
	_ "embed"
)

var (
	//go:embed doer.swagger.json
	APISwagger []byte
)
