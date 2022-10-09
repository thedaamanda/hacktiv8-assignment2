package responses

import . "assignment2/models"

type DataList struct {
	Count   int     `json:"count"`
	Data    []Order `json:"data"`
	Message string  `json:"message"`
}

type Data struct {
	Data    Order  `json:"data"`
	Message string `json:"message"`
}

type DataAffected struct {
	Count   int    `json:"count"`
	Message string `json:"message"`
}
