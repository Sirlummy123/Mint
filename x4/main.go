package main

// Extract seller info

import "fmt"
import "github.com/tidwall/gjson"
import E "github.com/xuri/excelize/v2"
import "os"
import "regexp"
import "strings"

func    main () {
	/***1***/
	AE31  , AE32 := E.OpenFile ("b.xlsx")
	if AE32 !=  nil {
		BE05 := fmt.Sprintf(`Data 'a.xlsx' read failed [%s]`, AE32.Error ())
		fmt.Println (BE05)
		return
	}
	/***2***/
	AF05 := E.NewFile ()
	AF10 := 0
	AF15 := ""
	for BE05 := 15401; BE05 <= 15950; BE05++{ AF10 = AF10 + 1
		/***1***/
		fmt.Println (fmt.Sprintf ("Processing: %d" , BE05))
		BE22 := fmt.Sprintf ("C%d", BE05)
		BE32  , BE42 := AE31.GetCellValue ("Sheet1", BE22)
		/***2***/
		if BE42 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE22, BE42.Error ())
			fmt.Println (CE05)
			return
		}
		BE32  = strings.Trim(BE32, " \t\r\n")
		BE60 := fmt.Sprintf ("../../Mint0/20: Listing Data Harvester/o/listing/%s", BE32)
		BE65  , BE70 := os.ReadFile (BE60)
		if BE70 !=  nil {
			CE05 := fmt.Sprintf (`Record read failed [%s]`, BE70.Error ())
			fmt.Println (CE05)
			return
		}
		BE75 := string (BE65)
		/***3***/
		BF01 := gjson.Get (BE75,"advert.images.0.url").String ()
		BF02 := gjson.Get (BE75,"seller.id").String (); for len(BF02) < 9 { BF02 = "0"+BF02 }
		BF02  = "user01" + BF02
		BF04 := gjson.Get (BE75,"seller.page_url").String ()
		if len(BF04) != 0 && BF04[0] != 'h' { BF04 = fmt.Sprintf ("https://jiji.ng%s", BF04) }
		/***6***/
		BF55 := regexp.MustCompile(`[a-zA-Z]+$`).FindString(BF01)
		BF60 := strings.Split(BE32,".")
		BF65 := fmt.Sprintf ("%s.%s", BF60[0], BF55)
		/***6***/
		BG61 := fmt.Sprintf ("B%d", AF10)
		BG62 := fmt.Sprintf ("C%d", AF10)
		BG63 := fmt.Sprintf ("D%d", AF10)
		BG64 := fmt.Sprintf ("E%d", AF10)
		BH01 := AF05.SetCellValue ("Sheet1", BG61, BF01)
		BH02 := AF05.SetCellValue ("Sheet1", BG62, BF02)
		BH03 := AF05.SetCellValue ("Sheet1", BG63, BF65)
		BH04 := AF05.SetCellValue ("Sheet1", BG64, BF04)
		if BH01 != nil { CE05 := fmt.Sprintf (`Cell '%s' wrtt failed [%s]`, BG61, BH01.Error ()); fmt.Println (CE05); return }
		if BH02 != nil { CE05 := fmt.Sprintf (`Cell '%s' wrtt failed [%s]`, BG62, BH02.Error ()); fmt.Println (CE05); return }
		if BH03 != nil { CE05 := fmt.Sprintf (`Cell '%s' wrtt failed [%s]`, BG63, BH03.Error ()); fmt.Println (CE05); return }
		if BH04 != nil { CE05 := fmt.Sprintf (`Cell '%s' wrtt failed [%s]`, BG64, BH04.Error ()); fmt.Println (CE05); return }
		AF15  = AF15 + fmt.Sprintf("%s\n out=%s\n",BF01,BF65)
	}
	/***3***/
	AF20 := AF05.SaveAs ("o.xlsx")
	if AF20 !=  nil {
		BE05 := fmt.Sprintf(`Data 'o.xlsx' wrtt failed [%s]`, AF20.Error ())
		fmt.Println (BE05)
		return
	}
	AF25 := os.WriteFile("o.data", []byte(AF15), 0600)
	if AF25 !=  nil {
		BE05 := fmt.Sprintf(`Data 'o.data' wrtt failed [%s]`, AF25.Error ())
		fmt.Println (BE05)
		return
	}
}
