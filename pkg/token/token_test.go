package token

import (
	"testing"
	"time"
)

const secret = "i1ydX9RtHyuJTrw7frcu"

func TestSign(t *testing.T) {
	tokenString, err := New(secret).Sign(123456789, "pingjin", 24*time.Hour)
	if err != nil {
		t.Error("sign error", err)
		return
	}
	t.Log(tokenString)
}

func TestParse(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEyMzQ1Njc4OSwiVXNlck5hbWUiOiJwaW5namluIiwiZXhwIjoxNjE1NDUxMDAzLCJpYXQiOjE2MTUzNjQ2MDMsIm5iZiI6MTYxNTM2NDYwM30.r7NSeQSTd66YQRsHaYmY8X53DjBtABp4DM-fqXWFfzU"
	user, err := New(secret).Parse(tokenString)
	if err != nil {
		t.Error("parse error", err)
		return
	}
	t.Log(user)
}

func BenchmarkSignAndParse(b *testing.B) {
	b.ResetTimer()
	token := New(secret)
	for i := 0; i < b.N; i++ {
		tokenString, _ := token.Sign(123456789, "pingjin", 24*time.Hour)
		token.Parse(tokenString)
	}
}
