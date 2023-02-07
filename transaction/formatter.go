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

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName

	UserTransactionFormatter := UserTransactionFormatter{}
	UserTransactionFormatter.ID = transaction.ID
	UserTransactionFormatter.Amount = transaction.Amount
	UserTransactionFormatter.Status = transaction.Status
	UserTransactionFormatter.CreatedAt = transaction.CreatedAt
	UserTransactionFormatter.Campaign = campaignFormatter

	return UserTransactionFormatter
}

func FormatUsersTransaction(transactions []Transaction) []UserTransactionFormatter {
	var formatUsersTransaction []UserTransactionFormatter
	for _, transaction := range transactions {
		formatUserTransaction := FormatUserTransaction(transaction)

		formatUsersTransaction = append(formatUsersTransaction, formatUserTransaction)
	}

	return formatUsersTransaction
}
