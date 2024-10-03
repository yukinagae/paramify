package main

import (
	"encoding/json"
	"log"
	"unsafe"
)

func main() {
	var i int64
	params := NewParams(
		true,
		"string",
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
		uintptr(unsafe.Pointer(&i)),
		1.1,
		2.2,
		// 1i,
		// 2i,
		FormatTypeA,
		Address{"street"},
		ContentString("content string"),
		map[string]any{"key": "value"},
		map[string]int64{"key": 1},
		map[string]*int64{"key": &i},
		map[string]Address{"key": {"street"}},
		map[string]*Address{"key": {"street"}},
		map[string][]string{"key": {"value1", "value2"}},
		map[string]map[string]float64{"key": {"key": 1.1}},
		[]string{"value1", "value2"},
		[][]string{{"value1", "value2"}, {"value3", "value4"}},
		[]Address{{"street1"}, {"street2"}},
		[]*Address{{"street1"}, {"street2"}},
		[]ContentUnion{ContentString("content string"), ContentArray{"content array1", "content array2"}, ContentStruct{"content struct"}},
		[]map[string]interface{}{{"key": "value"}},
		interface{}(nil),
		"any",
		WithParamsOptionalString("optional string"),
		WithParamsOptionalPString("optional pointer string"),
		WithParamsOptionalInt(1),
		WithParamsOptionalPInt(2),
		WithParamsOptionalPAddress(Address{"optional pointer address"}),
		WithParamsOptionalFormatType(FormatTypeA),
		WithParamsOptionalPFormatType(FormatTypeA),
		WithParamsOptionalContent(ContentStruct{"content struct"}),
		WithParamsOptionalMap(map[string]any{"key": "value"}),
		WithParamsOptionalArray([]string{"value1", "value2"}),
		WithParamsOptionalEmptyInterface("empty interface"),
		WithParamsOptionalAny("any"),
	)
	log.Printf("💖params: %#v", params)
	bs, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	log.Printf("💖json: %+v\n", string(bs))
}
