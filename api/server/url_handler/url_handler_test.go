package url_handler

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"sync"
	"testing"
	"time"
	"tinyUrl/config/env"
	"tinyUrl/models"
	"tinyUrl/services/cache"
	"tinyUrl/services/url"
	"tinyUrl/services/usecase"
	"tinyUrl/types/code"
	"tinyUrl/types/enums"
)

type urlHandlerTest struct {
	UrlUCase   url.UseCase
	CacheUCase cache.UseCase
}

func before() {
	env.InitEnvironment(enums.Testing)
	models.InitModels()
	usecase.InitUseCase()
}

func after() {
	if err := models.ClearDB(); err != nil {
		log.Fatalf("[Error] can not drop database: %s", err)
	}
	log.Print("[INFO] drop datebase successfully!")
	if err := models.ClearCache(); err != nil {
		log.Fatalf("[Error] can not flushall cache: %s", err)
	}
	log.Print("[INFO] flushall cache successfully!")
}

func TestCreateFreeUrlAPI(t *testing.T) {
	before()
	handler := &urlHandlerTest{UrlUCase: usecase.UrlUCase}

	originalUrl := "https://www.facebook.com/"
	url, err := handler.UrlUCase.CreateFreeUrl(originalUrl)
	if err != nil {
		t.Fatalf("could not create free tiny url: %s", err)
	}

	if url.OriginalURL != originalUrl {
		t.Fatalf("expect original url is %s but got %s", url.OriginalURL, originalUrl)
	}

	diff := url.ExpirationDate.Sub(url.CreationDate)
	if int(diff.Seconds()) != int((24 * time.Hour).Seconds()) {
		t.Fatalf("expect expired date is %s but got %s", url.CreationDate.Add(24*time.Hour), url.ExpirationDate)
	}
	after()
}

func TestCreateUrlAPI(t *testing.T) {
	before()
	handler := &urlHandlerTest{UrlUCase: usecase.UrlUCase}
	user := models.User{
		ID:           primitive.NewObjectID(),
		Name:         "nguyen",
		Email:        "pqnguyen1996@gmail.com",
		Password:     "",
		CreationDate: time.Now(),
		LastLogin:    time.Now(),
	}
	originalUrl := "https://www.facebook.com/"
	url, err := handler.UrlUCase.CreateUrl(&user, originalUrl, uint(2*24*time.Hour))
	if err != nil {
		t.Fatalf("could not create tiny url: %s", err)
	}

	if url.OriginalURL != originalUrl {
		t.Fatalf("expect original url is %s but got %s", url.OriginalURL, originalUrl)
	}

	diff := url.ExpirationDate.Sub(url.CreationDate)
	if int(diff.Seconds()) != int((2 * 24 * time.Hour).Seconds()) {
		t.Fatalf("expect expired date is %s but got %s", url.CreationDate.Add(2*24*time.Hour), url.ExpirationDate)
	}
	after()
}

func TestRedirectUrl(t *testing.T) {
	before()
	handler := &urlHandlerTest{
		UrlUCase:   usecase.UrlUCase,
		CacheUCase: usecase.CacheUCase,
	}
	user := models.User{
		ID:           primitive.NewObjectID(),
		Name:         "nguyen",
		Email:        "pqnguyen1996@gmail.com",
		Password:     "",
		CreationDate: time.Now(),
		LastLogin:    time.Now(),
	}
	originalUrl := "https://www.facebook.com/"
	url, err := handler.UrlUCase.CreateUrl(&user, originalUrl, uint(2*24*time.Hour))
	if err != nil {
		t.Fatalf("could not create tiny url: %s", err)
	}
	savedUrl, err := handler.UrlUCase.GetRedirectUrl(url.Hash)
	if err != nil {
		t.Fatalf("could not get tiny url: %s", err)
	}
	if savedUrl != originalUrl {
		t.Fatalf("saved url are different from original url: %s != %s", savedUrl, originalUrl)
	}
	redirectUrl, exists := handler.CacheUCase.GetOriginalUrl(url.Hash)
	if !exists {
		t.Fatal("original url wasn't cached")
	}
	if redirectUrl != originalUrl {
		t.Fatalf("redirect url url are different from original url: %s != %s", savedUrl, originalUrl)
	}
	after()
}

func TestExpiredUrl(t *testing.T) {
	before()
	handler := &urlHandlerTest{
		UrlUCase:   usecase.UrlUCase,
		CacheUCase: usecase.CacheUCase,
	}
	user := models.User{
		ID:           primitive.NewObjectID(),
		Name:         "nguyen",
		Email:        "pqnguyen1996@gmail.com",
		Password:     "",
		CreationDate: time.Now(),
		LastLogin:    time.Now(),
	}
	originalUrl := "https://www.facebook.com/"
	url, err := handler.UrlUCase.CreateUrl(&user, originalUrl, uint(1*time.Second))
	if err != nil {
		t.Fatalf("could not create tiny url: %s", err)
	}
	_, err = handler.UrlUCase.GetRedirectUrl(url.Hash)
	if err != nil {
		t.Fatalf("could not get tiny url: %s", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	time.AfterFunc(1*time.Second, func() {
		_, err = handler.UrlUCase.GetRedirectUrl(url.Hash)
		if err != nil && err != code.ErrTinyUrlExpired {
			wg.Done()
			t.Fatalf("expect error code is %s, but got %s", code.ErrTinyUrlExpired, err)
		}
		wg.Done()
	})

	time.AfterFunc(1*time.Second, func() {
		_, exists := handler.CacheUCase.GetOriginalUrl(url.Hash)
		if exists {
			wg.Done()
			t.Fatalf("expect expired url was clear but not")
		}
		wg.Done()
	})
	wg.Wait()
	after()
}

func BenchmarkRedirectUrl(t *testing.B) {
	before()
	handler := &urlHandlerTest{
		UrlUCase: usecase.UrlUCase,
	}
	for n := 1; n <= 1024; n *= 2 {
		for i := 0; i < t.N; i++ {
			t.StartTimer()
			user := models.User{
				ID:           primitive.NewObjectID(),
				Name:         "nguyen",
				Email:        "pqnguyen1996@gmail.com",
				Password:     "",
				CreationDate: time.Now(),
				LastLogin:    time.Now(),
			}
			originalUrl := "https://www.facebook.com/"
			url, _ := handler.UrlUCase.CreateUrl(&user, originalUrl, uint(1*time.Hour))
			redirectUrl, _ := handler.UrlUCase.GetRedirectUrl(url.Hash)
			if redirectUrl != originalUrl {
				t.Fatalf("expect redirect url is the same as original url, but got %s != %s", redirectUrl, originalUrl)
			}
			t.StopTimer()
		}
	}
	after()
}
