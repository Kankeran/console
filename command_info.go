package console

import (
	"fmt"
	"time"
)

type Command interface {
	RequiredInt(name, description string) Command
	RequiredInt64(name, description string) Command
	RequiredUint(name, description string) Command
	RequiredUint64(name, description string) Command
	RequiredBool(name, description string) Command
	RequiredString(name, description string) Command
	RequiredFloat64(name, description string) Command
	RequiredDuration(name, description string) Command
	RequiredSliceInt(name, description string) Command
	RequiredSliceInt64(name, description string) Command
	RequiredSliceUint(name, description string) Command
	RequiredSliceUint64(name, description string) Command
	RequiredSliceBool(name, description string) Command
	RequiredSliceString(name, description string) Command
	RequiredSliceFloat64(name, description string) Command
	RequiredSliceDuration(name, description string) Command

	OptionalInt(name, description string, value int) Command
	OptionalInt64(name, description string, value int64) Command
	OptionalUint(name, description string, value uint) Command
	OptionalUint64(name, description string, value uint64) Command
	OptionalBool(name, description string, value bool) Command
	OptionalString(name, description string, value string) Command
	OptionalFloat64(name, description string, value float64) Command
	OptionalDuration(name, description string, value time.Duration) Command
	OptionalSliceInt(name, description string, value []int) Command
	OptionalSliceInt64(name, description string, value []int64) Command
	OptionalSliceUint(name, description string, value []uint) Command
	OptionalSliceUint64(name, description string, value []uint64) Command
	OptionalSliceBool(name, description string, value []bool) Command
	OptionalSliceString(name, description string, value []string) Command
	OptionalSliceFloat64(name, description string, value []float64) Command
	OptionalSliceDuration(name, description string, value []time.Duration) Command
}

var (
	commandInfoMap = make(map[string]commonCommandInfo)
)

func RegisterCommand(name, description string, callback func(Input, Output) error) *commonCommandInfo {

	c := commonCommandInfo{
		Name:            name,
		Description:     description,
		ExecuteCallback: callback,
		flagsInfo:       make(map[string]commonFlagInfo),
	}
	commandInfoMap[c.Name] = c

	return &c
}

type commonCommandInfo struct {
	Name            string
	Description     string
	ExecuteCallback func(Input, Output) error
	flagsInfo       map[string]commonFlagInfo
}

type commonFlagInfo struct {
	isRequired  bool
	description string
	valueData   Value
}

func (c *commonCommandInfo) requiredFlagInfo(name, description string, valueData TypeOnlyValue) *commonCommandInfo {
	c.flagsInfo[name] = commonFlagInfo{
		isRequired:  true,
		description: description,
		valueData:   valueData,
	}

	return c
}

func (c *commonCommandInfo) RequiredInt(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "int"})
}

func (c *commonCommandInfo) RequiredInt64(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "int64"})
}
func (c *commonCommandInfo) RequiredUint(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "uint"})
}

func (c *commonCommandInfo) RequiredUint64(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "uint64"})
}

func (c *commonCommandInfo) RequiredBool(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "bool"})
}

func (c *commonCommandInfo) RequiredString(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "string"})
}

func (c *commonCommandInfo) RequiredFloat64(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "float64"})
}

func (c *commonCommandInfo) RequiredDuration(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "time.Duration"})
}

func (c *commonCommandInfo) RequiredSliceInt(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "[]int"})
}

func (c *commonCommandInfo) RequiredSliceInt64(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "[]int64"})
}

func (c *commonCommandInfo) RequiredSliceUint(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "[]uint"})
}

func (c *commonCommandInfo) RequiredSliceUint64(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "[]uint64"})
}

func (c *commonCommandInfo) RequiredSliceBool(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "[]bool"})
}

func (c *commonCommandInfo) RequiredSliceString(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "[]string"})
}

func (c *commonCommandInfo) RequiredSliceFloat64(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "[]float64"})
}

func (c *commonCommandInfo) RequiredSliceDuration(name, description string) Command {
	return c.requiredFlagInfo(name, description, typeOnlyValueData{dataType: "[]time.Duration"})
}

func (c *commonCommandInfo) optionalFlagInfo(name, description string, valueData Value) *commonCommandInfo {
	c.flagsInfo[name] = commonFlagInfo{
		isRequired:  false,
		description: description,
		valueData:   valueData,
	}

	return c
}

func (c *commonCommandInfo) OptionalInt(name, description string, value int) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "int", dataValue: value})
}

func (c *commonCommandInfo) OptionalInt64(name, description string, value int64) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "int64", dataValue: value})
}

func (c *commonCommandInfo) OptionalUint(name, description string, value uint) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "uint", dataValue: value})
}

func (c *commonCommandInfo) OptionalUint64(name, description string, value uint64) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "uint64", dataValue: value})
}

func (c *commonCommandInfo) OptionalBool(name, description string, value bool) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "bool", dataValue: value})
}

func (c *commonCommandInfo) OptionalString(name, description string, value string) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "string", dataValue: value})
}

func (c *commonCommandInfo) OptionalFloat64(name, description string, value float64) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "float64", dataValue: value})
}

func (c *commonCommandInfo) OptionalDuration(name, description string, value time.Duration) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "time.Duration", dataValue: value})
}

func (c *commonCommandInfo) OptionalSliceInt(name, description string, value []int) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "[]int", dataValue: value})
}

func (c *commonCommandInfo) OptionalSliceInt64(name, description string, value []int64) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "[]int64", dataValue: value})
}

func (c *commonCommandInfo) OptionalSliceUint(name, description string, value []uint) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "[]uint", dataValue: value})
}

func (c *commonCommandInfo) OptionalSliceUint64(name, description string, value []uint64) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "[]uint64", dataValue: value})
}

func (c *commonCommandInfo) OptionalSliceBool(name, description string, value []bool) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "[]bool", dataValue: value})
}

func (c *commonCommandInfo) OptionalSliceString(name, description string, value []string) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "[]string", dataValue: value})
}

func (c *commonCommandInfo) OptionalSliceFloat64(name, description string, value []float64) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "[]float64", dataValue: value})
}

func (c *commonCommandInfo) OptionalSliceDuration(name, description string, value []time.Duration) Command {
	return c.optionalFlagInfo(name, description, fullValueData{dataType: "[]time.Duration", dataValue: value})
}

type TypeOnlyValue interface {
	Value
}

type Value interface {
	String() string
	Type() string
	Get() any
}

type typeOnlyValueData struct {
	dataType string
}

func (v typeOnlyValueData) String() string {
	return fmt.Sprintf("type: %s", v.dataType)
}

func (v typeOnlyValueData) Type() string {
	return v.dataType
}

func (v typeOnlyValueData) Get() any {
	return nil
}

type fullValueData struct {
	dataType  string
	dataValue any
}

func (v fullValueData) String() string {
	return fmt.Sprintf("type: %s", v.dataType)
}

func (v fullValueData) Type() string {
	return v.dataType
}

func (v fullValueData) Get() any {
	return v.dataValue
}
