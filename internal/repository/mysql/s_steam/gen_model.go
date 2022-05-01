package s_steam

// SSteam 队伍信息表
//go:generate gormgen -structs SSteam -input .
type SSteam struct {
	Id         int64  //
	TeamName   string // 队伍名称
	TeamNum    int32  // 队伍编号
	TeamZbName string // 保障zb编号
	CreateTime string // 组建时间
	TeamSize   int32  // 队伍人数
	TeamContem string // 参训信息
	Deleted    int32  //
}
