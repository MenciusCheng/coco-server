package parse

import "testing"

func TestFuncString_ToCamel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "single word",
			args: args{s: "singleword"},
			want: "singleword",
		},
		{
			name: "hello world",
			args: args{s: "hello world"},
			want: "helloWorld",
		},
		{
			name: "Hello_world-Example",
			args: args{s: "Hello_world-Example"},
			want: "HelloWorldExample",
		},
		{
			name: "convert_this_to_camel",
			args: args{s: "convert_this_to_camel"},
			want: "convertThisToCamel",
		},
		{
			name: "this-is anotherTest!123",
			args: args{s: "this-is anotherTest!123"},
			want: "thisIsAnotherTest123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FuncString{}
			if got := f.ToCamel(tt.args.s); got != tt.want {
				t.Errorf("ToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuncString_ToLCamel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "single word",
			args: args{s: "singleword"},
			want: "singleword",
		},
		{
			name: "hello world",
			args: args{s: "hello world"},
			want: "helloWorld",
		},
		{
			name: "Hello_world-Example",
			args: args{s: "Hello_world-Example"},
			want: "helloWorldExample",
		},
		{
			name: "convert_this_to_camel",
			args: args{s: "convert_this_to_camel"},
			want: "convertThisToCamel",
		},
		{
			name: "this-is anotherTest!123",
			args: args{s: "this-is anotherTest!123"},
			want: "thisIsAnotherTest123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FuncString{}
			if got := f.ToLCamel(tt.args.s); got != tt.want {
				t.Errorf("ToLCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuncString_ToUCamel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "single word",
			args: args{s: "singleword"},
			want: "Singleword",
		},
		{
			name: "hello world",
			args: args{s: "hello world"},
			want: "HelloWorld",
		},
		{
			name: "Hello_world-Example",
			args: args{s: "Hello_world-Example"},
			want: "HelloWorldExample",
		},
		{
			name: "convert_this_to_camel",
			args: args{s: "convert_this_to_camel"},
			want: "ConvertThisToCamel",
		},
		{
			name: "this-is anotherTest!123",
			args: args{s: "this-is anotherTest!123"},
			want: "ThisIsAnotherTest123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FuncString{}
			if got := f.ToUCamel(tt.args.s); got != tt.want {
				t.Errorf("ToUCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuncString_ToSnake(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "single word",
			args: args{s: "singleword"},
			want: "singleword",
		},
		{
			name: "helloWorld",
			args: args{s: "helloWorld"},
			want: "hello_world",
		},
		{
			name: "HelloWorldExample",
			args: args{s: "HelloWorldExample"},
			want: "hello_world_example",
		},
		{
			name: "convert this to snake",
			args: args{s: "convert this to snake"},
			want: "convert_this_to_snake",
		},
		{
			name: "ThisIsAnotherTest123",
			args: args{s: "ThisIsAnotherTest123"},
			want: "this_is_another_test123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FuncString{}
			if got := f.ToSnake(tt.args.s); got != tt.want {
				t.Errorf("ToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}
