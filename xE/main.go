package main

// Extract list of image files to download

import "fmt"
import "github.com/tidwall/gjson"
import  E "github.com/xuri/excelize/v2"
import "os"

func    main () {
	AE05  , AE10 := E.OpenFile ("i.xlsx")
	if AE10 !=  nil {
		BE05 := fmt.Sprintf ("Input read failed [%s]", AE10.Error ())
		fmt.Println (BE05)
		return
	}
	Data := ""
	for BE05 := 1 ; BE05 <= 48594; BE05++ {
		BE11 := fmt.Sprintf ("C%d", BE05)
		BE15  , BE20 := AE05.GetCellValue ("Sheet1", BE11)
		if BE20 !=  nil {
			CE05 := fmt.Sprintf ("Cell %d id read failed [%s]", BE05, BE20.Error ())
			fmt.Println (CE05)
			return
		}
		BE25 := fmt.Sprintf ("../xB/data/%s", BE15)
		BE35  , BE40 := os.ReadFile (BE25)
		if BE40 !=  nil {
			CE05 := fmt.Sprintf ("Cell %d dt read failed [%s]", BE05, BE40.Error ())
			fmt.Println (CE05)
			return
		}
		BE45 := gjson.Get (string(BE35), "advert.images.0.url").String ()
		Data  = Data + "\n" + BE45
	}
	AE15 := os.WriteFile ("o.txt", []byte(Data), 0600)
	if AE15 !=  nil {
		BE05 := fmt.Sprintf ("Data write failed [%s]", AE15.Error ())
		fmt.Println (BE05)
		return
	}
}
