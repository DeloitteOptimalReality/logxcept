package trace

import "testing"

func TestTrace(t *testing.T) {
	tr := RandStringBytes(1)
	if len(tr) != 1 {
		t.Errorf("wrong length trace. got %v instead of 1", len(tr))
	}

	tr = RandStringBytes(0)
	if tr != "" {
		t.Fail()
	}

}
