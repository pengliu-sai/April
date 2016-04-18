package monster

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	. "protos/config"
)

const (
	monsterFileName       = "怪物表.xlsx"
	monsterSheetName      = "怪物"
	monsterGroupSheetName = "怪物群组"
)

const (
	menu_ID                     = "ID"
	menu_MonsterID              = "怪物ID"
	menu_MonsterName            = "怪物名称"
	menu_MonsterDesc            = "怪物描述"
	menu_MonsterHP              = "HP"
	menu_MonsterMP              = "MP"
	menu_MonsterStrength        = "力量"
	menu_MonsterAgility         = "敏捷"
	menu_MonsterPhysicalAttack  = "物理攻击"
	menu_MonsterPhysicalDefense = "物理防御"
	menu_MonsterMagicAttack     = "法术攻击"
	menu_MonsterMagicDefense    = "法术防御"
	menu_MonsterSkill_1         = "技能1"
	menu_MonsterSkill_2         = "技能2"
	menu_MonsterSkill_3         = "技能3"
	menu_MonsterSkill_4         = "技能4"
)

const (
	menu_group_ID          = "ID"
	menu_group_MonsterID   = "怪物群组ID"
	menu_group_MonsterName = "怪物群组名称"
	menu_group_MonsterData = "怪物群组数据"
	menu_group_MonsterDesc = "怪物群组描述"
)

func ExportMonsterConfig() *MonsterInfoList {
	monsterInfoList := MonsterInfoList{}

	excelFileName := os.Getenv("APRIL_PATH") + "design/策划配置表/" + monsterFileName
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	checkMonsterHead(xlFile.Sheets)
	monsterInfoLen := getMonsterInfoLen(xlFile.Sheets)
	if monsterInfoLen <= 0 {
		fmt.Printf("没有配置monster数据")
		return nil
	}
	monsterInfoList.MonsterList = make([]*MonsterInfo, monsterInfoLen)

	for _, sheet := range xlFile.Sheets {
		if sheet.Name == monsterSheetName {
			for index, row := range sheet.Rows {
				if index > 0 {
					monsterInfo := MonsterInfo{}
					for index, cell := range row.Cells {
						if index == 0 {
							id, _ := cell.Int64()
							monsterInfo.ID = &id
						} else if index == 1 {
							monsterID, _ := cell.Int64()
							monsterInfo.MonsterID = &monsterID
						} else if index == 2 {
							monsterName, _ := cell.String()
							monsterInfo.MonsterName = &monsterName
						} else if index == 3 {
							monsterDesc, _ := cell.String()
							monsterInfo.MonsterDesc = &monsterDesc
						} else if index == 4 {
							hp, _ := cell.Int64()
							monsterInfo.HP = &hp
						} else if index == 5 {
							mp, _ := cell.Int64()
							monsterInfo.MP = &mp
						} else if index == 6 {
							strength, _ := cell.Int64()
							monsterInfo.Strength = &strength
						} else if index == 7 {
							agility, _ := cell.Int64()
							monsterInfo.Agility = &agility
						} else if index == 8 {
							physicalAttack, _ := cell.Int64()
							monsterInfo.PhysicalAttack = &physicalAttack
						} else if index == 9 {
							physicalDefense, _ := cell.Int64()
							monsterInfo.PhysicalDefense = &physicalDefense
						} else if index == 10 {
							magicAttack, _ := cell.Int64()
							monsterInfo.MagicAttack = &magicAttack
						} else if index == 11 {
							magicDefense, _ := cell.Int64()
							monsterInfo.MagicDefense = &magicDefense
						}
					}
					monsterInfoList.MonsterList[index-1] = &monsterInfo
				}
			}
		}
	}

	return &monsterInfoList
}

func ExportMonsterGroupConfig() *MonsterGroupInfoList {
	monsterGroupInfoList := MonsterGroupInfoList{}

	excelFileName := os.Getenv("APRIL_PATH") + "design/策划配置表/" + monsterFileName
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	checkMonsterGroupHead(xlFile.Sheets)
	monsterGroupInfoLen := getMonsterGroupInfoLen(xlFile.Sheets)
	if monsterGroupInfoLen <= 0 {
		fmt.Printf("没有配置monster group数据")
		return nil
	}
	monsterGroupInfoList.MonsterGroupList = make([]*MonsterGroupInfo, monsterGroupInfoLen)

	for _, sheet := range xlFile.Sheets {
		if sheet.Name == monsterSheetName {
			for index, row := range sheet.Rows {
				if index > 0 {
					monsterGroupInfo := MonsterGroupInfo{}
					for index, cell := range row.Cells {
						if index == 0 {
							id, _ := cell.Int64()
							monsterGroupInfo.ID = &id
						} else if index == 1 {
							monsterGroupID, _ := cell.Int64()
							monsterGroupInfo.MonsterGroupID = &monsterGroupID
						} else if index == 2 {
							monsterGroupName, _ := cell.String()
							monsterGroupInfo.MonsterGroupName = &monsterGroupName
						} else if index == 3 {
							monsterGroupData, _ := cell.String()
							monsterGroupInfo.MonsterGroupData = &monsterGroupData
						} else if index == 4 {
							monsterGroupDesc, _ := cell.String()
							monsterGroupInfo.MonsterGroupDesc = &monsterGroupDesc
						}
					}
					monsterGroupInfoList.MonsterGroupList[index-1] = &monsterGroupInfo
				}
			}
		}
	}

	return &monsterGroupInfoList
}

