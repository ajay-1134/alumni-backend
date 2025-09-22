package config

import "time"

var JWTSecret = []byte("supersecretKey")

const TokenExpiry = time.Minute*30