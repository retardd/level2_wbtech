package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// Event представляет событие в календаре.
type Event struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Title  string `json:"title"`
}

// Cache представляет кэш для хранения событий.
type Cache struct {
	mu     sync.RWMutex
	events map[int]Event
}

// NewCache создает новый экземпляр кэша.
func NewCache() *Cache {
	return &Cache{
		events: make(map[int]Event),
	}
}

func (c *Cache) HandleCreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Проверка наличия события по ID
	if _, ok := c.events[event.ID]; ok {
		http.Error(w, "Event already exists", http.StatusConflict)
		return
	}

	// Валидация данных
	if event.UserID <= 0 {
		http.Error(w, "Invalid UserID", http.StatusBadRequest)
		return
	}

	if event.Date == "" {
		http.Error(w, "Invalid Date", http.StatusBadRequest)
		return
	}

	// Добавление события в кэш
	c.events[event.ID] = event
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": "Event created successfully"})
}

// HandleUpdateEvent обрабатывает запрос на обновление события.
func (c *Cache) HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.events[event.ID]; !ok {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	c.events[event.ID] = event

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": "Event updated successfully"})
}

// HandleDeleteEvent обрабатывает запрос на удаление события.
func (c *Cache) HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.events[event.ID]; !ok {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	delete(c.events, event.ID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": "Event deleted successfully"})
}

// HandleEventsForDay обрабатывает запрос на получение событий на день.
func (c *Cache) HandleEventsForDay(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр date из запроса
	date := r.URL.Query().Get("date")

	// Проверяем, что параметр date был передан
	if date == "" {
		http.Error(w, "Date parameter is required", http.StatusBadRequest)
		return
	}

	// Перебираем все события и фильтруем по указанной дате
	var eventsForDay []Event
	for _, event := range c.events {
		if event.Date == date {
			eventsForDay = append(eventsForDay, event)
		}
	}

	// Кодируем найденные события в формат JSON и отправляем клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventsForDay)
}

// HandleEventsForWeek обрабатывает запрос на получение событий на неделю.
func (c *Cache) HandleEventsForWeek(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр date из запроса
	date := r.URL.Query().Get("date")

	// Проверяем, что параметр date был передан
	if date == "" {
		http.Error(w, "Date parameter is required", http.StatusBadRequest)
		return
	}

	// Преобразуем строку date в объект времени
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	// Определяем начало и конец недели для указанной даты
	startOfWeek := dateTime.AddDate(0, 0, -int(dateTime.Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	// Перебираем все события и фильтруем по диапазону дат для текущей недели
	var eventsForWeek []Event
	for _, event := range c.events {
		eventDate, err := time.Parse("2006-01-02", event.Date)
		if err != nil {
			continue
		}
		if eventDate.After(startOfWeek) && eventDate.Before(endOfWeek) {
			eventsForWeek = append(eventsForWeek, event)
		}
	}

	// Кодируем найденные события в формат JSON и отправляем клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventsForWeek)
}

// HandleEventsForMonth обрабатывает запрос на получение событий на месяц.
func (c *Cache) HandleEventsForMonth(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр date из запроса
	date := r.URL.Query().Get("date")

	// Проверяем, что параметр date был передан
	if date == "" {
		http.Error(w, "Date parameter is required", http.StatusBadRequest)
		return
	}

	// Преобразуем строку date в объект времени
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	// Определяем начало и конец месяца для указанной даты
	startOfMonth := time.Date(dateTime.Year(), dateTime.Month(), 1, 0, 0, 0, 0, dateTime.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	// Перебираем все события и фильтруем по диапазону дат для текущего месяца
	var eventsForMonth []Event
	for _, event := range c.events {
		eventDate, err := time.Parse("2006-01-02", event.Date)
		if err != nil {
			continue
		}
		if eventDate.After(startOfMonth) && eventDate.Before(endOfMonth) {
			eventsForMonth = append(eventsForMonth, event)
		}
	}

	// Кодируем найденные события в формат JSON и отправляем клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventsForMonth)
}

func main() {
	// Инициализируем кэш для хранения данных
	cache := NewCache()

	// Инициализируем маршрутизатор и добавляем обработчики для методов API
	router := http.NewServeMux()
	router.HandleFunc("/create_event", cache.HandleCreateEvent)
	router.HandleFunc("/update_event", cache.HandleUpdateEvent)
	router.HandleFunc("/delete_event", cache.HandleDeleteEvent)
	router.HandleFunc("/events_for_day", cache.HandleEventsForDay)
	router.HandleFunc("/events_for_week", cache.HandleEventsForWeek)
	router.HandleFunc("/events_for_month", cache.HandleEventsForMonth)

	// Запускаем HTTP-сервер
	port := ":8080" // Укажите порт, на котором сервер должен слушать запросы
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
