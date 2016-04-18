package skill

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	. "protos/config"
)

const (
	skillFileName = "技能表.xlsx"
	sheetName     = "技能"
)

const (
	menu_ID          = "ID"
	menu_SkillID     = "技能ID"
	menu_SkillName   = "技能名称"
	menu_SkillDesc   = "技能描述"
	menu_SkillEffect = "技能特效"
)

func ExportSkillConfig() *SkillInfoList {
	skillInfoList := SkillInfoList{}

	excelFileName := os.Getenv("APRIL_PATH") + "design/策划配置表/" + skillFileName
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	checkSkillHead(xlFile.Sheets)
	skillInfoLen := getSkillInfoLen(xlFile.Sheets)
	if skillInfoLen <= 0 {
		fmt.Printf("没有配置skill数据")
		return nil
	}
	skillInfoList.SkillList = make([]*SkillInfo, skillInfoLen)

	for _, sheet := range xlFile.Sheets {
		if sheet.Name == sheetName {
			for index, row := range sheet.Rows {
				if index > 0 {
					skillInfo := SkillInfo{}
					for index, cell := range row.Cells {
						if index == 0 {
							id, _ := cell.Int64()
							skillInfo.ID = &id
						} else if index == 1 {
							skillID, _ := cell.Int64()
							skillInfo.SkillID = &skillID
						} else if index == 2 {
							skillName, _ := cell.String()
							skillInfo.SkillName = &skillName
						} else if index == 3 {
							skillDesc, _ := cell.String()
							skillInfo.SkillDesc = &skillDesc
						} else if index == 4 {
							skillEffect, _ := cell.String()
							skillInfo.SkillEffect = &skillEffect
						}
					}
					skillInfoList.SkillList[index-1] = &skillInfo
				}
			}
		}
	}

	return &skillInfoList
}

func checkSkillHead(sheets []*xlsx.Sheet) {
	for _, sheet := range sheets {
		if sheet.Name != sheetName {
			return
		}
		for index, row := range sheet.Rows {
			if index != 0 {
				return
			}

			headID, _ := row.Cells[0].String()
			if headID != menu_ID {
				log.Fatalln("技能表-ID未设置")
				return
			}

			headSkillID, _ := row.Cells[1].String()
			if headSkillID != menu_SkillID {
				log.Fatalln("技能表-技能ID未设置")
				return
			}

			headSkillName, _ := row.Cells[2].String()
			if headSkillName != menu_SkillName {
				log.Fatalln("技能表-技能名称未设置")
				return
			}

			headSkillDesc, _ := row.Cells[3].String()
			if headSkillDesc != menu_SkillDesc {
				log.Fatalln("技能表-技能描述未设置")
				return
			}

			headSkillEffect, _ := row.Cells[4].String()
			if headSkillEffect != menu_SkillEffect {
				log.Fatalln("技能表-技能特效未设置")
				return
			}
		}
	}
}

func getSkillInfoLen(sheets []*xlsx.Sheet) int {
	for _, sheet := range sheets {
		if sheet.Name == sheetName {
			return len(sheet.Rows) - 1
		}
	}
	return 0
}
