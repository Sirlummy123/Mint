package main

// Set apart ready data

import "fmt"
import  E "github.com/xuri/excelize/v2"
import "regexp"
import "strings"

func    main () {
	fmt.Println ("Running: a")
	AE05 := "../x9/i.xlsx"
	AE31  , AE32 := E.OpenFile (AE05)
	if AE32 !=  nil {
		BE05 := fmt.Sprintf (`Data 'i.xlsx' read failed [%s]`, AE32.Error ())
		fmt.Println (BE05)
		return
	}
	fmt.Println ("Running: b")
	AE41 := E.NewFile ()
	AE51 := E.NewFile ()
	var     AE55 *E.File
	var     AE60 int
	AE61 := 2
	AE62 := 1
	for BE05 := 1; BE05 <= 48594; BE05++{
		BF01 := fmt.Sprintf ("A%d", BE05)
		BF02 := fmt.Sprintf ("B%d", BE05)
		BF03 := fmt.Sprintf ("C%d", BE05)
		BF04 := fmt.Sprintf ("D%d", BE05)
		BF05 := fmt.Sprintf ("E%d", BE05)
		BF06 := fmt.Sprintf ("F%d", BE05)
		BF07 := fmt.Sprintf ("G%d", BE05)
		BF08 := fmt.Sprintf ("H%d", BE05)
		BF09 := fmt.Sprintf ("I%d", BE05)
		BF10 := fmt.Sprintf ("J%d", BE05)
		BF11 := fmt.Sprintf ("K%d", BE05)
		BF12 := fmt.Sprintf ("L%d", BE05)
		BF13 := fmt.Sprintf ("M%d", BE05)
		BF14 := fmt.Sprintf ("N%d", BE05)
		BF15 := fmt.Sprintf ("O%d", BE05)
		BF16 := fmt.Sprintf ("P%d", BE05)
		BF17 := fmt.Sprintf ("Q%d", BE05)
		BF18 := fmt.Sprintf ("R%d", BE05)
		BG01  , BH01 := AE31.GetCellValue ("Sheet1", BF01); if BH01 != nil { fmt.Println (BH01); return }
		BG02  , BH02 := AE31.GetCellValue ("Sheet1", BF02); if BH02 != nil { fmt.Println (BH02); return }
		BG03  , BH03 := AE31.GetCellValue ("Sheet1", BF03); if BH03 != nil { fmt.Println (BH03); return }
		BG04  , BH04 := AE31.GetCellValue ("Sheet1", BF04); if BH04 != nil { fmt.Println (BH04); return }
		BG05  , BH05 := AE31.GetCellValue ("Sheet1", BF05); if BH05 != nil { fmt.Println (BH05); return }
		BG06  , BH06 := AE31.GetCellValue ("Sheet1", BF06); if BH06 != nil { fmt.Println (BH06); return }
		BG07  , BH07 := AE31.GetCellValue ("Sheet1", BF07); if BH07 != nil { fmt.Println (BH07); return }
		BG08  , BH08 := AE31.GetCellValue ("Sheet1", BF08); if BH08 != nil { fmt.Println (BH08); return }
		BG09  , BH09 := AE31.GetCellValue ("Sheet1", BF09); if BH09 != nil { fmt.Println (BH09); return }
		BG10  , BH10 := AE31.GetCellValue ("Sheet1", BF10); if BH10 != nil { fmt.Println (BH10); return }
		BG11  , BH11 := AE31.GetCellValue ("Sheet1", BF11); if BH11 != nil { fmt.Println (BH11); return }
		BG12  , BH12 := AE31.GetCellValue ("Sheet1", BF12); if BH12 != nil { fmt.Println (BH12); return }
		BG13  , BH13 := AE31.GetCellValue ("Sheet1", BF13); if BH13 != nil { fmt.Println (BH13); return }
		BG14  , BH14 := AE31.GetCellValue ("Sheet1", BF14); if BH14 != nil { fmt.Println (BH14); return }
		BG15  , BH15 := AE31.GetCellValue ("Sheet1", BF15); if BH15 != nil { fmt.Println (BH15); return }
		BG16  , BH16 := AE31.GetCellValue ("Sheet1", BF16); if BH16 != nil { fmt.Println (BH16); return }
		BG17  , BH17 := AE31.GetCellValue ("Sheet1", BF17); if BH17 != nil { fmt.Println (BH17); return }
		BG18  , BH18 := AE31.GetCellValue ("Sheet1", BF18); if BH18 != nil { fmt.Println (BH18); return }
		BG15  , BG16  = Coordinate (BG14)
		fmt.Println (BG13)
		if regexp.MustCompile(`^\d{6,6}\-\d{9,9}\-\d{2,2}\.jpg$`).MatchString(BG13) == false && BG13 != "" {
			AE55  = AE41
			AE60  = AE61
			AE61  = AE61 + 1
			if BG11 == "" { BG11 = "N/A" }
		}  else {
			AE55  = AE51
			AE60  = AE62
			AE62  = AE62 + 1
		}
		
		BJ01 := fmt.Sprintf ("A%d", AE60)
		BJ02 := fmt.Sprintf ("B%d", AE60)
		BJ03 := fmt.Sprintf ("C%d", AE60)
		BJ04 := fmt.Sprintf ("D%d", AE60)
		BJ05 := fmt.Sprintf ("E%d", AE60)
		BJ06 := fmt.Sprintf ("F%d", AE60)
		BJ07 := fmt.Sprintf ("G%d", AE60)
		BJ08 := fmt.Sprintf ("H%d", AE60)
		BJ09 := fmt.Sprintf ("I%d", AE60)
		BJ10 := fmt.Sprintf ("J%d", AE60)
		BJ11 := fmt.Sprintf ("K%d", AE60)
		BJ12 := fmt.Sprintf ("L%d", AE60)
		BJ13 := fmt.Sprintf ("M%d", AE60)
		BJ14 := fmt.Sprintf ("N%d", AE60)
		BJ15 := fmt.Sprintf ("O%d", AE60)
		BJ16 := fmt.Sprintf ("P%d", AE60)
		BJ17 := fmt.Sprintf ("Q%d", AE60)
		BJ18 := fmt.Sprintf ("R%d", AE60)
		BL01 := AE55.SetCellValue ("Sheet1", BJ01, BG01); if BL01 != nil { fmt.Println (BL01); return }
		BL02 := AE55.SetCellValue ("Sheet1", BJ02, BG02); if BL02 != nil { fmt.Println (BL02); return }
		BL03 := AE55.SetCellValue ("Sheet1", BJ03, BG03); if BL03 != nil { fmt.Println (BL03); return }
		BL04 := AE55.SetCellValue ("Sheet1", BJ04, BG04); if BL04 != nil { fmt.Println (BL04); return }
		BL05 := AE55.SetCellValue ("Sheet1", BJ05, BG05); if BL05 != nil { fmt.Println (BL05); return }
		BL06 := AE55.SetCellValue ("Sheet1", BJ06, BG06); if BL06 != nil { fmt.Println (BL06); return }
		BL07 := AE55.SetCellValue ("Sheet1", BJ07, BG07); if BL07 != nil { fmt.Println (BL07); return }
		BL08 := AE55.SetCellValue ("Sheet1", BJ08, BG08); if BL08 != nil { fmt.Println (BL08); return }
		BL09 := AE55.SetCellValue ("Sheet1", BJ09, BG09); if BL09 != nil { fmt.Println (BL09); return }
		BL10 := AE55.SetCellValue ("Sheet1", BJ10, BG10); if BL10 != nil { fmt.Println (BL10); return }
		BL11 := AE55.SetCellValue ("Sheet1", BJ11, BG11); if BL11 != nil { fmt.Println (BL11); return }
		BL12 := AE55.SetCellValue ("Sheet1", BJ12, BG12); if BL12 != nil { fmt.Println (BL12); return }
		BL13 := AE55.SetCellValue ("Sheet1", BJ13, BG13); if BL13 != nil { fmt.Println (BL13); return }
		BL14 := AE55.SetCellValue ("Sheet1", BJ14, BG14); if BL14 != nil { fmt.Println (BL14); return }
		BL15 := AE55.SetCellValue ("Sheet1", BJ15, BG15); if BL15 != nil { fmt.Println (BL15); return }
		BL16 := AE55.SetCellValue ("Sheet1", BJ16, BG16); if BL16 != nil { fmt.Println (BL16); return }
		BL17 := AE55.SetCellValue ("Sheet1", BJ17, BG17); if BL17 != nil { fmt.Println (BL17); return }
		BL18 := AE55.SetCellValue ("Sheet1", BJ18, BG18); if BL18 != nil { fmt.Println (BL18); return }
	}
	AF71 := AE41.SetCellValue ("Sheet1", "A1", "user_name"     ); if AF71 != nil { fmt.Println (AF71); return }
	AF72 := AE41.SetCellValue ("Sheet1", "B1", "userName"      ); if AF72 != nil { fmt.Println (AF72); return }
	AF73 := AE41.SetCellValue ("Sheet1", "C1", "email"         ); if AF73 != nil { fmt.Println (AF73); return }
	AF74 := AE41.SetCellValue ("Sheet1", "D1", "phone"         ); if AF74 != nil { fmt.Println (AF74); return }
	AF75 := AE41.SetCellValue ("Sheet1", "E1", "password"      ); if AF75 != nil { fmt.Println (AF75); return }
	AF76 := AE41.SetCellValue ("Sheet1", "F1", "category_en"   ); if AF76 != nil { fmt.Println (AF76); return }
	AF77 := AE41.SetCellValue ("Sheet1", "G1", "category_ar"   ); if AF77 != nil { fmt.Println (AF77); return }
	AF78 := AE41.SetCellValue ("Sheet1", "H1", "subCategory_en"); if AF78 != nil { fmt.Println (AF78); return }
	AF79 := AE41.SetCellValue ("Sheet1", "I1", "subCategory_ar"); if AF79 != nil { fmt.Println (AF79); return }
	AF80 := AE41.SetCellValue ("Sheet1", "J1", "title"         ); if AF80 != nil { fmt.Println (AF80); return }
	AF81 := AE41.SetCellValue ("Sheet1", "K1", "description"   ); if AF81 != nil { fmt.Println (AF81); return }
	AF82 := AE41.SetCellValue ("Sheet1", "L1", "price"         ); if AF82 != nil { fmt.Println (AF82); return }
	AF83 := AE41.SetCellValue ("Sheet1", "M1", "images"        ); if AF83 != nil { fmt.Println (AF83); return }
	AF84 := AE41.SetCellValue ("Sheet1", "N1", "locationName"  ); if AF84 != nil { fmt.Println (AF84); return }
	AF85 := AE41.SetCellValue ("Sheet1", "O1", "latitude"      ); if AF85 != nil { fmt.Println (AF85); return }
	AF86 := AE41.SetCellValue ("Sheet1", "P1", "longitude"     ); if AF86 != nil { fmt.Println (AF86); return }
	AF87 := AE41.SetCellValue ("Sheet1", "Q1", "inputs"        ); if AF87 != nil { fmt.Println (AF87); return }
	AF88 := AE41.SetCellValue ("Sheet1", "R1", "postData"      ); if AF88 != nil { fmt.Println (AF88); return }

	AF10 := AE41.SaveAs ("x.xlsx")
	AF20 := AE51.SaveAs ("y.xlsx")
	if AF10 !=  nil {
		BE05 := fmt.Sprintf (`Data 'x.xlsx' crtn failed [%s]`, AF10.Error ())
		fmt.Println (BE05)
		return
	}
	if AF20 !=  nil {
		BE05 := fmt.Sprintf (`Data 'y.xlsx' crtn failed [%s]`, AF20.Error ())
		fmt.Println (BE05)
		return
	}
}
func    Coordinate (L string) (x, y string) {
	L = strings.ToLower (L)
	if        strings.Index (L, "abia"       ) >= 0 { x, y =  "5.5244",  "7.4934"
	} else if strings.Index (L, "adamawa"    ) >= 0 { x, y =  "9.2035", "12.4954"
	} else if strings.Index (L, "akwa ibom"  ) >= 0 { x, y =  "5.0377",  "7.9128"
	} else if strings.Index (L, "anambra"    ) >= 0 { x, y =  "6.2120",  "7.0690"
	} else if strings.Index (L, "bauchi"     ) >= 0 { x, y = "10.3103",  "9.8439"
	} else if strings.Index (L, "bayelsa"    ) >= 0 { x, y =  "4.9267",  "6.2676"
	} else if strings.Index (L, "benue"      ) >= 0 { x, y =  "7.7322",  "8.5391"
	} else if strings.Index (L, "borno"      ) >= 0 { x, y = "11.8333", "13.1500"
	} else if strings.Index (L, "cross river") >= 0 { x, y =  "4.9757",  "8.3417"
	} else if strings.Index (L, "delta"      ) >= 0 { x, y =  "6.2059",  "6.7305"
	} else if strings.Index (L, "ebonyi"     ) >= 0 { x, y =  "6.3249",  "8.1137"
	} else if strings.Index (L, "edo"        ) >= 0 { x, y =  "6.3350",  "5.6037"
	} else if strings.Index (L, "ekiti"      ) >= 0 { x, y =  "7.6211",  "5.2214"
	} else if strings.Index (L, "enugu"      ) >= 0 { x, y =  "6.5244",  "7.5106"
	} else if strings.Index (L, "gombe"      ) >= 0 { x, y = "10.2897", "11.1673"
	} else if strings.Index (L, "imo"        ) >= 0 { x, y =  "5.4840",  "7.0351"
	} else if strings.Index (L, "jigawa"     ) >= 0 { x, y = "11.7565",  "9.3390"
	} else if strings.Index (L, "kaduna"     ) >= 0 { x, y = "10.5222",  "7.4383"
	} else if strings.Index (L, "kano"       ) >= 0 { x, y = "12.0022",  "8.5920"
	} else if strings.Index (L, "katsina"    ) >= 0 { x, y = "12.9908",  "7.6018"
	} else if strings.Index (L, "kebbi"      ) >= 0 { x, y = "12.4534",  "4.1997"
	} else if strings.Index (L, "kogi"       ) >= 0 { x, y =  "7.8023",  "6.7337"
	} else if strings.Index (L, "kwara"      ) >= 0 { x, y =  "8.4966",  "4.5426"
	} else if strings.Index (L, "lagos"      ) >= 0 { x, y =  "6.6018",  "3.3515"
	} else if strings.Index (L, "nasarawa"   ) >= 0 { x, y =  "8.4934",  "8.5169"
	} else if strings.Index (L, "niger"      ) >= 0 { x, y =  "9.6139",  "6.5569"
	} else if strings.Index (L, "ogun"       ) >= 0 { x, y =  "7.1475",  "3.3619"
	} else if strings.Index (L, "ondo"       ) >= 0 { x, y =  "7.2571",  "5.2058"
	} else if strings.Index (L, "osun"       ) >= 0 { x, y =  "7.7719",  "4.5567"
	} else if strings.Index (L, "oyo"        ) >= 0 { x, y =  "7.3775",  "3.9470"
	} else if strings.Index (L, "plateau"    ) >= 0 { x, y =  "9.8965",  "8.8583"
	} else if strings.Index (L, "rivers"     ) >= 0 { x, y =  "4.8156",  "7.0498"
	} else if strings.Index (L, "sokoto"     ) >= 0 { x, y = "13.0059",  "5.2476"
	} else if strings.Index (L, "taraba"     ) >= 0 { x, y =  "8.8833", "11.3667"
	} else if strings.Index (L, "yobe"       ) >= 0 { x, y = "11.7467", "11.9608"
	} else if strings.Index (L, "zamfara"    ) >= 0 { x, y = "12.1704",  "6.6642"
	} else if strings.Index (L, "fct"        ) >= 0 { x, y =  "9.0765",  "7.3986"
	} else if strings.Index (L, "abuja"      ) >= 0 { x, y =  "9.0765",  "7.3986"
	} else { x, y = "0.0000", "0.0000"
	}
	return
}

