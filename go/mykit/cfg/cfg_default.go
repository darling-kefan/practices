package cfg

import (
	"time"
	"net/url"
)

// @todo 为什么不是var c *Config？因为所有方法的接收者是*Config而不是Config。

// c is the default Config instance used by Init and the package level funcs like
// String, MustString, and SetString.
var c Config

// Init populates the package's default Config and should be called only once.
// A Provider must be supplied which will return a map of key/value pairs to be
// loaded.
func Init(p Provider) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Get the provided configuration.
	m, err := p.Provide()
	if err != nil {
		return err
	}

	// Set it to the global instance.
	c.m = m

	return nil
}

func Log() string {
	return c.Log()
}

func String(key string) (string, error) {
	return c.String(key)
}

func MustString(key string) string {
	return c.MustString(key)
}

func SetString(key string, value string) {
	c.SetString(key, value)
}

func Int(key string) (int, error) {
	return c.Int(key)
}

func MustInt(key string) int {
	return c.MustInt(key)
}

func SetInt(key string, value int) {
	c.SetInt(key, value)
}

func Time(key string) (time.Time, error) {
	return c.Time(key)
}

func MustTime(key string) time.Time {
	return c.MustTime(key)
}

func SetTime(key string, value time.Time) {
	c.SetTime(key, value)
}

func Bool(key string) (bool, error) {
	return c.Bool(key)
}

func MustBool(key string) bool {
	return c.MustBool(key)
}

func SetBool(key string, value bool) {
	c.SetBool(key, value)
}

func URL(key string) (*url.URL, error) {
	return c.URL(key)
}

func MustURL(key string) *url.URL {
	return c.MustURL(key)
}

func SetURL(key string, value *url.URL) {
	c.SetURL(key, value)
}

func Duration(key string) (time.Duration, error) {
	return c.Duration(key)
}

func MustDuration(key string) time.Duration {
	return c.MustDuration(key)
}

func SetDuration(key string, value time.Duration) {
	c.SetDuration(key, value)
}
