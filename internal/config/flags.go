package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"reflect"
)

const (
	usage = "usage"
)

func SetupFlagsForConfig(cmd *cobra.Command, defaultInstance Config) error {

	// We just want to get some defaults for the help screen, don't care about errors here
	_ = defaultInstance.SetDefaults()

	v := reflect.ValueOf(defaultInstance)

	// Dereference if necessary
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	return visitKeys(
		&Key{},
		t,
		v,
		func (key *Key, field reflect.StructField, defaultValue reflect.Value) error {

			usageInfo := field.Tag.Get(usage)
			defaultValueInterface := defaultValue.Interface()

			// Config keys use "." to separate different structs (E.g. serverConfig.Database.Host)
			// Flags still use "-"
			flagKey := key.String("-")
			configKey := key.String(".")

			var err error
			switch field.Type.Kind() {

			case reflect.Slice:
				err = setupSliceFlag(cmd, usageInfo, flagKey, field, defaultValue)
				break

			case reflect.String:
				cmd.Flags().String(flagKey, defaultValueInterface.(string), usageInfo)
				break

			case reflect.Int:
				cmd.Flags().Int(flagKey, defaultValueInterface.(int), usageInfo)
				break

			case reflect.Int64:
				cmd.Flags().Int64(flagKey, defaultValueInterface.(int64), usageInfo)
				break

			case reflect.Bool:
				cmd.Flags().Bool(flagKey, defaultValueInterface.(bool), usageInfo)
				break

			default:
				err = fmt.Errorf("kind %s is not a supported flag type", field.Type.Kind().String())
				break
			}

			if err != nil {
				return err
			}

			err = viper.BindPFlag(configKey, cmd.Flags().Lookup(flagKey))
			if err != nil {
				return err
			}

			return nil
		})
}

func setupSliceFlag(cmd *cobra.Command, usageInfo string, key string, field reflect.StructField, defaultValue reflect.Value) (err error) {

	var defaultSlice []reflect.Value
	for i := 0; i < defaultValue.Len(); i++ {
		defaultSlice = append(defaultSlice, defaultValue.Index(i))
	}

	switch field.Type.Elem().Kind() {

	case reflect.String:
		cmd.Flags().StringSlice(key, makeStringSliceFrom(defaultSlice), usageInfo)
		break

	default:
		err = fmt.Errorf("kind []%s is not a supported flag type", field.Type.Elem().Kind().String())
		break
	}

	return err
}

func makeStringSliceFrom(slice []reflect.Value) []string {
	var result []string
	for _, v := range slice {
		result = append(result, v.Interface().(string))
	}
	return result
}