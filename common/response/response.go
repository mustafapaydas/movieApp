package response

type PageResponse struct {
	Data       any `json:"data"`
	Count      int `json:"count"`
	TotalPage  int `json:"totalPage"`
	StatusCode int
	Error      error
}
type DataResponse struct {
	Data       any `json:"data"`
	StatusCode int
	Error      error
}
