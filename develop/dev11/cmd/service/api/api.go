package api

type DtoEvent struct {
	Date    string `json:"date"`
	Info    string `json:"info"`
	EventID string `json:"event_id"`
}

type DtoCreateEvent struct {
	UserID string   `json:"user_id"`
	Event  DtoEvent `json:"event"`
}
