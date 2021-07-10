package config

import (
	"github.com/yukitsune/chameleon/pkg/errors"
	"reflect"
	"strings"
)

const (
	configKeyTag = "mapstructure"
)

type Key []string

func (k Key) String(sep string) string {
	return strings.Join(k, sep)
}

type keyVisitor func(k *Key, f reflect.StructField, v reflect.Value) error

func visitKeys(k *Key, t reflect.Type, v reflect.Value, fn keyVisitor) error {
	var errs errors.Errors
	for i := 0; i < t.NumField(); i++ {

		f := t.Field(i)
		var fv reflect.Value
		if v.Kind() == reflect.Ptr {
			fv = reflect.Indirect(v).FieldByName(f.Name)
		} else {
			fv = v.FieldByName(f.Name)
		}

		keyString := f.Tag.Get(configKeyTag)
		if keyString == "" {
			continue
		}

		if strings.Contains(keyString, ",") {
			tagParts := strings.Split(keyString, ",")
			keyString = tagParts[0]
		}

		key := append(*k, keyString)
		if f.Type.Kind() == reflect.Ptr && f.Type.Elem().Kind() == reflect.Struct {
			err := visitKeys(&key, f.Type.Elem(), fv, fn)
			if err != nil {
				errs = append(errs, err)
			}
			continue
		}

		err := fn(&key, f, fv)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
