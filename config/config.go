package config

import "time"

// Key used to hash tokens and passwords
var SECRET_KEY = "go_project_secret_key"

// JWT Exp life time
var JWTExpireLife = time.Now().Add(time.Hour * 24)

// JWT Identifier
var JWTName = "token"

// JWT Secure with HttpOnly
var JWTHTTPOnly = true
