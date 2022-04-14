package migrator

import "fmt"

func (c *Config) rootConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=disable",
		c.Host,
		c.Port,
		c.Username,
		c.Password,
	)
}

func (c *Config) connectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.Host,
		c.Port,
		c.DatabaseName,
		c.Username,
		c.Password,
	)
}
