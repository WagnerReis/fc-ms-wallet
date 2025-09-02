package create_account

import (
	"github.com/WagnerReis/fc-ms-wallet/internal/gateway"
)

type AddCreditDTO struct {
	ClientID  string  `json:"client_id"`
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type AddCreditUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewAddCreditUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *AddCreditUseCase {
	return &AddCreditUseCase{
		AccountGateway: a,
		ClientGateway:  c,
	}
}

func (uc *AddCreditUseCase) Execute(input AddCreditDTO) error {
	_, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return err
	}
	account, err := uc.AccountGateway.FindByID(input.AccountID)
	if err != nil {
		return err
	}

	account.Balance = input.Amount

	err = uc.AccountGateway.UpdateBalance(account)
	if err != nil {
		return err
	}
	return nil
}
