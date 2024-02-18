package helpers

type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PaginatedResult struct {
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	Items    []interface{} `json:"items"`
}

func PaginateResult(items []interface{}, params PaginationParams) PaginatedResult {
	startIndex := (params.Page - 1) * params.PageSize
	endIndex := startIndex + params.PageSize

	if endIndex > len(items) {
		endIndex = len(items)
	}

	paginatedItems := items[startIndex:endIndex]

	return PaginatedResult{
		Total:    len(items),
		Page:     params.Page,
		PageSize: params.PageSize,
		Items:    paginatedItems,
	}
}
