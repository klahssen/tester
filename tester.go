package tester

import (
	"reflect"
	"testing"
)

//T holds testing methods
type T struct {
	t *testing.T
}

//NewT returns an instance of a tester
func NewT(t *testing.T) *T {
	if t == nil {
		t = &testing.T{}
	}
	return &T{t: t}
}

//CheckError compares errors
func (t *T) CheckError(testID string, exp, err error) {
	if (err != nil && exp == nil) || (err == nil && exp != nil) || (err != nil && !reflect.DeepEqual(exp, err)) {
		t.t.Errorf("test '%s': expected %+v received %+v", testID, exp, err)
	}
}

//CheckVal compares 2 values
func (t *T) CheckVal(testID, valName string, exp, val float64) {
	if !reflect.DeepEqual(exp, val) {
		t.t.Errorf("test '%s': expected %s %+v received %+v", testID, valName, exp, val)
	}
}
