package model

type AuthMessage struct {
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
}

type SubscribeEventMessage struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	EventType string `json:"event_type"`
}

type UnsubscribeEventMessage struct {
	Id           int    `json:"id"`
	Type         string `json:"type"`
	Subscription int    `json:"subscription"`
}

type SubscribeTriggerMessage struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Trigger struct {
		Platform string `json:"platform"`
		EntityId string `json:"entity_id"`
		From     string `json:"from"`
		To       string `json:"to"`
	} `json:"trigger"`
}

type FireEventMessage struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	EventType string `json:"event_type"`
	EventData struct {
		DeviceId string `json:"device_id"`
		Type     string `json:"type"`
	} `json:"event_data"`
}
