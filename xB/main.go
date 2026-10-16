package main

// Create clones of data with phone no

import "fmt"
import "github.com/tidwall/gjson"
import "os"

func    main () {
	AE01 := `/home/qbrian/Local/Project/Mint0/20: Listing Data Harvester/o/listing`
	AE05  , AE10 := os.ReadDir (AE01)
	if AE10 !=  nil {
		BE05 := fmt.Sprintf (`Data listing failed [%s]`, AE10.Error ())
		fmt.Println (BE05)
		return
	}
	for _ , BE10 := range AE05 {
		BE15  , BE20 := os.ReadFile (AE01 + "/" + BE10.Name())
		if BE20 !=  nil {
			CE05 := fmt.Sprintf (`Data listing failed [%s]`, BE20.Error ())
			fmt.Println (CE05)
			return
		}
		BE25 := string(BE15)
		BE30 := gjson.Get (BE25, "seller.phone").String ()
		if BE30 ==  ""  { continue; }
		if len(BE30) != 11 { continue; }
		fmt.Println(BE30)
		BE35 := os.WriteFile ("/home/qbrian/Local/Project/Mint1/xB/data/"+BE10.Name(), BE15, 0600)
		if BE35 !=  nil {
			CE05 := fmt.Sprintf (`Data storing failed [%s]`, BE35.Error ())
			fmt.Println (CE05)
			return
		}
	}
}
