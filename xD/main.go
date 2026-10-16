package main

// Add category name to listing list

import "fmt"
import "github.com/tidwall/gjson"
import E "github.com/xuri/excelize/v2"
import "os"

func    main () {
	AE05  , AE10 := E.OpenFile ("i.xlsx")
	if AE10 !=  nil {
		BE05 := fmt.Sprintf ("File read failed [%s]", AE10.Error ())
		fmt.Println (BE05)
		return
	}
	for AE15 := 1; AE15 <= 48594; AE15 ++ {
		BE05 := fmt.Sprintf ("C%d", AE15)
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
		BE50 := gjson.Get (BE45, "advert.category_name").String ()
		BE55 := fmt.Sprintf ("D%d", AE15)
		AE05.SetCellValue ("Sheet1", BE55, BE50)
		fmt.Println (AE15)
	}
	AE05.SaveAs ("o.xlsx")
}
