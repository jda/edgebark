# Edge Bark
Card Access event notification handler for HID Edge access controllers

## Configuration

## Controller setup
Your access controller needs to be configured to send event notifications to edgebark.

First name: firstname
Last name: lastname
Door name: door

Event handler URLs follow this pattern: /event/*controller_id*/*category*/*type*?**PARAMS**

controller_id is whatever you call your controller, e.g.: msp-office1-floor1
category is one of access, alarm, door, or system

| Description | Category | Type | Params |
|---|---|---|---|---|
| Access granted | access | granted | door=$doorname&firstname=$firstname&lastname=$lastname  |
| Access granted with extra time  | access | granted_extratime | door=$doorname&firstname=$firstname&lastname=$lastname |
| Access denied by schedule | access | denied_schedule | door=$doorname&firstname=$firstname&lastname=$lastname |
| Access denied because of a invaid PIN  | access | denied_wrongpin | door=$doorname&firstname=$firstname&lastname=$lastname |
|  Access denied because of a expired card | access | denied_expiredcard | door=$doorname&firstname=$firstname&lastname=$lastname |
| Access denied because PIN is locked out  | access | denied_pinlocked | door=$doorname&firstname=$firstname&lastname=$lastname |
| Access denied because PIN expired  | access | denied_expiredpin |  door=$doorname&firstname=$firstname&lastname=$lastname |
| Access denied because card known but unassigned  | door | denied_cardunassigned | door=$doorname |
| Access denied because no card (tried to use PIN but card required)  | door | denied_nocard | door=$doorname |
| Access denied because no PIN (tried to use card but PIN required) | door | denied_nopin | door=$doorname |
| Door unlocked by schedule  | door | door_schedule_unlocked  | door=$doorname |
| Door locked by schedule | door | door_locked_schedule  | door=$doorname |
| Door open without access grant | door | alarm_forced | door=$doorname |
| Door held open past timeout | door | alarm_held | door=$doorname |
| Access granted manually through controller | door | granted_manual | door=$doorname |
| Door unlocked manually through controller  | door | door_unlocked | door=$doorname |
| Door locked manually through controller  | door | door_locked | door=$doorname |
| System time set | system | time_set |  |  
| Alarm acknowledged  | alarm | alarm_ack | | 
| Alarm on input A | alarm | alarm_inputa | | 
| Alarm on input B | alarm | alarm_inputb | | 
| Tamper alarm | alarm | alarm_tamper| | 



