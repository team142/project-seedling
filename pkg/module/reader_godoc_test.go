package module

import "testing"

func Test_commentContains(t *testing.T) {
	type args struct {
		doc        string
		lookingFor string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Simple false",
			args: args{
				doc:        `User is a user`,
				lookingFor: "#pk/n",
			},
			want: false,
		},
		{
			name: "Simple true",
			args: args{
				doc:        `#pk`,
				lookingFor: "#pk",
			},
			want: true,
		},
		{
			name: "Uppercase true",
			args: args{
				doc:        `#PK`,
				lookingFor: "#pk",
			},
			want: true,
		},
		{
			name: "Simple false no new line",
			args: args{
				doc: `User is a user
#pk`,
				lookingFor: "#pk\n",
			},
			want: false,
		},
		{
			name: "Simple true with a new line",
			args: args{
				doc: `User is a user
#pk
`,
				lookingFor: "#pk\n",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commentContains(tt.args.doc, tt.args.lookingFor); got != tt.want {
				t.Errorf("commentContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getComment(t *testing.T) {
	type args struct {
		doc        string
		lookingFor string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple test no result",
			args: args{
				doc:        `User is a user`,
				lookingFor: "#pk",
			},
			want: "",
		},
		{
			name: "Simple test with result",
			args: args{
				doc:        `#pk User is a user`,
				lookingFor: "#pk",
			},
			want: "#pk User is a user",
		},
		{
			name: "Simple test with result - multi lines",
			args: args{
				doc: `User is a user
#pk User is a user`,
				lookingFor: "#pk",
			},
			want: "#pk User is a user",
		},
		{
			name: "Simple test with result - multi lines - 2",
			args: args{
				doc: `User is a user
#pk User is a user
User is a user`,

				lookingFor: "#pk",
			},
			want: "#pk User is a user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getComment(tt.args.doc, tt.args.lookingFor); got != tt.want {
				t.Errorf("getComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
