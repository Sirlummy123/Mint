package main

// Compose pre-upload data

import "encoding/json"
import "fmt"
import "github.com/tidwall/gjson"
import E "github.com/xuri/excelize/v2"
import "os"
import "strings"

func    main () {
	/***1***/
	AE25  , AE30 := E.OpenFile ("a.xlsx")
	if AE30 !=  nil {
		BE05 := fmt.Sprintf(`Data 'a.xlsx' read failed [%s]`, AE30.Error ())
		fmt.Println (BE05)
		return
	}
	AE31  , AE32 := E.OpenFile ("b.xlsx")
	if AE32 !=  nil {
		BE05 := fmt.Sprintf(`Data 'a.xlsx' read failed [%s]`, AE32.Error ())
		fmt.Println (BE05)
		return
	}
	/***2***/
	AE35 := map[string]map[string]string {}
	for BE05 := 1; BE05 <= 529; BE05++ {
		/***1***/
		BE21 := fmt.Sprintf ("B%d", BE05)
		BE22 := fmt.Sprintf ("D%d", BE05)
		BE23 := fmt.Sprintf ("F%d", BE05)
		BE24 := fmt.Sprintf ("G%d", BE05)
		BE25 := fmt.Sprintf ("I%d", BE05)
		BE26 := fmt.Sprintf ("J%d", BE05)
		BE27 := fmt.Sprintf ("K%d", BE05)
		BE31  , BE41 := AE25.GetCellValue ("Sheet1", BE21)
		BE32  , BE42 := AE25.GetCellValue ("Sheet1", BE22)
		BE33  , BE43 := AE25.GetCellValue ("Sheet1", BE23)
		BE34  , BE44 := AE25.GetCellValue ("Sheet1", BE24)
		BE35  , BE45 := AE25.GetCellValue ("Sheet1", BE25)
		BE36  , BE46 := AE25.GetCellValue ("Sheet1", BE26)
		BE37  , BE47 := AE25.GetCellValue ("Sheet1", BE27)
		/***2***/
		if BE41 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE21, BE41.Error ())
			fmt.Println (CE05)
			return
		}
		if BE42 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE22, BE42.Error ())
			fmt.Println (CE05)
			return
		}
		if BE43 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE23, BE43.Error ())
			fmt.Println (CE05)
			return
		}
		if BE44 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE24, BE44.Error ())
			fmt.Println (CE05)
			return
		}
		if BE45 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE25, BE45.Error ())
			fmt.Println (CE05)
			return
		}
		if BE46 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE26, BE46.Error ())
			fmt.Println (CE05)
			return
		}
		if BE47 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE27, BE47.Error ())
			fmt.Println (CE05)
			return
		}
		BE31  = strings.Trim(BE31, " \t\r\n")
		BE32  = strings.Trim(BE32, " \t\r\n")
		BE33  = strings.Trim(BE33, " \t\r\n")
		BE34  = strings.Trim(BE34, " \t\r\n")
		BE35  = strings.Trim(BE35, " \t\r\n")
		BE36  = strings.Trim(BE36, " \t\r\n")
		BE37  = strings.Trim(BE37, " \t\r\n")
		if BE31 ==   "" { continue }
		/***3***/
		BF05 := map[string]string {}
		BF05["ct_id"] = BE32
		BF05["ct_en"] = BE33
		BF05["ct_ar"] = BE34
		BF05["sc_en"] = BE35
		BF05["sc_ar"] = BE36
		BF05["sc_at"] = BE37
		AE35[  BE31 ] = BF05
	}
	/***3***/
	AF05 := E.NewFile ()
	AF10 := 0
	for BE05 := 1; BE05 <= 48594; BE05++{ AF10 = AF10 + 1
		/***1***/
		fmt.Println (fmt.Sprintf ("Processing: %d" , BE05))
		BE21 := fmt.Sprintf ("B%d", BE05)
		BE22 := fmt.Sprintf ("C%d", BE05)
		BE23 := fmt.Sprintf ("D%d", BE05)
		BE31  , BE41 := AE31.GetCellValue ("Sheet1", BE21)
		BE32  , BE42 := AE31.GetCellValue ("Sheet1", BE22)
		BE33  , BE43 := AE31.GetCellValue ("Sheet1", BE23)
		/***2***/
		if BE41 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE21, BE41.Error ())
			fmt.Println (CE05)
			return
		}
		if BE42 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE22, BE42.Error ())
			fmt.Println (CE05)
			return
		}
		if BE43 !=  nil {
			CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BE23, BE43.Error ())
			fmt.Println (CE05)
			return
		}
		if BE33 ==   "" { continue }
		BE31  = strings.Trim (BE31, " \t\r\n")
		BE32  = strings.Trim (BE32, " \t\r\n")
		BE33  = strings.Trim (BE33, " \t\r\n")
		BE55 := AE35 [BE33]
