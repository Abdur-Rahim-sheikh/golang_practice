package utils

import "net/http"
import "encoding/json"

type PaginatedData struct {
	Data           any        `json:"data"`
	PaginationInfo Pagination `json:"pagination_info"`
}
type Pagination struct {
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
	TotalItems int64 `json:"total_items"`
	TotalPagse int64 `json:"total_pages"`
}

func SendData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func SendError(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}

func SendPage(w http.ResponseWriter, data any, page, limit, total_element int64) {
	paginated_info := Pagination{
		Page:       page,
		Limit:      limit,
		TotalItems: total_element,
		TotalPagse: total_element / limit,
	}
	paginated_data := PaginatedData{
		Data:           data,
		PaginationInfo: paginated_info,
	}
	SendData(w, http.StatusOK, paginated_data)
}
