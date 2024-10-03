package parser

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestResolveType(t *testing.T) {
	tests := []struct {
		name  string
		input ast.Expr
		want  string
	}{
		{
			name:  "primitive: bool",
			input: &ast.Ident{Name: "bool"},
			want:  "bool",
		},
		{
			name:  "primitive: string",
			input: &ast.Ident{Name: "string"},
			want:  "string",
		},
		{
			name:  "primitive: int",
			input: &ast.Ident{Name: "int"},
			want:  "int",
		},
		{
			name:  "primitive: int8",
			input: &ast.Ident{Name: "int8"},
			want:  "int8",
		},
		{
			name:  "primitive: int16",
			input: &ast.Ident{Name: "int16"},
			want:  "int16",
		},
		{
			name:  "primitive: int32",
			input: &ast.Ident{Name: "int32"},
			want:  "int32",
		},
		{
			name:  "primitive: int64",
			input: &ast.Ident{Name: "int64"},
			want:  "int64",
		},
		{
			name:  "primitive: uint",
			input: &ast.Ident{Name: "uint"},
			want:  "uint",
		},
		{
			name:  "primitive: uint8",
			input: &ast.Ident{Name: "uint8"},
			want:  "uint8",
		},
		{
			name:  "primitive: uint16",
			input: &ast.Ident{Name: "uint16"},
			want:  "uint16",
		},
		{
			name:  "primitive: uint32",
			input: &ast.Ident{Name: "uint32"},
			want:  "uint32",
		},
		{
			name:  "primitive: uint64",
			input: &ast.Ident{Name: "uint64"},
			want:  "uint64",
		},
		{
			name:  "primitive: uintptr",
			input: &ast.Ident{Name: "uintptr"},
			want:  "uintptr",
		},
		{
			name:  "primitive: byte",
			input: &ast.Ident{Name: "byte"},
			want:  "byte",
		},
		{
			name:  "primitive: rune",
			input: &ast.Ident{Name: "rune"},
			want:  "rune",
		},
		{
			name:  "primitive: float32",
			input: &ast.Ident{Name: "float32"},
			want:  "float32",
		},
		{
			name:  "primitive: float64",
			input: &ast.Ident{Name: "float64"},
			want:  "float64",
		},
		{
			name:  "primitive: complex64",
			input: &ast.Ident{Name: "complex64"},
			want:  "complex64",
		},
		{
			name:  "primitive: complex128",
			input: &ast.Ident{Name: "complex128"},
			want:  "complex128",
		},
		{
			name:  "alias: FormatType",
			input: &ast.Ident{Name: "FormatType"},
			want:  "FormatType",
		},
		{
			name:  "struct: Address",
			input: &ast.Ident{Name: "Address"},
			want:  "Address",
		},
		{
			name:  "interface: ContentUnion",
			input: &ast.Ident{Name: "ContentUnion"},
			want:  "ContentUnion",
		},
		{
			name:  "pointer: *string",
			input: &ast.StarExpr{X: &ast.Ident{Name: "string"}},
			want:  "*string",
		},
		{
			name:  "pointer: *int",
			input: &ast.StarExpr{X: &ast.Ident{Name: "int"}},
			want:  "*int",
		},
		{
			name:  "pointer: *Address",
			input: &ast.StarExpr{X: &ast.Ident{Name: "Address"}},
			want:  "*Address",
		},
		{
			name:  "array: []string",
			input: &ast.ArrayType{Elt: &ast.Ident{Name: "string"}},
			want:  "[]string",
		},
		{
			name:  "array: []Address",
			input: &ast.ArrayType{Elt: &ast.Ident{Name: "Address"}},
			want:  "[]Address",
		},
		{
			name:  "array: []*Address",
			input: &ast.ArrayType{Elt: &ast.StarExpr{X: &ast.Ident{Name: "Address"}}},
			want:  "[]*Address",
		},
		{
			name:  "map: map[string]any",
			input: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "any"}},
			want:  "map[string]any",
		},
		{
			name:  "map: map[string]int64",
			input: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "int64"}},
			want:  "map[string]int64",
		},
		{
			name:  "map: map[string]*int",
			input: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.StarExpr{X: &ast.Ident{Name: "int"}}},
			want:  "map[string]*int",
		},
		{
			name:  "map: map[string]map[string]float64",
			input: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "float64"}}},
			want:  "map[string]map[string]float64",
		},
		{
			name:  "map: map[string]Address",
			input: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "Address"}},
			want:  "map[string]Address",
		},
		{
			name:  "map: map[string]*Address",
			input: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.StarExpr{X: &ast.Ident{Name: "Address"}}},
			want:  "map[string]*Address",
		},
		{
			name:  "empty interface: interface{}",
			input: &ast.InterfaceType{},
			want:  "any",
		},
		{
			name:  "any: any",
			input: &ast.Ident{Name: "any"},
			want:  "any",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveType(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveType() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestIsPointer(t *testing.T) {
	tests := []struct {
		name  string
		input ast.Expr
		want  bool
	}{
		{
			name:  "should return false if input is not a pointer",
			input: &ast.Ident{},
			want:  false,
		},
		{
			name:  "should return false if input is a type with package",
			input: &ast.SelectorExpr{},
			want:  false,
		},
		{
			name:  "should return false if input is an array type",
			input: &ast.ArrayType{},
			want:  false,
		},
		{
			name:  "should return false if input is a map type",
			input: &ast.MapType{},
			want:  false,
		},
		{
			name:  "should return true if input is a pointer",
			input: &ast.StarExpr{},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPointer(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isPointer() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestIsInterface(t *testing.T) {
	tests := []struct {
		name  string
		input ast.Expr
		want  bool
	}{
		{
			name:  "should return false if input is not an interface",
			input: &ast.Ident{},
			want:  false,
		},
		{
			name:  "should return false if input is a pointer",
			input: &ast.StarExpr{},
			want:  false,
		},
		{
			name:  "should return false if input is a type with package",
			input: &ast.SelectorExpr{},
			want:  false,
		},
		{
			name:  "should return false if input is an array type",
			input: &ast.ArrayType{},
			want:  false,
		},
		{
			name:  "should return false if input is a map type",
			input: &ast.MapType{},
			want:  false,
		},
		{
			name:  "should return false if input is a struct type",
			input: &ast.StructType{},
			want:  false,
		},
		{
			name: "should return false if input is a struct type",
			input: &ast.Ident{
				//nolint:staticcheck
				Obj: &ast.Object{
					Decl: &ast.TypeSpec{
						Type: &ast.StructType{},
					},
				},
			},
			want: false,
		},
		{
			name: "should return true if input is an interface type",
			input: &ast.Ident{
				//nolint:staticcheck
				Obj: &ast.Object{
					Decl: &ast.TypeSpec{
						Type: &ast.InterfaceType{},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isInterface(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isInterface() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestIsOmitEmpty(t *testing.T) {
	tests := []struct {
		name  string
		input *ast.BasicLit
		want  bool
	}{
		{
			name: "should return false if tag is empty",
			input: &ast.BasicLit{
				Value: "",
			},
			want: false,
		},
		{
			name: "should return false if tag does not contain json",
			input: &ast.BasicLit{
				Value: "`validate:\"required\"`",
			},
			want: false,
		},
		{
			name: "should return false if tag does not contain json omitempty",
			input: &ast.BasicLit{
				Value: "`json:\"name\"`",
			},
			want: false,
		},
		{
			name: "should return true if tag contains json omitempty",
			input: &ast.BasicLit{
				Value: "`json:\"name,omitempty\"`",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isOmitEmpty(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isOmitEmpty() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