func checkMonsterHead(sheets []*xlsx.Sheet) {
	for _, sheet := range sheets {
		if sheet.Name != monsterSheetName {
			return
		}
		for index, row := range sheet.Rows {
			if index != 0 {
				return
			}

			headID, _ := row.Cells[0].String()
			if headID != menu_ID {
				log.Fatalln("怪物表-ID未设置")
				return
			}

			headMonsterID, _ := row.Cells[1].String()
			if headMonsterID != menu_MonsterID {
				log.Fatalln("怪物表-怪物ID未设置")
				return
			}

			headMonsterName, _ := row.Cells[2].String()
			if headMonsterName != menu_MonsterName {
				log.Fatalln("怪物表-怪物名称未设置")
				return
			}

			headMonsterDesc, _ := row.Cells[3].String()
			if headMonsterDesc != menu_MonsterDesc {
				log.Fatalln("怪物表-怪物描述未设置")
				return
			}

			headMonsterHP, _ := row.Cells[4].String()
			if headMonsterHP != menu_MonsterHP {
				log.Fatalln("怪物表-HP未设置")
				return
			}

			headMonsterMP, _ := row.Cells[5].String()
			if headMonsterMP != menu_MonsterMP {
				log.Fatalln("怪物表-MP未设置")
				return
			}

			headMonsterStrength, _ := row.Cells[6].String()
			if headMonsterStrength != menu_MonsterStrength {
				log.Fatalln("怪物表-力量未设置")
				return
			}

			headMonsterAgility, _ := row.Cells[7].String()
			if headMonsterAgility != menu_MonsterAgility {
				log.Fatalln("怪物表-敏捷未设置")
				return
			}

			headMonsterPhysicalAttack, _ := row.Cells[8].String()
			if headMonsterPhysicalAttack != menu_MonsterPhysicalAttack {
				log.Fatalln("怪物表-物理攻击未设置")
				return
			}

			headMonsterPhysicalDefense, _ := row.Cells[9].String()
			if headMonsterPhysicalDefense != menu_MonsterPhysicalDefense {
				log.Fatalln("怪物表-物理防御未设置")
				return
			}

			headMonsterMagicAttack, _ := row.Cells[10].String()
			if headMonsterMagicAttack != menu_MonsterMagicAttack {
				log.Fatalln("怪物表-法术攻击未设置")
				return
			}

			headMonsterMagicDefense, _ := row.Cells[11].String()
			if headMonsterMagicDefense != menu_MonsterMagicDefense {
				log.Fatalln("怪物表-法术防御未设置")
				return
			}

			//允许不设置技能

		}
	}
}

func checkMonsterGroupHead(sheets []*xlsx.Sheet) {
	for _, sheet := range sheets {
		if sheet.Name == monsterGroupSheetName {

			for index, row := range sheet.Rows {
				if index != 0 {
					return
				}

				headID, _ := row.Cells[0].String()
				if headID != menu_ID {
					log.Fatalln("怪物群组表-ID未设置")
					return
				}

				headMonsterGroupID, _ := row.Cells[1].String()
				if headMonsterGroupID != menu_group_MonsterID {
					log.Fatalln("怪物群组表-怪物群组ID未设置")
					return
				}

				headMonsterGroupName, _ := row.Cells[2].String()
				if headMonsterGroupName != menu_group_MonsterName {
					log.Fatalln("怪物群组表-怪物群组名称未设置")
					return
				}

				headMonsterGroupData, _ := row.Cells[3].String()
				if headMonsterGroupData != menu_group_MonsterData {
					log.Fatalln("怪物群组表-怪物群组数据未设置")
					return
				}

				headMonsterGroupDesc, _ := row.Cells[4].String()
				if headMonsterGroupDesc != menu_group_MonsterDesc {
					log.Fatalln("怪物群组表-怪物群组描述未设置")
					return
				}
			}
		}
	}
}

func getMonsterInfoLen(sheets []*xlsx.Sheet) int {
	for _, sheet := range sheets {
		if sheet.Name == monsterSheetName {
			return len(sheet.Rows) - 1
		}
	}
	return 0
}

func getMonsterGroupInfoLen(sheets []*xlsx.Sheet) int {
	for _, sheet := range sheets {
		if sheet.Name == monsterGroupSheetName {
			return len(sheet.Rows) - 1
		}
	}
	return 0
}
