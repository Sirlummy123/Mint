package main

// Classify landscape and portrait images

import "fmt"
import "image/png"
import "os"

func    main () {
	AE05  , AE10 := os.ReadDir ("../xF/p")
	if AE10 !=  nil {
		BE05 := fmt.Sprintf (`Images list failed [%s]`, AE10.Error ())
		fmt.Println (BE05)
		return
	}
	for _ , BE10 := range AE05 {
		fmt.Println (BE10.Name())
		BE11 := "../xF/p/"+BE10.Name()
		BE15  , BE20 := os.Open (BE11)
		if BE20 !=  nil {
			CE05 := fmt.Sprintf (`Image open failed [%s]`, BE20.Error ())
			fmt.Println (CE05)
			return
		}
		BE25  , BE30 := png.DecodeConfig (BE15)
		BE15.Close()
		if BE30 !=  nil {
			CE05 := fmt.Sprintf (`Image dcde failed [%s]`, BE30.Error ())
			fmt.Println (CE05)
			continue
		}
		if BE25.Width > BE25.Height {
			CE05 := os.Rename (BE11, "q/1/"+BE10.Name())
			if CE05 !=  nil {
				DE05 := fmt.Sprintf (`Image mve1 failed [%s]`, CE05.Error ())
				fmt.Println (DE05)
				return
			}
		}  else {
			CE05 := os.Rename (BE11, "q/2/"+BE10.Name())
			if CE05 !=  nil {
				DE05 := fmt.Sprintf (`Image mve2 failed [%s]`, CE05.Error ())
				fmt.Println (DE05)
				return
			}
		}
	}
}
