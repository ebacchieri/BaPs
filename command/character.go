package command

import (
	"fmt"
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/gdconf"
	"github.com/gucooing/BaPs/protocol/mx"
	"time"

	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/cdq"
)

const (
	characterPlayer  = -1
	characterUnknown = -2
)

func (c *Command) ApplicationCommandCharacter() {
	character := &cdq.Command{
		Name:        "character",
		AliasList:   []string{"c"},
		Description: "直接设置角色,角色不存在时将添加角色",
		Permissions: cdq.User,
		Options: []*cdq.CommandOption{
			{
				Name:        "uid",
				Description: "玩家游戏id",
				Required:    true,
				Alias:       "u",
			},
			{
				Name:        "characterId",
				Description: "角色id",
				Required:    true,
				Alias:       "id",
			},
			{
				Name:        "level",
				Description: "要设置的角色等级",
				Required:    false,
				Alias:       "l",
			},
			{
				Name:        "starGrade",
				Description: "要设置的角色星级",
				Required:    false,
				Alias:       "sg",
			},
			{
				Name:        "favorRank",
				Description: "要设置的角色好感度等级",
				Required:    false,
				Alias:       "fr",
			},
			{
				Name:        "max",
				Description: "角色属性设置满级",
				Required:    false,
				Alias:       "m",
			},
		},
		Handlers: cdq.AddHandlers(syncGateWay, c.character),
	}

	c.C.ApplicationCommand(character)
}

func (c *Command) character(ctx *cdq.Context) {
	// 玩家验证
	uid := ctx.GetFlags().Int64("uid")
	s := enter.GetSessionByAccountServerId(uid)
	if s == nil {
		ctx.Return(characterPlayer, fmt.Sprintf("玩家不在线或未注册 UID:%v", uid))
		return
	}

	m := ctx.GetFlags().Bool("max")
	level := ctx.GetFlags().Int32("level")
	starGrade := ctx.GetFlags().Int32("starGrade")
	favorRank := ctx.GetFlags().Int32("favorRank")

	set := func(characterInfo *sro.CharacterInfo) {
		if m {
			game.SetMaxCharacter(characterInfo)
			return
		}
		if level != 0 {
			game.SetCharacterLevel(characterInfo, level)
		}
		if starGrade != 0 {
			game.SetCharacterStarGrade(characterInfo, starGrade)
		}
		if favorRank != 0 {
			game.SetCharacterFavorRank(characterInfo, favorRank)
		}
	}

	switch ctx.GetFlags().String("characterId") {
	case "all", "All":
		for _, v := range game.GetCharacterInfoList(s) {
			set(v)
		}
	default:
		characterId := ctx.GetFlags().Int64("characterId")
		if gdconf.GetCharacterExcel(characterId) == nil {
			ctx.Return(characterUnknown, fmt.Sprintf("角色id错误characterId:%v", characterId))
			return
		}
		characterInfo := game.GetCharacterInfo(s, characterId)
		if characterInfo == nil {
			if game.AddCharacter(s, characterId) {
				characterInfo = game.GetCharacterInfo(s, characterId)
			} else {
				ctx.Return(characterUnknown, fmt.Sprintf("角色id错误characterId:%v", characterId))
				return
			}
		}
		set(characterInfo)
	}

	game.AddToast(s, &enter.Toast{
		Text:      "设置角色成功,请重新登录以刷新",
		BeginDate: mx.Now(),
		EndDate:   mx.Now().Add(30 * time.Second),
	})
	ctx.Return(cdq.ApiCodeOk, "设置角色成功")
}
