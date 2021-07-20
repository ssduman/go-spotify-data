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
	if reflect.DeepEqual(res, []string{"arr0", "arr4", "arr8", "arr12"}) == false || err != nil {
		t.Errorf("sliceColumn(arr, 0) = %s; want %s", res, []string{"arr0", "arr4", "arr8", "arr12"})
	}

	res, err = sliceColumn(arr, 3)
	if reflect.DeepEqual(res, []string{"arr3", "arr7", "arr11", "arr15"}) == false || err != nil {
		t.Errorf("sliceColumn(arr, 3) = %s; want %s", res, []string{"arr3", "arr7", "arr11", "arr15"})
	}

	res, err = sliceColumn(arr, 4)
	if reflect.DeepEqual(res, []string{}) != false || err == nil {
		t.Errorf("sliceColumn(arr, 4) = %s; want %s", res, []string{})
	}
}

func TestSliceStep(t *testing.T) {
	arr := make([]string, 6)
	for i := range arr {
		arr[i] = "arr" + strconv.Itoa(i)
	}

	res, err := sliceStep(arr, 0, 2)
	if reflect.DeepEqual(res, []string{"arr0", "arr2", "arr4"}) == false || err != nil {
		t.Errorf("sliceStep(arr, 0, 2) = %s; want %s", res, []string{"arr0", "arr2", "arr4"})
	}

	res, err = sliceStep(arr, 1, 1)
	if reflect.DeepEqual(res, []string{"arr1", "arr2", "arr3", "arr4", "arr5"}) == false || err != nil {
		t.Errorf("sliceStep(arr, 1, 1) = %s; want %s", res, []string{"arr1", "arr2", "arr3", "arr4", "arr5"})
	}

	res, err = sliceStep(arr, 6, 2)
	if reflect.DeepEqual(res, []string{}) != false || err == nil {
		t.Errorf("sliceStep(arr, 6, 2 = %s; want %s", res, []string{})
	}
}

func TestMs_to_hour(t *testing.T) {
	res := ms_to_hour(192715, false)
	if res != "03:12" {
		t.Errorf("Testms_to_hour(192715, false) = %s; want 03:12", res)
	}

	res = ms_to_hour(192715, true)
	if res != "00:03:12" {
		t.Errorf("Testms_to_hour(192715, false) = %s; want 00:03:12", res)
	}

	res = ms_to_hour(8627150, false)
	if res != "23:47" {
		t.Errorf("Testms_to_hour(8627150, false) = %s; want 23:47", res)
	}

	res = ms_to_hour(8627150, true)
	if res != "02:23:47" {
		t.Errorf("Testms_to_hour(8627150, false) = %s; want 02:23:47", res)
	}
}
