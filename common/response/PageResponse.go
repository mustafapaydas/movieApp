package response

type PageResponse struct {
	Data  any `json:"data"`
	Count int `json:"count"`
}
