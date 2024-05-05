package types

type ImagesCreateRequest struct{}

type ImageUploadRequest struct {
	Image string `json:"image"`
}

type ImageUploadResponse struct {
	Url string `json:"url"`
}
