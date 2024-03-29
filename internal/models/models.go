package models

import (
	"github.com/vk-rv/pvx"
	"time"
)

type Credentials struct {
	Password string `form:"password"`
	Username string `form:"username"`
}

type AdditionalClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Footer struct {
	MetaData string `json:"meta_data"`
}

type TokenData struct {
	Subject  string
	Duration time.Duration
	AdditionalClaims
	Footer
}

type ServiceClaims struct {
	pvx.RegisteredClaims
	AdditionalClaims
	Footer
}
