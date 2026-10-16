package main

// Rename files from original download names to listing ID

import "fmt"
import "github.com/tidwall/gjson"
import E "github.com/xuri/excelize/v2"
import "os"
import "strings"

func    main () {
	AE05  , AE10 := E.OpenFile ("i.xlsx")
	if AE10 !=  nil {
		BE05 := fmt.Sprintf ("File read failed [%s]", AE10.Error ())
		fmt.Println (BE05)
		return
	}
	for AE15 := 1; AE15 <= 51150; AE15 ++ {
		BE05 := fmt.Sprintf("C%d",AE15)
		BE15  , BE20 := AE05.GetCellValue ("Sheet1", BE05)
		if BE20 !=  nil {
			CE05 := fmt.Sprintf ("Cell read failed [%s]", BE20.Error ())
			fmt.Println (CE05)
			return
		}
		BE25 := fmt.Sprintf ("../xB/data/%s", BE15)
		BE35  , BE40 := os.ReadFile (BE25)
		if BE40 !=  nil {
			CE05 := fmt.Sprintf ("Data read failed [%s]", BE40.Error ())
			fmt.Println (CE05)
			return
		}
		BE45 := string (BE35)
		BE50 := gjson.Get (BE45, "advert.images.0.url").String ()
		if BE50 == "" { continue }
		BE55 := strings.Split (BE50, "/")
		BE60 := BE55[len(BE55)-1]
		BE65 := strings.Split (BE60, ".")
		BE70 := "jpg"
		if BE65[1] == "webp" { BE70 = "wbp" }
		BE71 := strings.Split (BE15, ".")
		BE75 := fmt.Sprintf ("o/%s.%s", BE71[0], BE70)
		BE80 := fmt.Sprintf ("../xE/o/download/%s", BE60)
		BE85  , BE90 := os.ReadFile (BE80)
		if BE90 !=  nil {
			CE05 := fmt.Sprintf ("Srcc read failed [%s]", BE90.Error ())
			fmt.Println (CE05)
			continue
		}
		BE95 := os.WriteFile (BE75, BE85, 0600)
		fmt.Println ("Writing to destination:", BE75)
		if BE95 !=  nil {
			CE05 := fmt.Sprintf ("Dstn wrtt failed [%s]", BE95.Error ())
			fmt.Println (CE05)
			continue
		}
	}
}