//		BE60 := fmt.Sprintf ("../../Mint0/20: Listing Data Harvester/o/listing/%s", BE32)
		BE60 := fmt.Sprintf ("../xB/data/%s", BE32)
		BE65  , BE70 := os.ReadFile (BE60)
		if BE70 !=  nil {
			CE05 := fmt.Sprintf (`Record read failed [%s]`, BE70.Error ())
			fmt.Println (CE05)
			return
		}
		BE75 := string (BE65)
		/***3***/
		BF01 := ""; BF02 := ""; BF03 := ""; BF04 := ""; BF05 := ""; BF06 := ""
		BF07 := ""; BF08 := ""; BF09 := ""; BF10 := ""; BF11 := ""; BF12 := ""
		BF13 := ""; BF14 := ""; BF15 := ""; BF16 := ""; BF17 := ""; BF18 := ""
		/***4***/
		BF01  = gjson.Get (BE75, "seller.name").String ()
		BF04  = gjson.Get (BE75, "seller.phone").String (); BF04 = strings.Trim (BF04, " ")
		if len(BF04) != 11 { BF04 = "08145798559" }
		BF02  = fmt.Sprintf ("user%s", BF04)
		BF03  = fmt.Sprintf ("%s@example.com", BF02)
		BF05  = "123456"
		BF06  = BE55 ["ct_en"]
		BF07  = BE55 ["ct_ar"]
		BF08  = BE55 ["sc_en"]
		BF09  = BE55 ["sc_ar"]
		BF10  = gjson.Get (BE75, "advert.title").String ()
		BF11  = gjson.Get (BE75, "advert.description").String ()
		BF12  = "0"
		if gjson.Get (BE75, "advert.price.value").Int () > 0 {
			CE05 := gjson.Get (BE75, "advert.price.value").Int () / 1500
			BF12  = fmt.Sprintf ("%d", CE05 )
		}
		BF13  = strings.ReplaceAll (BE32, ".json", ".jpg")
		BF14  = gjson.Get (BE75, "advert.region_text").String ()
		/***5***/
		BG01 := []map[string]any {}
		BG02 :=   map[string]any {}
		BG03 := []map[string]any {}
		BG25 := json.Unmarshal ([]byte(BE55["sc_at"]), &BG01)
		if BG25 !=  nil {
			CE05 := fmt.Sprintf (`Attributes unmarshal failed [%s]`, BG25.Error ())
			fmt.Println ( CE05 )
//			fmt.Println ( BE55["sc_at"], BE33, BE55, BE32 )
//			return
			continue
		}
		for _ , CE10 := range BG01 {
			CE15, _ := json.Marshal (CE10); CE20 := string (CE15)
			CE31 := gjson.Get (CE20, "name.en").String ()
			CE32 := gjson.Get (CE20, "name.ar").String ()
			CE33 := gjson.Get (CE20, "text"   ).String ()
			if CE32 == "" { continue }
			CE41 := map[string]any{}
			CE42 := map[string]any{}
			CE42[ "en" ] = CE31
			CE42[ "ar" ] = ""
			CE41["name"] = CE42
			CE41["text"] = CE33
			BG03  = append(BG03, CE41)
			CE55 := map[string]any{}
			CE60 := gjson.Get (BE75, "advert.attrs").Array ()
			for _ , DE10 := range CE60 {
				DE21 := DE10.Get ("name" ).String ()
				DE22 := DE10.Get ("value").String ()
				CE55[DE21] = DE22
			}
			if        CE32 == "__price__"    { BG02[CE31] = BF12
			} else if CE32 == "__location__" { BG02[CE31] = BF14
			} else if CE32 == "Make/Model"   { BG02[CE31] = fmt.Sprintf ("%s %s", CE55["Make"], CE55["Model"])
			} else /***********************/ { BG02[CE31] = CE55[CE32]
			}
		}
		BG42, _ := json.Marshal (BG02); BF18 = string(BG42)
		BG43, _ := json.Marshal (BG03); BF17 = string(BG43)
		/***6***/
		BG61 := fmt.Sprintf ("A%d", AF10)
		BG62 := fmt.Sprintf ("B%d", AF10)
		BG63 := fmt.Sprintf ("C%d", AF10)
		BG64 := fmt.Sprintf ("D%d", AF10)
		BG65 := fmt.Sprintf ("E%d", AF10)
		BG66 := fmt.Sprintf ("F%d", AF10)
		BG67 := fmt.Sprintf ("G%d", AF10)
		BG68 := fmt.Sprintf ("H%d", AF10)
		BG69 := fmt.Sprintf ("I%d", AF10)
		BG70 := fmt.Sprintf ("J%d", AF10)
		BG71 := fmt.Sprintf ("K%d", AF10)
		BG72 := fmt.Sprintf ("L%d", AF10)
		BG73 := fmt.Sprintf ("M%d", AF10)
		BG74 := fmt.Sprintf ("N%d", AF10)
		BG75 := fmt.Sprintf ("O%d", AF10)
		BG76 := fmt.Sprintf ("P%d", AF10)
		BG77 := fmt.Sprintf ("Q%d", AF10)
		BG78 := fmt.Sprintf ("R%d", AF10)
		BH01 := AF05.SetCellValue ("Sheet1", BG61, BF01)
		BH02 := AF05.SetCellValue ("Sheet1", BG62, BF02)
		BH03 := AF05.SetCellValue ("Sheet1", BG63, BF03)
		BH04 := AF05.SetCellValue ("Sheet1", BG64, BF04)
		BH05 := AF05.SetCellValue ("Sheet1", BG65, BF05)
		BH06 := AF05.SetCellValue ("Sheet1", BG66, BF06)
		BH07 := AF05.SetCellValue ("Sheet1", BG67, BF07)
		BH08 := AF05.SetCellValue ("Sheet1", BG68, BF08)
		BH09 := AF05.SetCellValue ("Sheet1", BG69, BF09)
		BH10 := AF05.SetCellValue ("Sheet1", BG70, BF10)
		BH11 := AF05.SetCellValue ("Sheet1", BG71, BF11)
		BH12 := AF05.SetCellValue ("Sheet1", BG72, BF12)
		BH13 := AF05.SetCellValue ("Sheet1", BG73, BF13)
		BH14 := AF05.SetCellValue ("Sheet1", BG74, BF14)
		BH15 := AF05.SetCellValue ("Sheet1", BG75, BF15)
		BH16 := AF05.SetCellValue ("Sheet1", BG76, BF16)
		BH17 := AF05.SetCellValue ("Sheet1", BG77, BF17)
		BH18 := AF05.SetCellValue ("Sheet1", BG78, BF18)
		if BH01 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG61, BH01.Error ()); fmt.Println (CE05); return }
		if BH02 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG62, BH02.Error ()); fmt.Println (CE05); return }
		if BH03 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG63, BH03.Error ()); fmt.Println (CE05); return }
		if BH04 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG64, BH04.Error ()); fmt.Println (CE05); return }
		if BH05 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG65, BH05.Error ()); fmt.Println (CE05); return }
		if BH06 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG66, BH06.Error ()); fmt.Println (CE05); return }
		if BH07 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG67, BH07.Error ()); fmt.Println (CE05); return }
		if BH08 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG68, BH08.Error ()); fmt.Println (CE05); return }
		if BH09 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG69, BH09.Error ()); fmt.Println (CE05); return }
		if BH10 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG70, BH10.Error ()); fmt.Println (CE05); return }
		if BH11 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG71, BH11.Error ()); fmt.Println (CE05); return }
		if BH12 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG72, BH12.Error ()); fmt.Println (CE05); return }
		if BH13 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG73, BH13.Error ()); fmt.Println (CE05); return }
		if BH14 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG74, BH14.Error ()); fmt.Println (CE05); return }
		if BH15 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG75, BH15.Error ()); fmt.Println (CE05); return }
		if BH16 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG76, BH16.Error ()); fmt.Println (CE05); return }
		if BH17 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG77, BH17.Error ()); fmt.Println (CE05); return }
		if BH18 != nil { CE05 := fmt.Sprintf (`Cell '%s' read failed [%s]`, BG78, BH18.Error ()); fmt.Println (CE05); return }
	}
	AF20 := AF05.SaveAs ("o.xlsx")
	if AF20 !=  nil {
		BE05 := fmt.Sprintf(`Data 'o.xlsx' wrtt failed [%s]`, AF20.Error ())
		fmt.Println (BE05)
		return
	}
}

