// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-08-29 10:08:04.85326469 -0700 PDT m=+1.565524498

package configs

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsDataCatalogConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementDataCatalogConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsDataCatalogConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookDataCatalogConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementDataCatalogConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_DataCatalogConfig(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookDataCatalogConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func testDecodeJson_DataCatalogConfig(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_DataCatalogConfig(val, result))
}

func testDecodeSlice_DataCatalogConfig(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_DataCatalogConfig(vStringSlice, result))
}

func TestDataCatalogConfig_GetPFlagSet(t *testing.T) {
	val := DataCatalogConfig{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestDataCatalogConfig_SetFlags(t *testing.T) {
	actual := DataCatalogConfig{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_storage-prefix", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("storage-prefix"); err == nil {
				assert.Equal(t, *new(string), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			cmdFlags.Set("storage-prefix", "1")
			if vString, err := cmdFlags.GetString("storage-prefix"); err == nil {
				testDecodeJson_DataCatalogConfig(t, fmt.Sprintf("%v", vString), &actual.StoragePrefix)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_metrics-scope", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("metrics-scope"); err == nil {
				assert.Equal(t, *new(string), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			cmdFlags.Set("metrics-scope", "1")
			if vString, err := cmdFlags.GetString("metrics-scope"); err == nil {
				testDecodeJson_DataCatalogConfig(t, fmt.Sprintf("%v", vString), &actual.MetricsScope)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_profiler-port", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("profiler-port"); err == nil {
				assert.Equal(t, *new(int), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			cmdFlags.Set("profiler-port", "1")
			if vInt, err := cmdFlags.GetInt("profiler-port"); err == nil {
				testDecodeJson_DataCatalogConfig(t, fmt.Sprintf("%v", vInt), &actual.ProfilerPort)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
