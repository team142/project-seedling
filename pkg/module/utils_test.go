package module

import "testing"

func Test_toSnakeCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Generic",
			args: args{
				input: "HelloThere",
			},
			want: "hello_there",
		},
		{
			name: "Generic spaced",
			args: args{
				input: "Hello There",
			},
			want: "hello_there",
		},
		{
			name: "Generic dashed",
			args: args{
				input: "Hello-There",
			},
			want: "hello_there",
		},
		{
			name: "Generic underscore",
			args: args{
				input: "Hello_There",
			},
			want: "hello_there",
		},
		{
			name: "Generic first small letter",
			args: args{
				input: "helloThere",
			},
			want: "hello_there",
		},
		{
			name: "Simple",
			args: args{
				input: "Hello",
			},
			want: "hello",
		},
		{
			name: "Simple with space",
			args: args{
				input: "Hello ",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toSnakeCase(tt.args.input); got != tt.want {
				t.Errorf("toSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toCamelCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test",
			args: args{
				input: "Test",
			},
			want: "test",
		},
		{
			name: "Test ",
			args: args{
				input: "Test ",
			},
			want: "test",
		},
		{
			name: "Test_",
			args: args{
				input: "Test_",
			},
			want: "test",
		},
		{
			name: "TestTest",
			args: args{
				input: "TestTest",
			},
			want: "testTest",
		},
		{
			name: "Test Test",
			args: args{
				input: "Test Test",
			},
			want: "testTest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toCamelCase(tt.args.input); got != tt.want {
				t.Errorf("toCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toPascalCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test",
			args: args{
				input: "Test",
			},
			want: "Test",
		},
		{
			name: "Test ",
			args: args{
				input: "Test ",
			},
			want: "Test",
		},
		{
			name: "Test_",
			args: args{
				input: "Test_",
			},
			want: "Test",
		},
		{
			name: "TestTest",
			args: args{
				input: "TestTest",
			},
			want: "TestTest",
		},
		{
			name: "Test Test",
			args: args{
				input: "Test Test",
			},
			want: "TestTest",
		},
		{
			name: "Test_Test",
			args: args{
				input: "Test_Test",
			},
			want: "TestTest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toPascalCase(tt.args.input); got != tt.want {
				t.Errorf("toPascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
