package main

import "encoding/json"
import "fmt"
import "github.com/tidwall/gjson"
import E "github.com/xuri/excelize/v2"
import "os"
import "regexp"
import "strings"

var     main_RE string = ` \([^a-zA-Z0-9]+\)$`
var     main_R1 string = ` \([^\(\)]+\)$`
var     main_R2 string = ` \([\p{Arabic}][\p{Arabic}\s،]*(\([\p{Arabic}\s،]*\))?[\p{Arabic}\s،]*\)`
var     main_R3 string = ` \([\p{Arabic}][\p{Arabic}\s،/]*(\([\p{Arabic}\s،/]*\))?[\p{Arabic}\s،/]*\)`

func    main () {
	/***1***/
	AE05  , AE10 := os.ReadFile ("i.json")
	if AE10 !=  nil {
		BE05 := fmt.Sprintf (`Data 'i.json' read failed [%s]`, AE10.Error ())
		fmt.Println (BE05)
		return
	}
	/***2***/
	AE15 := gjson.Get (string(AE05), "data.inputs").Array ()
	AE20 := map[string]map[string][]string {}
	for _ , BE10 := range AE15 {
		/***1***/
		BE21 := BE10.Get ("categoryName").String ()
		BE22 := BE10.Get ("subCategoryName").String ()
		BE23 := BE10.Get ("id").String ()
		BE24 := BE10.Get ("name").String ()
		/***2***/
		BF05 := regexp.MustCompile (main_RE).ReplaceAllString( BE21, "")
		BF10 := regexp.MustCompile (main_RE).FindString(BE21); BF10= strings.Trim (BF10, " ()")
		BF15 := regexp.MustCompile (main_RE).ReplaceAllString( BE22, "")
		BF20 := regexp.MustCompile (main_RE).FindString(BE22); BF20= strings.Trim (BF20, " ()")
		BF25 := regexp.MustCompile (main_RE).ReplaceAllString( BE24, "")
		BF30 := regexp.MustCompile (main_RE).FindString(BE24); BF30= strings.Trim (BF30, " ()")
		_= BF10
		_= BF20
		_= BE23
		_= BF30
		/***3***/
		if AE20[BF05] == nil { AE20[BF05] = map[string][]string {} }
		BG05 := AE20[BF05]
		if BG05[BF15] == nil { BG05[BF15] = [  ]string {} }
		BG10 := BG05[BF15]
		BG10  = append (BG10 , BF25)
		/***4***/
		BG05[BF15]  =   BG10
		AE20[BF05]  =   BG05
	}
	/***3***/
	AE25  , AE30 := E.OpenFile("i.xlsx")
	if AE30 !=  nil {
		BE05 := fmt.Sprintf (`Data 'i.xlsx' read failed [%s]`, AE30.Error ())
		fmt.Println (BE05)
		return
	}
	for BE05 := 1; BE05 <= 441; BE05++ {
		BE11 := fmt.Sprintf ("C%d" , BE05)
		BE12 := fmt.Sprintf ("F%d" , BE05)
		BE13 := fmt.Sprintf ("H%d" , BE05)
		BE21  , BE31 := AE25.GetCellValue ("Sheet1", BE11)
		BE22  , BE32 := AE25.GetCellValue ("Sheet1", BE12)
		if BE31 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE11, BE31.Error ())
			fmt.Println (CE05)
			return
		}
		if BE32 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE12, BE32.Error ())
			fmt.Println (CE05)
			return
		}
		BE25 := AE20 [ BE21 ]
		BE30 := BE25 [ BE22 ]
		BE35 := []map[string]any { }
		for _ , CE10 := range BE30 {
			CE15 := map[string]any {}
			CE15 ["name"] = map[string]string { "en": CE10, "ar": "" }
			CE15 ["type"] = "text"
			BE35  = append (BE35, CE15)
		}
		BE45, _ := json.Marshal(BE35);BE50 := string (BE45)
		BE55 := AE25.SetCellValue ("Sheet1", BE13, BE50)
		if BE55 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' wrtt failed [%s]`, BE13, BE55.Error ())
			fmt.Println (CE05)
			return
		}
	}
	/***4***/
	AE35 := AE25.SaveAs ("o.xlsx")
	if AE35 !=  nil {
		BE05 := fmt.Sprintf (`Data 'o.xlsx' wrtt failed [%s]`, AE35.Error ())
		fmt.Println (BE05)
		return
	}


/*
	AE25  , AE30 := json.Marshal ( AE20)
	if AE30 !=  nil {
		BE05 := fmt.Sprintf (`JSON marshal failed [%s]`, AE30.Error ())
		fmt.Println (BE05 )
		return
	}
	AE40 := os.WriteFile ("o.json", AE25, 0600)
	if AE40 !=  nil {
		BE05 := fmt.Sprintf (`Data 'o.json' read failed [%s]`, AE40.Error ())
		fmt.Println (BE05)
		return
	}
	fmt.Println ("Done")
*/
}
