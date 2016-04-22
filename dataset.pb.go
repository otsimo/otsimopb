// Code generated by protoc-gen-gogo.
// source: dataset.proto
// DO NOT EDIT!

package otsimopb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Column_Type int32

const (
	STRING      Column_Type = 0
	INTEGER     Column_Type = 1
	REAL        Column_Type = 2
	DATE        Column_Type = 3
	DATE_TIME   Column_Type = 4
	TIME_OF_DAY Column_Type = 5
)

var Column_Type_name = map[int32]string{
	0: "STRING",
	1: "INTEGER",
	2: "REAL",
	3: "DATE",
	4: "DATE_TIME",
	5: "TIME_OF_DAY",
}
var Column_Type_value = map[string]int32{
	"STRING":      0,
	"INTEGER":     1,
	"REAL":        2,
	"DATE":        3,
	"DATE_TIME":   4,
	"TIME_OF_DAY": 5,
}

func (x Column_Type) String() string {
	return proto.EnumName(Column_Type_name, int32(x))
}
func (Column_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorDataset, []int{0, 0} }

type Column struct {
	Type Column_Type `protobuf:"varint,1,opt,name=type,proto3,enum=apipb.Column_Type" json:"type,omitempty"`
	Name string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Column) Reset()                    { *m = Column{} }
func (m *Column) String() string            { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()               {}
func (*Column) Descriptor() ([]byte, []int) { return fileDescriptorDataset, []int{0} }

type TimeOfDay struct {
	Hours        int32 `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	Minutes      int32 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Seconds      int32 `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Milliseconds int32 `protobuf:"varint,4,opt,name=milliseconds,proto3" json:"milliseconds,omitempty"`
}

func (m *TimeOfDay) Reset()                    { *m = TimeOfDay{} }
func (m *TimeOfDay) String() string            { return proto.CompactTextString(m) }
func (*TimeOfDay) ProtoMessage()               {}
func (*TimeOfDay) Descriptor() ([]byte, []int) { return fileDescriptorDataset, []int{1} }

type RowValue struct {
	// Types that are valid to be assigned to Value:
	//	*RowValue_Str
	//	*RowValue_Int
	//	*RowValue_Real
	//	*RowValue_Date
	//	*RowValue_DateOfTime
	//	*RowValue_TimeOfDay
	Value isRowValue_Value `protobuf_oneof:"value"`
}

func (m *RowValue) Reset()                    { *m = RowValue{} }
func (m *RowValue) String() string            { return proto.CompactTextString(m) }
func (*RowValue) ProtoMessage()               {}
func (*RowValue) Descriptor() ([]byte, []int) { return fileDescriptorDataset, []int{2} }

type isRowValue_Value interface {
	isRowValue_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type RowValue_Str struct {
	Str string `protobuf:"bytes,1,opt,name=str,proto3,oneof"`
}
type RowValue_Int struct {
	Int int32 `protobuf:"varint,2,opt,name=int,proto3,oneof"`
}
type RowValue_Real struct {
	Real float32 `protobuf:"fixed32,3,opt,name=real,proto3,oneof"`
}
type RowValue_Date struct {
	Date int64 `protobuf:"varint,4,opt,name=date,proto3,oneof"`
}
type RowValue_DateOfTime struct {
	DateOfTime int64 `protobuf:"varint,5,opt,name=date_of_time,json=dateOfTime,proto3,oneof"`
}
type RowValue_TimeOfDay struct {
	TimeOfDay *TimeOfDay `protobuf:"bytes,6,opt,name=time_of_day,json=timeOfDay,oneof"`
}

func (*RowValue_Str) isRowValue_Value()        {}
func (*RowValue_Int) isRowValue_Value()        {}
func (*RowValue_Real) isRowValue_Value()       {}
func (*RowValue_Date) isRowValue_Value()       {}
func (*RowValue_DateOfTime) isRowValue_Value() {}
func (*RowValue_TimeOfDay) isRowValue_Value()  {}

func (m *RowValue) GetValue() isRowValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RowValue) GetStr() string {
	if x, ok := m.GetValue().(*RowValue_Str); ok {
		return x.Str
	}
	return ""
}

func (m *RowValue) GetInt() int32 {
	if x, ok := m.GetValue().(*RowValue_Int); ok {
		return x.Int
	}
	return 0
}

func (m *RowValue) GetReal() float32 {
	if x, ok := m.GetValue().(*RowValue_Real); ok {
		return x.Real
	}
	return 0
}

func (m *RowValue) GetDate() int64 {
	if x, ok := m.GetValue().(*RowValue_Date); ok {
		return x.Date
	}
	return 0
}

func (m *RowValue) GetDateOfTime() int64 {
	if x, ok := m.GetValue().(*RowValue_DateOfTime); ok {
		return x.DateOfTime
	}
	return 0
}

func (m *RowValue) GetTimeOfDay() *TimeOfDay {
	if x, ok := m.GetValue().(*RowValue_TimeOfDay); ok {
		return x.TimeOfDay
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RowValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RowValue_OneofMarshaler, _RowValue_OneofUnmarshaler, _RowValue_OneofSizer, []interface{}{
		(*RowValue_Str)(nil),
		(*RowValue_Int)(nil),
		(*RowValue_Real)(nil),
		(*RowValue_Date)(nil),
		(*RowValue_DateOfTime)(nil),
		(*RowValue_TimeOfDay)(nil),
	}
}

func _RowValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RowValue)
	// value
	switch x := m.Value.(type) {
	case *RowValue_Str:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Str)
	case *RowValue_Int:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Int))
	case *RowValue_Real:
		_ = b.EncodeVarint(3<<3 | proto.WireFixed32)
		_ = b.EncodeFixed32(uint64(math.Float32bits(x.Real)))
	case *RowValue_Date:
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Date))
	case *RowValue_DateOfTime:
		_ = b.EncodeVarint(5<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.DateOfTime))
	case *RowValue_TimeOfDay:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimeOfDay); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RowValue.Value has unexpected type %T", x)
	}
	return nil
}

