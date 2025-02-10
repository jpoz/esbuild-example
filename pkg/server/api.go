package server

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
)

// Quote represents an inspirational quote.
type Quote struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

// A collection of sample quotes.
var quotes = []Quote{
	{ID: 1, Text: "The only limit to our realization of tomorrow is our doubts of today.", Author: "Franklin D. Roosevelt"},
	{ID: 2, Text: "In the middle of every difficulty lies opportunity.", Author: "Albert Einstein"},
	{ID: 3, Text: "What you get by achieving your goals is not as important as what you become by achieving your goals.", Author: "Zig Ziglar"},
	{ID: 4, Text: "Life is 10% what happens to us and 90% how we react to it.", Author: "Charles R. Swindoll"},
	{ID: 5, Text: "The best way to predict the future is to invent it.", Author: "Alan Kay"},
	{ID: 6, Text: "You miss 100% of the shots you don't take.", Author: "Wayne Gretzky"},
	{ID: 7, Text: "Whether you think you can or you think you can't, you're right.", Author: "Henry Ford"},
	{ID: 8, Text: "Don't watch the clock; do what it does. Keep going.", Author: "Sam Levenson"},
	{ID: 9, Text: "Keep your eyes on the stars, and your feet on the ground.", Author: "Theodore Roosevelt"},
	{ID: 10, Text: "The harder I work, the luckier I get.", Author: "Samuel Goldwyn"},
}

// quoteHandler returns a random quote as a JSON response.
func (s *Server) quoteHandler(w http.ResponseWriter, _ *http.Request) {
	slog.Info("Handling /api/quote request")
	// Pick a random quote.
	quote := quotes[rand.Intn(len(quotes))]

	// Set the response header and encode the quote as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}
