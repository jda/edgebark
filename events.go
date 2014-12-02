package main

var eventMap map[string]string = map[string]string{
	// user access events
	"granted":            "Access granted to %s %s at door %s",
	"granted_extratime":  "Access granted to %s %s at door %s (with extra time)",
	"denied_schedule":    "Access granted to %s %s at door %s (denied by schedule)",
	"denied_wrongpin":    "Access denied to %s %s at door %s (wrong PIN)",
	"denied_expiredcard": "Access denied to %s %s at door %s (card expired)",
	"denied_pinlocked":   "Access denied to %s %s at door %s (PIN locked out)",
	"denied_expiredpin":  "Access denied to %s %s at door %s (PIN expired)",

	// door events
	"denied_cardunassigned":  "Access at door %s denied: unassigned card",
	"denied_nocard":          "Access at door %s denied: unknown card",
	"denied_nopin":           "Access at door %s denied: no PIN",
	"door_schedule_unlocked": "Door %s unlocked on schedule",
	"door_locked_schedule":   "Door %s locked on schedule",
	"alarm_forced":           "Door %s forced open",
	"alarm_held":             "Door %s held open",
	"granted_manual":         "Access to %s granted (manual/override)",
	"door_unlocked":          "Door %s unlocked",
	"door_locked":            "Door %s locked",

	// system events
	"alarm_ack":    "Controller %s alarm acknowledged",
	"alarm_inputa": "Input A alarm on controller %s",
	"alarm_inputb": "Input B alarm on controller %s",
	"alarm_tamper": "Tamper alarm on controller %s",
	"time_set":     "Controller %s time set",
}
