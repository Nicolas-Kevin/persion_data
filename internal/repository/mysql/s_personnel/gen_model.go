package s_personnel

// SPersonnel 人员表
//go:generate gormgen -structs SPersonnel -input .
type SPersonnel struct {
	Id               int64  //
	UserName         string // 姓名
	UserAge          int32  // 年龄
	UserSex          int32  // 性别（0：女，1：男）
	UserDate         string // 出生年月日
	PoliticalOutLook int32  // 社会面貌：0 :普通公民，1：团员，2：党员
	Company          string // 工作单位
	Position         string // 职称/等级
	Major            string // 专业
	UserLevel        string // 级别
	UserHeight       int32  // 身高cm
	UserWeight       int32  // 体重cm
	ShoeSize         int32  // 鞋号cm
	HeadNum          int32  // 头围cm
	BloodType        string // 血型cm
	IdNum            string // 身份证号
	BankNum          string // 银行卡号
	PhoneNum         string // 电话号
	UserTeam         int64  // 部队编号
	TrainStatus      string // 演训情况
	RewardStatus     string // 奖励情况
	Deleted          int32  //
}
