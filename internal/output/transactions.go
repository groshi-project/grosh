package output

import (
	"errors"
	"github.com/rodaine/table"
	"time"
)

func PrintTransactions(transactions []map[string]any) error {
	tbl := table.New("DATE", "AMOUNT", "CURRENCY", "DESCRIPTION", "UUID")
	for _, transaction := range transactions {
		transactionDateString, ok := transaction["date"].(string)
		if !ok {
			return errors.New("(bug) invalid type of transaction date field")
		}
		transactionDate, err := time.Parse(time.RFC3339, transactionDateString)
		if err != nil {
			return err
		}

		tbl.AddRow(
			transactionDate.Format(time.DateTime),
			transaction["amount"],
			transaction["currency"],
			transaction["description"],
			transaction["uuid"],
		)
	}
	tbl.Print()

	return nil
}
