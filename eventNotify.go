package main

import (
	"fmt"
	"github.com/tbruyelle/hipchat-go/hipchat"
	"log"
)

// user access events
func eventNotifyUser(e EdgeEvent) {
	c := hipchat.NewClient(cfg.HipchatKey)

	msg := fmt.Sprintf(eventMap[e.Type], e.FirstName, e.LastName, e.Door)
	log.Println(msg)

	color := "red"
	notify := true
	switch e.Type {
	case "granted", "granted_extratime":
		color = "green"
	}

	notif := &hipchat.NotificationRequest{Message: msg, Color: color, Notify: notify}
	resp, err := c.Room.Notification(cfg.HipchatRoom, notif)
	if err != nil {
		log.Printf("error sending hipchat notification: %s, %+v", err, resp)
	}
}

// door events
func eventNotifyDoor(e EdgeEvent) {
	c := hipchat.NewClient(cfg.HipchatKey)

	msg := fmt.Sprintf(eventMap[e.Type], e.Door)
	log.Println(msg)

	color := "red"
	switch e.Type {
	case "door_schedule_unlocked", "door_locked_schedule":
		color = "grey"
	case "granted_manual", "door_unlocked", "door_locked":
		color = "purple"
	}

	notif := &hipchat.NotificationRequest{Message: msg, Color: color}
	resp, err := c.Room.Notification(cfg.HipchatRoom, notif)
	if err != nil {
		log.Printf("error sending hipchat notification: %s, %+v", err, resp)
	}
}

// system events
func eventNotifySys(e EdgeEvent) {
	c := hipchat.NewClient(cfg.HipchatKey)

	msg := fmt.Sprintf(eventMap[e.Type], e.Controller)
	log.Println(msg)

	color := "red"
	notify := true
	switch e.Type {
	case "time_set":
		color = "grey"
		notify = false
	case "alarm_ack":
		color = "yellow"
		notify = false
	}

	notif := &hipchat.NotificationRequest{Message: msg, Color: color, Notify: notify}
	resp, err := c.Room.Notification(cfg.HipchatRoom, notif)
	if err != nil {
		log.Printf("error sending hipchat notification: %s, %+v", err, resp)
	}
}
