package common

//note : file ini untuk buat claims JWT
// claims JWT adalah data yang akan dienkripsi dan ditandatangani oleh JWT
// claims JWT ini nantinya akan disimpan di dalam payload JWT

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID                   int64  `json:"id"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Role                 string `json:"role"`
	jwt.RegisteredClaims        // ini untuk exp at
}
