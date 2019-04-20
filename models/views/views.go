package views

type CreateURLView struct {
	Url            string `json:"url" binding:"required"`
	ExpiryDuration uint   `json:"expiry_duration" binding:"exists"`
}

type RedirectUrlView struct {
	Hash string `uri:"hash" binding:"required"`
}
