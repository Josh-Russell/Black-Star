package main

import (
	"net/http"
	"time"
)

type Session struct {
	ID     string
	UserID string
	Expiry time.Time
}

func NewSession(w http.ResponseWriter) *Session {
	expiry := time.Now().Add(sessionLength)

	session := &Session{
		ID:     GenerateID("sess", sessionIDLength),
		Expiry: expiry,
	}

	cookie := http.Cookie{
		Name:    sessionCookieName,
		Value:   session.ID,
		Expires: session.Expiry,
	}

	http.SetCookie(w, &cookie)
	return session
}
