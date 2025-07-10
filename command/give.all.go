package command

import (
	"fmt"
	"time"

	"github.com/gucooing/BaPs/common/enter"
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/BaPs/gdconf"
	"github.com/gucooing/BaPs/pkg/alg"
	"github.com/gucooing/BaPs/protocol/proto"
	"github.com/gucooing/cdq"
)

type giveAllFunc func(num int64) []*sro.ParcelInfo

var giveAllFuncs = []giveAllFunc{
	GiveAllMaterial,
	GiveAllCharacter,
	GiveAllEquipment,
	GiveAllFurniture,
	GiveAllFavor,
	GiveAllEmblem,
	GiveAllSticker,
	GiveAllMemoryLobby,
}

var (
	gaAll         = "All"
	gaMaterial    = "Material"
	gaCharacter   = "Character"
	gaEquipment   = "Equipment"
	gaFurniture   = "Furniture"
	gaFavor       = "Favor"
	gaEmblem      = "Emblem"
	gaSticker     = "Sticker"
	gaMemoryLobby = "MemoryLobby"
)

const (
	gaPlayerUnknown = -1
	gaTypeUnknown   = -2
	gaSendMailErr   = -3
)

func (c *Command) ApplicationCommandGiveAll() {
	giveAll := &cdq.Command{
		Name:        "giveAll",
		AliasList:   []string{"ga"},
		Description: "获取某个类型的全部物品",
		Permissions: cdq.User,
		Options: []*cdq.CommandOption{
			{
				Name:        "uid",
				Description: "玩家游戏id",
				Required:    true,
				Alias:       "u",
			},
			{
				Name:        "type",
				Description: "需要给予物品的类型",
				Required:    true,
				Alias:       "t",
				ExpectedS: []string{
					gaAll,
					gaMaterial, "m",
					gaCharacter, "c",
					gaEquipment, "e",
					gaFurniture, "f",
					gaFavor, "fa",
					gaEmblem, "em",
					gaSticker, "s",
					gaMemoryLobby, "ml",
				},
			},
			{
				Name:        "num",
				Description: "需要给予物品的数量",
				Required:    false,
				Default:     "1",
				Alias:       "n",
			},
		},
		Handlers: cdq.AddHandlers(syncGateWay, c.giveALL),
	}

	c.C.ApplicationCommand(giveAll)
}

func (c *Command) giveALL(ctx *cdq.Context) {
	// 玩家验证
	uid := ctx.GetFlags().Int64("uid")
	s := enter.GetSessionByAccountServerId(uid)
	if s == nil {
		ctx.Return(gaPlayerUnknown, fmt.Sprintf("玩家不在线或未注册 UID:%v", uid))
		return
	}

	// 执行
	var parcelInfoList []*sro.ParcelInfo
	num := ctx.GetFlags().Int64("num")
	switch ctx.GetFlags().String("type") {
	case gaAll:
		parcelInfoList = GiveAllTypes(num)
	case gaMaterial, "m":
		parcelInfoList = GiveAllMaterial(num)
	case gaCharacter, "c":
		parcelInfoList = GiveAllCharacter(num)
	case gaEquipment, "e":
		parcelInfoList = GiveAllEquipment(num)
	case gaFurniture, "f":
		parcelInfoList = GiveAllFurniture(num)
	case gaFavor, "fa":
		parcelInfoList = GiveAllFavor(num)
	case gaEmblem, "em":
		parcelInfoList = GiveAllEmblem(num)
	case gaSticker, "s":
		parcelInfoList = GiveAllSticker(num)
	case gaMemoryLobby, "ml":
		parcelInfoList = GiveAllMemoryLobby(num)
	default:
		ctx.Return(gaTypeUnknown, "不存在此物品类型")
		return
	}

	mail := &sro.MailInfo{
		Sender:         "gucooing",
		Comment:        "请查收您的意外奖励",
		SendDate:       time.Now().Unix(),
		ExpireDate:     time.Now().Add(10 * time.Minute).Unix(),
		ParcelInfoList: parcelInfoList,
	}
	if game.AddMail(s, mail) {
		ctx.Return(cdq.ApiCodeOk, "请查询游戏内邮箱获取结果")
	} else {
		ctx.Return(gaSendMailErr, "游戏邮箱错误")
	}
}

