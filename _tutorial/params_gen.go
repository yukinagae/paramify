// Code generated by paramify -type=Params; DO NOT EDIT.
package main

func NewParams(
	bool bool,
	string string,
	int int,
	int8 int8,
	int16 int16,
	int32 int32,
	int64 int64,
	uint uint,
	uint8 uint8,
	uint16 uint16,
	uint32 uint32,
	uint64 uint64,
	byte byte,
	rune rune,
	uintptr uintptr,
	float32 float32,
	float64 float64,
	formatType FormatType,
	address Address,
	content ContentUnion,
	map1 map[string]any,
	map2 map[string]int64,
	map3 map[string]*int64,
	map4 map[string]Address,
	map5 map[string]*Address,
	map6 map[string][]string,
	map7 map[string]map[string]float64,
	arrayString []string,
	array2Dstring [][]string,
	arrayAddress []Address,
	arrayPaddress []*Address,
	arrayContent []ContentUnion,
	arrayMap []map[string]any,
	emptyInterface any,
	any any,
	opts ...ParamsOption,
) Params {
	params := Params{
		Bool:           bool,
		String:         string,
		Int:            int,
		Int8:           int8,
		Int16:          int16,
		Int32:          int32,
		Int64:          int64,
		Uint:           uint,
		Uint8:          uint8,
		Uint16:         uint16,
		Uint32:         uint32,
		Uint64:         uint64,
		Byte:           byte,
		Rune:           rune,
		Uintptr:        uintptr,
		Float32:        float32,
		Float64:        float64,
		FormatType:     formatType,
		Address:        address,
		Content:        content,
		Map1:           map1,
		Map2:           map2,
		Map3:           map3,
		Map4:           map4,
		Map5:           map5,
		Map6:           map6,
		Map7:           map7,
		ArrayString:    arrayString,
		Array2DString:  array2Dstring,
		ArrayAddress:   arrayAddress,
		ArrayPAddress:  arrayPaddress,
		ArrayContent:   arrayContent,
		ArrayMap:       arrayMap,
		EmptyInterface: emptyInterface,
		Any:            any,
	}
	for _, opt := range opts {
		opt.apply(&params)
	}
	return params
}

type ParamsOption interface {
	apply(*Params)
}

var (
	_ ParamsOption = (*paramsOptionalStringOption)(nil)
	_ ParamsOption = (*paramsOptionalPStringOption)(nil)
	_ ParamsOption = (*paramsOptionalIntOption)(nil)
	_ ParamsOption = (*paramsOptionalPIntOption)(nil)
	_ ParamsOption = (*paramsOptionalPAddressOption)(nil)
	_ ParamsOption = (*paramsOptionalFormatTypeOption)(nil)
	_ ParamsOption = (*paramsOptionalPFormatTypeOption)(nil)
	_ ParamsOption = (*paramsOptionalContentOption)(nil)
	_ ParamsOption = (*paramsOptionalMapOption)(nil)
	_ ParamsOption = (*paramsOptionalArrayOption)(nil)
	_ ParamsOption = (*paramsOptionalEmptyInterfaceOption)(nil)
	_ ParamsOption = (*paramsOptionalAnyOption)(nil)
)

type (
	paramsOptionalStringOption         string
	paramsOptionalPStringOption        string
	paramsOptionalIntOption            int
	paramsOptionalPIntOption           int
	paramsOptionalPAddressOption       Address
	paramsOptionalFormatTypeOption     FormatType
	paramsOptionalPFormatTypeOption    FormatType
	paramsOptionalContentOption        struct{ ContentUnion }
	paramsOptionalMapOption            map[string]any
	paramsOptionalArrayOption          []string
	paramsOptionalEmptyInterfaceOption struct{ any }
	paramsOptionalAnyOption            struct{ any }
)

func WithParamsOptionalString(optionalString string) ParamsOption {
	return paramsOptionalStringOption(optionalString)
}

func WithParamsOptionalPString(optionalPstring string) ParamsOption {
	return paramsOptionalPStringOption(optionalPstring)
}

func WithParamsOptionalInt(optionalInt int) ParamsOption {
	return paramsOptionalIntOption(optionalInt)
}

func WithParamsOptionalPInt(optionalPint int) ParamsOption {
	return paramsOptionalPIntOption(optionalPint)
}

func WithParamsOptionalPAddress(optionalPaddress Address) ParamsOption {
	return paramsOptionalPAddressOption(optionalPaddress)
}

func WithParamsOptionalFormatType(optionalFormatType FormatType) ParamsOption {
	return paramsOptionalFormatTypeOption(optionalFormatType)
}

func WithParamsOptionalPFormatType(optionalPformatType FormatType) ParamsOption {
	return paramsOptionalPFormatTypeOption(optionalPformatType)
}

func WithParamsOptionalContent(optionalContent ContentUnion) ParamsOption {
	return paramsOptionalContentOption{optionalContent}
}

func WithParamsOptionalMap(optionalMap map[string]any) ParamsOption {
	return paramsOptionalMapOption(optionalMap)
}

func WithParamsOptionalArray(optionalArray []string) ParamsOption {
	return paramsOptionalArrayOption(optionalArray)
}

func WithParamsOptionalEmptyInterface(optionalEmptyInterface any) ParamsOption {
	return paramsOptionalEmptyInterfaceOption{optionalEmptyInterface}
}

func WithParamsOptionalAny(optionalAny any) ParamsOption {
	return paramsOptionalAnyOption{optionalAny}
}

func (o paramsOptionalStringOption) apply(p *Params) {
	p.OptionalString = string(o)
}

func (o paramsOptionalPStringOption) apply(p *Params) {
	optionalPstring := string(o)
	p.OptionalPString = &optionalPstring
}

func (o paramsOptionalIntOption) apply(p *Params) {
	p.OptionalInt = int(o)
}

func (o paramsOptionalPIntOption) apply(p *Params) {
	optionalPint := int(o)
	p.OptionalPInt = &optionalPint
}

func (o paramsOptionalPAddressOption) apply(p *Params) {
	optionalPaddress := Address(o)
	p.OptionalPAddress = &optionalPaddress
}

func (o paramsOptionalFormatTypeOption) apply(p *Params) {
	p.OptionalFormatType = FormatType(o)
}

func (o paramsOptionalPFormatTypeOption) apply(p *Params) {
	optionalPformatType := FormatType(o)
	p.OptionalPFormatType = &optionalPformatType
}

func (o paramsOptionalContentOption) apply(p *Params) {
	p.OptionalContent = ContentUnion(o.ContentUnion)
}

func (o paramsOptionalMapOption) apply(p *Params) {
	p.OptionalMap = map[string]any(o)
}

func (o paramsOptionalArrayOption) apply(p *Params) {
	p.OptionalArray = []string(o)
}

func (o paramsOptionalEmptyInterfaceOption) apply(p *Params) {
	p.OptionalEmptyInterface = any(o.any)
}

func (o paramsOptionalAnyOption) apply(p *Params) {
	p.OptionalAny = any(o.any)
}
