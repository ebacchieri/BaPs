package gdconf

import (
	"encoding/json"
	"fmt"
	"github.com/gucooing/BaPs/config"
	"os"
	"runtime"
	"strings"
	"time"

	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

var GC *GameConfig

type GameConfig struct {
	dataPath string
	resPath  string
	managementDataUrl string
	gppFunc  []func()
	Excel    *sro.Excel
	GPP      *GPP
}

type GPP struct {
	CharacterExcel                      *CharacterExcel
	CafeInfoExcel                       *CafeInfoExcel
	ShopExcel                           *ShopExcel
	ShopInfoExcel                       *ShopInfoExcel
	ItemExcel                           *ItemExcel
	Emblem                              *Emblem
	AcademyFavorScheduleExcel           *AcademyFavorScheduleExcel
	AcademyMessangerExcel               *AcademyMessangerExcel
	GuideMissionExcel                   *GuideMissionExcel
	ScenarioMode                        *ScenarioMode
	ScenarioModeReward                  *ScenarioModeReward
	CharacterWeaponExcel                *CharacterWeaponExcel
	CharacterSkillListExcel             *CharacterSkillListExcel
	SkillExcel                          *SkillExcel
	RecipeIngredientExcel               *RecipeIngredientExcel
	CampaignStageExcel                  *CampaignStageExcel
	CampaignUnitExcel                   *CampaignUnitExcel
	WeekDungeonExcel                    *WeekDungeonExcel
	SchoolDungeonStage                  *SchoolDungeonStage
	AccountLevel                        *AccountLevel
	EquipmentExcel                      *EquipmentExcel
	EquipmentStatExcel                  *EquipmentStatExcel
	EquipmentLevelExcel                 *EquipmentLevelExcel
	FurnitureExcel                      *FurnitureExcel
	CafeRankExcel                       *CafeRankExcel
	CafeProductionExcel                 *CafeProductionExcel
	IdCardBackground                    *IdCardBackground
	WeekDungeonRewardExcel              *WeekDungeonRewardExcel
	CharacterLevelExcel                 *CharacterLevelExcel
	CharacterGear                       *CharacterGear
	CharacterPotential                  *CharacterPotential
	CharacterPotentialStat              *CharacterPotentialStat
	AcademyLocationExcel                *AcademyLocationExcel
	AcademyTicketExcel                  *AcademyTicketExcel
	AcademyZoneExcel                    *AcademyZoneExcel
	AcademyLocationRankExcel            *AcademyLocationRankExcel
	AcademyRewardExcel                  *AcademyRewardExcel
	SchoolDungeonReward                 *SchoolDungeonReward
	RaidSchedule                        *RaidSchedule
	RaidSeasonManageExcel               *RaidSeasonManageExcel
	RaidRankingRewardExcel              *RaidRankingRewardExcel
	RaidStageExcel                      *RaidStageExcel
	CharacterStatExcel                  *CharacterStatExcel
	RaidStageRewardExcel                *RaidStageRewardExcel
	MissionExcel                        *MissionExcel
	RaidStageSeasonRewardExcel          *RaidStageSeasonRewardExcel
	Attendance                          *Attendance
	MultiFloorRaidReward                *MultiFloorRaidReward
	MultiFloorRaidStage                 *MultiFloorRaidStage
	MultiFloorRaidSeasonManage          *MultiFloorRaidSeasonManage
	TimeAttackDungeonSeasonManageExcel  *TimeAttackDungeonSeasonManageExcel
	RaidEliminateSchedule               *RaidEliminateSchedule
	EliminateRaidSeasonManageExcel      *EliminateRaidSeasonManageExcel
	EliminateRaidRankingRewardExcel     *EliminateRaidRankingRewardExcel
	EliminateRaidStageExcel             *EliminateRaidStageExcel
	EliminateRaidStageRewardExcel       *EliminateRaidStageRewardExcel
	EliminateRaidStageSeasonRewardExcel *EliminateRaidStageSeasonRewardExcel
	GachaElementExcel                   *GachaElementExcel
	GachaElementRecursiveExcel          *GachaElementRecursiveExcel
	GoodsExcel                          *GoodsExcel
	TimeAttackDungeonGeasExcel          *TimeAttackDungeonGeasExcel
	TimeAttackDungeonRewardExcel        *TimeAttackDungeonRewardExcel
	StrategyMap                         map[string]*StrategyMap
	ShopRefreshExcel                    *ShopRefreshExcel
	ArenaSeasonExcel                    *ArenaSeasonExcel
	ArenaNPCInfo                        *ArenaNPCInfo
	FavorLevel                          *FavorLevel
	MailInfoMap                         map[string]*MailInfo
	ProdIndex                           *ProdIndex
	ServerInfo                          *ServerInfo
	EventContentMissionExcel            *EventContentMissionExcel
	EventContentStageExcel              *EventContentStageExcel
	EventContentStageRewardExcel        *EventContentStageRewardExcel
	StickerPageContentExcel             *StickerPageContentExcel
	MemoryLobbyExcel                    *MemoryLobbyExcel
}

func LoadGameConfig() *GameConfig {
	gc := new(GameConfig)
	GC = gc
	gc.managementDataUrl = config.GetOtherAddr().GetManagementDataUrl()
	gc.dataPath = config.GetDataPath()
	gc.resPath = config.GetResourcesPath()
	logger.Info("开始读取资源文件")
	startTime := time.Now()
	gc.LoadExcel()
	gc.gpp()
	endTime := time.Now()
	runtime.GC()
	logger.Info("读取资源完成,用时:%s", endTime.Sub(startTime))
	return gc
}

func (g *GameConfig) gpp() {
	// 验证文件夹是否存在
	if dirInfo, err := os.Stat(g.dataPath); err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("找不到文件夹:%s,err:%s", g.dataPath, err)
		panic(info)
	}
	g.dataPath += "/"
	g.GPP = &GPP{}

	g.gppFunc = []func(){
		g.gppCafeInfoExcelTable,
		g.gppCharacterExcelTable,
		g.gppShopExcel,
		g.gppShopInfoExcel,
		g.gppItemExcel,
		g.gppEmblemExcel,
		g.gppAcademyFavorScheduleExcelTable,
		g.gppAcademyMessangerExcelTable,
		g.gppGuideMissionExcel,
		g.gppScenarioModeExcel,
		g.gppScenarioModeRewardExcel,
		g.gppCharacterWeaponExcelTable,
		g.gppCharacterSkillListExcelTable,
		g.gppSkillExcelTable,
		g.gppRecipeIngredientExcelTable,
		g.gppCampaignStageExcelTable,
		g.gppCampaignUnitExcelTable,
		g.gppWeekDungeonExcelTable,
		g.gppSchoolDungeonStageExcel,
		g.gppAccountLevelExcel,
		g.gppEquipmentExcel,
		g.gppEquipmentStatExcel,
		g.gppEquipmentLevelExcel,
		g.gppFurnitureExcelTable,
		g.gppCafeRankExcelTable,
		g.gppCafeProductionExcelTable,
		g.gppIdCardBackgroundExcel,
		g.gppWeekDungeonRewardExcelTable,
		g.gppCharacterLevelExcelTable,
		g.gppCharacterGearExcel,
		g.gppCharacterPotentialExcel,
		g.gppCharacterPotentialStatExcel,
		g.gppAcademyLocationExcelTable,
		g.gppAcademyTicketExcelTable,
		g.gppAcademyZoneExcelTable,
		g.gppAcademyLocationRankExcelTable,
		g.gppAcademyRewardExcelTable,
		g.gppSchoolDungeonRewardExcel,
		g.gppRaidSeasonManageExcelTable,
		g.gppRaidRankingRewardExcelTable,
		g.gppRaidStageExcelTable,
		g.gppCharacterStatExcelTable,
		g.gppRaidStageRewardExcelTable,
		g.gppMissionExcelTable,
		g.gppRaidStageSeasonRewardExcelTable,
		g.gppMultiFloorRaidRewardExcel,
		g.gppMultiFloorRaidStageExcel,
		g.gppMultiFloorRaidSeasonManageExcel,
		g.gppTimeAttackDungeonSeasonManageExcelTable,
		g.gppEliminateRaidSeasonManageExcelTable,
		g.gppEliminateRaidRankingRewardExcelTable,
		g.gppEliminateRaidStageExcelTable,
		g.gppEliminateRaidStageRewardExcelTable,
		g.gppEliminateRaidStageSeasonRewardExcelTable,
		g.gppGachaElementExcel,
		g.gppGachaElementRecursiveExcel,
		g.gppGoodsExcel,
		g.gppTimeAttackDungeonGeasExcel,
		g.gppTimeAttackDungeonRewardExcel,
		g.gppShopRefreshExcel,
		g.gppArenaSeasonExcelTable,
		g.gppFavorLevelExcel,
		g.gppEventContentMissionExcelTable,
		g.gppEventContentStageExcelTable,
		g.gppEventContentStageRewardExcelTable,
		g.gppStickerPageContentExcel,
		g.gppMemoryLobbyExcel,

		// data
		g.loadRaidSchedule,
		g.loadAttendance,
		g.loadRaidEliminateSchedule,
		g.loadStrategyMap,
		g.loadArenaNPC,
		g.loadMailInfo,
		g.loadProdIndex,
		g.loadServerInfo,
	}

	for _, fn := range g.gppFunc {
		fn()
	}
}

func (g *GameConfig) GetExcel() *sro.Excel {
	if g == nil {
		return nil
	}
	return g.Excel
}

func (g *GameConfig) GetGPP() *GPP {
	if g == nil {
		return nil
	}
	return g.GPP
}

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format("2006-01-02 15:04:05"))
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	str := strings.Trim(string(data), "\"")
	parsedTime, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		parsedTime, err = time.Parse(time.RFC3339, str)
		if err != nil {
			return err
		}
		return err
	}
	*t = Time(parsedTime)
	return nil
}

func (t *Time) Time() time.Time {
	return time.Time(*t)
}
