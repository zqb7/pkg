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
	// Self *TestObj
}

func TestGen(t *testing.T) {
	type args struct {
		t    any
		wont any
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{t: []int{}, wont: []int{0}}},
		{name: "2", args: args{t: []*int{}, wont: []*int{new(int)}}},
		{name: "3", args: args{t: &TestObj{}, wont: &TestObj{A: 0, B: "", C: new(int), D: new(string)}}},
		{name: "4", args: args{t: TestObj{}, wont: TestObj{A: 0, B: "", C: new(int), D: new(string)}}},
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
