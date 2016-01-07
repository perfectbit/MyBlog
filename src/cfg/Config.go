package cfg

import(
	s "system"
	c "controllers"
)

type Config struct {
	Pages []s.Route
}

func NewConfig() Config {
	var pages = []s.Route{
		s.Route{
			"/",
			&c.Home{},
			"home.html"},
		}

	var config = Config{pages}

	return config
}
