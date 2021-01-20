package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// Default parameter namespace
const (
	DefaultMaxValueLength = 10000
)

// Parameter store keys
var (
	ParamStoreKeyMaxValueLength = []byte("MaxValueLength")
)

// String implements stringer interface
func (params Params) String() string {
	out, _ := yaml.Marshal(params)
	return string(out)
}

// ParamKeyTable for slashing module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams create a new Params object
func NewParams(
	maxValueLength uint32,
) Params {
	return Params{
		MaxValueLength: maxValueLength,
	}
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyMaxValueLength, &p.MaxValueLength, validateMaxValueLength),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultMaxValueLength,
	)
}

func validateMaxValueLength(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("max value length must be greater than zero: %d", v)
	}

	return nil
}
