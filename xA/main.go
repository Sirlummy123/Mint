package main

// Extract user's with phone no

import "encoding/json"
import "fmt"
import "github.com/tidwall/gjson"
import "os"
import "regexp"

func    main () {
	AE05 := `/home/qbrian/Local/Project/Mint0/20: Listing Data Harvester/o/listing`
	AE15  , AE20 := os.ReadDir (AE05)
	if AE20 !=  nil {
		BE05 := fmt.Sprintf (`Dir read failed [%s]`, AE20.Error ())
		fmt.Println (BE05)
		return
	}
	AE25 := map[string]string {}
	for _ , BE10 := range AE15 {
		if regexp.MustCompile(`\-03\.json$`).MatchString(BE10.Name()) != true { continue }
		BE15 := fmt.Sprintf (`%s/%s`, AE05, BE10.Name())
		BE25  , BE30 := os.ReadFile (BE15)
		if BE30 !=  nil {
			CE05 := fmt.Sprintf (`Entry read failed [%s]`, BE30.Error ())
			fmt.Println (CE05)
			return
		}
		BE35 := string (BE25)
		BE41 := gjson.Get (BE35, "seller.phone").String ()
		BE42 := gjson.Get (BE35, "seller.page_url").String ()
		if BE41 == "" { continue }
		if AE25[BE41] != "" { continue }
		if regexp.MustCompile(`^https`).MatchString(BE42) == false { BE42 = "https://jiji.ng/"+BE42 }
		AE25[BE41] = BE42
		fmt.Println (BE41)
	}
	AF05  , AF10 := json.MarshalIndent (AE25, "", "    ")
	if AF10 !=  nil {
		BE05 := fmt.Sprintf (`JSON marshal failed [%s]`, AF10.Error ())
		fmt.Println (BE05)
		return
	}
	AF15 := os.WriteFile ("data.xt", AF05, 0600)
	if AF15 !=  nil {
		BE05 := fmt.Sprintf (`JSON write failed [%s]`, AF15.Error ())
		fmt.Println (BE05)
		return
	}
}
