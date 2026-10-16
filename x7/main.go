package main

// Cluster images together

import "encoding/json"
import "fmt"
import "image"
import "image/draw"
import "image/png"
import "os"
import "slices"
import "strings"

func    main () {
	AE01 := "../x6/q/1"
	AE05  , AE10 := os.ReadDir (AE01)
	if AE10 !=  nil {
		BE05 := fmt.Sprintf (`Images load failed [%s]`, AE10.Error ())
		fmt.Println (BE05)
		return
	}
	AE15 := [  ]string {}
	for _, BE10 := range AE05 { AE15 = append (AE15, BE10.Name()) }
	slices.Sort(AE15)
	AE20 := [][]string {}
	AE25 := len(AE15)/ 16
	for BE05 := 1; BE05 <= AE25; BE05++ {
		BF01 := ((BE05-1)*16) +  0
		BF02 := ((BE05-1)*16) +  1
		BF03 := ((BE05-1)*16) +  2
		BF04 := ((BE05-1)*16) +  3
		BF05 := ((BE05-1)*16) +  4
		BF06 := ((BE05-1)*16) +  5
		BF07 := ((BE05-1)*16) +  6
		BF08 := ((BE05-1)*16) +  7
		BF09 := ((BE05-1)*16) +  8
		BF10 := ((BE05-1)*16) +  9
		BF11 := ((BE05-1)*16) + 10
		BF12 := ((BE05-1)*16) + 11
		BF13 := ((BE05-1)*16) + 12
		BF14 := ((BE05-1)*16) + 13
		BF15 := ((BE05-1)*16) + 14
		BF16 := ((BE05-1)*16) + 15
		BG05 := []string {
			AE15[BF01], AE15[BF02], AE15[BF03], AE15[BF04], AE15[BF05], AE15[BF06],
			AE15[BF07], AE15[BF08], AE15[BF09], AE15[BF10], AE15[BF11], AE15[BF12],
			AE15[BF13], AE15[BF14], AE15[BF15], AE15[BF16],
		}
		AE20  = append (AE20, BG05)
	}
	AF05 := false
	for _ , BE10 := range AE20 {
		fmt.Println (BE10[0])
		if BE10[0] == "000388-031284199-03.png" { AF05 = true }
		if AF05 != true { continue }
		BF01 := fmt.Sprintf ("%s/%s", AE01, BE10[ 0])
		BF02 := fmt.Sprintf ("%s/%s", AE01, BE10[ 1])
		BF03 := fmt.Sprintf ("%s/%s", AE01, BE10[ 2])
		BF04 := fmt.Sprintf ("%s/%s", AE01, BE10[ 3])
		BF05 := fmt.Sprintf ("%s/%s", AE01, BE10[ 4])
		BF06 := fmt.Sprintf ("%s/%s", AE01, BE10[ 5])
		BF07 := fmt.Sprintf ("%s/%s", AE01, BE10[ 6])
		BF08 := fmt.Sprintf ("%s/%s", AE01, BE10[ 7])
		BF09 := fmt.Sprintf ("%s/%s", AE01, BE10[ 8])
		BF10 := fmt.Sprintf ("%s/%s", AE01, BE10[ 9])
		BF11 := fmt.Sprintf ("%s/%s", AE01, BE10[10])
		BF12 := fmt.Sprintf ("%s/%s", AE01, BE10[11])
		BF13 := fmt.Sprintf ("%s/%s", AE01, BE10[12])
		BF14 := fmt.Sprintf ("%s/%s", AE01, BE10[13])
		BF15 := fmt.Sprintf ("%s/%s", AE01, BE10[14])
		BF16 := fmt.Sprintf ("%s/%s", AE01, BE10[15])
		BG01  , BH01 := os.Open   (BF01); if BH01 != nil { fmt.Println (BH01.Error()); return }
		BG02  , BH02 := os.Open   (BF02); if BH02 != nil { fmt.Println (BH02.Error()); return }
		BG03  , BH03 := os.Open   (BF03); if BH03 != nil { fmt.Println (BH03.Error()); return }
		BG04  , BH04 := os.Open   (BF04); if BH04 != nil { fmt.Println (BH04.Error()); return }
		BG05  , BH05 := os.Open   (BF05); if BH05 != nil { fmt.Println (BH05.Error()); return }
		BG06  , BH06 := os.Open   (BF06); if BH06 != nil { fmt.Println (BH06.Error()); return }
		BG07  , BH07 := os.Open   (BF07); if BH07 != nil { fmt.Println (BH07.Error()); return }
		BG08  , BH08 := os.Open   (BF08); if BH08 != nil { fmt.Println (BH08.Error()); return }
		BG09  , BH09 := os.Open   (BF09); if BH09 != nil { fmt.Println (BH09.Error()); return }
		BG10  , BH10 := os.Open   (BF10); if BH10 != nil { fmt.Println (BH10.Error()); return }
		BG11  , BH11 := os.Open   (BF11); if BH11 != nil { fmt.Println (BH11.Error()); return }
		BG12  , BH12 := os.Open   (BF12); if BH12 != nil { fmt.Println (BH12.Error()); return }
		BG13  , BH13 := os.Open   (BF13); if BH13 != nil { fmt.Println (BH13.Error()); return }
		BG14  , BH14 := os.Open   (BF14); if BH14 != nil { fmt.Println (BH14.Error()); return }
		BG15  , BH15 := os.Open   (BF15); if BH15 != nil { fmt.Println (BH15.Error()); return }
		BG16  , BH16 := os.Open   (BF16); if BH16 != nil { fmt.Println (BH16.Error()); return }
		BI01  , BJ01 := png.Decode(BG01); if BJ01 != nil { fmt.Println (BJ01.Error()); return }
		BI02  , BJ02 := png.Decode(BG02); if BJ02 != nil { fmt.Println (BJ02.Error()); return }
		BI03  , BJ03 := png.Decode(BG03); if BJ03 != nil { fmt.Println (BJ03.Error()); return }
		BI04  , BJ04 := png.Decode(BG04); if BJ04 != nil { fmt.Println (BJ04.Error()); return }
		BI05  , BJ05 := png.Decode(BG05); if BJ05 != nil { fmt.Println (BJ05.Error()); return }
		BI06  , BJ06 := png.Decode(BG06); if BJ06 != nil { fmt.Println (BJ06.Error()); return }
		BI07  , BJ07 := png.Decode(BG07); if BJ07 != nil { fmt.Println (BJ07.Error()); return }
		BI08  , BJ08 := png.Decode(BG08); if BJ08 != nil { fmt.Println (BJ08.Error()); return }
		BI09  , BJ09 := png.Decode(BG09); if BJ09 != nil { fmt.Println (BJ09.Error()); return }
		BI10  , BJ10 := png.Decode(BG10); if BJ10 != nil { fmt.Println (BJ10.Error()); return }
		BI11  , BJ11 := png.Decode(BG11); if BJ11 != nil { fmt.Println (BJ11.Error()); return }
		BI12  , BJ12 := png.Decode(BG12); if BJ12 != nil { fmt.Println (BJ12.Error()); return }
		BI13  , BJ13 := png.Decode(BG13); if BJ13 != nil { fmt.Println (BJ13.Error()); return }
		BI14  , BJ14 := png.Decode(BG14); if BJ14 != nil { fmt.Println (BJ14.Error()); return }
		BI15  , BJ15 := png.Decode(BG15); if BJ15 != nil { fmt.Println (BJ15.Error()); return }
		BI16  , BJ16 := png.Decode(BG16); if BJ16 != nil { fmt.Println (BJ16.Error()); return }
		BG01.Close()
		BG02.Close()
		BG03.Close()
		BG04.Close()
		BG05.Close()
		BG06.Close()
		BG07.Close()
		BG08.Close()
		BG09.Close()
		BG10.Close()
		BG11.Close()
		BG12.Close()
		BG13.Close()
		BG14.Close()
		BG15.Close()
		BG16.Close()
		BK05 := image.NewRGBA (image.Rect(0,0,6000,6000))
		draw.Draw (BK05, image.Rect(   0,   0,1500,1500), BI01, BI01.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(1500,   0,3000,1500), BI02, BI02.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(3000,   0,4500,1500), BI03, BI03.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(4500,   0,6000,1500), BI04, BI04.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(   0,1500,1500,3000), BI05, BI05.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(1500,1500,3000,3000), BI06, BI06.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(3000,1500,4500,3000), BI07, BI07.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(4500,1500,6000,3000), BI08, BI08.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(   0,3000,1500,4500), BI09, BI09.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(1500,3000,3000,4500), BI10, BI10.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(3000,3000,4500,4500), BI11, BI11.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(4500,3000,6000,4500), BI12, BI12.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(   0,4500,1500,6000), BI13, BI13.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(1500,4500,3000,6000), BI14, BI14.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(3000,4500,4500,6000), BI15, BI15.Bounds().Min, draw.Src)
		draw.Draw (BK05, image.Rect(4500,4500,6000,6000), BI16, BI16.Bounds().Min, draw.Src)
		BL05 := fmt.Sprintf ("o/1/%s", BE10[0])
		BL10 := fmt.Sprintf ("p/1/%s", BE10[0]); BL10 = strings.ReplaceAll (BL10, ".png", ".json")
		BL15  , BL20 := os.OpenFile (BL05, os.O_CREATE|os.O_RDWR, 0600)
		if BL20 !=  nil {
			CE05 := fmt.Sprintf (`Images encd failed [%s]`, BL20.Error ())
			fmt.Println (CE05)
			return
		}
		BL25 := png.Encode  (BL15, BK05)
		if BL25 !=  nil {
			CE05 := fmt.Sprintf (`Images sve1 failed [%s]`, BL25.Error ())
			fmt.Println (CE05)
			return
		}
		BL15.Close ()
		BM05 := map[string]any{}
		BN01  , BP01 := os.Open         (BF01); if BP01 != nil { fmt.Println (BP01.Error()); return }
		BN02  , BP02 := os.Open         (BF02); if BP02 != nil { fmt.Println (BP02.Error()); return }
		BN03  , BP03 := os.Open         (BF03); if BP03 != nil { fmt.Println (BP03.Error()); return }
		BN04  , BP04 := os.Open         (BF04); if BP04 != nil { fmt.Println (BP04.Error()); return }
		BN05  , BP05 := os.Open         (BF05); if BP05 != nil { fmt.Println (BP05.Error()); return }
		BN06  , BP06 := os.Open         (BF06); if BP06 != nil { fmt.Println (BP06.Error()); return }
		BN07  , BP07 := os.Open         (BF07); if BP07 != nil { fmt.Println (BP07.Error()); return }
		BN08  , BP08 := os.Open         (BF08); if BP08 != nil { fmt.Println (BP08.Error()); return }
		BN09  , BP09 := os.Open         (BF09); if BP09 != nil { fmt.Println (BP09.Error()); return }
		BN10  , BP10 := os.Open         (BF10); if BP10 != nil { fmt.Println (BP10.Error()); return }
		BN11  , BP11 := os.Open         (BF11); if BP11 != nil { fmt.Println (BP11.Error()); return }
		BN12  , BP12 := os.Open         (BF12); if BP12 != nil { fmt.Println (BP12.Error()); return }
		BN13  , BP13 := os.Open         (BF13); if BP13 != nil { fmt.Println (BP13.Error()); return }
		BN14  , BP14 := os.Open         (BF14); if BP14 != nil { fmt.Println (BP14.Error()); return }
		BN15  , BP15 := os.Open         (BF15); if BP15 != nil { fmt.Println (BP15.Error()); return }
		BN16  , BP16 := os.Open         (BF16); if BP16 != nil { fmt.Println (BP16.Error()); return }
		BQ01  , BR01 := png.DecodeConfig(BN01); if BR01 != nil { fmt.Println (BR01.Error()); return }
		BQ02  , BR02 := png.DecodeConfig(BN02); if BR02 != nil { fmt.Println (BR02.Error()); return }
		BQ03  , BR03 := png.DecodeConfig(BN03); if BR03 != nil { fmt.Println (BR03.Error()); return }
		BQ04  , BR04 := png.DecodeConfig(BN04); if BR04 != nil { fmt.Println (BR04.Error()); return }
		BQ05  , BR05 := png.DecodeConfig(BN05); if BR05 != nil { fmt.Println (BR05.Error()); return }
		BQ06  , BR06 := png.DecodeConfig(BN06); if BR06 != nil { fmt.Println (BR06.Error()); return }
		BQ07  , BR07 := png.DecodeConfig(BN07); if BR07 != nil { fmt.Println (BR07.Error()); return }
		BQ08  , BR08 := png.DecodeConfig(BN08); if BR08 != nil { fmt.Println (BR08.Error()); return }
		BQ09  , BR09 := png.DecodeConfig(BN09); if BR09 != nil { fmt.Println (BR09.Error()); return }
		BQ10  , BR10 := png.DecodeConfig(BN10); if BR10 != nil { fmt.Println (BR10.Error()); return }
		BQ11  , BR11 := png.DecodeConfig(BN11); if BR11 != nil { fmt.Println (BR11.Error()); return }
		BQ12  , BR12 := png.DecodeConfig(BN12); if BR12 != nil { fmt.Println (BR12.Error()); return }
		BQ13  , BR13 := png.DecodeConfig(BN13); if BR13 != nil { fmt.Println (BR13.Error()); return }
		BQ14  , BR14 := png.DecodeConfig(BN14); if BR14 != nil { fmt.Println (BR14.Error()); return }
		BQ15  , BR15 := png.DecodeConfig(BN15); if BR15 != nil { fmt.Println (BR15.Error()); return }
		BQ16  , BR16 := png.DecodeConfig(BN16); if BR16 != nil { fmt.Println (BR16.Error()); return }
		BN01.Close()
		BN02.Close()
		BN03.Close()
		BN04.Close()
		BN05.Close()
		BN06.Close()
		BN07.Close()
		BN08.Close()
		BN09.Close()
		BN10.Close()
		BN11.Close()
		BN12.Close()
		BN13.Close()
		BN14.Close()
		BN15.Close()
		BN16.Close()
		BM05["01_d"] = []int {BQ01.Width, BQ01.Height}
		BM05["02_d"] = []int {BQ02.Width, BQ02.Height}
		BM05["03_d"] = []int {BQ03.Width, BQ03.Height}
		BM05["04_d"] = []int {BQ04.Width, BQ04.Height}
		BM05["05_d"] = []int {BQ05.Width, BQ05.Height}
		BM05["06_d"] = []int {BQ06.Width, BQ06.Height}
		BM05["07_d"] = []int {BQ07.Width, BQ07.Height}
		BM05["08_d"] = []int {BQ08.Width, BQ08.Height}
		BM05["09_d"] = []int {BQ09.Width, BQ09.Height}
		BM05["10_d"] = []int {BQ10.Width, BQ10.Height}
		BM05["11_d"] = []int {BQ11.Width, BQ11.Height}
		BM05["12_d"] = []int {BQ12.Width, BQ12.Height}
		BM05["13_d"] = []int {BQ13.Width, BQ13.Height}
		BM05["14_d"] = []int {BQ14.Width, BQ14.Height}
		BM05["15_d"] = []int {BQ15.Width, BQ15.Height}
		BM05["16_d"] = []int {BQ16.Width, BQ16.Height}
		BM05["01_n"] = BE10[ 0]
		BM05["02_n"] = BE10[ 1]
		BM05["03_n"] = BE10[ 2]
		BM05["04_n"] = BE10[ 3]
		BM05["05_n"] = BE10[ 4]
		BM05["06_n"] = BE10[ 5]
		BM05["07_n"] = BE10[ 6]
		BM05["08_n"] = BE10[ 7]
		BM05["09_n"] = BE10[ 8]
		BM05["10_n"] = BE10[ 9]
		BM05["11_n"] = BE10[10]
		BM05["12_n"] = BE10[11]
		BM05["13_n"] = BE10[12]
		BM05["14_n"] = BE10[13]
		BM05["15_n"] = BE10[14]
		BM05["16_n"] = BE10[15]
		BS05  , BS10 := json.MarshalIndent (BM05, "", "    ")
		if BS10 !=  nil {
			CE05 := fmt.Sprintf (`Images mrsh failed [%s]`, BS10.Error ())
			fmt.Println (CE05)
			return
		}
		BS25 := os.WriteFile (BL10, BS05, 0600)
		if BS25 !=  nil {
			CE05 := fmt.Sprintf (`Images sve2 failed [%s]`, BS25.Error ())
			fmt.Println (CE05)
			return
		}
	}
}
