package logic

import (
	"github.com/funny/link"
	"protos"
	. "tools"
	"protos/game"
	"protos/config"
)

//请求游戏副本章节列表
func gameFBChapterList(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive fbChapterListC2S message")
	roleID := g_gameSession_RoleID[gameSession]

	if roleID <= 0 {
		//未登陆角色
		ERR("role canot login")
		return
	}

	send_msg := protos.MarshalProtoMsg(&game.Game_FBChapterListS2C{
		FbChapterList: globalConfig.FbChapterInfoList.FbChapterList,
	})

	INFO("Game send fbChapterListS2C message")
	gameSession.Send(send_msg)
}

//请求章节的子列表
func gameFBSectionListByChapterID(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive fbSectionListC2S message")
	roleID := g_gameSession_RoleID[gameSession]

	if roleID <= 0 {
		//未登陆角色
		ERR("role canot login")
		return
	}

	rev_msg := msg.Body.(*game.Game_FBSectionListC2S)
	chapterID := *rev_msg.ChapterID

	_fbList := globalConfig.FbInfoList.FbList
	num := 0
	for _, v := range _fbList {
		if *v.ChapterID == chapterID {
			num++
		}
	}

	fbInfoList := make([]*config.FBInfo, num)
	index := 0
	for _, v := range _fbList {
		if *v.ChapterID == chapterID {
			fbInfoList[index] = v
			index++

		}
	}



	send_msg := protos.MarshalProtoMsg(&game.Game_FBSectionListS2C{
		FbInfoList: fbInfoList,
	})

	INFO("Game send fbSectionListS2C message")
	gameSession.Send(send_msg)
}

//通过某个小节请求副本详细信息
func gameFBInfoBySectionID(gameSession *link.Session, msg protos.ProtoMsg) {

}

//挑战某个小节副本
func gameFBBattleBySectionID(gameSession *link.Session, msg protos.ProtoMsg) {

}
