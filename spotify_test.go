package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSliceColumn(t *testing.T) {
	arr := make([][]string, 4)
	for i := range arr {
		arr[i] = make([]string, 4)
		for j := range arr[i] {
			arr[i][j] = "arr" + strconv.Itoa(i*4+j)
		}
	}

	res, err := sliceColumn(arr, 0)
	correct := []string{"arr0", "arr4", "arr8", "arr12"}
	if reflect.DeepEqual(res, correct) == false || err != nil {
		t.Errorf("sliceColumn(arr, 0) = %s; want %s", res, correct)
	}

	res, err = sliceColumn(arr, 3)
	correct = []string{"arr3", "arr7", "arr11", "arr15"}
	if reflect.DeepEqual(res, correct) == false || err != nil {
		t.Errorf("sliceColumn(arr, 3) = %s; want %s", res, correct)
	}

	res, err = sliceColumn(arr, 4)
	correct = []string{}
	if reflect.DeepEqual(res, correct) != false || err == nil {
		t.Errorf("sliceColumn(arr, 4) = %s; want %s", res, correct)
	}
}

func TestSliceStep(t *testing.T) {
	arr := make([]string, 6)
	for i := range arr {
		arr[i] = "arr" + strconv.Itoa(i)
	}

	res, err := sliceStep(arr, 0, 2)
	correct := []string{"arr0", "arr2", "arr4"}
	if reflect.DeepEqual(res, correct) == false || err != nil {
		t.Errorf("sliceStep(arr, 0, 2) = %s; want %s", res, correct)
	}

	res, err = sliceStep(arr, 1, 1)
	correct = []string{"arr1", "arr2", "arr3", "arr4", "arr5"}
	if reflect.DeepEqual(res, correct) == false || err != nil {
		t.Errorf("sliceStep(arr, 1, 1) = %s; want %s", res, correct)
	}

	res, err = sliceStep(arr, 6, 2)
	correct = []string{}
	if reflect.DeepEqual(res, correct) != false || err == nil {
		t.Errorf("sliceStep(arr, 6, 2) = %s; want %s", res, correct)
	}
}

func TestMs_to_hour(t *testing.T) {
	res := ms_to_hour(192715, false)
	correct := "03:12"
	if res != correct {
		t.Errorf("ms_to_hour(192715, false) = %s; want %s", res, correct)
	}

	res = ms_to_hour(192715, true)
	correct = "00:03:12"
	if res != correct {
		t.Errorf("ms_to_hour(192715, true) = %s; want %s", res, correct)
	}

	res = ms_to_hour(8627150, false)
	correct = "23:47"
	if res != correct {
		t.Errorf("ms_to_hour(8627150, false) = %s; want %s", res, correct)
	}

	res = ms_to_hour(8627150, true)
	correct = "02:23:47"
	if res != correct {
		t.Errorf("ms_to_hour(8627150, true) = %s; want %s", res, correct)
	}
}
