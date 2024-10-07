package paramify

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Field represents a struct field with additional metadata.
type Field struct {
	Name        string // The name of the field.
	Type        string // The type of the field.
	IsPointer   bool   // Indicates if the field is a pointer.
	IsInterface bool   // Indicates if the field is an interface.
	IsAny       bool   // Indicates if the field is an any type.
}

// Fields represents the required and optional fields of a struct.
type Fields struct {
	Required []Field // Required fields.
	Optional []Field // Optional fields.
}

// ParsePackage loads and returns the package in the specified directory.
func ParsePackage(directory string) (*packages.Package, error) {
	conf := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedCompiledGoFiles | packages.NeedFiles,
		Dir:  directory,
	}

	// Load the packages based on the configuration
	pkgs, err := packages.Load(conf)
	if err != nil {
		return nil, fmt.Errorf("couldn't load packages in %s: %v", directory, err)
	}

	// Check if any packages were loaded
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no packages found in directory: %s", directory)
	}

	// Return the first package found
	return pkgs[0], nil
}

// ValuesOfType returns the required and optional fields of the specified type in the given package.
func ValuesOfType(pkg *packages.Package, typeName string) (*Fields, error) {
	var (
		requiredFields []Field
		optionalFields []Field
	)
	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(node ast.Node) bool {
			typeSpec, ok := node.(*ast.TypeSpec)
			if !ok || typeSpec.Name.Name != typeName {
				return true
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				return true
			}

			for _, field := range structType.Fields.List {
				value := convertField(field)

				// determine if the field is required or optional based on the struct tag omitempty option
				if isOmitEmpty(field.Tag) {
					optionalFields = append(optionalFields, value)
				} else {
					requiredFields = append(requiredFields, value)
				}
			}

			return false // Stop inspecting further nodes in this file
		})
	}

	if len(requiredFields) == 0 && len(optionalFields) == 0 {
		return nil, fmt.Errorf("no values defined for type: %s", typeName)
	}

	return &Fields{
		Required: requiredFields,
		Optional: optionalFields,
	}, nil
}

// convertField converts an AST field to a Field struct with additional metadata.
func convertField(field *ast.Field) Field {
	value := Field{
		Name: field.Names[0].Name,
		Type: resolveType(field.Type),
	}

	// Check if the field is a pointer
	if isPointer(field.Type) {
		value.IsPointer = true
	}

	// Check if the field is an interface
	if isInterface(field.Type) {
		value.IsInterface = true
	}

	// Check if the field is an any type
	if value.Type == "any" {
		value.IsAny = true
	}

	return value
}

// isPointer checks if the given AST expression is a pointer type.
func isPointer(expr ast.Expr) bool {
	_, ok := expr.(*ast.StarExpr)
	return ok
}

// isInterface checks if the given AST expression is an interface type.
func isInterface(expr ast.Expr) bool {
	switch t := expr.(type) {
	case *ast.Ident:
		if t.Obj != nil && t.Obj.Decl != nil {
			if decl, ok := t.Obj.Decl.(*ast.TypeSpec); ok {
				_, ok := decl.Type.(*ast.InterfaceType)
				return ok
			}
		}
	}
	return false
}

// resolveType returns the string representation of an AST expression type.
func resolveType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		// Basic types (i.e. string, int, Address)
		return t.Name
	case *ast.StarExpr:
		// Pointer (i.e. *string, *Address)
		return "*" + resolveType(t.X)
	case *ast.SelectorExpr:
		// Types with package (i.e. pkg.Type)
		return resolveType(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		// Array (i.e. []string)
		return "[]" + resolveType(t.Elt)
	case *ast.MapType:
		// Map (i.e. map[string]any)
		return "map[" + resolveType(t.Key) + "]" + resolveType(t.Value)
	case *ast.InterfaceType:
		// Interface (i.e. interface{})
		return "any"
	default:
		// Handle unexpected types
		return fmt.Sprintf("unknown type: %T", expr)
	}
}

// isOmitEmpty checks if the given struct tag has the "omitempty" option for the "json" key.
func isOmitEmpty(tag *ast.BasicLit) bool {
	if tag == nil {
		return false
	}

	tagContent := strings.Trim(tag.Value, "`")
	structTag := reflect.StructTag(tagContent)
	jsonTag := structTag.Get("json")
	if jsonTag == "" {
		return false
	}

	jsonTagValues := strings.Split(jsonTag, ",")
	return len(jsonTagValues) > 1 && jsonTagValues[1] == "omitempty"
}
