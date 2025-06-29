package infrastructure

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/brkss/dextrace/internal/domain"
)





type NightscoutRepository struct {
	nightscoutURL 	string
	token 			string
}


func NewNightscoutRepository(URL string, token string) *NightscoutRepository {
	return &NightscoutRepository{
		nightscoutURL: URL,
		token: token,
	}
}


func (r *NightscoutRepository) PushData(data []domain.GetDataResponse) error {
	client := &http.Client{}
	
	resp, err := client.Get(r.nightscoutURL + "/api/v1/entries.json?count=1")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var existingEntries []domain.NightscoutEntry;
	if err := json.NewDecoder(resp.Body).Decode(&existingEntries); err != nil {
		return err
	}

	var latestTimestamp time.Time
	if len(existingEntries) > 0 {
		latestTimestamp, _ = time.Parse(time.RFC3339, existingEntries[0].DateString)
	}

	
	var newEntries []domain.NightscoutEntry
	for _, d := range data {
		t, err := time.Parse(time.RFC3339, d.Timestamp)
		if err != nil {
			// Try parsing as Unix timestamp in milliseconds
			timestampInt := int64(0)
			_, err = fmt.Sscanf(d.Timestamp, "%d", &timestampInt)
			if err != nil {
				continue
			}
			t = time.Unix(0, timestampInt*int64(time.Millisecond))
		}
		
		if err != nil {
			continue
		}
		if t.After(latestTimestamp) {
			newEntries = append(newEntries, domain.NightscoutEntry{
				Type:      "sgv",
				SGV:       d.Value,
				Date:      t.UnixNano() / int64(time.Millisecond),
				DateString: d.Timestamp,
			})
		}
	}

	
	if len(newEntries) == 0 {
		return nil;
	}

	payload, err := json.Marshal(newEntries)
	if err != nil {
		return err;
	}

	fmt.Println("token : ", r.token)
	req, err := http.NewRequest("POST", r.nightscoutURL+"/api/v1/entries.json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if r.token != "" {
		h := sha1.New()
		h.Write([]byte(r.token))
		hashed := hex.EncodeToString(h.Sum(nil))
		req.Header.Set("api-secret", hashed)
			
	}

	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	errMsg := fmt.Sprintf("Failed to push data, status: %s\n", resp.Status)
	return errors.New(errMsg)

	
}