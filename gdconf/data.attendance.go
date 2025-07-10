package gdconf

import (
	"encoding/json"
	"fmt"
	"github.com/gucooing/BaPs/pkg/logger"
	"os"
)

type Attendance struct {
	AttendanceMap map[int64]*AttendanceInfo
}

type AttendanceInfo struct {
	Id                int64                       `json:"Id"`
	Type              string                      `json:"Type"`
	AccountType       string                      `json:"AccountType"`
	AccountLevelLimit int64                       `json:"AccountLevelLimit"`
	Title             string                      `json:"Title"`
	TitleImagePath    string                      `json:"TitleImagePath"`
	CountRule         string                      `json:"CountRule"`
	CountReset        string                      `json:"CountReset"`
	BookSize          int64                       `json:"BookSize"`
	StartTime         Time                        `json:"StartDate"`
	StartableEndTime  Time                        `json:"StartableEndDate"`
	EndTime           Time                        `json:"EndDate"`
	MailType          string                      `json:"MailType"`
	AttendanceReward  map[int64]*AttendanceReward `json:"AttendanceReward"`
}

type AttendanceReward struct {
	RewardParcelType int32 `json:"RewardParcelType"`
	RewardId         int64 `json:"RewardId"`
	RewardAmoun      int64 `json:"RewardAmoun"`
}

func (g *GameConfig) loadAttendance() {
	g.GetGPP().Attendance = &Attendance{
		AttendanceMap: make(map[int64]*AttendanceInfo),
	}
	name := "Attendance.json"
	file, err := os.ReadFile(g.dataPath + name)
	if err != nil {
		logger.Error("文件:%s 读取失败,err:%s", name, err)
		return
	}
	if err := json.Unmarshal(file, &g.GetGPP().Attendance.AttendanceMap); err != nil {
		logger.Error("文件:%s 解析失败,err:%s", name, err)
		return
	}

	for _, v := range g.GetGPP().Attendance.AttendanceMap {
		// 检查排期奖励是否合法
		if v.BookSize != int64(len(v.AttendanceReward)) {
			panic(fmt.Sprintf("登录奖励错误 排期奖励不足,Id:%v", v.Id))
		}
	}
	logger.Info("登录奖励排期读取成功文件:%s 读取成功,解析数量:%v", name, len(g.GetGPP().Attendance.AttendanceMap))
}

func GetAttendanceMap() map[int64]*AttendanceInfo {
	return GC.GetGPP().Attendance.AttendanceMap
}

func GetAttendanceInfo(id int64) *AttendanceInfo {
	return GC.GetGPP().Attendance.AttendanceMap[id]
}
