package s_train

// STrain 演训表
//go:generate gormgen -structs STrain -input .
type STrain struct {
	Id           int32  //
	TrainName    string // 演训名称
	TrainData    string // 时间年月日
	TrainOrg     string // 地点（市）
	TrainNum     int32  // 参赛人员数量
	TrainProject string // 科目
	TrainPlace   string // 组织方
	TrainTime    string // 宣传图片，视频
	Deleted      int32  //
}
