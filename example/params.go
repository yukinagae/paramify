package main

//go:generate go run ../paramify.go -type=Params
type Params struct {
	//
	// required
	//
	Bool    bool    `json:"bool"`
	String  string  `json:"string"`
	Int     int     `json:"int"`
	Int8    int8    `json:"int8"`
	Int16   int16   `json:"int16"`
	Int32   int32   `json:"int32"`
	Int64   int64   `json:"int64"`
	Uint    uint    `json:"uint"`
	Uint8   uint8   `json:"uint8"`
	Uint16  uint16  `json:"uint16"`
	Uint32  uint32  `json:"uint32"`
	Uint64  uint64  `json:"uint64"`
	Byte    byte    `json:"byte"`
	Rune    rune    `json:"rune"`
	Uintptr uintptr `json:"uintptr"`
	Float32 float32 `json:"float32"`
	Float64 float64 `json:"float64"`
	// Complex64       complex64    `json:"complex64"`
	// Complex128      complex128   `json:"complex128"`
	FormatType     FormatType                    `json:"format_type"`
	Address        Address                       `json:"address"`
	Content        ContentUnion                  `json:"content"`
	Map1           map[string]any                `json:"map1"`
	Map2           map[string]int64              `json:"map2"`
	Map3           map[string]*int64             `json:"map3"`
	Map4           map[string]Address            `json:"map4"`
	Map5           map[string]*Address           `json:"map5"`
	Map6           map[string][]string           `json:"map6"`
	Map7           map[string]map[string]float64 `json:"map7"`
	ArrayString    []string                      `json:"array_string"`
	Array2DString  [][]string                    `json:"array_2d_string"`
	ArrayAddress   []Address                     `json:"array_address"`
	ArrayPAddress  []*Address                    `json:"array_p_address"`
	ArrayContent   []ContentUnion                `json:"array_content"`
	ArrayMap       []map[string]interface{}      `json:"array_map"`
	EmptyInterface interface{}                   `json:"empty_interface"`
	Any            any                           `json:"any"`
	//
	// optional
	//
	OptionalString         string         `json:"optional_string,omitempty"`
	OptionalPString        *string        `json:"optional_p_string,omitempty"`
	OptionalInt            int            `json:"optional_int,omitempty"`
	OptionalPInt           *int           `json:"optional_p_int,omitempty"`
	OptionalPAddress       *Address       `json:"optional_p_address,omitempty"`
	OptionalFormatType     FormatType     `json:"optional_format_type,omitempty"`
	OptionalPFormatType    *FormatType    `json:"optional_p_format_type,omitempty"`
	OptionalContent        ContentUnion   `json:"optional_content,omitempty"`
	OptionalMap            map[string]any `json:"optional_map,omitempty"`
	OptionalArray          []string       `json:"optional_array,omitempty"`
	OptionalEmptyInterface interface{}    `json:"optional_empty_interface,omitempty"`
	OptionalAny            any            `json:"optional_any,omitempty"`
}

type FormatType string

const (
	FormatTypeA FormatType = "A"
)

type Address struct {
	Street string `json:"street"`
}

type (
	// ContentUnion = ContentString | ContentArray | ContentStruct
	ContentUnion  interface{ IsContentUnion() }
	ContentString string
	ContentArray  []string
	ContentStruct struct {
		Name string `json:"name"`
	}
)

var (
	_ ContentUnion = (*ContentString)(nil)
	_ ContentUnion = (*ContentArray)(nil)
	_ ContentUnion = (*ContentStruct)(nil)
)

func (ContentString) IsContentUnion() {}
func (ContentArray) IsContentUnion()  {}
func (ContentStruct) IsContentUnion() {}
