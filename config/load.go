package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/go-openapi/swag"
)

const tagName = "env"

// Load TODO
func Load() (*Config, error) {
	var c = Config{}

	t := reflect.TypeOf(c)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v := reflect.ValueOf(&c).Elem().FieldByName(f.Name)
		tagData := strings.Split(f.Tag.Get(tagName), ",")

		if len(tagData) == 0 {
			return nil, errors.New("invalid tag format")
		}

		var (
			envKey             = tagData[0]
			envValue, envFound = os.LookupEnv(envKey)
			required           = true
			isPtr              = false
		)

		if v.Kind() == reflect.String {
			// The field is a string
		} else if v.Kind() == reflect.Ptr && v.Type().Elem().Kind() == reflect.String {
			// The field is a *string
			required = false
			isPtr = true
		} else {
			// We don't support that type yet :(
			return nil, fmt.Errorf("warning: struct field %s must be of type string or *string", f.Name)
		}

		if required && !envFound {
			err := fmt.Errorf("env %s is required but not set", envKey)
			return nil, err
		}

		if !v.CanSet() {
			return nil, fmt.Errorf("cannot set field %s", f.Name)
		}

		// Expand the environment variables before setting
		envValue = os.ExpandEnv(envValue)
		os.Setenv(envKey, envValue)

		if isPtr {
			if envFound {
				v.Set(reflect.ValueOf(swag.String(envValue)))
			}
		} else {
			v.Set(reflect.ValueOf(envValue))
		}
	}

	return &c, nil
}
