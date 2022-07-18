package middleware

import (
	"myapp/data"

	"github.com/fengsterooni/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
