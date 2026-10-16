package main

// Create listing list

import "fmt"
import E "github.com/xuri/excelize/v2"
import "os"
import "slices"
import "sort"
import "strings"
import "time"

func    main () {
	fmt.Println ("Phase 1 in progress...")
	AE05  , AE10 := os.ReadDir ("../xB/data")
	if AE10 !=  nil {
		BE05 := fmt.Sprintf ("Listing read failed [%s]", AE10.Error ())
		fmt.Println (BE05)
		return
	}
	fmt.Println ("Phase 2 in progress...")
	AE15 := []string{ }
	for _ , BE10 := range AE05{	
		BE15 := strings.Split (BE10.Name(), "-")
		if slices.Contains(Shortlist,BE15[0]) != true{ continue }
		if BE15[2] != "03.json" { continue }
		AE15 = append(AE15, BE10.Name ())
	}
	fmt.Println ("Phase 3 in progress...")
	sort.Strings (AE15)
	AE25 := map[string][]string {}
	for _ , BE10 := range AE15 {
		BE15 := strings.Split (BE10, "-")
		BE20 := AE25[BE15[0]]
		if BE20 ==  nil { BE20 = []string {} }
		if len(BE20) >= 550 { continue }
		BE20 = append (BE20, BE10)
		AE25[BE15[0]] = BE20
	}
	AE35 := E.NewFile ()
	AE40 := fmt.Sprintf ("%d.xlsx", time.Now().Unix())
	defer AE35.SaveAs (AE40)
	AE45 := 0;
	fmt.Println ("Phase 4 in progress...")
	for BE05 , BE10 := range AE25 {
		for _ , CE10 := range BE10 {
			AE45  = AE45 + 1;
			CE15 := fmt.Sprintf ("B%d", AE45)
			CE20 := fmt.Sprintf ("C%d", AE45)
			CF05 := AE35.SetCellValue ("Sheet1", CE15, BE05)
			CF10 := AE35.SetCellValue ("Sheet1", CE20, CE10)
			if CF05 !=  nil {
				DF05 := fmt.Sprintf ("Data set failed [c%d: %s]", AE45, CF05.Error ())
				fmt.Println (DF05)
				continue
			}
			if CF10 !=  nil {
				DF05 := fmt.Sprintf ("Data set failed [c%d: %s]", AE45, CF10.Error ())
				fmt.Println (DF05)
				continue
			}
		}
	}
}
var Shortlist = []string{
        "000172",
        "000250",
        "000303",
        "000015",
        "000019",
        "000168",
        "000280",
        "000302",
        "000312",
        "000067",
        "000022",
        "000263",
        "000279",
        "000301",
        "000382",
        "000016",
        "000061",
        "000256",
        "000257",
        "000274",
        "000386",
        "000014",
        "000021",
        "000037",
        "000056",
        "000025",
        "000065",
        "000044",
        "000051",
        "000046",
        "000033",
        "000270",
        "000288",
        "000289",
        "000042",
        "000053",
        "000045",
        "000057",
        "000017",
        "000253",
        "000278",
        "000285",
        "000397",
        "000348",
        "000349",
        "000055",
        "000398",
        "000399",
        "000254",
        "000258",
        "000287",
        "000310",
        "000043",
        "000252",
        "000387",
        "000030",
        "000152",
        "000389",
        "000395",
        "000012",
        "000249",
        "000020",
        "000026",
        "000036",
        "000385",
        "000060",
        "000165",
        "000011",
        "000393",
        "000294",
        "000074",
        "000291",
        "000169",
        "000255",
        "000031",
        "000040",
        "000281",
        "000296",
        "000300",
        "000304",
        "000299",
        "000388",
        "000018",
        "000029",
        "000173",
        "000284",
        "000259",
        "000297",
        "000383",
        "000384",
        "000035",
        "000038",
        "000062",
}

