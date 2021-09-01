package mapsorter

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_getSortedKeysSliceFromMap(t *testing.T) {
	tests := []struct {
		name      string
		mapToSort interface{}
		sortBy    By
		sortAs    As
		reverse   bool
		top       int
		wantRes   []interface{}
		wantErr   bool
	}{
		{
			name: "01_asString_byKeys",
			mapToSort: map[string]int{
				"a":   1,
				"aa":  2,
				"abc": 3,
				"aaa": 4,
			},
			sortBy:  ByKeys,
			sortAs:  AsString,
			reverse: false,
			top:     -1,
			wantRes: []interface{}{"a", "aa", "aaa", "abc"},
			wantErr: false,
		},
		{
			name: "02_asString_byValues",
			mapToSort: map[string]string{
				"a":   "1",
				"aa":  "2",
				"abc": "3",
				"aaa": "4",
			},
			sortBy:  ByValues,
			sortAs:  AsString,
			reverse: false,
			top:     -1,
			wantRes: []interface{}{"a", "aa", "abc", "aaa"},
			wantErr: false,
		},
		{
			name: "03_asString_byKeys_reverse_top3",
			mapToSort: map[string]string{
				"a":   "1",
				"aa":  "2",
				"abc": "3",
				"aaa": "4",
			},
			sortBy:  ByKeys,
			sortAs:  AsString,
			reverse: true,
			top:     3,
			wantRes: []interface{}{"abc", "aaa", "aa"},
		},
		{
			name: "04_asString_byValues_equals",
			mapToSort: map[string]string{
				"a":   "1",
				"aa":  "2",
				"abc": "2",
				"aaa": "4",
			},
			sortBy:  ByValues,
			sortAs:  AsString,
			reverse: false,
			top:     -1,
			wantRes: []interface{}{"a", "aa", "abc", "aaa"},
		},
		{
			name: "05_asString_error",
			mapToSort: map[string]int{
				"a":   1,
				"aa":  2,
				"abc": 3,
				"aaa": 4,
			},
			sortBy:  ByValues,
			sortAs:  AsString,
			reverse: false,
			top:     -1,
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "06_asDatetime_byKeys_reverse_top3",
			mapToSort: map[time.Time]int{
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC): 1,
				time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC): 2,
				time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC): 3,
				time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC): 4,
			},
			sortBy:  ByKeys,
			sortAs:  AsDatetime,
			reverse: true,
			top:     3,
			wantRes: []interface{}{
				time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "07_asDatetime_byValues_with_parsing_equals",
			mapToSort: map[int]string{
				1: "2020-01-01",
				3: "2020-01-03",
				4: "2020-01-03",
				2: "2020-01-02",
				5: "2020-01-04",
			},
			sortBy:  ByValues,
			sortAs:  AsDatetime,
			reverse: false,
			top:     -1,
			wantRes: []interface{}{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name: "08_asDatetime_byValues_with_parsing_error",
			mapToSort: map[int]string{
				1: "hello",
				3: "2020-01-03",
				4: "2020-01-03",
				2: "2020-01-02",
				5: "2020-01-04",
			},
			sortBy:  ByValues,
			sortAs:  AsDatetime,
			reverse: false,
			top:     -1,
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "09_asStringByLength_byKeys_reverse_top3",
			mapToSort: map[string]string{
				"a":    "1",
				"ab":   "2",
				"abcd": "3",
				"abc":  "4",
			},
			sortBy:  ByKeys,
			sortAs:  AsStringByLength,
			reverse: true,
			top:     3,
			wantRes: []interface{}{"abcd", "abc", "ab"},
		},
		{
			name: "10_asStringByLength_byValues_equals",
			mapToSort: map[string]string{
				"a":    "1",
				"ab":   "22",
				"abcd": "22",
				"abc":  "333",
			},
			sortBy:  ByValues,
			sortAs:  AsStringByLength,
			reverse: false,
			top:     -1,
			wantRes: []interface{}{"a", "ab", "abcd", "abc"},
		},
		{
			name: "11_asInt_byKeys_reverse_top3",
			mapToSort: map[int]string{
				1: "1",
				4: "2",
				3: "3",
				2: "4",
			},
			sortBy:  ByKeys,
			sortAs:  AsInt,
			reverse: true,
			top:     3,
			wantRes: []interface{}{4, 3, 2},
		},
		{
			name: "12_asInt_byValues_equals",
			mapToSort: map[string]string{
				"a":    "3",
				"ab":   "2",
				"abcd": "2",
				"abc":  "1",
			},
			sortBy:  ByValues,
			sortAs:  AsInt,
			reverse: false,
			top:     -1,
			wantRes: []interface{}{"abc", "ab", "abcd", "a"},
		},
		{
			name: "13_asInt_byValues_error",
			mapToSort: map[interface{}]interface{}{
				"a":    "hello",
				"ab":   "2",
				"abcd": "2",
				"abc":  "1",
			},
			sortBy:  ByValues,
			sortAs:  AsInt,
			reverse: false,
			top:     -1,
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "14_asFloat_byKeys_reverse_top3",
			mapToSort: map[float64]string{
				0.1: "1",
				0.4: "2",
				0.3: "3",
				0.2: "4",
			},
			sortBy:  ByKeys,
			sortAs:  AsFloat,
			reverse: true,
			top:     3,
			wantRes: []interface{}{0.4, 0.3, 0.2},
		},
		{
			name: "15_asFloat_byValues_equals",
			mapToSort: map[interface{}]interface{}{
				"a":    "0.3",
				"ab":   "0.2",
				"abcd": "0.2",
				"abc":  "0.1",
			},
			sortBy:  ByValues,
			sortAs:  AsFloat,
			reverse: false,
			top:     -1,
			wantRes: []interface{}{"abc", "ab", "abcd", "a"},
		},
		{
			name: "16_asFloat_byValues_error",
			mapToSort: map[interface{}]interface{}{
				"a":    "hello",
				"ab":   "2",
				"abcd": "2",
				"abc":  "1",
			},
			sortBy:  ByValues,
			sortAs:  AsInt,
			reverse: false,
			top:     -1,
			wantRes: nil,
			wantErr: true,
		},
		{
			name:      "17_not_map_error",
			mapToSort: 1,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Sort(tt.mapToSort, tt.sortBy, tt.sortAs, tt.reverse, tt.top)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Sort() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_sorter_Sort(t *testing.T) {
	mapToSort := map[string]string{
		"a":    "1",
		"ab":   "22",
		"abcd": "22",
		"abc":  "333",
	}
	t.Run("01_struct_Sort()", func(t *testing.T) {
		got, err := Map(mapToSort).
			ByValues().
			ByKeys().
			AsString().
			AsInt().
			AsFloat().
			AsDatetime().
			AsStringByLength().
			Reverse().Forward().
			Top(1).All().
			Sort()

		if err != nil {
			t.Errorf("sorter.Sort() error = %v", err)
			return
		}
		want := []interface{}{"a", "ab", "abc", "abcd"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("sorter.Sort() = %v, want %v", got, want)
		}
	})
	t.Run("02_struct_MustSort()_panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustSort did not panic")
			}
		}()
		Map(mapToSort).
			AsFloat().
			MustSort()
	})
	t.Run("03_struct_MustSort()_not_panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("MustSort did panic but should not: %v", r)
			}
		}()
		Map(mapToSort).
			AsString().
			MustSort()
	})
}

