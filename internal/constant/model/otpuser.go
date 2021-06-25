package model

type Otpuser struct {
	Phonenumber string `json:"phonenumber" binding:"required"`
	Otpcode     int    `json:"otpcode"`
}
