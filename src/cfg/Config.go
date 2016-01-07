package cfg

import(
	s "system"
	c "controllers"
)

type Config struct {
	Pages []s.Route
}

func GetConfig() Config {
	return config
}


var config = Config{[]s.Route{
	s.Route{
		"/",
		&c.Home{},
		"home.html"},
	}}
