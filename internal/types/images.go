package types

type ImagesCreateRequest struct{}

type ImageUploadRequest struct {
	Image string `json:"id"`
}

type ImageUploadResponse struct {
	Url string `json:"url"`
}
