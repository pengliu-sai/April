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
	fbChapterSheetName  = "章节"
	fbSheetName = "副本"
)

const (
	menu_chapter_ID = "ID"
	menu_chapter_ChapterID = "章节ID"
	menu_chapter_ChapterName = "章节名称"
	menu_chapter_ChapterDesc = "章节描述"
	menu_chapter_ChapterBG = "章节背景"
)

const (
	menu_ID           = "ID"
	menu_FBID         = "副本ID"
	menu_FBName       = "副本名称"
	menu_FBDesc       = "副本描述"
	menu_MonsterGroup = "怪物群组"
	menu_ItemGroup    = "副本掉落群组"
	menu_ChapterID    = "章节ID"
)

func ExportFBChapterConfig() *FBChapterInfoList {
	fbChapterInfoList := FBChapterInfoList{}
	excelFileName := os.Getenv("APRIL_PATH") + "design/策划配置表/" + fbFileName
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	checkFBChapterHead(xlFile.Sheets)
	fbChapterInfoLen := getFBChapterInfoLen(xlFile.Sheets)
	if fbChapterInfoLen <= 0 {
		fmt.Printf("没有配置fb chapter数据")
		return nil
	}
	fbChapterInfoList.FbChapterList = make([]*FBChapterInfo, fbChapterInfoLen)

	for _, sheet := range xlFile.Sheets {
		if sheet.Name == fbChapterSheetName {
			for index, row := range sheet.Rows {
				if index > 0 {
					fbChapterInfo := FBChapterInfo{}
					for index, cell := range row.Cells {
						if index == 0 {
							id, _ := cell.Int64()
							fbChapterInfo.ID = &id
						} else if index == 1 {
							chapterID, _ := cell.Int64()
							fbChapterInfo.ChapterID = &chapterID
						} else if index == 2 {
							chapterName, _ := cell.String()
							fbChapterInfo.ChapterName = &chapterName
						} else if index == 3 {
							chapterDesc, _ := cell.String()
							fbChapterInfo.ChapterDesc = &chapterDesc
						} else if index == 4 {
							chapterBG, _ := cell.String()
							fbChapterInfo.ChapterBG = &chapterBG
						}
					}
					fbChapterInfoList.FbChapterList[index-1] = &fbChapterInfo
				}
			}
		}
	}

	return &fbChapterInfoList
}

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
		if sheet.Name == fbSheetName {
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
						} else if index == 6 {
							chapterID, _ := cell.Int64()
							fbInfo.ChapterID = &chapterID
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
		if sheet.Name == fbSheetName {

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

				headChapterID, _ := row.Cells[6].String()
				if headChapterID != menu_ChapterID {
					log.Fatalln("副本表-章节ID未设置")
					return
				}

			}
		}
	}
}

func checkFBChapterHead(sheets []*xlsx.Sheet) {
	for _, sheet := range sheets {
		if sheet.Name == fbChapterSheetName {

			for index, row := range sheet.Rows {
				if index != 0 {
					return
				}

				headID, _ := row.Cells[0].String()
				if headID != menu_chapter_ID {
					log.Fatalln("章节表-ID未设置")
					return
				}

				headFBChatperID, _ := row.Cells[1].String()
				if headFBChatperID != menu_chapter_ChapterID {
					log.Fatalln("章节表-章节ID未设置")
					return
				}

				headFBChapterName, _ := row.Cells[2].String()
				if headFBChapterName != menu_chapter_ChapterName {
					log.Fatalln("章节表-章节名称未设置")
					return
				}

				headFBChapterDesc, _ := row.Cells[3].String()
				if headFBChapterDesc != menu_chapter_ChapterDesc {
					log.Fatalln("章节表-章节描述未设置")
					return
				}

				headFBChapterBG, _ := row.Cells[4].String()
				if headFBChapterBG != menu_chapter_ChapterBG {
					log.Fatalln("章节表-章节背景未设置")
					return
				}
			}
		}
	}
}

func getFBInfoLen(sheets []*xlsx.Sheet) int {
	for _, sheet := range sheets {
		if sheet.Name == fbSheetName {
			return len(sheet.Rows) - 1
		}
	}
	return 0
}


func getFBChapterInfoLen(sheets []*xlsx.Sheet) int {
	for _, sheet := range sheets {
		if sheet.Name == fbChapterSheetName {
			return len(sheet.Rows) - 1
		}
	}
	return 0
}
