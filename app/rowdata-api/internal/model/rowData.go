package model

import "time"

type RowDataUploadInfo struct {
	Weather string
	Temp    int
	WindDir string
	Loc     string
	Coach   string
	Remark  string
}

type TrainingSummary struct {
	TrainingId      uint      //训练ID主键
	TrainingName    string    //训练名称
	TrainDate       time.Time //训练日期`
	EventGender     string    //项目选手性别
	EventPeopleType string    //项目总人数与类型
	EventScale      string    //项目量级
	Event           string    //项目类型
	Weather         string    //天气
	Temp            int       //气温
	WindDir         string    //风向
	Loc             string    //训练地点
	Coach           string    //教练员名称
	SampleCount     int       //采样桨频的数目
	Remark          string    //备注
}

type AthleteTrainingData struct {
	AthleteTrainingId uint    //运动员训练数据ID主键
	TrainingId        uint    //外键 训练ID
	Name              string  //运动员姓名
	Gender            string  //性别
	Seat              int     //运动员座位号
	Side              int     //运动员两侧方位
	Height            float32 //运动员身高
	Weight            float32 //运动员体重
	OarInboard        float32
	OarLength         float32
	OarBladeLength    float32
}

type SampleMetrics struct {
	Id                uint //训练采样指标主键
	AthleteTrainingId uint //外键 运动员训练数据Id（二级训练数据id）

	DataSample                  string  //训练采样的采样名
	StrokeRate                  float64 //桨频(str/min)
	DriveTime                   float64 //拉桨时间(s)
	Rhythm                      float64 // (%)
	CatchAngle                  float64 // (deg)
	FinishAngle                 float64 // (deg)
	TotalAngle                  float64 // (deg)
	CatchSlip                   float64 // (deg)
	ReleaseSlip                 float64 // (deg)
	EffectiveAnglePercent       float64 // (%)
	MaxBladeDepth               float64 // (deg)
	BladeEfficiency             float64 // (%)
	RowingPower                 float64 // (W)
	WorkPerStroke               float64 // (J)
	RelativeWpS                 float64 // (%)
	EffectiveAngleDegree        float64 // (deg)
	TargetAngle                 float64 // 冠军指标
	TargetForce                 float64 // 冠军指标
	TargetWpS                   float64 // 冠军指标
	AngleDivTarget              float64
	ForceDivTarget              float64
	WpSDivTarget                float64
	AverageVelocity             float64 // (m/s)
	BladeSpecificImpulse        float64
	TimeOver2000m               float64 //这个具体怎么存储 time.time 还是 float64 还需要考虑
	MaxForce                    float64 // (N)
	AverageForce                float64 // (N)
	RatioAverDivMaxForce        float64 // (%)
	PositionOfPeakForce         float64 // (% of SL)
	CatchForceGradient          float64 // (deg)
	FinishForceGradient         float64 // (deg)
	MaxHandleVelocity           float64 // (m/s)
	HDF                         float64
	LegsDrive                   float64 // (m)
	LegsMaxSpeed                float64 // (m/s)
	CatchFactor                 float64 // (ms)
	RowingStyleFactor           float64 // (%)
	ReleaseWash                 float64 // (deg)
	AverForceDivWeight          float64 // (N/kg)
	VseatAtCatch                float64 // (m/s)
	HandleTravelAtEntryForce    float64 // (cm)
	HandleTravelAt70PerForce    float64 // (cm)
	HandleTravelAt0As           float64 // (cm)
	SeatTravelAtEntryForce      float64 // (cm)
	SeatTravelAt70PerForce      float64 // (cm)
	SeatTravelAt0As             float64 // (cm)
	DTravelAtEntryForcePercent  float64 // (%)
	DTravelAt70PerForcePercent  float64 // (%)
	DTravelAt0AsPercent         float64 // (%)
	DTravelAtEntryForceDistance float64 // (cm)
	DTravelAt70PerForceDistance float64 // (cm)
	DTravelAt0AsDistance        float64 // (cm)
	SeatOnRecovery              float64
	VertAtCatch                 float64
	EntryForce                  float64
	ForceUpto70Per              float64
	MaxVseat                    float64
	PeakForce                   float64
	ForceFrom70Per              float64
	VertAtFinish                float64
	ForceAtFinish               float64

	AverageBoatSpeed        float64 // (m/s)
	MinimalBoatSpeed        float64 // (m/s)
	MaximalBoatSpeed        float64 // (m/s)
	DistancePerStroke       float64 // (m)
	DragFactor              float64
	WindForwardCompRelWater float64 // (m/s)
	WindDirectionRelWater   float64 // (deg)
	Time250m                float64
	BoatSpeedEfficiency     float64 // (%)
	TimeAtWaterTemp25Deg    float64
	BoatSpeedVariation      float64 // (%)
	WindSpeedRelBoat        float64 // (m/s)
	WindDirectionRelBoat    float64 // (deg)
	DragFactorF             float64 //(F)
	DragFactorPprop         float64
	DragFactorPold          float64
	DragFactorPtot          float64
	AccelerationMinimun     float64 // (m/s^2)
	AccelerationMaximum     float64 // (m/s^2)
	ModelSpeed              float64
	EffectiveWorkPerStroke  float64 // (%)
	ModelDPS                float64
	PropulsivePower         float64 // (W)
	DriveMaximalAt          float64 // (%)
	FirstPeak               float64 // (m/s2)
	ZeroBeforeCatch         float64 // (%)
	MinimalFromCatch        float64 // (%)
	ZeroAfterCatch          float64 // (%)
	StdDeviation            float64 // (m/s)

	OarAngle              [51]float64
	HandleForce           [51]float64
	VerticalAngleBoatRoll [51]float64
	LegsVelocity          [51]float64
	HandleSpeed           [51]float64
	HDFFig                [51]float64
	BladeDF               [51]float64
	Velocity              [51]float64
	BoatAcceleration      [51]float64
	VelocityRel           [51]float64
}
