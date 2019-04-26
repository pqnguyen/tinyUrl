package cache

type UseCase interface {
	GetOriginalUrl(hash string) (string, bool)
}
