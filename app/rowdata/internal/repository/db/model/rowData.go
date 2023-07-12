package model

import "time"

type TrainingSummary struct {
	TrainingId      uint      `gorm:"<-:false;primary_key;AUTO_INCREMENT;comment:'训练ID主键'"`
	TrainingName    string    `gorm:"type:varchar(50);not null;comment:'训练名称'"`
	TrainingDate    time.Time `gorm:"type:Date;not null;comment:'训练日期'"`
	EventGender     string    `gorm:"type:char(1);not null;comment:'项目选手性别'"`
	EventPeopleType string    `gorm:"type:char(2);not null;comment:'项目总人数与类型'"`
	EventScale      string    `gorm:"type:char(1);not null;comment:'项目量级'"`
	Event           string    `gorm:"type:varchar(50);not null;comment:'项目类型'"`
	Weather         string    `gorm:"type:varchar(50);not null;comment:'天气'"`
	Temp            int       `gorm:"not null;comment:'气温'"`
	WindDir         string    `gorm:"type:varchar(50);not null;comment:'风向'"`
	Loc             string    `gorm:"type:varchar(50);not null;comment:'训练地点'"`
	Coach           string    `gorm:"type:varchar(50);not null;comment:'教练员名称'"`
	SampleCount     int       `gorm:"not null;comment:'采样桨频的数目'"`
	Remark          string    `gorm:"type:varchar(50);comment:'备注'"`
}

type AthleteTrainingData struct {
	AthleteTrainingId uint    `gorm:"<-:false;primary_key;AUTO_INCREMENT;comment:'运动员数据ID主键'"`
	TrainingId        uint    `gorm:"type:bigint;not null;comment:'训练数据id'"`
	Name              string  `gorm:"type:varchar(50);not null;comment:'运动员姓名'"`
	Gender            string  `gorm:"type:char(2);not null;comment:'性别'"`
	Seat              int     `gorm:"not null;comment:'运动员座位'"`
	Side              int     `gorm:"not null;comment:'运动员方位'"`
	Height            float32 `gorm:"not null;comment:'运动员身高'"`
	Weight            float32 `gorm:"not null;comment:'运动员体重'"`
	OarInboard        float32 `gorm:"not null"`
	OarLength         float32 `gorm:"not null"`
	OarBladeLength    float32
}

