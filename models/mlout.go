package models

import (
	"fmt"
	"time"
)

type Mlout struct {
	Id int64 `gorm:"primary key;AUTO_INCREMENT"`
	Time time.Time
	StationName string `gorm:"type:varchar(50)"`
	SubstanceType int8 `gorm:"not null"`
	PredictionValue float32
}

func ListMlouts (beginTimePoint string, hour string, substanceType string) [] Mlout{

	fmt.Println(beginTimePoint)
	const base_format = "2006-01-02 15:04:05"
	beginTime, _:= time.Parse(base_format, beginTimePoint)
	fmt.Println(beginTime.Format(base_format))

	duration, _ := time.ParseDuration("24h")
	endTime := beginTime.Add(duration)
	endTimePoint := endTime.Format(base_format)
	fmt.Println(beginTimePoint)
	fmt.Println(endTimePoint)
	var substance string
	if substanceType == "PM25" {
		substance = "1"
	} else {
		substance = "2"
	}
	var mlouts [] Mlout
	db.Where("time >= ? AND time <= ? AND substance_type = ?", beginTimePoint, endTimePoint, substance).Find(&mlouts)
	return mlouts
}