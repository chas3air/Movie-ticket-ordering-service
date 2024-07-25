package config

import (
	"go_psql/internal/models"
)

const TemplatesPath = "go_psql/web/templates"
const DataBaseName = "models"
const UsersTableName = "users"
const MoviesTableName = "movies"

const LimitTime = 60
const CookieName = "session_user"

var SessionTable = map[string]models.Session{} // cookie - session{Login, last time activity}
var UsersTable = map[string]models.Customer{}  // Login - customer
