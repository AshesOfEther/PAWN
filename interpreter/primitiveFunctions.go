package interpreter

import (
	"fmt"
	"reflect"
	"unicode"
	"unicode/utf8"
)

var pawnValueType = reflect.TypeFor[PawnValue]()

func newPrimitiveFunction[F any](function F) PawnFunctionPrimitive {
	functionType := reflect.TypeFor[F]()
	if functionType.Kind() != reflect.Func {
		panic(fmt.Sprintf("expected function, got %#v", function))
	}
	if functionType.NumOut() > 1 {
		panic(fmt.Sprintf("expected function with 0 or 1 return values, got %v", functionType.NumOut()))
	}

	positionalArgumentCount := functionType.NumIn()
	var namedArgumentStructType reflect.Type = nil
	var namedArgumentTypes map[string]reflect.Type = nil
	positionalArgumentTypes := []reflect.Type{}
	acceptsBody := false

	if positionalArgumentCount != 0 && functionType.In(positionalArgumentCount - 1) == reflect.MapOf(reflect.TypeFor[string](), pawnValueType) {
		acceptsBody = true
		positionalArgumentCount--
	}
	
	if positionalArgumentCount != 0 {
		if namedArguments := functionType.In(positionalArgumentCount - 1); namedArguments.Kind() == reflect.Struct && namedArguments.Name() == "" {
			namedArgumentTypes = map[string]reflect.Type{}
			namedArgumentStructType = functionType.In(positionalArgumentCount - 1)
			positionalArgumentCount--
			
			for i := range namedArguments.NumField() {
				field := namedArguments.Field(i)
				if field.Type != pawnValueType && !(field.Type.Kind() == reflect.Pointer && field.Type.Implements(pawnValueType)) {
					panic(fmt.Sprintf("expected named argument types to be PawnValue or a pointer to an implementation of PawnValue, got argument '%s' of type %s", field.Name, field.Type.String()))
				}
				namedArgumentTypes[fromStructFieldCase(field.Name)] = field.Type.Elem()
			}
		}
	}
	
	for i := range positionalArgumentCount {
		argumentType := functionType.In(i)
		if !isValidPawnType(argumentType) {
			panic(fmt.Sprintf("expected positional argument types to be PawnValue or implement it, got argument #%d is of type %s", i, argumentType.String()))
		}
		positionalArgumentTypes = append(positionalArgumentTypes, argumentType)
	}

	return PawnFunctionPrimitive{&PawnFunctionPrimitiveInner{
		reflect.ValueOf(function),
		positionalArgumentTypes,
		namedArgumentStructType,
		namedArgumentTypes,
		acceptsBody,
	}}
}

func isValidPawnType(t reflect.Type) bool {
	return t == pawnValueType || t.Implements(pawnValueType)
}

func getTypeName(t reflect.Type) string {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	switch t {
		case reflect.TypeFor[PawnObject]():
			return "object"
		case reflect.TypeFor[PawnBoolean]():
			return "boolean"
		case reflect.TypeFor[PawnInt]():
			return "integer"
		case reflect.TypeFor[PawnFloat]():
			return "float"
		case reflect.TypeFor[PawnString]():
			return "string"
		case reflect.TypeFor[PawnList]():
			return "list"
		case reflect.TypeFor[PawnFunctionPrimitive](), reflect.TypeFor[PawnFunctionUser]():
			return "function"
		case reflect.TypeFor[PawnValue]():
			return "any"
		default:
			panic("tried to call `getTypeName` with a non-PawnValue type")
	}
}

func toStructFieldCase(s string) string {
	firstChar, byteLength := utf8.DecodeRuneInString(s)
	return fmt.Sprintf("%v%s", string(unicode.ToUpper(firstChar)), s[byteLength:])
}

func fromStructFieldCase(s string) string {
	firstChar, byteLength := utf8.DecodeRuneInString(s)
	return fmt.Sprintf("%v%s", string(unicode.ToLower(firstChar)), s[byteLength:])
}
