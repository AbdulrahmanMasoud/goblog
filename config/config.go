package config

type Config struct {
	App    App
	Server Server
	Auth   Auth
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type Auth struct {
	Email    string
	Password string
}
