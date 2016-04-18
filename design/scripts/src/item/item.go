package item

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	. "protos/config"
)

const (
	itemFileName       = "物品表.xlsx"
	itemSheetName      = "物品"
	itemGroupSheetName = "物品群组"
)

const (
	menu_ID       = "ID"
	menu_ItemID   = "物品ID"
	menu_ItemName = "物品名称"
	menu_ItemDesc = "物品描述"
)

const (
	menu_group_ID       = "ID"
	menu_group_ItemID   = "物品群组ID"
	menu_group_ItemName = "物品群组名称"
	menu_group_ItemData = "物品数据"
	menu_group_ItemDesc = "物品群组描述"
)

func ExportItemConfig() *ItemInfoList {
	itemInfoList := ItemInfoList{}

	excelFileName := os.Getenv("APRIL_PATH") + "design/策划配置表/" + itemFileName
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	checkItemHead(xlFile.Sheets)

	itemInfoLen := getItemInfoLen(xlFile.Sheets)
	if itemInfoLen <= 0 {
		fmt.Println("没有配置item数据")
		return nil
	}
	itemInfoList.ItemList = make([]*ItemInfo, itemInfoLen)

	for _, sheet := range xlFile.Sheets {
		if sheet.Name == itemSheetName {
			for index, row := range sheet.Rows {
				if index > 0 {
					itemInfo := ItemInfo{}
					for index, cell := range row.Cells {
						if index == 0 {
							id, _ := cell.Int64()
							itemInfo.ID = &id
						} else if index == 1 {
							itemID, _ := cell.Int64()
							itemInfo.ItemID = &itemID
						} else if index == 2 {
							itemName, _ := cell.String()
							itemInfo.ItemName = &itemName
						} else if index == 3 {
							itemDesc, _ := cell.String()
							itemInfo.ItemDesc = &itemDesc
						}
					}
					itemInfoList.ItemList[index-1] = &itemInfo
				}
			}
		}
	}

	return &itemInfoList
}

func ExportItemGroupConfig() *ItemGroupInfoList {
	itemGroupInfoList := ItemGroupInfoList{}

	excelFileName := os.Getenv("APRIL_PATH") + "design/策划配置表/" + itemFileName
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	checkItemGroupHead(xlFile.Sheets)

	itemGroupInfoLen := getItemGroupInfoLen(xlFile.Sheets)
	if itemGroupInfoLen <= 0 {
		fmt.Println("没有配置item group数据")
		return nil
	}
	itemGroupInfoList.ItemGroupList = make([]*ItemGroupInfo, itemGroupInfoLen)

	for _, sheet := range xlFile.Sheets {
		if sheet.Name == itemGroupSheetName {
			for index, row := range sheet.Rows {
				if index > 0 {
					itemGroupInfo := ItemGroupInfo{}
					for index, cell := range row.Cells {
						if index == 0 {
							id, _ := cell.Int64()
							itemGroupInfo.ID = &id
						} else if index == 1 {
							itemGroupID, _ := cell.Int64()
							itemGroupInfo.ItemGroupID = &itemGroupID
						} else if index == 2 {
							itemGroupName, _ := cell.String()
							itemGroupInfo.ItemGroupName = &itemGroupName
						} else if index == 3 {
							itemGroupData, _ := cell.String()
							itemGroupInfo.ItemGroupData = &itemGroupData
						} else if index == 4 {
							itemGroupDesc, _ := cell.String()
							itemGroupInfo.ItemGroupDesc = &itemGroupDesc
						}
					}
					itemGroupInfoList.ItemGroupList[index-1] = &itemGroupInfo
				}
			}
		}
	}

	return &itemGroupInfoList
}

func checkItemHead(sheets []*xlsx.Sheet) {
	for _, sheet := range sheets {
		if sheet.Name == itemSheetName {

			for index, row := range sheet.Rows {
				if index != 0 {
					return
				}

				headID, _ := row.Cells[0].String()
				if headID != menu_ID {
					log.Fatalln("物品表-ID未设置")
					return
				}

				headItemID, _ := row.Cells[1].String()
				if headItemID != menu_ItemID {
					log.Fatalln("物品表-副本ID未设置")
					return
				}

				headItemName, _ := row.Cells[2].String()
				if headItemName != menu_ItemName {
					log.Fatalln("物品表-物品名称未设置")
					return
				}

				headItemDesc, _ := row.Cells[3].String()
				if headItemDesc != menu_ItemDesc {
					log.Fatalln("物品表-物品描述未设置")
					return
				}
			}
		}
	}
}

func checkItemGroupHead(sheets []*xlsx.Sheet) {
	for _, sheet := range sheets {
		if sheet.Name == itemGroupSheetName {

			for index, row := range sheet.Rows {
				if index != 0 {
					return
				}

				headGroupID, _ := row.Cells[0].String()
				if headGroupID != menu_group_ID {
					log.Fatalln("物品群组表-ID未设置")
					return
				}

				headGroupItemID, _ := row.Cells[1].String()
				if headGroupItemID != menu_group_ItemID {
					log.Fatalln("物品群组表-物品群组ID未设置")
					return
				}

				headGroupItemName, _ := row.Cells[2].String()
				if headGroupItemName != menu_group_ItemName {
					log.Fatalln("物品群组表-物品群组名称未设置")
					return
				}

				headGroupItemData, _ := row.Cells[3].String()
				if headGroupItemData != menu_group_ItemData {
					log.Fatalln("物品群组表-物品群组数据未设置")
					return
				}

				headGroupItemDesc, _ := row.Cells[4].String()
				if headGroupItemDesc != menu_group_ItemDesc {
					log.Fatalln("物品群组表-怪物群组描述未设置")
					return
				}
			}
		}
	}
}

func getItemInfoLen(sheets []*xlsx.Sheet) int {
	for _, sheet := range sheets {
		if sheet.Name == itemSheetName {
			return len(sheet.Rows) - 1
		}
	}
	return 0
}

func getItemGroupInfoLen(sheets []*xlsx.Sheet) int {
	for _, sheet := range sheets {
		if sheet.Name == itemGroupSheetName {
			return len(sheet.Rows) - 1
		}
	}
	return 0
}
