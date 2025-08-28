package createtransaction

import (
	"github.com/WagnerReis/fc-ms-wallet/internal/entity"
	"github.com/WagnerReis/fc-ms-wallet/internal/gateway"
	"github.com/WagnerReis/fc-ms-wallet/pkg/events"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type createtransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
	EventDispatcher    *events.EventDispatcher
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
	eventDispatcher events.EventDispatcher,
	transactionCreated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
		EventDispatcher:    &eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*createtransactionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}
	accountTo, err := uc.AccountGateway.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}
	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	output := &createtransactionOutputDTO{
		ID: transaction.ID,
	}

	uc.TransactionCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	return output, nil
}
