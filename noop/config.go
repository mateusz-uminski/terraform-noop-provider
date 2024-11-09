package noop

type Config struct {
	Dir string
}

func (c *Config) Meta() (interface{}, error) {
	return c, nil
}
