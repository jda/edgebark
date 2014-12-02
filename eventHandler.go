package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func newEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	controllerId := params["id"]
	category := params["category"]
	t := params["type"]

	// validate category
	switch category {
	case "access", "alarm", "door", "system":
	default:
		log.Printf("Unknown category %s from %s", category, r.RemoteAddr)
		fmt.Fprintln(w, "E_WRONG_CAT")
		return
	}

	// validate controller ID (in the future, check if ID is known to us/in DB)
	if controllerId == "" {
		log.Printf("No controller ID from %s", r.RemoteAddr)
		fmt.Fprintln(w, "E_NO_ID")
		return
	}

	// process by type
	switch t {
	case "denied_nocard":
	case "denied_nopin":
	case "granted":
	case "granted_extratime":
	case "denied_schedule":
	case "denied_wrongpin":
	case "denied_expiredcard":
	case "denied_pinlocked":
	case "denied_cardunassigned":
	case "denied_expiredpin":
	case "alarm_ack":
	case "door_schedule_unlocked":
	case "door_locked_schedule":
	case "alarm_forced":
	case "alarm_held":
	case "alarm_inputa":
	case "alarm_inputb":
	case "time_set":
	case "granted_manual":
	case "door_unlocked":
	case "door_locked":
	default:
		log.Printf("Unknown type %s from %s", t, r.RemoteAddr)
	}

	fmt.Fprintln(w, "OK")
}
