package views

type CreateFreeURLView struct {
	Url            string `json:"url" binding:"required"`
	ExpiryDuration uint   `json:"expiry_duration" binding:"exists"`
}

type CreateURLView struct {
	Url            string `json:"url" binding:"required"`
	ExpiryDuration uint   `json:"expiry_duration" binding:"exists"`
}

type RedirectUrlView struct {
	Hash string `uri:"hash" binding:"required"`
}
