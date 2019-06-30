/*
@Time : 2019/5/14 10:40
@Author : Lukebryan
@File : tokenutil.go
@Software: GoLand
*/
package utils

import "testing"

func TestGentToken(t *testing.T) {
	type args struct {
		userID    string
		wechat_id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{
				"1",
				"wxid_o9b704atomvj12"},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GentToken(tt.args.userID, tt.args.wechat_id); got != tt.want {
				t.Errorf("GentToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
