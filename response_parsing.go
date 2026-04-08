package sockudo

import (
	"encoding/json"
)

// Channel represents the information about a channel from the Sockudo API.
type Channel struct {
	Name              string
	Occupied          bool `json:"occupied,omitempty"`
	UserCount         int  `json:"user_count,omitempty"`
	SubscriptionCount int  `json:"subscription_count,omitempty"`
}

// ChannelsList represents a list of channels received by the Sockudo API.
type ChannelsList struct {
	Channels map[string]ChannelListItem `json:"channels"`
}

// ChannelListItem represents an item within ChannelsList
type ChannelListItem struct {
	UserCount int `json:"user_count"`
}

type TriggerChannelsList struct {
	Channels map[string]TriggerChannelListItem `json:"channels"`
}

type TriggerChannelListItem struct {
	UserCount         *int `json:"user_count,omitempty"`
	SubscriptionCount *int `json:"subscription_count,omitempty"`
}

type TriggerBatchChannelsList struct {
	Batch []TriggerBatchChannelListItem `json:"batch"`
}

type TriggerBatchChannelListItem struct {
	UserCount         *int `json:"user_count,omitempty"`
	SubscriptionCount *int `json:"subscription_count,omitempty"`
}

// Users represents a list of users in a presence-channel
type Users struct {
	List []User `json:"users"`
}

type HistoryPage struct {
	Items      []HistoryItem      `json:"items"`
	Direction  string             `json:"direction"`
	Limit      int                `json:"limit"`
	HasMore    bool               `json:"has_more"`
	NextCursor *string            `json:"next_cursor"`
	Bounds     HistoryBounds      `json:"bounds"`
	Continuity HistoryContinuity  `json:"continuity"`
}

type HistoryItem struct {
	StreamID        string                 `json:"stream_id"`
	Serial          int64                  `json:"serial"`
	PublishedAtMS   int64                  `json:"published_at_ms"`
	MessageID       *string                `json:"message_id"`
	EventName       *string                `json:"event_name"`
	OperationKind   string                 `json:"operation_kind"`
	PayloadSizeByte int                    `json:"payload_size_bytes"`
	Message         map[string]interface{} `json:"message"`
}

type HistoryBounds struct {
	StartSerial *int64 `json:"start_serial"`
	EndSerial   *int64 `json:"end_serial"`
	StartTimeMS *int64 `json:"start_time_ms"`
	EndTimeMS   *int64 `json:"end_time_ms"`
}

type HistoryContinuity struct {
	StreamID                   *string `json:"stream_id"`
	OldestAvailableSerial      *int64  `json:"oldest_available_serial"`
	NewestAvailableSerial      *int64  `json:"newest_available_serial"`
	OldestAvailablePublishedMS *int64  `json:"oldest_available_published_at_ms"`
	NewestAvailablePublishedMS *int64  `json:"newest_available_published_at_ms"`
	RetainedMessages           int64   `json:"retained_messages"`
	RetainedBytes              int64   `json:"retained_bytes"`
	Complete                   bool    `json:"complete"`
	TruncatedByRetention       bool    `json:"truncated_by_retention"`
}

// User represents a user and contains their ID.
type User struct {
	ID string `json:"id"`
}

/*
MemberData represents what to assign to a channel member, consisting of a
`UserID` and any custom `UserInfo`.
*/
type MemberData struct {
	UserID   string            `json:"user_id"`
	UserInfo map[string]string `json:"user_info,omitempty"`
}

func unmarshalledTriggerChannelsList(response []byte) (*TriggerChannelsList, error) {
	channels := &TriggerChannelsList{}
	err := json.Unmarshal(response, channels)

	if err != nil {
		return nil, err
	}

	return channels, nil
}

func unmarshalledTriggerBatchChannelsList(response []byte) (*TriggerBatchChannelsList, error) {
	channels := &TriggerBatchChannelsList{}
	err := json.Unmarshal(response, channels)

	if err != nil {
		return nil, err
	}

	return channels, nil
}

func unmarshalledChannelsList(response []byte) (*ChannelsList, error) {
	channels := &ChannelsList{}
	err := json.Unmarshal(response, channels)

	if err != nil {
		return nil, err
	}

	return channels, nil
}

func unmarshalledChannel(response []byte, name string) (*Channel, error) {
	channel := &Channel{Name: name}
	err := json.Unmarshal(response, channel)

	if err != nil {
		return nil, err
	}

	return channel, nil
}

func unmarshalledChannelUsers(response []byte) (*Users, error) {
	users := &Users{}
	err := json.Unmarshal(response, users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func unmarshalledHistory(response []byte) (*HistoryPage, error) {
	page := &HistoryPage{}
	err := json.Unmarshal(response, page)
	if err != nil {
		return nil, err
	}
	return page, nil
}
