package output

import (
	go_groshi "github.com/groshi-project/go-groshi"
	"github.com/rodaine/table"
	"time"
)

func PrintTransactions(transactions []*go_groshi.Transaction) error {
	tbl := table.New("DATE", "AMOUNT", "CURRENCY", "DESCRIPTION", "UUID")
	for _, transaction := range transactions {
		tbl.AddRow(
			transaction.Timestamp.Format(time.DateTime),
			transaction.Amount/100,
			transaction.Currency,
			transaction.Description,
			transaction.UUID,
		)
	}
	tbl.Print()

	return nil
}