func _RowValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RowValue)
	switch tag {
	case 1: // value.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &RowValue_Str{x}
		return true, err
	case 2: // value.int
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &RowValue_Int{int32(x)}
		return true, err
	case 3: // value.real
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &RowValue_Real{math.Float32frombits(uint32(x))}
		return true, err
	case 4: // value.date
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &RowValue_Date{int64(x)}
		return true, err
	case 5: // value.date_of_time
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &RowValue_DateOfTime{int64(x)}
		return true, err
	case 6: // value.time_of_day
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TimeOfDay)
		err := b.DecodeMessage(msg)
		m.Value = &RowValue_TimeOfDay{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RowValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RowValue)
	// value
	switch x := m.Value.(type) {
	case *RowValue_Str:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Str)))
		n += len(x.Str)
	case *RowValue_Int:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int))
	case *RowValue_Real:
		n += proto.SizeVarint(3<<3 | proto.WireFixed32)
		n += 4
	case *RowValue_Date:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Date))
	case *RowValue_DateOfTime:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.DateOfTime))
	case *RowValue_TimeOfDay:
		s := proto.Size(x.TimeOfDay)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Row struct {
	Values []*RowValue `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *Row) Reset()                    { *m = Row{} }
func (m *Row) String() string            { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()               {}
func (*Row) Descriptor() ([]byte, []int) { return fileDescriptorDataset, []int{3} }

type DataSet struct {
	Label   string    `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Columns []*Column `protobuf:"bytes,5,rep,name=columns" json:"columns,omitempty"`
	Rows    []*Row    `protobuf:"bytes,6,rep,name=rows" json:"rows,omitempty"`
}

func (m *DataSet) Reset()                    { *m = DataSet{} }
func (m *DataSet) String() string            { return proto.CompactTextString(m) }
func (*DataSet) ProtoMessage()               {}
func (*DataSet) Descriptor() ([]byte, []int) { return fileDescriptorDataset, []int{4} }

