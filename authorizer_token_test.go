package wx

import (
	"net/http"
	"reflect"
	"testing"
)

var testClient = NewWeChatClient()

func TestWeChatClient_AuthorizerToken(t *testing.T) {
	type fields struct {
		httpProxy  string
		httpsProxy string
		httpClient *http.Client
	}
	type args struct {
		req *AuthorizerTokenRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AuthorizerTokenResponse
		wantErr bool
	}{
		{ 
			args: args{
				req: &AuthorizerTokenRequest{
					ComponentAccessToken:   "",
					ComponentAppID:         "",
					AuthorizerAppID:        "",
					AuthorizerRefreshToken: "",
				},
			},
		},
	}

		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.AuthorizerToken(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthorizerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizerToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}