func GiveAllTypes(num int64) []*sro.ParcelInfo {
	parcelInfoList := make([]*sro.ParcelInfo, 0)
	for _, v := range giveAllFuncs {
		parcelInfoList = append(parcelInfoList, v(num)...)
	}

	return parcelInfoList
}

func GiveAllMaterial(num int64) []*sro.ParcelInfo {
	list := make([]*sro.ParcelInfo, 0)
	for _, conf := range gdconf.GetItemExcelCategoryMapByCategory("Material") {
		list = append(list, &sro.ParcelInfo{
			Type: int32(proto.ParcelType_Item),
			Id:   conf.Id,
			Num:  alg.MinInt64(num, 9999),
		})
	}

	return list
}

func GiveAllCharacter(num int64) []*sro.ParcelInfo {
	list := make([]*sro.ParcelInfo, 0)
	for _, conf := range gdconf.GetCharacterMap() {
		list = append(list, &sro.ParcelInfo{
			Type: int32(proto.ParcelType_Character),
			Id:   conf.Id,
			Num:  1,
		})
	}
	return list
}

func GiveAllEquipment(num int64) []*sro.ParcelInfo {
	list := make([]*sro.ParcelInfo, 0)
	nums := num
	for _, conf := range gdconf.GetEquipmentExcelMap() {
		if conf.MaxLevel < 10 {
			num = alg.MinInt64(nums, 9999)
		} else {
			num = alg.MinInt64(nums, 20)
		}
		list = append(list, &sro.ParcelInfo{
			Type: int32(proto.ParcelType_Equipment),
			Id:   conf.Id,
			Num:  num,
		})
	}

	return list
}

func GiveAllFurniture(num int64) []*sro.ParcelInfo {
	list := make([]*sro.ParcelInfo, 0)
	num = alg.MinInt64(num, 3)
	for _, conf := range gdconf.GetFurnitureExcelTableMap() {
		list = append(list, &sro.ParcelInfo{
			Type: int32(proto.ParcelType_Furniture),
			Id:   conf.Id,
			Num:  num,
		})
	}

	return list
}

func GiveAllFavor(num int64) []*sro.ParcelInfo {
	list := make([]*sro.ParcelInfo, 0)
	num = alg.MinInt64(num, 9999)
	for _, conf := range gdconf.GetItemExcelCategoryMapByCategory("Favor") {
		list = append(list, &sro.ParcelInfo{
			Type: int32(proto.ParcelType_Item),
			Id:   conf.Id,
			Num:  num,
		})
	}

	return list
}

func GiveAllEmblem(num int64) []*sro.ParcelInfo {
	list := make([]*sro.ParcelInfo, 0)
	for _, conf := range gdconf.GetEmblemExcelList() {
		list = append(list, &sro.ParcelInfo{
			Type: int32(proto.ParcelType_Emblem),
			Id:   conf.Id,
		})
	}
	return list
}

func GiveAllSticker(num int64) []*sro.ParcelInfo {
	list := make([]*sro.ParcelInfo, 0)
	for _, conf := range gdconf.GetStickerPageContentExcelList() {
		list = append(list, &sro.ParcelInfo{
			Type: int32(proto.ParcelType_Sticker),
			Id:   conf.Id,
			Num:  1,
		})
	}
	return list
}

func GiveAllMemoryLobby(num int64) []*sro.ParcelInfo {
	list := make([]*sro.ParcelInfo, 0)
	for _, conf := range gdconf.GetMemoryLobbyExcelList() {
		list = append(list, &sro.ParcelInfo{
			Type: int32(proto.ParcelType_MemoryLobby),
			Id:   conf.Id,
			Num:  1,
		})
	}
	return list
}
