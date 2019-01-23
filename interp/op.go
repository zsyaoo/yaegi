package interp

// Code generated by 'go run ../cmd/genop/genop.go'. DO NOT EDIT.

import "reflect"

// Arithmetic operators

func add(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.String:
		v0 := genValue(n.child[0])
		v1 := genValue(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetString(v0(f).String() + v1(f).String())
			return next
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) + v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) + v1(f))
			return next
		}
	case reflect.Float32, reflect.Float64:
		v0 := genValueFloat(n.child[0])
		v1 := genValueFloat(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetFloat(v0(f) + v1(f))
			return next
		}
	case reflect.Complex64, reflect.Complex128:
		v0 := genValue(n.child[0])
		v1 := genValue(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetComplex(v0(f).Complex() + v1(f).Complex())
			return next
		}
	}
}

func and(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) & v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) & v1(f))
			return next
		}
	}
}

func andnot(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) &^ v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) &^ v1(f))
			return next
		}
	}
}

func mul(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) * v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) * v1(f))
			return next
		}
	case reflect.Float32, reflect.Float64:
		v0 := genValueFloat(n.child[0])
		v1 := genValueFloat(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetFloat(v0(f) * v1(f))
			return next
		}
	case reflect.Complex64, reflect.Complex128:
		v0 := genValue(n.child[0])
		v1 := genValue(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetComplex(v0(f).Complex() * v1(f).Complex())
			return next
		}
	}
}

func or(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) | v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) | v1(f))
			return next
		}
	}
}

func quo(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) / v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) / v1(f))
			return next
		}
	case reflect.Float32, reflect.Float64:
		v0 := genValueFloat(n.child[0])
		v1 := genValueFloat(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetFloat(v0(f) / v1(f))
			return next
		}
	case reflect.Complex64, reflect.Complex128:
		v0 := genValue(n.child[0])
		v1 := genValue(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetComplex(v0(f).Complex() / v1(f).Complex())
			return next
		}
	}
}

func rem(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) % v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) % v1(f))
			return next
		}
	}
}

func shl(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) << v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) << v1(f))
			return next
		}
	}
}

func shr(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) >> v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) >> v1(f))
			return next
		}
	}
}

func sub(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) - v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) - v1(f))
			return next
		}
	case reflect.Float32, reflect.Float64:
		v0 := genValueFloat(n.child[0])
		v1 := genValueFloat(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetFloat(v0(f) - v1(f))
			return next
		}
	case reflect.Complex64, reflect.Complex128:
		v0 := genValue(n.child[0])
		v1 := genValue(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetComplex(v0(f).Complex() - v1(f).Complex())
			return next
		}
	}
}

func xor(n *Node) {
	i := n.findex
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetInt(v0(f) ^ v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			f.data[i].SetUint(v0(f) ^ v1(f))
			return next
		}
	}
}

// Assign operators

func addAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.String:
		v0 := genValue(n.child[0])
		v1 := genValue(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetString(v0(f).String() + v1(f).String())
			return next
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) + v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) + v1(f))
			return next
		}
	case reflect.Float32, reflect.Float64:
		v0 := genValueFloat(n.child[0])
		v1 := genValueFloat(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetFloat(v0(f) + v1(f))
			return next
		}
	case reflect.Complex64, reflect.Complex128:
		v0 := genValue(n.child[0])
		v1 := genValue(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetComplex(v0(f).Complex() + v1(f).Complex())
			return next
		}
	}
}

func andAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) & v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) & v1(f))
			return next
		}
	}
}

func andnotAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) &^ v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) &^ v1(f))
			return next
		}
	}
}

func mulAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) * v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) * v1(f))
			return next
		}
	case reflect.Float32, reflect.Float64:
		v0 := genValueFloat(n.child[0])
		v1 := genValueFloat(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetFloat(v0(f) * v1(f))
			return next
		}
	case reflect.Complex64, reflect.Complex128:
		v0 := genValue(n.child[0])
		v1 := genValue(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetComplex(v0(f).Complex() * v1(f).Complex())
			return next
		}
	}
}

func orAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) | v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) | v1(f))
			return next
		}
	}
}

func quoAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) / v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) / v1(f))
			return next
		}
	case reflect.Float32, reflect.Float64:
		v0 := genValueFloat(n.child[0])
		v1 := genValueFloat(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetFloat(v0(f) / v1(f))
			return next
		}
	case reflect.Complex64, reflect.Complex128:
		v0 := genValue(n.child[0])
		v1 := genValue(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetComplex(v0(f).Complex() / v1(f).Complex())
			return next
		}
	}
}

func remAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) % v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) % v1(f))
			return next
		}
	}
}

func shlAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) << v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) << v1(f))
			return next
		}
	}
}

func shrAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) >> v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) >> v1(f))
			return next
		}
	}
}

func subAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) - v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) - v1(f))
			return next
		}
	case reflect.Float32, reflect.Float64:
		v0 := genValueFloat(n.child[0])
		v1 := genValueFloat(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetFloat(v0(f) - v1(f))
			return next
		}
	case reflect.Complex64, reflect.Complex128:
		v0 := genValue(n.child[0])
		v1 := genValue(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetComplex(v0(f).Complex() - v1(f).Complex())
			return next
		}
	}
}

func xorAssign(n *Node) {
	next := getExec(n.tnext)
	typ := n.typ.TypeOf()
	value := genValue(n.child[0])

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v0 := genValueInt(n.child[0])
		v1 := genValueInt(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetInt(v0(f) ^ v1(f))
			return next
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v0 := genValueUint(n.child[0])
		v1 := genValueUint(n.child[1])
		n.exec = func(f *Frame) Builtin {
			value(f).SetUint(v0(f) ^ v1(f))
			return next
		}
	}
}
