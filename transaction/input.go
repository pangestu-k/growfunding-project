package transaction

import "growfunding/user"

type GetTransactionCampaignInput struct {
	ID   int       `uri:"id" binding:"required"`
	User user.User `json:"user"`
}
