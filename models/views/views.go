package views

type CreateFreeURLView struct {
	Url string `json:"url" binding:"required"`
}

type CreateURLView struct {
	Url            string `json:"url" binding:"required"`
	ExpiryDuration uint   `json:"expiry_duration" binding:"exists"`
}

type RedirectUrlView struct {
	Hash string `uri:"hash" binding:"required"`
}

type PasswordLoginView struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PasswordRegisterView struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GoogleLoginView struct {
	IdToken string `json:"idToken" binding:"required"`
}
