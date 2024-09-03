package parse

import "testing"

func TestFuncCompare_EqStr(t *testing.T) {
	type args struct {
		s1 interface{}
		s2 interface{}
	}

	var app interface{}
	app = "app"

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s1: "app",
				s2: app,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &FuncCompare{}
			if got := c.EqStr(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("EqStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
