package app

// 返回自定义状态码
const (
	Success          = 0
	PermissionDenied = 403
	NotFound         = 404
	Fail             = 500
	AuthFail         = 401
)

// 自定义的一些错误消息的返回
const (
	AuthFailMessage         = "AuthFailed"
	PermissionDeniedMessage = "PermissionDenied"
)
