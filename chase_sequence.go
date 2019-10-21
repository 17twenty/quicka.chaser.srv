package main

import (
	"time"
)

type chaseMethod int64

const (
	chaseMethodUnknown chaseMethod = iota
	chaseMethodEmail
	chaseMethodTextMessage
)

const interval = time.Second

// const interval = time.Hour * 24

var (
	chaseSequence = []chaseStage{
		{
			time.Duration(1 * interval), // Add a day (day 2)
			"This is your day 2 chase message",
			false,
			chaseMethodTextMessage,
		},
		{
			time.Duration(2 * interval),
			"This is your day 3 chase message",
			false,
			chaseMethodEmail,
		},
		{
			time.Duration(3 * interval),
			"This is your day 4 chase message",
			false,
			chaseMethodTextMessage,
		},
		{
			time.Duration(4 * interval),
			"This is your day 5 chase message",
			false,
			chaseMethodEmail,
		},
		{
			time.Duration(6 * interval),
			"This is your day 7 chase message",
			false,
			chaseMethodTextMessage,
		},
		{
			time.Duration(8 * interval),
			"This is your day 9 chase message",
			false,
			chaseMethodEmail,
		},
		{
			time.Duration(10 * interval),
			"This is your day 11 chase message",
			false,
			chaseMethodEmail,
		},
		{
			time.Duration(12 * interval),
			"This is your day 13 chase message",
			false,
			chaseMethodTextMessage,
		},
		{
			time.Duration(16 * interval),
			"This is your day 17 chase message",
			false,
			chaseMethodEmail,
		},
		{
			time.Duration(18 * interval),
			"This is your day 19 chase message",
			false,
			chaseMethodEmail,
		},
		{
			time.Duration(23 * interval),
			"This is your day 24 chase message",
			false,
			chaseMethodTextMessage,
		},
		{
			time.Duration(28 * interval),
			"This is your day 29 chase message",
			false,
			chaseMethodEmail,
		},
		{
			time.Duration(33 * interval),
			"This is your day 34 chase message",
			false,
			chaseMethodEmail,
		},
		{
			time.Duration(38 * interval),
			"This is your day 39 chase message",
			false,
			chaseMethodEmail,
		},
		// End of chase
	}
)

// chaseStage is initialised as an array of stages that are folllowed
// iff an invoice has languished in a state for too long
type chaseStage struct {
	triggerTime    time.Duration // duration as it's relative for each dueDate
	template       string        // TODO: Replace this with a load of the actual email text/sms text
	notifyInvoicer bool          // Do we want to keep our customers aware of updates?
	method         chaseMethod
}

func (x chaseMethod) String() string {
	return map[chaseMethod]string{
		chaseMethodUnknown:     "chaseMethodUnknown",
		chaseMethodEmail:       "chaseMethodEmail",
		chaseMethodTextMessage: "chaseMethodTextMessage",
	}[x]
}