// Benchmarking with boilerplate sorting

const (
	maxBenchmarkIntValue  = 10000
	maxBenchmarkMapLength = 1000
)

func createMapToSort() map[int]string {
	var key int
	res := make(map[int]string)
	for i := 0; i < maxBenchmarkMapLength; i++ {
		key = rand.Intn(maxBenchmarkIntValue)
		res[key] = uuid.NewString()
	}
	return res
}

func sortBoilerplateByIntKeys(in map[int]string) []int {
	s := make([]int, len(in))
	i := 0
	for key := range in {
		s[i] = key
		i++
	}
	sort.Ints(s)
	return s
}

func sortBoilerplateByLenStringsWithEquals(in map[int]string) []int {
	type kv struct {
		k int
		v string
	}
	s := make([]kv, len(in))
	i := 0
	for key, value := range in {
		s[i] = kv{key, value}
		i++
	}
	sort.Slice(s, func(i, j int) bool {
		if s[i].v == s[j].v {
			return s[i].k < s[j].k
		}
		return s[i].v < s[j].v
	})
	res := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		res[i] = s[i].k
	}
	return res
}

// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkSort_ByIntKeys-12    	    1474	    734563 ns/op	  471943 B/op	    4935 allocs/op
func BenchmarkSort_ByIntKeys(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort(createMapToSort(), ByKeys, AsInt, false, -1)
	}
}

// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkSort_ByIntKeys_Boilerplate-12    	    3286	    363895 ns/op	  194580 B/op	    2043 allocs/op
func BenchmarkSort_ByIntKeys_Boilerplate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortBoilerplateByIntKeys(createMapToSort())
	}
}

// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkSort_ByLenStringsWithEquals-12    	     535	   2137583 ns/op	  547821 B/op	   23852 allocs/op
func BenchmarkSort_ByLenStringsWithEquals(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort(createMapToSort(), ByValues, AsStringByLength, false, -1)
	}
}

// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkSort_ByLenStringsWithEquals_Boilerplate-12    	    1975	    509752 ns/op	  219252 B/op	    2046 allocs/op
func BenchmarkSort_ByLenStringsWithEquals_Boilerplate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortBoilerplateByLenStringsWithEquals(createMapToSort())
	}
}
