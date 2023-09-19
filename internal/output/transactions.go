package output

import (
	go_groshi "github.com/groshi-project/go-groshi"
	"github.com/rodaine/table"
	"time"
)

func PrintTransactions(transactions []*go_groshi.Transaction) error {
	tbl := table.New("UUID", "DATE", "AMOUNT", "CURRENCY", "DESCRIPTION")
	for _, transaction := range transactions {
		tbl.AddRow(
			transaction.UUID,
			transaction.Timestamp.In(time.Local).Format(time.DateTime),
			transaction.Amount/100,
			transaction.Currency,
			transaction.Description,
		)
	}
	tbl.Print()

	return nil
}
