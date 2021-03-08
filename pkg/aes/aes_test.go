package aes

import "testing"

const (
	key = "IgkibX71IEf382PT"
	iv  = "IgkibX71IEf382PT"
)

func TestEncrypt(t *testing.T) {
	t.Log(New(key, iv).Encrypt("anthonyzero"))
}

func TestDecrypt(t *testing.T) {
	t.Log(New(key, iv).Decrypt("Jel4N029LEB5HHLdwGXFPQ=="))
}

func BenchmarkEncryptAndDecrypt(b *testing.B) {
	b.ResetTimer()
	aes := New(key, iv)
	for i := 0; i < b.N; i++ {
		encryptString, _ := aes.Encrypt("hello world")
		aes.Decrypt(encryptString)
	}
}