func init() {
	proto.RegisterType((*Column)(nil), "apipb.Column")
	proto.RegisterType((*TimeOfDay)(nil), "apipb.TimeOfDay")
	proto.RegisterType((*RowValue)(nil), "apipb.RowValue")
	proto.RegisterType((*Row)(nil), "apipb.Row")
	proto.RegisterType((*DataSet)(nil), "apipb.DataSet")
	proto.RegisterEnum("apipb.Column_Type", Column_Type_name, Column_Type_value)
}
func (m *Column) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Column) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintDataset(data, i, uint64(m.Type))
	}
	if len(m.Name) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintDataset(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	return i, nil
}

func (m *TimeOfDay) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TimeOfDay) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Hours != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintDataset(data, i, uint64(m.Hours))
	}
	if m.Minutes != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintDataset(data, i, uint64(m.Minutes))
	}
	if m.Seconds != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintDataset(data, i, uint64(m.Seconds))
	}
	if m.Milliseconds != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintDataset(data, i, uint64(m.Milliseconds))
	}
	return i, nil
}

func (m *RowValue) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RowValue) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		nn1, err := m.Value.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *RowValue_Str) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0xa
	i++
	i = encodeVarintDataset(data, i, uint64(len(m.Str)))
	i += copy(data[i:], m.Str)
	return i, nil
}
func (m *RowValue_Int) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x10
	i++
	i = encodeVarintDataset(data, i, uint64(m.Int))
	return i, nil
}
func (m *RowValue_Real) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x1d
	i++
	i = encodeFixed32Dataset(data, i, uint32(math.Float32bits(float32(m.Real))))
	return i, nil
}
func (m *RowValue_Date) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x20
	i++
	i = encodeVarintDataset(data, i, uint64(m.Date))
	return i, nil
}
func (m *RowValue_DateOfTime) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x28
	i++
	i = encodeVarintDataset(data, i, uint64(m.DateOfTime))
	return i, nil
}
func (m *RowValue_TimeOfDay) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.TimeOfDay != nil {
		data[i] = 0x32
		i++
		i = encodeVarintDataset(data, i, uint64(m.TimeOfDay.Size()))
		n2, err := m.TimeOfDay.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *Row) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Row) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, msg := range m.Values {
			data[i] = 0xa
			i++
			i = encodeVarintDataset(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *DataSet) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DataSet) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Label) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintDataset(data, i, uint64(len(m.Label)))
		i += copy(data[i:], m.Label)
	}
	if len(m.Columns) > 0 {
		for _, msg := range m.Columns {
			data[i] = 0x2a
			i++
			i = encodeVarintDataset(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Rows) > 0 {
		for _, msg := range m.Rows {
			data[i] = 0x32
			i++
			i = encodeVarintDataset(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Dataset(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Dataset(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDataset(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Column) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovDataset(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDataset(uint64(l))
	}
	return n
}

func (m *TimeOfDay) Size() (n int) {
	var l int
	_ = l
	if m.Hours != 0 {
		n += 1 + sovDataset(uint64(m.Hours))
	}
	if m.Minutes != 0 {
		n += 1 + sovDataset(uint64(m.Minutes))
	}
	if m.Seconds != 0 {
		n += 1 + sovDataset(uint64(m.Seconds))
	}
	if m.Milliseconds != 0 {
		n += 1 + sovDataset(uint64(m.Milliseconds))
	}
	return n
}

func (m *RowValue) Size() (n int) {
	var l int
	_ = l
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *RowValue_Str) Size() (n int) {
	var l int
	_ = l
	l = len(m.Str)
	n += 1 + l + sovDataset(uint64(l))
	return n
}
func (m *RowValue_Int) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovDataset(uint64(m.Int))
	return n
}
func (m *RowValue_Real) Size() (n int) {
	var l int
	_ = l
	n += 5
	return n
}
func (m *RowValue_Date) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovDataset(uint64(m.Date))
	return n
}
func (m *RowValue_DateOfTime) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovDataset(uint64(m.DateOfTime))
	return n
}
func (m *RowValue_TimeOfDay) Size() (n int) {
	var l int
	_ = l
	if m.TimeOfDay != nil {
		l = m.TimeOfDay.Size()
		n += 1 + l + sovDataset(uint64(l))
	}
	return n
}
func (m *Row) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovDataset(uint64(l))
		}
	}
	return n
}

