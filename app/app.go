package app

type Application struct {
	Env      *Env
	Postgres Postgres
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	app.Postgres.ClosePostgresDBConnection()
}
