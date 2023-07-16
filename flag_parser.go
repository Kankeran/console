package console

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Input interface {
	ParseInt(variable *int, key string)
	ParseInt64(variable *int64, key string)
	ParseUint(variable *uint, key string)
	ParseUint64(variable *uint64, key string)
	ParseBool(variable *bool, key string)
	ParseString(variable *string, key string)
	ParseFloat64(variable *float64, key string)
	ParseDuration(variable *time.Duration, key string)

	ParseIntSlice(variable *[]int, key string)
	ParseInt64Slice(variable *[]int64, key string)
	ParseUintSlice(variable *[]uint, key string)
	ParseUint64Slice(variable *[]uint64, key string)
	ParseStringSlice(variable *[]string, key string)
	ParseFloat64Slice(variable *[]float64, key string)
	ParseDurationSlice(variable *[]time.Duration, key string)
}

type Output interface {
	io.Writer
}

type FlagParser struct {
	flags     map[string][]string
	flagsInfo map[string]commonFlagInfo
}

func (f *FlagParser) getFlagValueData(key string) Value {
	return f.flagsInfo[key].valueData
}

func (f *FlagParser) writeWarning(key, curType, expectedType string) {
	fmt.Printf("[WARNING] default value for '%s' is of type '%s', should be '%s'\n", key, curType, expectedType)
}

func (f *FlagParser) ParseInt(variable *int, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().(int)
		if !ok {
			f.writeWarning(key, defValData.Type(), "int")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	i64, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}
	*variable = int(i64)
}

func (f *FlagParser) ParseInt64(variable *int64, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().(int64)
		if !ok {
			f.writeWarning(key, defValData.Type(), "int64")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	i64, err := strconv.ParseInt(vals[0], 10, 64)
	if err != nil {
		panic(err)
	}
	*variable = i64
}

func (f *FlagParser) ParseUint(variable *uint, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().(uint)
		if !ok {
			f.writeWarning(key, defValData.Type(), "uint")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	u64, err := strconv.ParseUint(vals[0], 10, 64)
	if err != nil {
		panic(err)
	}
	*variable = uint(u64)
}

func (f *FlagParser) ParseUint64(variable *uint64, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().(uint64)
		if !ok {
			f.writeWarning(key, defValData.Type(), "uint64")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	u64, err := strconv.ParseUint(vals[0], 10, 64)
	if err != nil {
		panic(err)
	}
	*variable = u64
}

func (f *FlagParser) ParseBool(variable *bool, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().(bool)
		if !ok {
			f.writeWarning(key, defValData.Type(), "bool")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	b, err := strconv.ParseBool(vals[0])
	if err != nil {
		panic(err)
	}
	*variable = b
}

func (f *FlagParser) ParseString(variable *string, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().(string)
		if !ok {
			f.writeWarning(key, defValData.Type(), "string")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	*variable = vals[0]
}

func (f *FlagParser) ParseFloat64(variable *float64, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().(float64)
		if !ok {
			f.writeWarning(key, defValData.Type(), "float64")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	f64, err := strconv.ParseFloat(vals[0], 64)
	if err != nil {
		panic(err)
	}
	*variable = f64
}

func (f *FlagParser) ParseDuration(variable *time.Duration, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().(time.Duration)
		if !ok {
			f.writeWarning(key, defValData.Type(), "time.Duration")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	d, err := time.ParseDuration(vals[0])
	if err != nil {
		panic(err)
	}
	*variable = d
}

func (f *FlagParser) ParseIntSlice(variable *[]int, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().([]int)
		if !ok {
			f.writeWarning(key, defValData.Type(), "[]int")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	for _, v := range vals {
		i64, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		*variable = append(*variable, int(i64))
	}
}

func (f *FlagParser) ParseInt64Slice(variable *[]int64, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().([]int64)
		if !ok {
			f.writeWarning(key, defValData.Type(), "[]int64")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	for _, v := range vals {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		*variable = append(*variable, i64)
	}
}

func (f *FlagParser) ParseUintSlice(variable *[]uint, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().([]uint)
		if !ok {
			f.writeWarning(key, defValData.Type(), "[]uint")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	for _, v := range vals {
		u64, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			panic(err)
		}
		*variable = append(*variable, uint(u64))
	}
}

func (f *FlagParser) ParseUint64Slice(variable *[]uint64, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().([]uint64)
		if !ok {
			f.writeWarning(key, defValData.Type(), "[]uint64")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	for _, v := range vals {
		u64, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			panic(err)
		}
		*variable = append(*variable, u64)
	}
}

func (f *FlagParser) ParseStringSlice(variable *[]string, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().([]string)
		if !ok {
			f.writeWarning(key, defValData.Type(), "[]string")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	*variable = append(*variable, vals...)
}

func (f *FlagParser) ParseFloat64Slice(variable *[]float64, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().([]float64)
		if !ok {
			f.writeWarning(key, defValData.Type(), "[]float64")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	for _, v := range vals {
		f64, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
		*variable = append(*variable, f64)
	}
}

func (f *FlagParser) ParseDurationSlice(variable *[]time.Duration, key string) {
	vals := f.flags[key]
	defValData := f.getFlagValueData(key)
	if defValData.Get() != nil {
		typedVal, ok := defValData.Get().([]time.Duration)
		if !ok {
			f.writeWarning(key, defValData.Type(), "[]time.Duration")
		}
		if len(vals) == 0 {
			*variable = typedVal
			return
		}
	}
	for _, v := range vals {
		d, err := time.ParseDuration(v)
		if err != nil {
			panic(err)
		}
		*variable = append(*variable, d)
	}
}
