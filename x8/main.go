package main

// Disassemble processed images

import "encoding/json"
import "fmt"
import "image"
import "image/jpeg"
import "os"
import "slices"
import "strings"

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}
func    main () {
	AE05  , AE10 := os.ReadDir ("j")
	if AE10 !=  nil {
		BE05 := fmt.Sprintf (`Dir read failed [%s]`, AE10.Error ())
		fmt.Println (BE05)
		return
	}
	AE15 := []string {}
	for _ , BE10 := range AE05 { AE15 = append (AE15, BE10.Name()) }
	slices.Sort(AE15)
	for _ , BE10 := range AE15 {
		fmt.Println ( BE10 )
		BE15 := fmt.Sprintf (`i/%s`, BE10); BE15 = strings.ReplaceAll (BE15, ".jpg", ".json")
		BE25  , BE30 := os.ReadFile (BE15)
		if BE30 !=  nil && os.IsNotExist(BE30) == false {
			CE05 := fmt.Sprintf (`File read failed [%s]`, BE30.Error ())
			fmt.Println (CE05)
			return
		}
		if BE30 !=  nil && os.IsNotExist(BE30) {
			fmt.Println ("skipping", BE15)
			continue
		}
		BE35 := map[string]any{}
		BE36 := &DB{}
		BE40 := json.Unmarshal(BE25,&BE35)
		if BE40 !=  nil {
			CE05 := fmt.Sprintf (`Data unmarshal failed 1 [%s]`, BE40.Error ())
			fmt.Println (CE05)
			return
		}
		BE42 := json.Unmarshal(BE25,&BE36)
		if BE42 !=  nil {
			CE05 := fmt.Sprintf (`Data unmarshal failed 2 [%s]`, BE42.Error ())
			fmt.Println (CE05)
			return
		}
//		fmt.Println (string(BE25), BE36)
		BE45 := fmt.Sprintf (`j/%s`, BE10)
		BF01  , BG01 := os.Open    (BE45); if BG01 != nil { fmt.Println (BG01.Error ()); return }
		BF02  , BG02 := os.Open    (BE45); if BG02 != nil { fmt.Println (BG02.Error ()); return }
		BF03  , BG03 := os.Open    (BE45); if BG03 != nil { fmt.Println (BG03.Error ()); return }
		BF04  , BG04 := os.Open    (BE45); if BG04 != nil { fmt.Println (BG04.Error ()); return }
		BF05  , BG05 := os.Open    (BE45); if BG05 != nil { fmt.Println (BG05.Error ()); return }
		BF06  , BG06 := os.Open    (BE45); if BG06 != nil { fmt.Println (BG06.Error ()); return }
		BF07  , BG07 := os.Open    (BE45); if BG07 != nil { fmt.Println (BG07.Error ()); return }
		BF08  , BG08 := os.Open    (BE45); if BG08 != nil { fmt.Println (BG08.Error ()); return }
		BF09  , BG09 := os.Open    (BE45); if BG09 != nil { fmt.Println (BG09.Error ()); return }
		BF10  , BG10 := os.Open    (BE45); if BG10 != nil { fmt.Println (BG10.Error ()); return }
		BF11  , BG11 := os.Open    (BE45); if BG11 != nil { fmt.Println (BG11.Error ()); return }
		BF12  , BG12 := os.Open    (BE45); if BG12 != nil { fmt.Println (BG12.Error ()); return }
		BF13  , BG13 := os.Open    (BE45); if BG13 != nil { fmt.Println (BG13.Error ()); return }
		BF14  , BG14 := os.Open    (BE45); if BG14 != nil { fmt.Println (BG14.Error ()); return }
		BF15  , BG15 := os.Open    (BE45); if BG15 != nil { fmt.Println (BG15.Error ()); return }
		BF16  , BG16 := os.Open    (BE45); if BG16 != nil { fmt.Println (BG16.Error ()); return }
		BH01  , BI01 := jpeg.Decode(BF01); if BI01 != nil { fmt.Println (BI01.Error ()); return }
		BH02  , BI02 := jpeg.Decode(BF02); if BI02 != nil { fmt.Println (BI02.Error ()); return }
		BH03  , BI03 := jpeg.Decode(BF03); if BI03 != nil { fmt.Println (BI03.Error ()); return }
		BH04  , BI04 := jpeg.Decode(BF04); if BI04 != nil { fmt.Println (BI04.Error ()); return }
		BH05  , BI05 := jpeg.Decode(BF05); if BI05 != nil { fmt.Println (BI05.Error ()); return }
		BH06  , BI06 := jpeg.Decode(BF06); if BI06 != nil { fmt.Println (BI06.Error ()); return }
		BH07  , BI07 := jpeg.Decode(BF07); if BI07 != nil { fmt.Println (BI07.Error ()); return }
		BH08  , BI08 := jpeg.Decode(BF08); if BI08 != nil { fmt.Println (BI08.Error ()); return }
		BH09  , BI09 := jpeg.Decode(BF09); if BI09 != nil { fmt.Println (BI09.Error ()); return }
		BH10  , BI10 := jpeg.Decode(BF10); if BI10 != nil { fmt.Println (BI10.Error ()); return }
		BH11  , BI11 := jpeg.Decode(BF11); if BI11 != nil { fmt.Println (BI11.Error ()); return }
		BH12  , BI12 := jpeg.Decode(BF12); if BI12 != nil { fmt.Println (BI12.Error ()); return }
		BH13  , BI13 := jpeg.Decode(BF13); if BI13 != nil { fmt.Println (BI13.Error ()); return }
		BH14  , BI14 := jpeg.Decode(BF14); if BI14 != nil { fmt.Println (BI14.Error ()); return }
		BH15  , BI15 := jpeg.Decode(BF15); if BI15 != nil { fmt.Println (BI15.Error ()); return }
		BH16  , BI16 := jpeg.Decode(BF16); if BI16 != nil { fmt.Println (BI16.Error ()); return }
		BF01.Close()
		BF02.Close()
		BF03.Close()
		BF04.Close()
		BF05.Close()
		BF06.Close()
		BF07.Close()
		BF08.Close()
		BF09.Close()
		BF10.Close()
		BF11.Close()
		BF12.Close()
		BF13.Close()
		BF14.Close()
		BF15.Close()
		BF16.Close()
		BJ01 := BH01.(SubImager)
		BJ02 := BH02.(SubImager)
		BJ03 := BH03.(SubImager)
		BJ04 := BH04.(SubImager)
		BJ05 := BH05.(SubImager)
		BJ06 := BH06.(SubImager)
		BJ07 := BH07.(SubImager)
		BJ08 := BH08.(SubImager)
		BJ09 := BH09.(SubImager)
		BJ10 := BH10.(SubImager)
		BJ11 := BH11.(SubImager)
		BJ12 := BH12.(SubImager)
		BJ13 := BH13.(SubImager)
		BJ14 := BH14.(SubImager)
		BJ15 := BH15.(SubImager)
		BJ16 := BH16.(SubImager)
		BK01 := BJ01.SubImage(image.Rect(   0,   0,1500,   0+BE36.I_01_d[1]))
		BK02 := BJ02.SubImage(image.Rect(1500,   0,3000,   0+BE36.I_02_d[1]))
		BK03 := BJ03.SubImage(image.Rect(3000,   0,4500,   0+BE36.I_03_d[1]))
		BK04 := BJ04.SubImage(image.Rect(4500,   0,6000,   0+BE36.I_04_d[1]))
		BK05 := BJ05.SubImage(image.Rect(   0,1500,1500,1500+BE36.I_05_d[1]))
		BK06 := BJ06.SubImage(image.Rect(1500,1500,3000,1500+BE36.I_06_d[1]))
		BK07 := BJ07.SubImage(image.Rect(3000,1500,4500,1500+BE36.I_07_d[1]))
		BK08 := BJ08.SubImage(image.Rect(4500,1500,6000,1500+BE36.I_08_d[1]))
		BK09 := BJ09.SubImage(image.Rect(   0,3000,1500,3000+BE36.I_09_d[1]))
		BK10 := BJ10.SubImage(image.Rect(1500,3000,3000,3000+BE36.I_10_d[1]))
		BK11 := BJ11.SubImage(image.Rect(3000,3000,4500,3000+BE36.I_11_d[1]))
		BK12 := BJ12.SubImage(image.Rect(4500,3000,6000,3000+BE36.I_12_d[1]))
		BK13 := BJ13.SubImage(image.Rect(   0,4500,1500,4500+BE36.I_13_d[1]))
		BK14 := BJ14.SubImage(image.Rect(1500,4500,3000,4500+BE36.I_14_d[1]))
		BK15 := BJ15.SubImage(image.Rect(3000,4500,4500,4500+BE36.I_15_d[1]))
		BK16 := BJ16.SubImage(image.Rect(4500,4500,6000,4500+BE36.I_16_d[1]))
		BK51 := "o/"+BE35["01_n"].(string); BK51=strings.ReplaceAll(BK51,".png",".jpg")
		BK52 := "o/"+BE35["02_n"].(string); BK52=strings.ReplaceAll(BK52,".png",".jpg")
		BK53 := "o/"+BE35["03_n"].(string); BK53=strings.ReplaceAll(BK53,".png",".jpg")
		BK54 := "o/"+BE35["04_n"].(string); BK54=strings.ReplaceAll(BK54,".png",".jpg")
		BK55 := "o/"+BE35["05_n"].(string); BK55=strings.ReplaceAll(BK55,".png",".jpg")
		BK56 := "o/"+BE35["06_n"].(string); BK56=strings.ReplaceAll(BK56,".png",".jpg")
		BK57 := "o/"+BE35["07_n"].(string); BK57=strings.ReplaceAll(BK57,".png",".jpg")
		BK58 := "o/"+BE35["08_n"].(string); BK58=strings.ReplaceAll(BK58,".png",".jpg")
		BK59 := "o/"+BE35["09_n"].(string); BK59=strings.ReplaceAll(BK59,".png",".jpg")
		BK60 := "o/"+BE35["10_n"].(string); BK60=strings.ReplaceAll(BK60,".png",".jpg")
		BK61 := "o/"+BE35["11_n"].(string); BK61=strings.ReplaceAll(BK61,".png",".jpg")
		BK62 := "o/"+BE35["12_n"].(string); BK62=strings.ReplaceAll(BK62,".png",".jpg")
		BK63 := "o/"+BE35["13_n"].(string); BK63=strings.ReplaceAll(BK63,".png",".jpg")
		BK64 := "o/"+BE35["14_n"].(string); BK64=strings.ReplaceAll(BK64,".png",".jpg")
		BK65 := "o/"+BE35["15_n"].(string); BK65=strings.ReplaceAll(BK65,".png",".jpg")
		BK66 := "o/"+BE35["16_n"].(string); BK66=strings.ReplaceAll(BK66,".png",".jpg")
		BL01  , BM01 := os.OpenFile (BK51, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM01 != nil { fmt.Println (BM01.Error ()); return }
		BL02  , BM02 := os.OpenFile (BK52, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM02 != nil { fmt.Println (BM02.Error ()); return }
		BL03  , BM03 := os.OpenFile (BK53, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM03 != nil { fmt.Println (BM03.Error ()); return }
		BL04  , BM04 := os.OpenFile (BK54, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM04 != nil { fmt.Println (BM04.Error ()); return }
		BL05  , BM05 := os.OpenFile (BK55, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM05 != nil { fmt.Println (BM05.Error ()); return }
		BL06  , BM06 := os.OpenFile (BK56, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM06 != nil { fmt.Println (BM06.Error ()); return }
		BL07  , BM07 := os.OpenFile (BK57, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM07 != nil { fmt.Println (BM07.Error ()); return }
		BL08  , BM08 := os.OpenFile (BK58, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM08 != nil { fmt.Println (BM08.Error ()); return }
		BL09  , BM09 := os.OpenFile (BK59, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM09 != nil { fmt.Println (BM09.Error ()); return }
		BL10  , BM10 := os.OpenFile (BK60, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM10 != nil { fmt.Println (BM10.Error ()); return }
		BL11  , BM11 := os.OpenFile (BK61, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM11 != nil { fmt.Println (BM11.Error ()); return }
		BL12  , BM12 := os.OpenFile (BK62, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM12 != nil { fmt.Println (BM12.Error ()); return }
		BL13  , BM13 := os.OpenFile (BK63, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM13 != nil { fmt.Println (BM13.Error ()); return }
		BL14  , BM14 := os.OpenFile (BK64, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM14 != nil { fmt.Println (BM14.Error ()); return }
		BL15  , BM15 := os.OpenFile (BK65, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM15 != nil { fmt.Println (BM15.Error ()); return }
		BL16  , BM16 := os.OpenFile (BK66, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600); if BM16 != nil { fmt.Println (BM16.Error ()); return }
		BN01 := jpeg.Encode (BL01, BK01, nil); if BN01 != nil { fmt.Println (BN01.Error ()); return }
		BN02 := jpeg.Encode (BL02, BK02, nil); if BN02 != nil { fmt.Println (BN02.Error ()); return }
		BN03 := jpeg.Encode (BL03, BK03, nil); if BN03 != nil { fmt.Println (BN03.Error ()); return }
		BN04 := jpeg.Encode (BL04, BK04, nil); if BN04 != nil { fmt.Println (BN04.Error ()); return }
		BN05 := jpeg.Encode (BL05, BK05, nil); if BN05 != nil { fmt.Println (BN05.Error ()); return }
		BN06 := jpeg.Encode (BL06, BK06, nil); if BN06 != nil { fmt.Println (BN06.Error ()); return }
		BN07 := jpeg.Encode (BL07, BK07, nil); if BN07 != nil { fmt.Println (BN07.Error ()); return }
		BN08 := jpeg.Encode (BL08, BK08, nil); if BN08 != nil { fmt.Println (BN08.Error ()); return }
		BN09 := jpeg.Encode (BL09, BK09, nil); if BN09 != nil { fmt.Println (BN09.Error ()); return }
		BN10 := jpeg.Encode (BL10, BK10, nil); if BN10 != nil { fmt.Println (BN10.Error ()); return }
		BN11 := jpeg.Encode (BL11, BK11, nil); if BN11 != nil { fmt.Println (BN11.Error ()); return }
		BN12 := jpeg.Encode (BL12, BK12, nil); if BN12 != nil { fmt.Println (BN12.Error ()); return }
		BN13 := jpeg.Encode (BL13, BK13, nil); if BN13 != nil { fmt.Println (BN13.Error ()); return }
		BN14 := jpeg.Encode (BL14, BK14, nil); if BN14 != nil { fmt.Println (BN14.Error ()); return }
		BN15 := jpeg.Encode (BL15, BK15, nil); if BN15 != nil { fmt.Println (BN15.Error ()); return }
		BN16 := jpeg.Encode (BL16, BK16, nil); if BN16 != nil { fmt.Println (BN16.Error ()); return }
		BL01.Close ()
		BL02.Close ()
		BL03.Close ()
		BL04.Close ()
		BL05.Close ()
		BL06.Close ()
		BL07.Close ()
		BL08.Close ()
		BL09.Close ()
		BL10.Close ()
		BL11.Close ()
		BL12.Close ()
		BL13.Close ()
		BL14.Close ()
		BL15.Close ()
		BL16.Close ()
	}
}
type    DB struct {
	I_01_d [2]int `json:"01_d"`
	I_01_n string `json:"01_n"`
	I_02_d [2]int `json:"02_d"`
	I_02_n string `json:"02_n"`
	I_03_d [2]int `json:"03_d"`
	I_03_n string `json:"03_n"`
	I_04_d [2]int `json:"04_d"`
	I_04_n string `json:"04_n"`
	I_05_d [2]int `json:"05_d"`
	I_05_n string `json:"05_n"`
	I_06_d [2]int `json:"06_d"`
	I_06_n string `json:"06_n"`
	I_07_d [2]int `json:"07_d"`
	I_07_n string `json:"07_n"`
	I_08_d [2]int `json:"08_d"`
	I_08_n string `json:"08_n"`
	I_09_d [2]int `json:"09_d"`
	I_09_n string `json:"09_n"`
	I_10_d [2]int `json:"10_d"`
	I_10_n string `json:"10_n"`
	I_11_d [2]int `json:"11_d"`
	I_11_n string `json:"11_n"`
	I_12_d [2]int `json:"12_d"`
	I_12_n string `json:"12_n"`
	I_13_d [2]int `json:"13_d"`
	I_13_n string `json:"13_n"`
	I_14_d [2]int `json:"14_d"`
	I_14_n string `json:"14_n"`
	I_15_d [2]int `json:"15_d"`
	I_15_n string `json:"15_n"`
	I_16_d [2]int `json:"16_d"`
	I_16_n string `json:"16_n"`
}
