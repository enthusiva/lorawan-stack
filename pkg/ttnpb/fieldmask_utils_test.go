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

package ttnpb_test

import (
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestTopLevelFields(t *testing.T) {
	a := assertions.New(t)
	paths := []string{
		"a",
		"b",
		"b.c",
		"b.c.d",
		"c.d",
	}
	a.So(TopLevelFields(paths), should.Resemble, []string{"a", "b", "c"})
}

func TestBottomLevelFields(t *testing.T) {
	a := assertions.New(t)
	paths := []string{
		"a",
		"b",
		"b.c",
		"d.e",
		"f.g.h",
		"f.g.h.i",
		"f.g.h.i.j",
		"f.g.h.i.k",
	}
	a.So(BottomLevelFields(paths), should.HaveSameElementsDeep, []string{
		"a",
		"b.c",
		"d.e",
		"f.g.h.i.j",
		"f.g.h.i.k",
	})
}

func TestHasOnlyAllowedFields(t *testing.T) {
	a := assertions.New(t)
	allowed := []string{
		"a",
		"b.c",
		"d.e",
	}

	{
		requested := []string{
			"a",
			"b.c",
			"b.c.d", // lower level allowed
		}
		a.So(HasOnlyAllowedFields(requested, allowed...), should.BeTrue)
	}

	{
		requested := []string{
			"a",
			"e.f",
		}
		a.So(HasOnlyAllowedFields(requested, allowed...), should.BeFalse)
	}

	{
		requested := []string{
			"a",
			"d", // higher level not allowed
		}
		a.So(HasOnlyAllowedFields(requested, allowed...), should.BeFalse)
	}
}

func TestHasAnyField(t *testing.T) {
	a := assertions.New(t)
	requested := []string{
		"a",
		"b.c",
		"d",
	}
	a.So(HasAnyField(requested, "x", "a"), should.BeTrue)
	a.So(HasAnyField(requested, "x.y", "b"), should.BeFalse)
	a.So(HasAnyField(requested, "x", "b.c"), should.BeTrue)
	a.So(HasAnyField(requested, "x", "b.c.d"), should.BeTrue)
	a.So(HasAnyField(requested, "d"), should.BeTrue)
	a.So(HasAnyField(requested, "d.e", "b"), should.BeTrue)
}

func TestFlattenPaths(t *testing.T) {
	a := assertions.New(t)
	paths := []string{
		"a",
		"a.b",
		"a.b.c",
		"a.b.c.d",
		"e.f",
		"x.y",
		"x.y.z",
	}
	a.So(FlattenPaths(paths, []string{"a.b", "x", "notfound"}), should.Resemble, []string{"a", "a.b", "e.f", "x"})
}

func TestContainsField(t *testing.T) {
	a := assertions.New(t)
	a.So(ContainsField("a.b", []string{"a.b", "c"}), should.BeTrue)
	a.So(ContainsField("x", []string{"a.b", "c"}), should.BeFalse)
}

func TestAllowedFields(t *testing.T) {
	a := assertions.New(t)
	paths := []string{
		"x",
		"c.d",
	}
	allowedPaths := []string{
		"a",
		"a.b",
		"c",
		"c.d",
		"c.d.e",
	}
	a.So(AllowedFields(paths, allowedPaths), should.Resemble, []string{"c.d"})
}

func TestAllowedBottomLevelFields(t *testing.T) {
	a := assertions.New(t)
	paths := []string{
		"x",
		"c",
	}
	allowedPaths := []string{
		"a",
		"a.b",
		"c",
		"c.d",
		"c.e",
	}
	a.So(AllowedBottomLevelFields(paths, allowedPaths), should.HaveSameElementsDeep, []string{"c.d", "c.e"})
}

func TestExcludeFields(t *testing.T) {
	a := assertions.New(t)
	paths := []string{
		"a.b.c",
		"c",
		"c.d",
		"e",
		"e.f",
	}
	excludePaths := []string{
		"a",
		"c.d",
	}
	a.So(ExcludeFields(paths, excludePaths...), should.HaveSameElementsDeep, []string{"c", "e", "e.f"})
}

func TestExcludeSubFields(t *testing.T) {
	a := assertions.New(t)
	paths := []string{
		"a.b.c",
		"c",
		"c.d.e",
		"c.d",
		"c.d.f.g",
		"e",
		"e.f",
	}
	excludePaths := []string{
		"a",
		"c.d",
	}
	a.So(ExcludeSubFields(paths, excludePaths...), should.HaveSameElementsDeep, []string{"c", "c.d", "e", "e.f"})
}

func TestAllFields(t *testing.T) {
	a := assertions.New(t)
	paths := []string{
		"a.b.c",
		"c",
		"c.d",
		"e",
	}
	addPaths := []string{
		"a.b.c",
		"a.b.c.d.f",
		"a.b.c.d.e",
		"a.b.c.g",
		"c",
		"e.f",
		"f",
	}
	a.So(AddFields(paths, addPaths...), should.HaveSameElementsDeep, []string{
		"a.b.c",
		"c",
		"e",
		"f",
	})
}

func TestIncludeFields(t *testing.T) {
	t.Parallel()

	a := assertions.New(t)
	paths := []string{
		"a.b.c",
		"b",
		"c.d",
		"c.d.e",
		"c.e",
		"c.e.f",
	}
	a.So(IncludeFields(paths, "c.d"), should.Resemble, []string{"c.d", "c.d.e"})
	a.So(IncludeFields(paths, "c.d", "c.e.f"), should.Resemble, []string{"c.d", "c.d.e", "c.e.f"})
	a.So(IncludeFields(paths, "a", "b", "c"), should.Resemble, paths)
	a.So(IncludeFields(paths, "c.e.g"), should.Resemble, []string{})
}

func TestFieldsWithoutWrappers(t *testing.T) {
	t.Parallel()

	a := assertions.New(t)
	paths := []string{
		"a",
		"a.b.value",
		"a.c.value",
		"a.c.d",
	}
	a.So(FieldsWithoutWrappers(paths), should.Resemble, []string{"a", "a.c.value", "a.c.d"})
}

func TestFieldMaskPathsSet(t *testing.T) {
	t.Parallel()

	a := assertions.New(t)
	testCases := []struct {
		Name   string
		Paths  []string
		Result map[string]struct{}
	}{
		{
			Name:   "Empty",
			Paths:  []string{},
			Result: map[string]struct{}{},
		},
		{
			Name:   "Nil",
			Paths:  nil,
			Result: map[string]struct{}{},
		},
		{
			Name: "Valid",
			Paths: []string{
				"a",
				"a.b.value",
				"a.c.value",
				"a.c.d",
			},
			Result: map[string]struct{}{
				"a":         {},
				"a.b.value": {},
				"a.c.value": {},
				"a.c.d":     {},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			actual := FieldMaskPathsSet(tc.Paths)
			a.So(actual, should.Resemble, tc.Result)
		})
	}
}

func TestFieldMaskPathsSetContainsAll(t *testing.T) {
	t.Parallel()

	testSet := map[string]struct{}{
		"a":         {},
		"a.b.value": {},
		"a.c.value": {},
	}

	a := assertions.New(t)
	testCases := []struct {
		Name                string
		Set                 map[string]struct{}
		Subset              []string
		ExpectedContainsAll bool
		ExpectedMissing     string
	}{
		{
			Name:                "Empty",
			Set:                 map[string]struct{}{},
			Subset:              []string{},
			ExpectedContainsAll: true,
			ExpectedMissing:     "",
		},
		{
			Name:                "EmptySet",
			Set:                 map[string]struct{}{},
			Subset:              []string{"a"},
			ExpectedContainsAll: false,
			ExpectedMissing:     "a",
		},
		{
			Name:                "EmptySubset",
			Set:                 testSet,
			Subset:              []string{},
			ExpectedContainsAll: true,
			ExpectedMissing:     "",
		},
		{
			Name:                "Nil",
			Set:                 nil,
			Subset:              nil,
			ExpectedContainsAll: true,
			ExpectedMissing:     "",
		},
		{
			Name:                "NilSet",
			Set:                 nil,
			Subset:              []string{"a"},
			ExpectedContainsAll: false,
			ExpectedMissing:     "a",
		},
		{
			Name:                "NilSubset",
			Set:                 testSet,
			Subset:              nil,
			ExpectedContainsAll: true,
			ExpectedMissing:     "",
		},
		{
			Name: "MissingSingle",
			Set:  testSet,
			Subset: []string{
				"a",
				"a.b.value",
				"a.c.value",
				"a.c.d",
			},
			ExpectedContainsAll: false,
			ExpectedMissing:     "a.c.d",
		},
		{
			Name: "MissingMultiple",
			Set:  testSet,
			Subset: []string{
				"a",
				"a.b.value",
				"a.c.value",
				"a.c.d",
				"a.c.e",
			},
			ExpectedContainsAll: false,
			ExpectedMissing:     "a.c.d",
		},
		{
			Name: "All",
			Set:  testSet,
			Subset: []string{
				"a",
				"a.b.value",
				"a.c.value",
			},
			ExpectedContainsAll: true,
			ExpectedMissing:     "",
		},
		{
			Name: "Subset",
			Set:  testSet,
			Subset: []string{
				"a",
				"a.b.value",
			},
			ExpectedContainsAll: true,
			ExpectedMissing:     "",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			actualContainsAll, actualMissing := FieldMaskPathsSetContainsAll(tc.Set, tc.Subset...)
			a.So(actualContainsAll, should.Equal, tc.ExpectedContainsAll)
			a.So(actualMissing, should.Equal, tc.ExpectedMissing)
		})
	}
}
