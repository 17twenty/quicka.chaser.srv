package main

import (
	"time"

	"github.com/17twenty/quicka.my/pkg/quicka"
	"github.com/gofrs/uuid"
)

type invoice struct {
	CurrentStage  int
	DueDate       time.Time
	LastContacted time.Time
	InvoiceID     uuid.UUID
	ContactName   string
	ContactNumber string
	ContactEmail  string
	Status        quicka.InvoiceStatus
}

var (
	inThreeMinutes = time.Now().Add(time.Minute * 3)
	inTwoMinute    = time.Now().Add(time.Minute * 2)
	inOneMinute    = time.Now().Add(time.Minute * 1)
)
