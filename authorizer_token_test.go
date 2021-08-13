package wx

import (
	"fmt"
	"net/http"
	"testing"
)

var testClient = NewWeChatClient()

func init() {
	testClient.SetHttpClient(&http.Client{})
}

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

			fmt.Println(got, err)
		})
	}
}