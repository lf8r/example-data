// Copyright (C) Subhajit DasGupta 2022

package data3

import (
	"github.com/lf8r/dbgen-common/pkg/common"
	"github.com/lf8r/example-data/pkg/data"
)

type PersonT struct {
	common.Resource
	Byte                  int8
	Word                  int16
	Long                  int32
	Dlong                 int64
	Float32               float32
	Float64               float64
	UByte                 uint8
	UWord                 uint16
	ULong                 uint32
	UDlong                uint64
	Age                   int
	Good                  bool
	GoodP1                *bool
	GoodP2                *bool
	Address               data.Address
	AddressP2             *data.Address
	AddressP3             *data.Address
	ByteP                 *int8
	WordP                 *int16
	LongP                 *int32
	DwordP                *int64
	IntP                  *int
	UbyteP                *uint8
	UwordP                *uint16
	UlongP                *uint32
	UdwordP               *uint64
	UintP                 *uint
	Child1                string
	Pchild1               *string
	Ints                  []int
	Ints8                 []int8
	Ints16                []int16
	Ints32                []int32
	Ints64                []int64
	Uints                 []uint
	Uints8                []uint8
	Uints16               []uint16
	Uints32               []uint32
	Uints64               []uint64
	Strs1                 []string
	Bools1                []bool
	Addresses             []data.Address
	Ints8P                []*int8
	Ints16P               []*int16
	Ints32P               []*int32
	Ints64P               []*int64
	IntsP                 []*int
	Uints8P               []*uint8
	Uints16P              []*uint16
	Uints32P              []*uint32
	Uints64P              []*uint64
	UintsP                []*uint
	BoolsP                []*bool
	Strings1P             []*string
	Strings2P             []*string
	AddressesP            []*data.Address
	AddressMap            map[string]data.Address
	AgeMap                map[int]data.Address
	StrStrMap             map[string]string
	IntStrMap             map[int]string
	IntBoolMap            map[int]bool
	StringBoolMap         map[string]bool
	StringInt8Map         map[string]int8
	StringInt16Map        map[string]int16
	StringInt32Map        map[string]int32
	StringInt64Map        map[string]int64
	StringIntMap          map[string]int
	StringUint8Map        map[string]uint8
	StringUint16Map       map[string]uint16
	StringUint32Map       map[string]uint32
	StringUint64Map       map[string]uint64
	StringUintMap         map[string]uint
	IntInt8Map            map[int]int8
	IntInt16Map           map[int]int16
	IntInt32Map           map[int]int32
	IntInt64Map           map[int]int64
	IntIntMap             map[int]int
	IntUint8Map           map[int]uint8
	IntUint16Map          map[int]uint16
	IntUint32Map          map[int]uint32
	IntUint64Map          map[int]uint64
	IntUintMap            map[int]uint
	IntFloat32Map         map[int]float32
	IntFloat64Map         map[int]float64
	StringFloat32Map      map[string]float32
	StringFloat64Map      map[string]float64
	StringToAddressPtrMap map[string]*data.Address
	IntToAddressPtrMap    map[int]*data.Address
}
