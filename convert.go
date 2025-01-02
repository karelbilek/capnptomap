package capnptomap

import (
	"fmt"
	"math"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/schemas"
	"github.com/karelbilek/capnptomap/schema"
)

func findInRegistry(registry *schemas.Registry, typeID uint64) (schema.Node, error) {
	data, err := registry.Find(typeID)
	if err != nil {
		return schema.Node{}, err
	}
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return schema.Node{}, err
	}
	req, err := schema.ReadRootCodeGeneratorRequest(msg)
	if err != nil {
		return schema.Node{}, err
	}
	nodes, err := req.Nodes()
	if err != nil {
		return schema.Node{}, err
	}

	for i := 0; i < nodes.Len(); i++ {
		n := nodes.At(i)
		if n.Id() == typeID {
			return n, nil
		}
	}
	return schema.Node{}, fmt.Errorf("type not found")
}

func ToMap(registry *schemas.Registry, typeID uint64, s capnp.Struct) (map[string]any, error) {
	return extractStruct(registry, typeID, s)
}

func extractStruct(registry *schemas.Registry, typeID uint64, s capnp.Struct) (map[string]any, error) {
	node, err := findInRegistry(registry, typeID)
	if err != nil {
		return nil, err
	}
	fields, err := node.StructNode().Fields()
	if err != nil {
		return nil, err
	}
	m := map[string]any{}
	for i := 0; i < fields.Len(); i++ {
		f := fields.At(i)
		typ, err := f.Slot().Type()
		if err != nil {
			return nil, err
		}
		dv, err := f.Slot().DefaultValue()
		if err != nil {
			return nil, err
		}

		name, err := f.Name()
		if err != nil {
			return nil, err
		}
		switch typ.Which() {
		case schema.Type_Which_bool:
			v := s.Bit(capnp.BitOffset(f.Slot().Offset()))
			d := dv.Bool()
			m[name] = v != d // != acts as XOR

		case schema.Type_Which_int8:
			v := int8(s.Uint8(capnp.DataOffset(f.Slot().Offset())))
			d := dv.Int8()
			m[name] = v != d
		case schema.Type_Which_int16:
			v := int16(s.Uint16(capnp.DataOffset(f.Slot().Offset() * 2)))
			d := dv.Int16()
			m[name] = int64(v ^ d)
		case schema.Type_Which_int32:
			v := int32(s.Uint32(capnp.DataOffset(f.Slot().Offset() * 4)))
			d := dv.Int32()
			m[name] = int64(v ^ d)
		case schema.Type_Which_int64:
			v := int64(s.Uint64(capnp.DataOffset(f.Slot().Offset() * 8)))
			d := dv.Int64()
			m[name] = v ^ d
		case schema.Type_Which_uint8:
			v := s.Uint8(capnp.DataOffset(f.Slot().Offset()))
			d := dv.Uint8()
			m[name] = uint64(v ^ d)
		case schema.Type_Which_uint16:
			v := s.Uint16(capnp.DataOffset(f.Slot().Offset() * 2))
			d := dv.Uint16()
			m[name] = uint64(v ^ d)
		case schema.Type_Which_enum:
			v := s.Uint16(capnp.DataOffset(f.Slot().Offset() * 2))
			d := dv.Enum()
			m[name] = uint64(v ^ d)
		case schema.Type_Which_uint32:
			v := s.Uint32(capnp.DataOffset(f.Slot().Offset() * 4))
			d := dv.Uint32()
			m[name] = uint64(v ^ d)
		case schema.Type_Which_uint64:
			v := s.Uint64(capnp.DataOffset(f.Slot().Offset() * 8))
			d := dv.Uint64()
			m[name] = v ^ d
		case schema.Type_Which_float32:
			v := s.Uint32(capnp.DataOffset(f.Slot().Offset() * 4))
			d := math.Float32bits(dv.Float32())
			m[name] = float64(math.Float32frombits(v ^ d))
		case schema.Type_Which_float64:
			v := s.Uint64(capnp.DataOffset(f.Slot().Offset() * 8))
			d := math.Float64bits(dv.Float64())
			m[name] = math.Float64frombits(v ^ d)
		case schema.Type_Which_text:
			p, err := s.Ptr(uint16(f.Slot().Offset()))
			if err != nil {
				return nil, err
			}
			var b []byte
			if p.IsValid() {
				b = p.TextBytes()
			} else {
				b, _ = dv.TextBytes()
			}
			m[name] = string(b)
		case schema.Type_Which_data:
			p, err := s.Ptr(uint16(f.Slot().Offset()))
			if err != nil {
				return nil, err
			}
			var b []byte
			if p.IsValid() {
				b = p.Data()
			} else {
				b, _ = dv.Data()
			}
			m[name] = b
		case schema.Type_Which_structType:
			p, err := s.Ptr(uint16(f.Slot().Offset()))
			if err != nil {
				return nil, err
			}
			ss := p.Struct()
			if !ss.IsValid() {
				p, _ = dv.StructValue()
				ss = p.Struct()
			}
			m[name], err = extractStruct(registry, typ.StructType().TypeId(), ss)
			if err != nil {
				return nil, err
			}
		case schema.Type_Which_list:
			p, err := s.Ptr(uint16(f.Slot().Offset()))
			if err != nil {
				return nil, err
			}
			l := p.List()
			if !l.IsValid() {
				p, _ = dv.List()
				l = p.List()
			}
			m[name], err = extractList(registry, typ, l)
			if err != nil {
				return nil, err
			}
		case schema.Type_Which_interface:
			return nil, fmt.Errorf("interfaces not supported")
		case schema.Type_Which_anyPointer:
			return nil, fmt.Errorf("anypointer not supported")
		default:
			return nil, fmt.Errorf("unknown field type %v", typ.Which())
		}
	}
	return m, nil
}

