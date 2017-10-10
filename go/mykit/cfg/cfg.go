package cfg

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Config is a goroutine safe configuration store, with a map of values
// set from a config Provider.
type Config struct {
	m  map[string]string
	mu sync.RWMutex
}

// Provider is implemented by the user to provide the configuration as a map.
// There are currently three Providers implemented, EnvProvider, MapProvider and FileProvider.
type Provider interface {
	Provide() (map[string]string, error)
}

// New populates a new Config from a Provider. It will return an error if there
// was any problem reading from the Provider.
func New(p Provider) (*Config, error) {
	m, err := p.Provide()
	if err != nil {
		return nil, err
	}

	c := &Config{m: m}

	return c, nil
}

// Log returns a string to help with logging your configuration. It excludes
// any values whose key contains the string "PASS".
func (c *Config) Log() string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var buf bytes.Buffer
	for k, v := range c.m {
		if !strings.Contains(k, "PASS") {
			buf.WriteString(k + "=" + v + "\n")
		}
	}

	return buf.String()
}

// String returns the value of the given key as a string. It will return an
// error if key was not found.
func (c *Config) String(key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		return "", fmt.Errorf("Unknown key %s", key)
	}

	return value, nil
}

// MustString returns the value of the given key as a string. It will panic if
// the key was not found.
func (c *Config) MustString(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		panic(fmt.Sprintf("Unknown key %s", key))
	}

	return value
}

// SetString adds or modifies the configuration for the specified key and value.
func (c *Config) SetString(key string, value string) {
	c.mu.RLock()
	{
		c.m[key] = value
	}
	c.mu.RUnlock()
}

// Int returns the value of the given key as an int. It will return an error if
// the key was not found or the value can't be converted to an int.
func (c *Config) Int(key string) (int, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		return 0, fmt.Errorf("Unknown key %s", key)
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return val, nil
}

// MustInt returns the value of the given key as an int. It will panic if the
// key was not found or the value can't be converted to an int.
func (c *Config) MustInt(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		panic(fmt.Sprintf("Unknown key %s", key))
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("key %q value is not an int", key))
	}

	return val
}

// SetInt adds or modifies the configuration for the specified key and value.
func (c *Config) SetInt(key string, value int) {
	c.mu.RLock()
	{
		c.m[key] = strconv.Itoa(value)
	}
	c.mu.RUnlock()
}

// Time returns the value of the given key as a Time. It will return an error
// if the key was not found or the value can't be converted to a Time.
func (c *Config) Time(key string) (time.Time, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		return time.Time{}, fmt.Errorf("Unknown key %s", key)
	}

	tv, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return tv, err
	}

	return tv, nil
}

// MustTime returns the value of the given key as a Time. It will panic if the
// key was not found or the value can't be converted to a Time.
func (c *Config) MustTime(key string) time.Time {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		panic(fmt.Sprintf("Unknown key %s", key))
	}

	tv, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(fmt.Sprintf("key %q value is not a Time", key))
	}

	return tv
}

// SetTime adds or modifies the configuration for the specified key and value.
func (c *Config) SetTime(key string, value time.Time) {
	c.mu.RLock()
	{
		c.m[key] = value.Format(time.RFC3339)
	}
	c.mu.RUnlock()
}

// Bool returns the bool value of a given key as a bool. It will return an
// error if the key was not found or the value can't be converted to a bool.
func (c *Config) Bool(key string) (bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		return false, fmt.Errorf("Unknown key %s", key)
	}

	value = strings.ToLower(value)
	if value == "yes" || value == "on" {
		value = "true"
	} else if value == "no" || value == "off" {
		value = "false"
	}

	bv, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}

	return bv, nil
}

// MustBool returns the bool value of a given key as a bool. It will panic if
// the key was not found or the value can't be converted to a bool.
func (c *Config) MustBool(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		panic(fmt.Sprintf("Unknown key %s", key))
	}

	value = strings.ToLower(value)
	if value == "yes" || value == "on" {
		value = "true"
	} else if value == "no" || value == "off" {
		value = "false"
	}

	val, err := strconv.ParseBool(value)
	if err != nil {
		panic(fmt.Sprintf("key %q value is not a bool", key))
	}

	return val
}

// SetBool adds or modifies the configuration for the specified key and value.
func (c *Config) SetBool(key string, value bool) {
	c.mu.RLock()
	{
		c.m[key] = strconv.FormatBool(value)
	}
	c.mu.RUnlock()
}

// URL returns the value of the given key as a URL. It will return an error if
// the key was not found or the value can't be converted to a URL.
func (c *Config) URL(key string) (*url.URL, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		return nil, fmt.Errorf("Unknown key %s", key)
	}

	val, err := url.Parse(value)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// MustURL returns the value of the given key as a URL. It will panic if the
// key was not found or the value can't be converted to a URL.
func (c *Config) MustURL(key string) *url.URL {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		panic(fmt.Sprintf("Unknown key %s", key))
	}

	val, err := url.Parse(value)
	if err != nil {
		panic(fmt.Sprintf("key %q value is not a URL", key))
	}

	return val
}

// SetURL adds or modifies the configuration for the specified key and value.
func (c *Config) SetURL(key string, value *url.URL) {
	c.mu.RLock()
	{
		c.m[key] = value.String()
	}
	c.mu.RUnlock()
}

// Duration returns the value of the given key as a Duration. It will return an
// error if the key was not found or the value can't be converted to a Duration.
func (c *Config) Duration(key string) (time.Duration, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		return time.Duration(0), fmt.Errorf("Unknown key %s", key)
	}

	val, err := time.ParseDuration(value)
	if err != nil {
		return val, err
	}

	return val, nil
}

// MustDuration returns the value of the given key as a Duration. It will panic
// if the key was not found or the value can't be converted into a Duration.
func (c *Config) MustDuration(key string) time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, found := c.m[key]
	if !found {
		panic(fmt.Sprintf("Unknown key %s", key))
	}

	val, err := time.ParseDuration(value)
	if err != nil {
		panic(fmt.Sprintf("key %q value is not a time.Duration", key))
	}

	return val
}

// SetDuration adds or modifies the configuration for a given duration at a
// specific key.
func (c *Config) SetDuration(key string, value time.Duration) {
	c.mu.RLock()
	{
		c.m[key] = value.String()
	}
	c.mu.RUnlock()
}
