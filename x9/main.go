package main

// Upload image to mint

import "bytes"
import "io"
import "fmt"
import "github.com/tidwall/gjson"
import E "github.com/xuri/excelize/v2"
import "os"
import "mime/multipart"
import "net/http"
import "path/filepath"

func    main () {
	AE05 := "i.xlsx"
	AE31  , AE32 := E.OpenFile (AE05)
	if AE32 !=  nil {
		BE05 := fmt.Sprintf(`Data 'i.xlsx' read failed [%s]`, AE32.Error ())
		fmt.Println (BE05)
		return
	}
	fmt.Println ("Running...")
	for BE05 := 1; BE05 <= 48594; BE05++{
		BE15 := fmt.Sprintf ("M%d", BE05)
		BE20 := fmt.Sprintf ("N%d", BE05)
		BF01  , BG01 := AE31.GetCellValue ("Sheet1", BE15); if BG01 != nil { fmt.Println (BG01); return }
		BF02  , BG02 := AE31.GetCellValue ("Sheet1", BE20); if BG02 != nil { fmt.Println (BG02); return }
		if BF01 == "" { continue }
		if BF02 == "" { continue }
		BG55 := fmt.Sprintf ("../x8/o/%s", BF01)
		_, BG70 := os.Stat (BG55)
		if BG70 !=  nil && os.IsNotExist(BG70) == false {
			CE05 := fmt.Sprintf (`Image existence check failed [%s]`, BG70.Error ())
			fmt.Println (CE05)
			return
		}
		if BG70 !=  nil { continue }
		fmt.Println (fmt.Sprintf ("Processing: %d" , BE05))
		BH05  , BH10 := uploadImage ("https://prodapi.mintclassifieds.com/admin/uploadMultipleImages", BG55)
		if BH05 !=  nil {
			CE05 := fmt.Sprintf (`Image upload failed [%s]`, BH05.Error ())
			fmt.Println (CE05)
			return
		}
		BH15 := gjson.Get (BH10, "data.0.image").String ()
		BH25 := AE31.SetCellValue ("Sheet1", BE15, BH15); if BH25 != nil { fmt.Println (BH25); return }
		AF32 := AE31.SaveAs ("i.xlsx")
		if AF32 !=  nil {
			BE05 := fmt.Sprintf(`Data 'o.xlsx' read failed [%s]`, AF32.Error ())
			fmt.Println (BE05)
			return
		}
	}
}
func uploadImage(url, filePath string) (error, string) {
	// 1. Open the local image file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err), ""
	}
	defer file.Close()

	// 2. Create a buffer to hold the multipart request body
	body := &bytes.Buffer{}

	// 3. Initialize a new multipart writer bound to the buffer
	writer := multipart.NewWriter(body)

	// 4. Create the form file part with the field name "image"
	part, err := writer.CreateFormFile("image", filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("could not create form file: %w", err), ""
	}

	// 5. Copy the image file content directly into the form part
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("could not copy file content: %w", err), ""
	}

	// Optional: Add extra text fields to the form if needed
	// _ = writer.WriteField("description", "My profile picture")

	// 6. Close the writer to finalize boundaries before sending
	err = writer.Close()
	if err != nil {
		return fmt.Errorf("could not close multipart writer: %w", err), ""
	}

	// 7. Create the HTTP POST request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err), ""
	}

	// CRITICAL: Set the content type header including the boundary string
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjZhMGQ2ZTczNmNkOTcyMmNkZTlmNDA5ZCIsImlhdCI6MTc4NTExMjAwOH0.i9EP0nTcHodjj1S4ZOB5eAA8wtQcVqzxrxG_7zjSasE")

	// 8. Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err), ""
	}
	defer resp.Body.Close()

	// 9. Process the server response
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server returned status %s: %s", resp.Status, string(respBody)), ""
	}
	respBody, _ := io.ReadAll(resp.Body)

	return nil, string(respBody)
}
