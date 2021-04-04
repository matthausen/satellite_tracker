package models

type (
	// Response type of the API
	Response struct {
		Info      Info       `json:"info"`
		Positions []Position `json:"positions"`
	}

	// Info type of the satellite
	Info struct {
		Satname          string `json:"satname"`
		Satid            int    `json:"satid"`
		TransactionCount int    `json:"transactionscount"`
	}

	// Position type of the satellite
	Position struct {
		Latitude  float64 `json:"satlatitude"`
		Longitude float64 `json:"satlongitude"`
		Altitude  float64 `json:"sataltitude"`
		Timestamp int
	}
)
