package output

import (
	groshi "github.com/groshi-project/go-groshi"
	"github.com/rodaine/table"
	"time"
)

func Transactions(transactions []*groshi.Transaction, displayUUID bool) error {
	headers := []any{"DATE", "AMOUNT", "CURRENCY", "DESCRIPTION"}
	if displayUUID {
		headers = append(headers, "UUID")
	}

	tbl := table.New(headers...)

	for _, transaction := range transactions {
		cols := []any{
			transaction.Timestamp.In(time.Local).Format(time.DateTime),
			transaction.Amount / 100,
			transaction.Currency,
			transaction.Description,
		}
		if displayUUID {
			cols = append(cols, transaction.UUID)
		}
		tbl.AddRow(cols...)
	}

	tbl.Print()
	return nil
}
