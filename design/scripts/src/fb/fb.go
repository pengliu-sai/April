package fb

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	. "protos/config"
)

const (
	fbFileName = "副本表.xlsx"
	sheetName  = "副本"
)

const (
	menu_ID           = "ID"
	menu_FBID         = "副本ID"
	menu_FBName       = "副本名称"
	menu_FBDesc       = "副本描述"
	menu_MonsterGroup = "怪物群组"
	menu_ItemGroup    = "副本掉落群组"
)

func ExportFBConfig() *FBInfoList {
	fbInfoList := FBInfoList{}

	excelFileName := os.Getenv("APRIL_PATH") + "design/策划配置表/" + fbFileName
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	checkFBHead(xlFile.Sheets)
	fbInfoLen := getFBInfoLen(xlFile.Sheets)
	if fbInfoLen <= 0 {
		fmt.Printf("没有配置fb数据")
		return nil
	}
	fbInfoList.FbList = make([]*FBInfo, fbInfoLen)

	for _, sheet := range xlFile.Sheets {
		if sheet.Name == sheetName {
			for index, row := range sheet.Rows {
				if index > 0 {
					fbInfo := FBInfo{}
					for index, cell := range row.Cells {
						if index == 0 {
							id, _ := cell.Int64()
							fbInfo.ID = &id
						} else if index == 1 {
							fbID, _ := cell.Int64()
							fbInfo.FBID = &fbID
						} else if index == 2 {
							fbName, _ := cell.String()
							fbInfo.FBName = &fbName
						} else if index == 3 {
							fbDesc, _ := cell.String()
							fbInfo.FBDesc = &fbDesc
						} else if index == 4 {
							monsterGroupID, _ := cell.String()
							fbInfo.MonsterGroupID = &monsterGroupID
						} else if index == 5 {
							itemGroupID, _ := cell.String()
							fbInfo.ItemGroupID = &itemGroupID
						}
					}
					fbInfoList.FbList[index-1] = &fbInfo
				}
			}
		}
	}

	return &fbInfoList
}

func checkFBHead(sheets []*xlsx.Sheet) {
	for _, sheet := range sheets {
		if sheet.Name == sheetName {

			for index, row := range sheet.Rows {
				if index != 0 {
					return
				}

				headID, _ := row.Cells[0].String()
				if headID != menu_ID {
					log.Fatalln("副本表-ID未设置")
					return
				}

				headFBID, _ := row.Cells[1].String()
				if headFBID != menu_FBID {
					log.Fatalln("副本表-副本ID未设置")
					return
				}

				headFBName, _ := row.Cells[2].String()
				if headFBName != menu_FBName {
					log.Fatalln("副本表-副本名称未设置")
					return
				}

				headFBDesc, _ := row.Cells[3].String()
				if headFBDesc != menu_FBDesc {
					log.Fatalln("副本表-副本描述未设置")
					return
				}

				headMonsterGroup, _ := row.Cells[4].String()
				if headMonsterGroup != menu_MonsterGroup {
					log.Fatalln("副本表-怪物群组未设置")
					return
				}

				headItemGroup, _ := row.Cells[5].String()
				if headItemGroup != menu_ItemGroup {
					log.Fatalln("副本表-物品群组未设置")
					return
				}

			}
		}
	}
}

func getFBInfoLen(sheets []*xlsx.Sheet) int {
	for _, sheet := range sheets {
		if sheet.Name == sheetName {
			return len(sheet.Rows) - 1
		}
	}
	return 0
}