type SampleMetrics struct {
	Id                uint `gorm:"<-:false;primary_key;AUTO_INCREMENT;comment:'训练采样指标主键'"`
	AthleteTrainingId uint `gorm:"not null"`

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
	DragFactor              float64 // (P)
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

	OarAngle_0  float64
	OarAngle_1  float64
	OarAngle_2  float64
	OarAngle_3  float64
	OarAngle_4  float64
	OarAngle_5  float64
	OarAngle_6  float64
	OarAngle_7  float64
	OarAngle_8  float64
	OarAngle_9  float64
	OarAngle_10 float64
	OarAngle_11 float64
	OarAngle_12 float64
	OarAngle_13 float64
	OarAngle_14 float64
	OarAngle_15 float64
	OarAngle_16 float64
	OarAngle_17 float64
	OarAngle_18 float64
	OarAngle_19 float64
	OarAngle_20 float64
	OarAngle_21 float64
	OarAngle_22 float64
	OarAngle_23 float64
	OarAngle_24 float64
	OarAngle_25 float64
	OarAngle_26 float64
	OarAngle_27 float64
	OarAngle_28 float64
	OarAngle_29 float64
	OarAngle_30 float64
	OarAngle_31 float64
	OarAngle_32 float64
	OarAngle_33 float64
	OarAngle_34 float64
	OarAngle_35 float64
	OarAngle_36 float64
	OarAngle_37 float64
	OarAngle_38 float64
	OarAngle_39 float64
	OarAngle_40 float64
	OarAngle_41 float64
	OarAngle_42 float64
	OarAngle_43 float64
	OarAngle_44 float64
	OarAngle_45 float64
	OarAngle_46 float64
	OarAngle_47 float64
	OarAngle_48 float64
	OarAngle_49 float64
	OarAngle_50 float64

	HandleForce_0  float64
	HandleForce_1  float64
	HandleForce_2  float64
	HandleForce_3  float64
	HandleForce_4  float64
	HandleForce_5  float64
	HandleForce_6  float64
	HandleForce_7  float64
	HandleForce_8  float64
	HandleForce_9  float64
	HandleForce_10 float64
	HandleForce_11 float64
	HandleForce_12 float64
	HandleForce_13 float64
	HandleForce_14 float64
	HandleForce_15 float64
	HandleForce_16 float64
	HandleForce_17 float64
	HandleForce_18 float64
	HandleForce_19 float64
	HandleForce_20 float64
	HandleForce_21 float64
	HandleForce_22 float64
	HandleForce_23 float64
	HandleForce_24 float64
	HandleForce_25 float64
	HandleForce_26 float64
	HandleForce_27 float64
	HandleForce_28 float64
	HandleForce_29 float64
	HandleForce_30 float64
	HandleForce_31 float64
	HandleForce_32 float64
	HandleForce_33 float64
	HandleForce_34 float64
	HandleForce_35 float64
	HandleForce_36 float64
	HandleForce_37 float64
	HandleForce_38 float64
	HandleForce_39 float64
	HandleForce_40 float64
	HandleForce_41 float64
	HandleForce_42 float64
	HandleForce_43 float64
	HandleForce_44 float64
	HandleForce_45 float64
	HandleForce_46 float64
	HandleForce_47 float64
	HandleForce_48 float64
	HandleForce_49 float64
	HandleForce_50 float64

	VerticalAngleBoatRoll_0  float64
	VerticalAngleBoatRoll_1  float64
	VerticalAngleBoatRoll_2  float64
	VerticalAngleBoatRoll_3  float64
	VerticalAngleBoatRoll_4  float64
	VerticalAngleBoatRoll_5  float64
	VerticalAngleBoatRoll_6  float64
	VerticalAngleBoatRoll_7  float64
	VerticalAngleBoatRoll_8  float64
	VerticalAngleBoatRoll_9  float64
	VerticalAngleBoatRoll_10 float64
	VerticalAngleBoatRoll_11 float64
	VerticalAngleBoatRoll_12 float64
	VerticalAngleBoatRoll_13 float64
	VerticalAngleBoatRoll_14 float64
	VerticalAngleBoatRoll_15 float64
	VerticalAngleBoatRoll_16 float64
	VerticalAngleBoatRoll_17 float64
	VerticalAngleBoatRoll_18 float64
	VerticalAngleBoatRoll_19 float64
	VerticalAngleBoatRoll_20 float64
	VerticalAngleBoatRoll_21 float64
	VerticalAngleBoatRoll_22 float64
	VerticalAngleBoatRoll_23 float64
	VerticalAngleBoatRoll_24 float64
	VerticalAngleBoatRoll_25 float64
	VerticalAngleBoatRoll_26 float64
	VerticalAngleBoatRoll_27 float64
	VerticalAngleBoatRoll_28 float64
	VerticalAngleBoatRoll_29 float64
	VerticalAngleBoatRoll_30 float64
	VerticalAngleBoatRoll_31 float64
	VerticalAngleBoatRoll_32 float64
	VerticalAngleBoatRoll_33 float64
	VerticalAngleBoatRoll_34 float64
	VerticalAngleBoatRoll_35 float64
	VerticalAngleBoatRoll_36 float64
	VerticalAngleBoatRoll_37 float64
	VerticalAngleBoatRoll_38 float64
	VerticalAngleBoatRoll_39 float64
	VerticalAngleBoatRoll_40 float64
	VerticalAngleBoatRoll_41 float64
	VerticalAngleBoatRoll_42 float64
	VerticalAngleBoatRoll_43 float64
	VerticalAngleBoatRoll_44 float64
	VerticalAngleBoatRoll_45 float64
	VerticalAngleBoatRoll_46 float64
	VerticalAngleBoatRoll_47 float64
	VerticalAngleBoatRoll_48 float64
	VerticalAngleBoatRoll_49 float64
	VerticalAngleBoatRoll_50 float64

	LegsVelocity_0  float64
	LegsVelocity_1  float64
	LegsVelocity_2  float64
	LegsVelocity_3  float64
	LegsVelocity_4  float64
	LegsVelocity_5  float64
	LegsVelocity_6  float64
	LegsVelocity_7  float64
	LegsVelocity_8  float64
	LegsVelocity_9  float64
	LegsVelocity_10 float64
	LegsVelocity_11 float64
	LegsVelocity_12 float64
	LegsVelocity_13 float64
	LegsVelocity_14 float64
	LegsVelocity_15 float64
	LegsVelocity_16 float64
	LegsVelocity_17 float64
	LegsVelocity_18 float64
	LegsVelocity_19 float64
	LegsVelocity_20 float64
	LegsVelocity_21 float64
	LegsVelocity_22 float64
	LegsVelocity_23 float64
	LegsVelocity_24 float64
	LegsVelocity_25 float64
	LegsVelocity_26 float64
	LegsVelocity_27 float64
	LegsVelocity_28 float64
	LegsVelocity_29 float64
	LegsVelocity_30 float64
	LegsVelocity_31 float64
	LegsVelocity_32 float64
	LegsVelocity_33 float64
	LegsVelocity_34 float64
	LegsVelocity_35 float64
	LegsVelocity_36 float64
	LegsVelocity_37 float64
	LegsVelocity_38 float64
	LegsVelocity_39 float64
	LegsVelocity_40 float64
	LegsVelocity_41 float64
	LegsVelocity_42 float64
	LegsVelocity_43 float64
	LegsVelocity_44 float64
	LegsVelocity_45 float64
	LegsVelocity_46 float64
	LegsVelocity_47 float64
	LegsVelocity_48 float64
	LegsVelocity_49 float64
	LegsVelocity_50 float64

	HandleSpeed_0  float64
	HandleSpeed_1  float64
	HandleSpeed_2  float64
	HandleSpeed_3  float64
	HandleSpeed_4  float64
	HandleSpeed_5  float64
	HandleSpeed_6  float64
	HandleSpeed_7  float64
	HandleSpeed_8  float64
	HandleSpeed_9  float64
	HandleSpeed_10 float64
	HandleSpeed_11 float64
	HandleSpeed_12 float64
	HandleSpeed_13 float64
	HandleSpeed_14 float64
	HandleSpeed_15 float64
	HandleSpeed_16 float64
	HandleSpeed_17 float64
	HandleSpeed_18 float64
	HandleSpeed_19 float64
	HandleSpeed_20 float64
	HandleSpeed_21 float64
	HandleSpeed_22 float64
	HandleSpeed_23 float64
	HandleSpeed_24 float64
	HandleSpeed_25 float64
	HandleSpeed_26 float64
	HandleSpeed_27 float64
	HandleSpeed_28 float64
	HandleSpeed_29 float64
	HandleSpeed_30 float64
	HandleSpeed_31 float64
	HandleSpeed_32 float64
	HandleSpeed_33 float64
	HandleSpeed_34 float64
	HandleSpeed_35 float64
	HandleSpeed_36 float64
	HandleSpeed_37 float64
	HandleSpeed_38 float64
	HandleSpeed_39 float64
	HandleSpeed_40 float64
	HandleSpeed_41 float64
	HandleSpeed_42 float64
	HandleSpeed_43 float64
	HandleSpeed_44 float64
	HandleSpeed_45 float64
	HandleSpeed_46 float64
	HandleSpeed_47 float64
	HandleSpeed_48 float64
	HandleSpeed_49 float64
	HandleSpeed_50 float64

	HDF_0  float64
	HDF_1  float64
	HDF_2  float64
	HDF_3  float64
	HDF_4  float64
	HDF_5  float64
	HDF_6  float64
	HDF_7  float64
	HDF_8  float64
	HDF_9  float64
	HDF_10 float64
	HDF_11 float64
	HDF_12 float64
	HDF_13 float64
	HDF_14 float64
	HDF_15 float64
	HDF_16 float64
	HDF_17 float64
	HDF_18 float64
	HDF_19 float64
	HDF_20 float64
	HDF_21 float64
	HDF_22 float64
	HDF_23 float64
	HDF_24 float64
	HDF_25 float64
	HDF_26 float64
	HDF_27 float64
	HDF_28 float64
	HDF_29 float64
	HDF_30 float64
	HDF_31 float64
	HDF_32 float64
	HDF_33 float64
	HDF_34 float64
	HDF_35 float64
	HDF_36 float64
	HDF_37 float64
	HDF_38 float64
	HDF_39 float64
	HDF_40 float64
	HDF_41 float64
	HDF_42 float64
	HDF_43 float64
	HDF_44 float64
	HDF_45 float64
	HDF_46 float64
	HDF_47 float64
	HDF_48 float64
	HDF_49 float64
	HDF_50 float64

	BladeDF_0  float64
	BladeDF_1  float64
	BladeDF_2  float64
	BladeDF_3  float64
	BladeDF_4  float64
	BladeDF_5  float64
	BladeDF_6  float64
	BladeDF_7  float64
	BladeDF_8  float64
	BladeDF_9  float64
	BladeDF_10 float64
	BladeDF_11 float64
	BladeDF_12 float64
	BladeDF_13 float64
	BladeDF_14 float64
	BladeDF_15 float64
	BladeDF_16 float64
	BladeDF_17 float64
	BladeDF_18 float64
	BladeDF_19 float64
	BladeDF_20 float64
	BladeDF_21 float64
	BladeDF_22 float64
	BladeDF_23 float64
	BladeDF_24 float64
	BladeDF_25 float64
	BladeDF_26 float64
	BladeDF_27 float64
	BladeDF_28 float64
	BladeDF_29 float64
	BladeDF_30 float64
	BladeDF_31 float64
	BladeDF_32 float64
	BladeDF_33 float64
	BladeDF_34 float64
	BladeDF_35 float64
	BladeDF_36 float64
	BladeDF_37 float64
	BladeDF_38 float64
	BladeDF_39 float64
	BladeDF_40 float64
	BladeDF_41 float64
	BladeDF_42 float64
	BladeDF_43 float64
	BladeDF_44 float64
	BladeDF_45 float64
	BladeDF_46 float64
	BladeDF_47 float64
	BladeDF_48 float64
	BladeDF_49 float64
	BladeDF_50 float64

	Velocity_0  float64
	Velocity_1  float64
	Velocity_2  float64
	Velocity_3  float64
	Velocity_4  float64
	Velocity_5  float64
	Velocity_6  float64
	Velocity_7  float64
	Velocity_8  float64
	Velocity_9  float64
	Velocity_10 float64
	Velocity_11 float64
	Velocity_12 float64
	Velocity_13 float64
	Velocity_14 float64
	Velocity_15 float64
	Velocity_16 float64
	Velocity_17 float64
	Velocity_18 float64
	Velocity_19 float64
	Velocity_20 float64
	Velocity_21 float64
	Velocity_22 float64
	Velocity_23 float64
	Velocity_24 float64
	Velocity_25 float64
	Velocity_26 float64
	Velocity_27 float64
	Velocity_28 float64
	Velocity_29 float64
	Velocity_30 float64
	Velocity_31 float64
	Velocity_32 float64
	Velocity_33 float64
	Velocity_34 float64
	Velocity_35 float64
	Velocity_36 float64
	Velocity_37 float64
	Velocity_38 float64
	Velocity_39 float64
	Velocity_40 float64
	Velocity_41 float64
	Velocity_42 float64
	Velocity_43 float64
	Velocity_44 float64
	Velocity_45 float64
	Velocity_46 float64
	Velocity_47 float64
	Velocity_48 float64
	Velocity_49 float64
	Velocity_50 float64

	BoatAcceleration_0  float64
	BoatAcceleration_1  float64
	BoatAcceleration_2  float64
	BoatAcceleration_3  float64
	BoatAcceleration_4  float64
	BoatAcceleration_5  float64
	BoatAcceleration_6  float64
	BoatAcceleration_7  float64
	BoatAcceleration_8  float64
	BoatAcceleration_9  float64
	BoatAcceleration_10 float64
	BoatAcceleration_11 float64
	BoatAcceleration_12 float64
	BoatAcceleration_13 float64
	BoatAcceleration_14 float64
	BoatAcceleration_15 float64
	BoatAcceleration_16 float64
	BoatAcceleration_17 float64
	BoatAcceleration_18 float64
	BoatAcceleration_19 float64
	BoatAcceleration_20 float64
	BoatAcceleration_21 float64
	BoatAcceleration_22 float64
	BoatAcceleration_23 float64
	BoatAcceleration_24 float64
	BoatAcceleration_25 float64
	BoatAcceleration_26 float64
	BoatAcceleration_27 float64
	BoatAcceleration_28 float64
	BoatAcceleration_29 float64
	BoatAcceleration_30 float64
	BoatAcceleration_31 float64
	BoatAcceleration_32 float64
	BoatAcceleration_33 float64
	BoatAcceleration_34 float64
	BoatAcceleration_35 float64
	BoatAcceleration_36 float64
	BoatAcceleration_37 float64
	BoatAcceleration_38 float64
	BoatAcceleration_39 float64
	BoatAcceleration_40 float64
	BoatAcceleration_41 float64
	BoatAcceleration_42 float64
	BoatAcceleration_43 float64
	BoatAcceleration_44 float64
	BoatAcceleration_45 float64
	BoatAcceleration_46 float64
	BoatAcceleration_47 float64
	BoatAcceleration_48 float64
	BoatAcceleration_49 float64
	BoatAcceleration_50 float64

	VelocityRel_0  float64
	VelocityRel_1  float64
	VelocityRel_2  float64
	VelocityRel_3  float64
	VelocityRel_4  float64
	VelocityRel_5  float64
	VelocityRel_6  float64
	VelocityRel_7  float64
	VelocityRel_8  float64
	VelocityRel_9  float64
	VelocityRel_10 float64
	VelocityRel_11 float64
	VelocityRel_12 float64
	VelocityRel_13 float64
	VelocityRel_14 float64
	VelocityRel_15 float64
	VelocityRel_16 float64
	VelocityRel_17 float64
	VelocityRel_18 float64
	VelocityRel_19 float64
	VelocityRel_20 float64
	VelocityRel_21 float64
	VelocityRel_22 float64
	VelocityRel_23 float64
	VelocityRel_24 float64
	VelocityRel_25 float64
	VelocityRel_26 float64
	VelocityRel_27 float64
	VelocityRel_28 float64
	VelocityRel_29 float64
	VelocityRel_30 float64
	VelocityRel_31 float64
	VelocityRel_32 float64
	VelocityRel_33 float64
	VelocityRel_34 float64
	VelocityRel_35 float64
	VelocityRel_36 float64
	VelocityRel_37 float64
	VelocityRel_38 float64
	VelocityRel_39 float64
	VelocityRel_40 float64
	VelocityRel_41 float64
	VelocityRel_42 float64
	VelocityRel_43 float64
	VelocityRel_44 float64
	VelocityRel_45 float64
	VelocityRel_46 float64
	VelocityRel_47 float64
	VelocityRel_48 float64
	VelocityRel_49 float64
	VelocityRel_50 float64
}

func (*TrainingSummary) TableName() string {
	return "training_summary"
}

func (*AthleteTrainingData) TableName() string {
	return "athlete_training_data"
}

func (*SampleMetrics) TableName() string {
	return "sample_metrics"
}
