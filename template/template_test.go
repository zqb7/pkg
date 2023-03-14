package template

import (
	"encoding/json"
	"testing"
)

type TestObj struct {
	A int
	B string
	C *int
	D *string
	E []int
	// Self *TestObj
}

type TestObj2 struct {
	a *int
	A *int
	B struct {
		A2 []int
		A3 []*int
	}
}

func TestGen(t *testing.T) {
	type args struct {
		t    interface{}
		wont interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{t: []int{}, wont: []int{0}}},
		{name: "2", args: args{t: []*int{}, wont: []*int{new(int)}}},
		{name: "3", args: args{t: &TestObj{}, wont: &TestObj{A: 0, B: "", C: new(int), D: new(string), E: []int{0}}}},
		{name: "4", args: args{t: TestObj{}, wont: TestObj{A: 0, B: "", C: new(int), D: new(string), E: []int{0}}}},
		{name: "5", args: args{t: TestObj2{}, wont: TestObj2{A: new(int), B: struct {
			A2 []int
			A3 []*int
		}{A2: []int{0}, A3: []*int{new(int)}}}}},
		{name: "6", args: args{t: &TestObj2{}, wont: TestObj2{A: new(int), B: struct {
			A2 []int
			A3 []*int
		}{A2: []int{0}, A3: []*int{new(int)}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Gen(tt.args.t)
			gotByte, err := json.Marshal(got)
			if err != nil {
				t.Fatal(err)
			}
			wontByte, err := json.Marshal(tt.args.wont)
			if err != nil {
				t.Fatal(err)
			}
			if string(gotByte) != string(wontByte) {
				t.Fatalf("wont:%+v got:%+v,", tt.args.t, got)
			}
			t.Logf("got:%s want:%s", gotByte, wontByte)
		})
	}
}
