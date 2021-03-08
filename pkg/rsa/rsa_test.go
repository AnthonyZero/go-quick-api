package rsa

import (
	"testing"
)

const (
	publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1O3p0JN0/RrP7eY3f81i
zPf16FS0WMNGCJkd+y5c6yBzUvN0IEeoxiIWIBhoMKH0pzlzBg0rfttojSodOgNo
m/UCAzAYEgdIsNee5LSN/7e0T2/QvsIAHINuA8gI8fGoGiSA2TEzpUo6aVXwhZT3
4GGRdrSJ+m4iVk/Kt95tavBNk+NDVSeb5xAjxBchT5BjAMMlE0ffGZb0MMjjO5+e
9Tn8f99M2VMqpzXHXZzv1ABmqufzS20iWcSvnjhWcJ9hiKwO8Z30GgJyACmml+HM
xLYEFN9h2MWYgxLm9Z0rLMrWwMM+E2rCs8tsxAD5sO9RZMJPl1C0FIsMR53ngqbz
owIDAQAB
-----END PUBLIC KEY-----`

	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEA1O3p0JN0/RrP7eY3f81izPf16FS0WMNGCJkd+y5c6yBzUvN0
IEeoxiIWIBhoMKH0pzlzBg0rfttojSodOgNom/UCAzAYEgdIsNee5LSN/7e0T2/Q
vsIAHINuA8gI8fGoGiSA2TEzpUo6aVXwhZT34GGRdrSJ+m4iVk/Kt95tavBNk+ND
VSeb5xAjxBchT5BjAMMlE0ffGZb0MMjjO5+e9Tn8f99M2VMqpzXHXZzv1ABmqufz
S20iWcSvnjhWcJ9hiKwO8Z30GgJyACmml+HMxLYEFN9h2MWYgxLm9Z0rLMrWwMM+
E2rCs8tsxAD5sO9RZMJPl1C0FIsMR53ngqbzowIDAQABAoIBAQCO1RE1ItUlO6kj
Un0ENAgEqojAUqGvsT33Yo7kAZO+/cOeb0UEqk0iq5bf7L9ncBynWDg6ZPc6X3/g
wdFdKxAvHck9zjM3VL+EMP+bNyrR0K8ZYk5Kx+Q/PEK+Mp8dfRdgggAUsZaNWB+a
rVVspiMo1wo28KBl5x8NevTnJkOLqXAyB7UyLWqnOL1fb988lZvZPR7ZUYroVIZa
pyXtZcafIJeKyQ3bvWI5+eFqOe61Z4Bx1+TpfZ3fKfSDW0vhxzNqaimOa8jSXtMJ
jMeOctL4nZ0TPo/jS3I+XlaH4ZQlFLuUWGscpxwfEeBN23I8HRLkZXJsw66yvRN3
s4bUKPXRAoGBAP/3oSZAECvfsYYzs76tnrAmR/0GxCqgguxDlWn5DowQzdWFOdHC
ZbTo/hUVoMSQnO1EKCFlnBS+wg/3TuIzUO0ewC1aeT7qHbOMDl0zKbNpS2Z9/j+U
zro+qz7XmkWolMCfmDrCrw9CtCxcMSII+ajbI8SAgFVMz9XnDt+xW9E9AoGBANT0
4F6kCUJTEyqf2+v84tjQ2wGIF6XtZPU9JR806zeMyahQ9F6z3hY8BYb0tIy5b3uJ
VlJ9TG1qg/t59TWxIq43mYSUJHe0aJi3ilooObQtHlhPu8nwmmX47sX0PyG2hMoD
kBVxTpTDmBaDz7O9uBnlMXJN5qEygctaixpEbmZfAoGBAMBA9kEMjRjnAyeRXcgy
D6aumhNqKZz6wltCx864yjxZwsBFOJBcOpgPCAg+HmqFU9jCAIJVF05dmNT1I8Ky
WG5BUoa+FaMzpOtenstRylh/Far9pyGKW1t4BpdEyRLY9CFZvbUk1OfZagqHlD/E
DgDN16eX/MwUzWYUDg/l3tjhAoGBAKGip/ZNjVWRFpggs9z/mfK1O7WC5Wgksp9N
ZLK2CN6l9p3RrFmBLk00C4HulGfHi+15RVLhFbRqx3iFje/N3iPbwaMWikNtZIKd
tN5Pb9To9gJTqpZRD+/cLOeFRrHBBjMK1z7fPKS/fN2B+JFVq7nD827t3+J0In4F
4FT0odMDAoGBAJk3ELB/FHY8xzZ4jF1wG/a1CK681Xm6SuU5KIELDSAUNoou6OPG
mS8gU20MMPAeV2z7khyDcSxlHsUyL73eLeaakbQov9NMW7cc99XX4wnP4W7FRpmr
QbHmKuHIRFHCFv+XX8c0aK2mDZMUlzJdy4FgD/YCEZ7kZMZKyvZW/ZuV
-----END RSA PRIVATE KEY-----`
)

func TestEncrypt(t *testing.T) {
	//公钥加密 同一个content每次输出不一样
	str, err := NewPublic(publicKey).Encrypt("123456")
	if err != nil {
		t.Error("rsa public encrypt error", err)
		return
	}

	t.Log(str)
}

func TestDecrypt(t *testing.T) {
	decryptStr := "Y7af_Zq5NJN4xwHcVBiFmdPJRx4-IiRujKpapWG9O6YAYlBnM7AH50PLdj-0VnKZpllidgpUXuFlFa4_IcvQcJuo3u6oHvK-qmpISyS02adf4VORVQn3fUcceJMF2DQM86PvbojqQu91eATKKcOQim9AtjHa2wSNc68YLJJjt2NVD8FrLAWBCpLcIHmr7JffyerELnflyn29yz4O5FOdrZPwCfk-04nIQZ4mGBYf583rYjmZYkF_7TuBWMO_uCvv4_vSeZdDB03b2HYlF1JwnPrlX25lxtzmg5xUlkMmaRIGtvbqTDvse5yrkMbIAvN28GteGNH7pcMPvklbkRqj7A=="
	//私钥 解密
	str, err := NewPrivate(privateKey).Decrypt(decryptStr)
	if err != nil {
		t.Error("rsa private decrypt error", err)
		return
	}

	t.Log(str)
}

func BenchmarkEncryptAndDecrypt(b *testing.B) {
	b.ResetTimer()
	rsaPublic := NewPublic(publicKey)
	rsaPrivate := NewPrivate(privateKey)
	for i := 0; i < b.N; i++ {
		encryptString, _ := rsaPublic.Encrypt("123456")
		rsaPrivate.Decrypt(encryptString)
	}
}
