package mapsorter

// cSpell:disable
import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"time"

	"github.com/mehanizm/dateparse"
)

type By uint8

const (
	ByKeys By = iota
	ByValues
)

type As uint8

const (
	AsString As = iota
	AsStringByLength
	AsInt
	AsFloat
	AsDatetime
)

func Sort(mapToSortInterface interface{}, sortBy By,
	sortAs As, reverse bool, top int) (res []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("unknown panic: %v", r)
			}
			res = nil
		}
	}()
	mapToSort := make(map[interface{}]interface{})
	if reflect.ValueOf(mapToSortInterface).Kind() == reflect.Map {
		for _, key := range reflect.ValueOf(mapToSortInterface).MapKeys() {
			value := reflect.ValueOf(mapToSortInterface).MapIndex(key)
			mapToSort[key.Interface()] = value.Interface()
		}
	} else {
		return nil, errors.New("not map type to sort")
	}
	sliceSortHelper := make([][]interface{}, len(mapToSort))
	i := 0
	for key, value := range mapToSort {
		sliceSortHelper[i] = []interface{}{key, value}
		i++
	}
	switch sortAs {
	case AsString:
		sort.Slice(sliceSortHelper, func(i, j int) bool {
			if reverse {
				i, j = j, i
			}
			if sortBy == ByValues &&
				sliceSortHelper[i][sortBy].(string) == sliceSortHelper[j][sortBy].(string) {
				return fmt.Sprintf("%v", sliceSortHelper[i][ByKeys]) < fmt.Sprintf("%v", sliceSortHelper[j][ByKeys])
			}
			return sliceSortHelper[i][sortBy].(string) < sliceSortHelper[j][sortBy].(string)
		})
	case AsStringByLength:
		sort.Slice(sliceSortHelper, func(i, j int) bool {
			if reverse {
				i, j = j, i
			}
			if sortBy == ByValues &&
				len(sliceSortHelper[i][sortBy].(string)) == len(sliceSortHelper[j][sortBy].(string)) {
				return fmt.Sprintf("%v", sliceSortHelper[i][ByKeys]) < fmt.Sprintf("%v", sliceSortHelper[j][ByKeys])
			}
			return len(sliceSortHelper[i][sortBy].(string)) < len(sliceSortHelper[j][sortBy].(string))
		})
	case AsInt:
		sort.Slice(sliceSortHelper, func(i, j int) bool {
			if reverse {
				i, j = j, i
			}
			var iInt, jInt int
			if _, ok := sliceSortHelper[i][sortBy].(int); !ok {
				iInt, err = strconv.Atoi(sliceSortHelper[i][sortBy].(string))
				if err != nil {
					panic(fmt.Errorf("cannot convert string to int: %w", err))
				}
				jInt, err = strconv.Atoi(sliceSortHelper[j][sortBy].(string))
				if err != nil {
					panic(fmt.Errorf("cannot convert string to int: %w", err))
				}
			} else {
				iInt, _ = sliceSortHelper[i][sortBy].(int)
				jInt, _ = sliceSortHelper[j][sortBy].(int)
			}
			if sortBy == ByValues && iInt == jInt {
				return fmt.Sprintf("%v", sliceSortHelper[i][ByKeys]) < fmt.Sprintf("%v", sliceSortHelper[j][ByKeys])
			}
			return iInt < jInt
		})
	case AsFloat:
		sort.Slice(sliceSortHelper, func(i, j int) bool {
			if reverse {
				i, j = j, i
			}
			var iFloat, jFloat float64
			if _, ok := sliceSortHelper[i][sortBy].(float64); !ok {
				iFloat, err = strconv.ParseFloat(sliceSortHelper[i][sortBy].(string), 64)
				if err != nil {
					panic(fmt.Errorf("cannot convert string to float: %w", err))
				}
				jFloat, err = strconv.ParseFloat(sliceSortHelper[j][sortBy].(string), 64)
				if err != nil {
					panic(fmt.Errorf("cannot convert string to float: %w", err))
				}
			} else {
				iFloat, _ = sliceSortHelper[i][sortBy].(float64)
				jFloat, _ = sliceSortHelper[j][sortBy].(float64)
			}
			if sortBy == ByValues && iFloat == jFloat {
				return fmt.Sprintf("%v", sliceSortHelper[i][ByKeys]) < fmt.Sprintf("%v", sliceSortHelper[j][ByKeys])
			}
			return iFloat < jFloat
		})
	case AsDatetime:
		sort.Slice(sliceSortHelper, func(i, j int) bool {
			if reverse {
				i, j = j, i
			}
			var iDt, jDt time.Time
			if _, ok := sliceSortHelper[i][sortBy].(time.Time); !ok {
				iDt, err = dateparse.ParseAny(sliceSortHelper[i][sortBy].(string),
					dateparse.RetryAmbiguousDateWithSwap(true),
					dateparse.PreferMonthFirst(false))
				if err != nil {
					panic(fmt.Errorf("cannot convert string to datetime: %w", err))
				}
				jDt, err = dateparse.ParseAny(sliceSortHelper[j][sortBy].(string),
					dateparse.RetryAmbiguousDateWithSwap(true),
					dateparse.PreferMonthFirst(false))
				if err != nil {
					panic(fmt.Errorf("cannot convert string to datetime: %w", err))
				}
			} else {
				iDt, _ = sliceSortHelper[i][sortBy].(time.Time)
				jDt, _ = sliceSortHelper[j][sortBy].(time.Time)
			}
			if sortBy == ByValues && iDt == jDt {
				return fmt.Sprintf("%v", sliceSortHelper[i][ByKeys]) < fmt.Sprintf("%v", sliceSortHelper[j][ByKeys])
			}
			return iDt.Before(jDt)
		})
	}
	res = make([]interface{}, len(mapToSort))
	for i, item := range sliceSortHelper {
		res[i] = item[0]
	}
	if top < len(mapToSort) && top > 0 {
		res = res[:top]
	}
	return
}
