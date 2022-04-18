package examples

import (
	"context"
)

// CustomCredential 自定义认证
type CustomCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c CustomCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "620f9a6fb47b05e7",
		"appkey": "65356533-3539-6464-6330-65653738",
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c CustomCredential) RequireTransportSecurity() bool {
	return false
}
