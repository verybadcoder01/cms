package internal

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

// SessionAuthCheck проверяет, актуальна ли текущая сессия модера
func SessionAuthCheck(c *fiber.Ctx) int {
	session, exists := c.GetReqHeaders()["Session"]
	if !exists {
		log.Println("session header not found")
		return http.StatusBadRequest
	}
	info, exists := sessions.Load(session)
	if !exists {
		log.Printf("user not authorized on access route, session %v", info)
		return http.StatusForbidden
	}
	if info.(Session).isExpired() {
		log.Printf("session %v has expired", info)
		sessions.Delete(session)
		return http.StatusUnauthorized
	}
	return 0
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s Session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func (s Session) untilExpiration() int64 {
	if s.isExpired() {
		return 0
	}
	return s.expiry.Unix() - time.Now().Unix()
}
