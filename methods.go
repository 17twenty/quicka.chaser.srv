package main

import "log"

func sendTextMessage(i invoice) {
	if i.ContactNumber == "" {
		log.Println("Couldn't text", i.ContactName, "- no number listed")
	}
	log.Println("Sending a text message to", i.ContactName, "on", i.ContactNumber)
	log.Println("[", chaseSequence[i.CurrentStage].template, "]")
}

func sendEmail(i invoice) {
	if i.ContactNumber == "" {
		log.Println("Couldn't email", i.ContactName, "- no email listed")
	}
	log.Println("Sending an Email to", i.ContactName, "on", i.ContactEmail)
	log.Println("[", chaseSequence[i.CurrentStage].template, "]")
}
