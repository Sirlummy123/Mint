package main

import "fmt"
import "github.com/tidwall/gjson"
import E "github.com/xuri/excelize/v2"
import "os"
import "time"

func    main () {
	AE05  , AE10 := os.ReadFile ("i.json")
	if AE10 !=  nil {
		BE05 := fmt.Sprintf (`Data read failed [%s]`, AE10.Error ())
		fmt.Println (BE05)
		return
	}
	AE15 := string(AE05)
	AE20 := gjson.Get (AE15, "data").Array ()
	AE25 := []map[string]string{}
	for _ , BE10 := range AE20 {
		BE21 := BE10.Get ("categoryId._id").String ()
		BE22 := BE10.Get ("categoryId.name.en").String ()
		BE23 := BE10.Get ("categoryId.name.ar").String ()
		BE24 := BE10.Get ("_id").String ()
		BE25 := BE10.Get ("name.en").String ()
		BE26 := BE10.Get ("name.ar").String ()
		BE35 := map[string]string {}
		BE35 ["ct/id"] = BE21
		BE35 ["ct/en"] = BE22
		BE35 ["ct/ar"] = BE23
		BE35 ["sc/id"] = BE24
		BE35 ["sc/en"] = BE25
		BE35 ["sc/ar"] = BE26
		AE25  = append (AE25, BE35)
	}
	AE35 := E.NewFile ()
	AE40 := fmt.Sprintf ("%d.xlsx", time.Now().Unix())
	defer AE35.SaveAs (AE40)
	for BE05 , BE10 := range AE25 {
		BE21 := fmt.Sprintf ("A%d", BE05+1)
		BE22 := fmt.Sprintf ("B%d", BE05+1)
		BE23 := fmt.Sprintf ("C%d", BE05+1)
		BE24 := fmt.Sprintf ("D%d", BE05+1)
		BE25 := fmt.Sprintf ("E%d", BE05+1)
		BE26 := fmt.Sprintf ("F%d", BE05+1)
		BE27 := fmt.Sprintf ("G%d", BE05+1)
		AE35.SetCellValue ("Sheet1", BE21, fmt.Sprintf("%d",BE05+1))
		AE35.SetCellValue ("Sheet1", BE22, BE10["ct/id"])
		AE35.SetCellValue ("Sheet1", BE23, BE10["ct/en"])
		AE35.SetCellValue ("Sheet1", BE24, BE10["ct/ar"])
		AE35.SetCellValue ("Sheet1", BE25, BE10["sc/id"])
		AE35.SetCellValue ("Sheet1", BE26, BE10["sc/en"])
		AE35.SetCellValue ("Sheet1", BE27, BE10["sc/ar"])
	}
}
