package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/WagnerReis/fc-ms-wallet/internal/database"
	"github.com/WagnerReis/fc-ms-wallet/internal/event"
	"github.com/WagnerReis/fc-ms-wallet/internal/event/handler"
	"github.com/WagnerReis/fc-ms-wallet/internal/usecase/create_account"
	"github.com/WagnerReis/fc-ms-wallet/internal/usecase/create_client"
	"github.com/WagnerReis/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com/WagnerReis/fc-ms-wallet/internal/web"
	"github.com/WagnerReis/fc-ms-wallet/internal/web/webserver"
	"github.com/WagnerReis/fc-ms-wallet/pkg/events"
	"github.com/WagnerReis/fc-ms-wallet/pkg/kafka"
	"github.com/WagnerReis/fc-ms-wallet/pkg/uow"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))

	balanceUpdatedEvent := event.NewBalanceUpdated()
	eventDispatcher.Register("BalanceUpdated", handler.NewBalanceUpdatedKafkaHandler(kafkaProducer))

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	addCreditUseCase := create_account.NewAddCreditUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, *eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase, *addCreditUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/accounts/credit", accountHandler.AddCredit)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}