func (m *DataSet) Size() (n int) {
	var l int
	_ = l
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovDataset(uint64(l))
	}
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovDataset(uint64(l))
		}
	}
	if len(m.Rows) > 0 {
		for _, e := range m.Rows {
			l = e.Size()
			n += 1 + l + sovDataset(uint64(l))
		}
	}
	return n
}

func sovDataset(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDataset(x uint64) (n int) {
	return sovDataset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Column) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Column: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Column: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (Column_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataset(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDataset
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimeOfDay) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimeOfDay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeOfDay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Hours |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minutes", wireType)
			}
			m.Minutes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Minutes |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Seconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Milliseconds", wireType)
			}
			m.Milliseconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Milliseconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDataset(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDataset
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RowValue) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RowValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RowValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Str", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &RowValue_Str{string(data[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &RowValue_Int{v}
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Real", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Value = &RowValue_Real{float32(math.Float32frombits(v))}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &RowValue_Date{v}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateOfTime", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &RowValue_DateOfTime{v}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeOfDay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TimeOfDay{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &RowValue_TimeOfDay{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataset(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDataset
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Row) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Row: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Row: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &RowValue{})
			if err := m.Values[len(m.Values)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataset(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDataset
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DataSet) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, &Column{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rows = append(m.Rows, &Row{})
			if err := m.Rows[len(m.Rows)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataset(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDataset
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDataset(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataset
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDataset
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDataset
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDataset(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDataset = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataset   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorDataset = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x52, 0x4b, 0x8e, 0xd3, 0x40,
	0x10, 0x8d, 0x63, 0x3b, 0x9f, 0x0a, 0x21, 0x56, 0x8b, 0x85, 0xc5, 0x22, 0x42, 0x5e, 0x00, 0x1b,
	0x3c, 0x52, 0x10, 0x07, 0xc8, 0x90, 0x30, 0x44, 0x82, 0x09, 0xea, 0x58, 0x48, 0xb0, 0x89, 0xec,
	0xa4, 0x93, 0x58, 0xb2, 0xdd, 0x96, 0xdd, 0x66, 0x94, 0x5b, 0xb0, 0xe5, 0x34, 0x6c, 0x67, 0x39,
	0x47, 0xe0, 0x73, 0x11, 0xaa, 0xda, 0x36, 0xa3, 0x59, 0xb4, 0x5c, 0xef, 0xbd, 0xaa, 0xae, 0x7a,
	0xed, 0x82, 0xf1, 0x3e, 0x54, 0x61, 0x29, 0x94, 0x9f, 0x17, 0x52, 0x49, 0x66, 0x87, 0x79, 0x9c,
	0x47, 0x4f, 0x5f, 0x1d, 0x63, 0x75, 0xaa, 0x22, 0x7f, 0x27, 0xd3, 0x8b, 0xa3, 0x3c, 0xca, 0x0b,
	0xad, 0x46, 0xd5, 0x41, 0x23, 0x0d, 0x74, 0x54, 0x57, 0x79, 0x3f, 0x0c, 0xe8, 0xbd, 0x95, 0x49,
	0x95, 0x66, 0xec, 0x39, 0x58, 0xea, 0x9c, 0x0b, 0xd7, 0x78, 0x66, 0xbc, 0x7c, 0x3c, 0x63, 0xbe,
	0xbe, 0xcf, 0xaf, 0x45, 0x3f, 0x40, 0x85, 0x6b, 0x9d, 0x31, 0xb0, 0xb2, 0x30, 0x15, 0x6e, 0x17,
	0xf3, 0x86, 0x5c, 0xc7, 0xde, 0x06, 0x2c, 0xca, 0x60, 0x00, 0xbd, 0x4d, 0xc0, 0x57, 0xd7, 0x57,
	0x4e, 0x87, 0x8d, 0xa0, 0xbf, 0xba, 0x0e, 0x96, 0x57, 0x4b, 0xee, 0x18, 0x6c, 0x00, 0x16, 0x5f,
	0xce, 0x3f, 0x38, 0x5d, 0x8a, 0x16, 0xf3, 0x60, 0xe9, 0x98, 0x6c, 0x0c, 0x43, 0x8a, 0xb6, 0xc1,
	0xea, 0xe3, 0xd2, 0xb1, 0xd8, 0x04, 0x46, 0x14, 0x6d, 0xd7, 0xef, 0xb6, 0x8b, 0xf9, 0x17, 0xc7,
	0xf6, 0xce, 0x30, 0x0c, 0xe2, 0x54, 0xac, 0x0f, 0x8b, 0xf0, 0xcc, 0x9e, 0x80, 0x7d, 0x92, 0x55,
	0x51, 0xea, 0xf1, 0x6c, 0x5e, 0x03, 0xe6, 0x42, 0x3f, 0x8d, 0xb3, 0x4a, 0x89, 0x52, 0x8f, 0x63,
	0xf3, 0x16, 0x92, 0x52, 0x8a, 0x9d, 0xcc, 0xf6, 0xa5, 0x6b, 0xd6, 0x4a, 0x03, 0x99, 0x07, 0x8f,
	0xd2, 0x38, 0x49, 0xe2, 0x56, 0xb6, 0xb4, 0xfc, 0x80, 0xf3, 0x7e, 0x1a, 0x30, 0xe0, 0xf2, 0xe6,
	0x73, 0x98, 0x54, 0x64, 0xd8, 0x2c, 0x55, 0xa1, 0x1b, 0x0f, 0xdf, 0x77, 0x38, 0x01, 0xe2, 0xe2,
	0x4c, 0xd5, 0x4d, 0x89, 0x43, 0x80, 0x23, 0x5a, 0x85, 0x08, 0x13, 0xdd, 0xaf, 0x8b, 0xa4, 0x46,
	0xc4, 0xe2, 0x8f, 0x12, 0xba, 0x8d, 0x49, 0x2c, 0x21, 0x1a, 0x82, 0xbe, 0x5b, 0x79, 0xd8, 0x2a,
	0xf4, 0xe8, 0xda, 0x8d, 0x0a, 0xc4, 0xae, 0x0f, 0xe4, 0x9b, 0xcd, 0x60, 0x44, 0x1a, 0xe5, 0xec,
	0xc3, 0xb3, 0xdb, 0xc3, 0x94, 0xd1, 0xcc, 0x69, 0xfe, 0xcb, 0xff, 0x97, 0xc1, 0xa2, 0xa1, 0x6a,
	0xc1, 0x65, 0x1f, 0xec, 0x6f, 0x34, 0xb4, 0xe7, 0x83, 0x89, 0x06, 0xd8, 0x0b, 0xe8, 0x69, 0x4c,
	0xef, 0x66, 0x62, 0xf9, 0xa4, 0x29, 0x6f, 0xcd, 0xf1, 0x46, 0xf6, 0x4e, 0xd0, 0x5f, 0xe0, 0x3e,
	0x6d, 0x04, 0xf9, 0xb0, 0x93, 0x30, 0x12, 0x49, 0xed, 0x98, 0xd7, 0x00, 0x6f, 0xea, 0xef, 0xf4,
	0x2e, 0x94, 0x38, 0x2c, 0x5d, 0x35, 0x7e, 0xb0, 0x21, 0xbc, 0x55, 0xd9, 0x14, 0x9f, 0x41, 0xde,
	0x94, 0x38, 0x2f, 0x65, 0xc1, 0x7d, 0x43, 0xae, 0xf9, 0xcb, 0x37, 0xb7, 0xbf, 0xa7, 0x9d, 0x3b,
	0x3c, 0xb7, 0x7f, 0xa6, 0xc6, 0x1d, 0x9e, 0x5f, 0x78, 0xbe, 0xff, 0x9d, 0x76, 0x60, 0x82, 0x4b,
	0xeb, 0x4b, 0x55, 0xc6, 0xa9, 0xf4, 0x8f, 0x45, 0xbe, 0xfb, 0x64, 0x7c, 0x1d, 0xd4, 0x30, 0x8f,
	0xa2, 0x9e, 0x5e, 0xd8, 0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x44, 0x9a, 0x37, 0xe5, 0xf7,
	0x02, 0x00, 0x00,
}
