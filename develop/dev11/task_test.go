package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// перед тестом нужно запустить сервер

func TestServer(t *testing.T) {
	tests := []struct {
		name     string
		event    Event
		url      string
		expected string
	}{
		{name: "test_event_created",
			event: Event{
				ID:     3,
				UserID: 123,
				Date:   "2024-02-15",
				Title:  "Test event2",
			},
			url:      "http://localhost:8080/create_event",
			expected: `{"result":"Event created successfully"}`},
		{name: "test_event_update",
			event: Event{
				ID:     3,
				UserID: 1233213,
				Date:   "2024-02-15",
				Title:  "Test event2_22",
			},
			url:      "http://localhost:8080/update_event",
			expected: `{"result":"Event updated successfully"}`},
		{name: "test_event_delete",
			event: Event{
				ID:     3,
				UserID: 1233213,
				Date:   "2024-02-15",
				Title:  "Test event2_22",
			},
			url:      "http://localhost:8080/delete_event",
			expected: `{"result":"Event deleted successfully"}`},
	}

	for _, test := range tests {
		eventJSON, err := json.Marshal(test.event)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.Post(test.url, "application/json", bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Проверка статуса ответа
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		// Проверка тела ответа
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		if strings.TrimSpace(string(body)) != strings.TrimSpace(test.expected) {
			t.Errorf("expected body %s, got %s", test.expected, string(body))
		}

		err = resp.Body.Close()
		if err != nil {
			return
		}
	}

}
