package core

import (
	"encoding/json"
	"net/http"

	"github.com/Thiti-Dev/scraperor/graph/model"
)


func ScrapeHandle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var body model.ScrapeInput
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				http.Error(w, "JSON body required", http.StatusBadRequest)
				return
			}
			// Validation [Manual]
			if body.Pointer == "" || body.TargetURL == "" {
				http.Error(w, "JSON body required both key of target_url,pointer", http.StatusBadRequest)
				return
			}

			// SCRAPING PROCESS
			engineEntity := CreateEngineEntity()
			engineEntity.AddTask(body.TargetURL, body.Pointer)

			elements := engineEntity.Scrape()
			// ─────────────────────────────────────────────────────────────────

			// ─── FORMATTING RESP ─────────────────────────────────────────────

			response_data := make(map[string]interface{})
			response_data["success"] = true
			response_data["elements"] = elements
			json_resp, _ := json.Marshal(response_data)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(json_resp)

			// ─────────────────────────────────────────────────────────────────

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 - Method not allowed!"))
		}
	}
}