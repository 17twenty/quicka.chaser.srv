package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/17twenty/quicka.my/pkg/quicka"
	"github.com/gofrs/uuid"
)

var (
	invoiceList = []invoice{
		{
			CurrentStage:  0,
			DueDate:       time.Now().Add(20 * time.Second),
			LastContacted: time.Now(),
			InvoiceID:     uuid.Must(uuid.FromString("652B6768-DF5F-4408-84BE-067D432ABAF2")),
			ContactName:   "Nick Glynn",
			ContactNumber: "+61401886425",
			ContactEmail:  "nick@quicka.co",
			Status:        quicka.InvoiceAwaitingTerms,
		},
		{
			CurrentStage:  0,
			DueDate:       time.Now().Add(5 * time.Second),
			LastContacted: time.Now(),
			InvoiceID:     uuid.Must(uuid.FromString("00012345-DF5F-4408-84BE-067D432ABAF2")),
			ContactName:   "Muz Riquelme",
			ContactNumber: "+61401245678",
			ContactEmail:  "muz@quicka.co",
			Status:        quicka.InvoiceAwaitingTerms,
		},
	}
	startTime = time.Now()
)

func printSecondsSince() {
	log.Println(fmt.Sprintf("%.2f", time.Now().Sub(startTime).Minutes()))
}

// isDue tells us that we've received no action and the invoice is now due
// but no terms have been agreed to
func isDue(i *invoice) bool {
	if i.Status == quicka.InvoiceAwaitingTerms && time.Now().After(i.DueDate) {
		return true
	}
	return false
}

// performChase will follow the sequence of chasing
// the invoice until we reach the end of our efforts (45 days)
func performChase(i *invoice) {

	// Work out relative time based on dueDate

	triggerTime := i.DueDate.Add(chaseSequence[i.CurrentStage].triggerTime)

	if time.Now().After(triggerTime) {
		log.Println("Moving to next stage")

		switch chaseSequence[i.CurrentStage].method {
		case chaseMethodEmail:
			sendEmail(*i)
		case chaseMethodTextMessage:
			sendTextMessage(*i)
		}
		// Progress to next step and fail if NFA
		i.CurrentStage++
		if i.CurrentStage >= len(chaseSequence) {
			log.Println("Failed to collect invoice", i.InvoiceID)
			i.Status = quicka.InvoiceFailedToCollect
			return
		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	log.Println("Starting Quicka Chaser")
	// Declare flags before parse
	flag.Parse()

	// Nevar exit, nevar forget!
	for {
		time.Sleep(time.Second)
		// Still waiting
		for i := range invoiceList {
			// Look for
			if isDue(&invoiceList[i]) {
				performChase(&invoiceList[i])
			}
		}

		// TODO: Same logic for invoices needing payment
		// log.Println("Looking for failed to collect invoices")
		// A failure to collect was marked in the system
		// An instalment/invoice will need to be rescheduled

	}
}
