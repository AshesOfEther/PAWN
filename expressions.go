package main

import ()

func add(Value1 PawnValue, Value2 PawnValue) PawnValue {
	var todo PawnString

	switch v1 := Value1.(type) {
	case PawnInt:
		switch v2 := Value2.(type) {
		case PawnInt:
			return PawnInt{v1.v + v2.v}
		case PawnFloat:
			return PawnFloat{float64(v1.v) + v2.v}
		default:
			return todo
		}

	case PawnFloat:
		switch v2 := Value2.(type) {
		case PawnInt:
			return PawnFloat{v1.v + float64(v2.v)}
		case PawnFloat:
			return PawnFloat{v1.v + v2.v}
		default:
			return todo
		}
	default:
		return todo
	}
}
