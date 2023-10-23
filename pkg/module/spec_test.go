package module

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestTypeSpec_Generate(t *testing.T) {
	type fields struct {
		TypeSpec *ast.StructType
		Fields   []Field
		Ignored  bool
		Version  string
		APIPath  string
	}
	type args struct {
		template string
		path     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *File
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "BasicTest - no template",
			fields: fields{
				TypeSpec: nil,
				Fields:   nil,
				Ignored:  false,
				Version:  "",
				APIPath:  "",
			},
			args: args{
				template: "",
				path:     "",
			},
			want:    nil,
			wantErr: ErrorNoTemplate,
		},
		{
			name: "BasicTest - simple template",
			fields: fields{
				TypeSpec: nil,
				Fields:   nil,
				Ignored:  false,
				Version:  "",
				APIPath:  "",
			},
			args: args{
				template: "hi",
				path:     "hi",
			},
			want: &File{
				Path:    "hi",
				Content: "hi",
			},
			wantErr: nil,
		},
		{
			name: "BasicTest - template with Version",
			fields: fields{
				TypeSpec: nil,
				Fields:   nil,
				Ignored:  false,
				Version:  "111",
				APIPath:  "",
			},
			args: args{
				template: "{{.Version}}",
				path:     "hi",
			},
			want: &File{
				Path:    "hi",
				Content: "111",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TypeSpec{
				TypeSpec: tt.fields.TypeSpec,
				Fields:   tt.fields.Fields,
				Ignored:  tt.fields.Ignored,
				Version:  tt.fields.Version,
				APIPath:  tt.fields.APIPath,
			}
			got, err := ts.Generate(tt.args.template, tt.args.path, tt.name)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GenerateHandler() error = %v, wantErr %v", got, tt.want)
			}
			// We set the want name to the name of the test
			if tt.want != nil {
				tt.want.Name = tt.name
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GenerateHandler() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
