// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package goproto_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/smarty/assertions"
	"github.com/spf13/cast"
	"go.thethings.network/lorawan-stack/v3/pkg/goproto"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/protobuf/types/known/structpb"
)

type jsonMarshaler struct {
	Text string
}

func (m jsonMarshaler) MarshalJSON() ([]byte, error) {
	return bytes.ToUpper([]byte(`"` + m.Text + `"`)), nil
}

func (m *jsonMarshaler) UnmarshalJSON(b []byte) error {
	m.Text = string(bytes.ToLower(bytes.Trim(b, `"`)))
	return nil
}

func TestStructProto(t *testing.T) {
	a := assertions.New(t)

	ptr := "ptr"
	m := map[string]any{
		"foo":            "bar",
		"ptr":            &ptr,
		"answer":         42,
		"answer.precise": 42.0,
		"works":          true,
		"empty":          nil,
		"list":           []string{"a", "b", "c"},
		"map":            map[string]string{"foo": "bar"},
		"eui":            types.EUI64{1, 2, 3, 4, 5, 6, 7, 8},
		"jsonMarshaler":  &jsonMarshaler{Text: "testtext"},
	}
	s, err := goproto.Struct(m)
	a.So(err, should.BeNil)
	sm, err := goproto.Map(s)
	a.So(err, should.BeNil)
	for k, v := range m {
		a.So(s.Fields, should.ContainKey, k)
		a.So(sm, should.ContainKey, k)
		if v == nil {
			a.So(s.Fields[k].Kind, should.HaveSameTypeAs, &structpb.Value_NullValue{})
			a.So(sm[k], should.BeNil)
			continue
		}

		rv := reflect.Indirect(reflect.ValueOf(v))

		switch kind := rv.Kind(); kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:

			var vt float64
			a.So(s.Fields[k].Kind, should.HaveSameTypeAs, &structpb.Value_NumberValue{})
			a.So(sm[k], should.HaveSameTypeAs, vt)
			a.So(sm[k], should.Equal, cast.ToFloat64(rv.Interface()))

		case reflect.Bool:
			var vt bool
			a.So(s.Fields[k].Kind, should.HaveSameTypeAs, &structpb.Value_BoolValue{})
			a.So(sm[k], should.HaveSameTypeAs, vt)
			a.So(sm[k], should.Equal, rv.Bool())

		case reflect.String:
			var vt string
			a.So(s.Fields[k].Kind, should.HaveSameTypeAs, &structpb.Value_StringValue{})
			a.So(sm[k], should.HaveSameTypeAs, vt)
			a.So(sm[k], should.Equal, rv.String())

		case reflect.Slice, reflect.Array:
			var vt []any
			a.So(s.Fields[k].Kind, should.HaveSameTypeAs, &structpb.Value_ListValue{})
			a.So(sm[k], should.HaveSameTypeAs, vt)
			if a.So(sm[k], should.HaveLength, rv.Len()) {
				// TODO find a way to compare these values
				//smv := reflect.ValueOf(sm[k])
				//for i := 0; i < rv.Len(); i++ {
				//a.So(smv.Index(i).Interface(), should.Resemble, rv.Index(i).Interface())
				//}
			}

		case reflect.Struct, reflect.Map:
			var vt map[string]any
			a.So(s.Fields[k].Kind, should.HaveSameTypeAs, &structpb.Value_StructValue{})
			a.So(sm[k], should.HaveSameTypeAs, vt)
			if kind == reflect.Map {
				a.So(sm[k], should.HaveLength, rv.Len())
			}

		default:
			panic("Unmatched kind: " + rv.Kind().String())
		}
		pv, err := goproto.Value(rv.Interface())
		if a.So(err, should.BeNil) {
			a.So(s.Fields[k], should.Resemble, pv)

			gv, err := goproto.Interface(pv)
			if a.So(err, should.BeNil) {
				a.So(sm[k], should.Resemble, gv)
			}
		}
	}
}

func TestRecursiveStructures(t *testing.T) {
	t.Parallel()

	recursiveStruct := &structpb.Struct{Fields: make(map[string]*structpb.Value)}
	recursiveStruct.Fields["test"] = &structpb.Value{
		Kind: &structpb.Value_StructValue{
			StructValue: recursiveStruct,
		},
	}
	recursiveList := &structpb.ListValue{Values: make([]*structpb.Value, 1)}
	recursiveList.Values[0] = &structpb.Value{
		Kind: &structpb.Value_ListValue{
			ListValue: recursiveList,
		},
	}
	recursiveValueStruct := &structpb.Value{
		Kind: &structpb.Value_StructValue{
			StructValue: recursiveStruct,
		},
	}
	recursiveValueList := &structpb.Value{
		Kind: &structpb.Value_ListValue{
			ListValue: recursiveList,
		},
	}

	recursiveMap := make(map[string]any)
	recursiveMap["test"] = recursiveMap
	recursiveSlice := make([]any, 1)
	recursiveSlice[0] = recursiveSlice
	type recursiveGoStruct struct {
		self *recursiveGoStruct
	}
	recursiveGoStructValue := &recursiveGoStruct{}
	recursiveGoStructValue.self = recursiveGoStructValue

	t.Run("Map", func(t *testing.T) {
		t.Parallel()

		a := assertions.New(t)
		_, err := goproto.Map(recursiveStruct)
		a.So(err, should.NotBeNil)
	})

	t.Run("Slice", func(t *testing.T) {
		t.Parallel()

		a := assertions.New(t)
		_, err := goproto.Slice(recursiveList)
		a.So(err, should.NotBeNil)
	})

	t.Run("Interface", func(t *testing.T) {
		t.Parallel()

		a := assertions.New(t)
		_, err := goproto.Interface(recursiveValueStruct)
		a.So(err, should.NotBeNil)
		_, err = goproto.Interface(recursiveValueList)
		a.So(err, should.NotBeNil)
	})

	t.Run("Struct", func(t *testing.T) {
		t.Parallel()

		a := assertions.New(t)
		_, err := goproto.Struct(recursiveMap)
		a.So(err, should.NotBeNil)
	})

	t.Run("List", func(t *testing.T) {
		t.Parallel()

		a := assertions.New(t)
		_, err := goproto.List(recursiveSlice)
		a.So(err, should.NotBeNil)
	})

	t.Run("Value", func(t *testing.T) {
		t.Parallel()

		a := assertions.New(t)
		_, err := goproto.Value(recursiveSlice)
		a.So(err, should.NotBeNil)
		_, err = goproto.Value(recursiveMap)
		a.So(err, should.NotBeNil)
		_, err = goproto.Value(recursiveGoStructValue)
		a.So(err, should.NotBeNil)
	})
}
