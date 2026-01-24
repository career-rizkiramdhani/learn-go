# Config Module Context

Purpose: application configuration and DB initialization.

Key files:
- `config/app.go` - application config struct and loading
- `config/config_db.go` - database initialization

Notes:
- Responsible for loading `.env` and setting `config.Cfg` and `config.DB`.
