package transaction

import "time"

type TransactionCampaignFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func TransactionFormatter(transaction Transaction) TransactionCampaignFormatter {
	transactionFormatter := TransactionCampaignFormatter{}
	transactionFormatter.ID = transaction.ID
	transactionFormatter.Name = transaction.User.Name
	transactionFormatter.Amount = transaction.Amount
	transactionFormatter.CreatedAt = transaction.CreatedAt

	return transactionFormatter
}

func TransactionsFormatter(transactions []Transaction) []TransactionCampaignFormatter {
	var transactionsFormatter []TransactionCampaignFormatter

	for _, transaction := range transactions {
		transactionFormatter := TransactionFormatter(transaction)

		transactionsFormatter = append(transactionsFormatter, transactionFormatter)
	}

	return transactionsFormatter
}
