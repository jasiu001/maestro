package spreadsheets

import (
	"log"

	. "github.com/jasiu001/maestro/list"
	"google.golang.org/api/sheets/v4"
)

func CreateData() []RowData {

	client := GetClient()
	config := InitConfig()

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(config.Spreadsheet.Id, config.Spreadsheet.Scope).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	var data []RowData
	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			data = append(data, InitData(row))
		}
	} else {
		log.Fatalln("No data found.")
	}

	return data
}
