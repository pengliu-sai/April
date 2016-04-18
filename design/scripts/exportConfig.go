package main

import (
	"fb"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"item"
	"log"
	"monster"
	"os"
	. "protos/config"
	"skill"
)

var (
	configFileName string
)

//读取excel数据, 导出成proto数据
func main() {
	configFileName = os.Getenv("APRIL_PATH") + "design/data/config.data"

	globalConfig := GlobalConfigInfo{}
	//导出chapter配置
	globalConfig.FbChapterInfoList = fb.ExportFBChapterConfig()
	//导出fb配置
	globalConfig.FbInfoList = fb.ExportFBConfig()
	//导出物品配置
	globalConfig.ItemInfoList = item.ExportItemConfig()
	//导出物品群组配置
	globalConfig.ItemGroupInfoList = item.ExportItemGroupConfig()
	//导出技能表
	globalConfig.SkillInfoList = skill.ExportSkillConfig()
	//导出怪物表
	globalConfig.MonsterInfoList = monster.ExportMonsterConfig()
	//导出怪物群组表
	globalConfig.MonsterGroupInfoList = monster.ExportMonsterGroupConfig()
	//保存到文件
	saveToFile(&globalConfig)

	//加载到内存
	//loadToMemory()
}

//序列化到文件
func saveToFile(config *GlobalConfigInfo) {
	data, err := proto.Marshal(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.OpenFile(configFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		log.Fatalln("写入文件失败")
	}

	log.Println("saveToFile success!")
}

//反序列化
func loadToMemory() {
	fileByte, err := ioutil.ReadFile(configFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	globalConfig := GlobalConfigInfo{}
	err = proto.Unmarshal(fileByte, &globalConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("globalConfig: ", *globalConfig.FbInfoList.FbList[0].FBName)

	log.Println("loadToMemory success!")
}
