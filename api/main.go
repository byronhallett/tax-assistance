// Package handler contains an HTTP Cloud Function.
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type data struct {
	QueryResult struct {
		Parameters struct {
			Number       int64 `json:"number"`
			UnitCurrency struct {
				Amount int64 `json:"amount"`
			} `json:"unit-currency"`
		} `json:"parameters"`
	} `json:"queryResult"`
}

// Handler does thing
func Handler(w http.ResponseWriter, r *http.Request) {
	var d data
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, err)
		return
	}
	var value int64 = 0
	if d.QueryResult.Parameters.Number != "" {
		value, _ = strconv.ParseInt(d.QueryResult.Parameters.Number, 10, 64)
	}
	if d.QueryResult.Parameters.UnitCurrency.Amount != "" {
		value, _ = strconv.ParseInt(d.QueryResult.Parameters.UnitCurrency.Amount, 10, 64)
	}
	// TODO: choose between the number and currency amount fields
	// TODO: calculate withholdable amount
	fmt.Fprintf(w, "Withhold $%d. Any further calculations?", value)
}
