/*
@Time : 2019/5/14 9:30
@Author : Lukebryan
@File : httputil.go
@Software: GoLand
*/
package utils

import (
	"testing"
)

func TestPostFormRequest(t *testing.T) {
	type args struct {
		path     string
		dataByte map[string]interface{}
		headers  map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{
				"http://192.168.1.156:8989/api/v1/account/batchcheckloginstatus",
				map[string]interface{}{"wxid_list":"wxid_8axqag6ij0kt12"},
				map[string]string{"Authorization":GentToken("1","")}},
			"",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostFormRequest(tt.args.path, tt.args.dataByte, tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostFormRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PostFormRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
