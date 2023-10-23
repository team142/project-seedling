package module

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_getProjectInformation(t *testing.T) {
	getwd, err := os.Getwd()
	if err != nil {
		return
	}

	moduleArray := strings.Split(reflect.TypeOf(Config{}).PkgPath(), "/")
	module := moduleArray[0]

	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name: "core test",
			args: args{
				filePath: getwd,
			},
			want:    module,
			want1:   "pkg/module",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getProjectInformation(tt.args.filePath, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("getProjectInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getProjectInformation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getProjectInformation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
