package main

import "fmt"
import "github.com/tidwall/gjson"
import E "github.com/xuri/excelize/v2"
import "io"
import "net/http"
import "strings"

func    main () {
	AE05  , AE10 := E.OpenFile ("i.xlsx")
	if AE10 !=  nil {
		BE05 := fmt.Sprintf ("File read failed [%s]", AE10.Error ())
		fmt.Println (BE05)
		return
	}
	for BE15 := 2; BE15 <= 164; BE15 ++ {
		BE20 := fmt.Sprintf ("C%d", BE15)
		BE22 := fmt.Sprintf ("D%d", BE15)
		BE25  , BE30 := AE05.GetCellValue ("Sheet1", BE20)
		if BE30 !=  nil {
			CE05 := fmt.Sprintf ("Data read failed [%s]", BE30.Error ())
			fmt.Println (CE05)
			return
		}
		BE25  = strings.ReplaceAll (BE25, "https://jiji.ng/", "")
		BE35 := fmt.Sprintf (`https://jiji.ng/api_web/v1/listing?slug=%s&init_page=true&page=1&webp=true`, BE25)
		BE45  , BE50 := http.Get (BE35)
		if BE50 !=  nil {
			CE05 := fmt.Sprintf ("Request load failed [%s]", BE50.Error ())
			fmt.Println (CE05)
			return
		}
		BE55  , BE60 := io.ReadAll (BE45.Body)
		if BE60 !=  nil {
			CE05 := fmt.Sprintf ("Request read failed [%s]", BE60.Error ())
			fmt.Println (CE05)
			return
		}
		BE45.Body.Close ()
		BE65 := string (BE55)
		BE75 := gjson.Get (BE65, "adverts_list.adverts.0.category_id").String ()
		fmt.Println (BE75)
		BE85 := AE05.SetCellValue ("Sheet1", BE22, BE75)
		if BE85 !=  nil {
			CE05 := fmt.Sprintf ("Data wrtt failed [%s]", BE85.Error ())
			fmt.Println (CE05)
			return
		}
	}
	AE20 := AE05.SaveAs ("o.xlsx")
	if AE20 !=  nil {
		BE05 := fmt.Sprintf ("File save failed [%s]", AE20.Error ())
		fmt.Println (BE05)
		return
	}
}
