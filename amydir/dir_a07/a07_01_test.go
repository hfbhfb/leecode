package dira07

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Event struct {
	EventType string `json:"event_type"`
}

type ClickEvent struct {
	Event  `json:",inline"`
	OnPosX int64 `json:"on_pos_x"`
	OnPosY int64 `json:"on_pos_y"`
}

type DragEvent struct {
	Event  `json:",inline"`
	OnPosX int64 `json:"on_pos_x"`
	OnPosY int64 `json:"on_pos_y"`
	ToPosX int64 `json:"to_pos_x"`
	ToPosY int64 `json:"to_pos_y"`
}

func TestXxx(t *testing.T) {
	event1 := ClickEvent{
		Event:  Event{EventType: "click"},
		OnPosX: 0,
		OnPosY: 0,
	}
	b1, _ := json.Marshal(&event1)

	event2 := DragEvent{
		Event:  Event{EventType: "drag"},
		OnPosX: 1,
		OnPosY: 2,
		ToPosX: 3,
		ToPosY: 4,
	}
	b2, _ := json.Marshal(&event2)

	fmt.Println(string(b1)) // {"event_type":"click","on_pos_x":0,"on_pos_y":0}
	fmt.Println(string(b2)) // {"event_type":"drag","on_pos_x":1,"on_pos_y":2,"to_pos_x":3,"to_pos_y":4}
}
