package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type EdgeEvent struct {
	Controller string
	Category   string
	Type       string
	Door       string
	FirstName  string
	LastName   string
	Recieved   time.Time
}

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
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "E_BAD_CAT")
		return
	}

	// validate controller ID (in the future, check if ID is known to us/in DB)
	if controllerId == "" {
		log.Printf("No controller ID from %s", r.RemoteAddr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "E_NO_ID")
		return
	}

	// process by type
	ev := EdgeEvent{Controller: controllerId, Category: category, Recieved: time.Now()}
	switch t {

	// user access events with cardholder fname, lname
	case "granted", "granted_extratime", "denied_schedule", "denied_wrongpin",
		"denied_expiredcard", "denied_pinlocked", "denied_expiredpin":
		ev.Type = t
		ev.Door = r.FormValue("door")
		if ev.Door == "" {
			log.Printf("Got event %s without door form %s", t, r.RemoteAddr)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "E_NO_DOOR")
			return
		}

		ev.FirstName = r.FormValue("firstname")
		ev.LastName = r.FormValue("lastname")

		go eventNotifyUser(ev)

	// door events without cardholder fname, lname
	case "denied_cardunassigned", "denied_nocard", "denied_nopin",
		"door_schedule_unlocked", "door_locked_schedule", "alarm_forced",
		"alarm_held", "granted_manual", "door_unlocked", "door_locked":
		ev.Type = t
		ev.Door = r.FormValue("door")
		if ev.Door == "" {
			log.Printf("Got event %s without door form %s", t, r.RemoteAddr)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "E_NO_DOOR")
			return
		}

		go eventNotifyDoor(ev)

	// system events
	case "alarm_ack", "alarm_inputa", "alarm_inputb", "alarm_tamper", "time_set":
		ev.Type = t
		ev.Door = r.FormValue("door")

		go eventNotifySys(ev)

	// unknown event
	default:
		log.Printf("Unknown type %s from %s", t, r.RemoteAddr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "E_BAD_EVENT")
		return
	}

	fmt.Fprintln(w, "ACK")
}
