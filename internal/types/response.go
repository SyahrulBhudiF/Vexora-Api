package types

type WebResponse[T any] struct {
	Data         T             `json:"data"`
	Paging       *PageMetadata `json:"paging,omitempty"`
	Message      string        `json:"message,omitempty"`
	ShouldNotify bool          `json:"should_notify"`
	Success      bool          `json:"success"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
