package wx

import "net/http"

// app document https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/weapp-wxverify/secwxaapi_wxaauth.html

const ApiSecAuthidentityTree = "https://api.weixin.qq.com/wxa/sec/authidentitytree"

func (client *WeChatClient) SecAuthidentityTree(data *SecAuthidentityTreeRequest) (*SecAuthidentityTreeResponse, error) {
	req := &CommonRequest{}
	req.WithURL(ApiSecAuthidentityTree).
		WithMethod(http.MethodGet).
		WithContentType(MineJson).
		WithData(data)
	rsp := &struct {
		CommonResponse
		SecAuthidentityTreeResponse
	}{}
	if err := client.DoRequest(req, rsp); err != nil {
		return nil, err
	}
	if err := rsp.Error(); err != nil {
		return nil, err
	}
	return &rsp.SecAuthidentityTreeResponse, nil
}

func (client *WeChatClient) BuildSecAuthidentityTreeRequest() *SecAuthidentityTreeRequest {
	return &SecAuthidentityTreeRequest{}
}

type SecAuthidentityTreeRequest struct {
	AccessToken string `position:"query" name:"access_token" json:"-"`
}

type SecAuthidentityLeafInfo struct {
	Requirement string `json:"requirement"` //
}
type SecAuthidentityRootInfo struct {
	Type int64 `json:"type"`
}

type SecAuthidentity struct {
	Name     string                   `json:"name"`                //名称
	NodeID   int64                    `json:"node_id"`             //节点编号
	NodeList []*SecAuthidentity       `json:"node_list"`           //节点列表
	LeafInfo *SecAuthidentityLeafInfo `json:"leaf_info,"`          //叶子节点信息
	RootType *SecAuthidentityRootInfo `json:"root_type,omitempty"` // 根节点信息
}
type SecAuthidentityTreeResponse struct {
	IdentityTreeList []*SecAuthidentity `json:"identity_tree_list"`
}
