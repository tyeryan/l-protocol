package app

import (
	"github.com/tyeryan/l-protocol/event"
)

var (
	AppUserTopic = event.Topic{
		Domain: "app",
		Name:   "AppUser",
		ID:     "uapp.appuser",
		Events: map[string]string{
			"Created":    "uapp.appuser.created",
			"Disabled":   "uapp.appuser.disabled",
			"PinChanged": "uapp.appuser.pinchanged",
			"PinReset":   "uapp.appuser.pinreset",
			"Unlocked":   "uapp.appuser.unlocked",
		},
		ProtoPackage:  "github.com/tyeryan/l-protocol/go/lapp",
		ProtoName:     "AppUser",
		Keys:          []string{"UserID"},
		PartitionKeys: nil,
	}
)
