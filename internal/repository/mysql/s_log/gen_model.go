package s_log

import "time"

// SLog 操作日志表
//go:generate gormgen -structs SLog -input .
type SLog struct {
	Id        int32  //
	OperaUser int32  // 具体操作
	OperaTime string // 操作时间
	OperaId   string // 操作人
}