func extractList(registry *schemas.Registry, typ schema.Type, l capnp.List) ([]any, error) {
	elem, err := typ.List().ElementType()
	if err != nil {
		return nil, err
	}
	if !l.IsValid() {
		return nil, nil
	}
	n := l.Len()
	res := make([]any, n, n)
	switch elem.Which() {
	case schema.Type_Which_bool:
		for i := 0; i < n; i++ {
			res[i] = capnp.BitList(l).At(i)
		}
	case schema.Type_Which_int8:
		for i := 0; i < n; i++ {
			res[i] = int64(capnp.Int8List(l).At(i))
		}
	case schema.Type_Which_int16:
		for i := 0; i < n; i++ {
			res[i] = int64(capnp.Int16List(l).At(i))
		}
	case schema.Type_Which_int32:
		for i := 0; i < n; i++ {
			res[i] = int64(capnp.Int32List(l).At(i))
		}
	case schema.Type_Which_int64:
		for i := 0; i < n; i++ {
			res[i] = capnp.Int64List(l).At(i)
		}
	case schema.Type_Which_uint8:
		for i := 0; i < n; i++ {
			res[i] = uint64(capnp.UInt8List(l).At(i))
		}
	case schema.Type_Which_uint16, schema.Type_Which_enum:
		for i := 0; i < n; i++ {
			res[i] = uint64(capnp.UInt16List(l).At(i))
		}
	case schema.Type_Which_uint32:
		for i := 0; i < n; i++ {
			res[i] = uint64(capnp.UInt32List(l).At(i))
		}
	case schema.Type_Which_uint64:
		for i := 0; i < n; i++ {
			res[i] = capnp.UInt64List(l).At(i)
		}
	case schema.Type_Which_float32:
		for i := 0; i < n; i++ {
			res[i] = float64(capnp.Float32List(l).At(i))
		}
	case schema.Type_Which_float64:
		for i := 0; i < n; i++ {
			res[i] = capnp.Float64List(l).At(i)
		}
	case schema.Type_Which_text:
		for i := 0; i < n; i++ {
			s, err := capnp.TextList(l).At(i)
			if err != nil {
				return nil, err
			}
			res[i] = s
		}
	case schema.Type_Which_data:
		for i := 0; i < n; i++ {
			b, err := capnp.DataList(l).At(i)
			if err != nil {
				return nil, err
			}
			res[i] = b
		}
	case schema.Type_Which_list:
		for i := 0; i < n; i++ {
			p, err := capnp.PointerList(l).At(i)
			if err != nil {
				return nil, err
			}
			ls, err := extractList(registry, elem, p.List())
			if err != nil {
				return nil, err
			}
			res[i] = ls
			if ls == nil {
				res[i] = []any{} // burtsushi toml has problem with nil array...
			}
		}
	case schema.Type_Which_structType:
		for i := 0; i < n; i++ {
			mp, err := extractStruct(registry, elem.StructType().TypeId(), l.Struct(i))
			if err != nil {
				return nil, err
			}
			res[i] = mp
		}
	case schema.Type_Which_interface:
		return nil, fmt.Errorf("interfaces not supported")
	case schema.Type_Which_anyPointer:
		return nil, fmt.Errorf("anypointer not supported")
	default:
		return nil, fmt.Errorf("unknown field type %v", typ.Which())
	}
	return res, nil
}
