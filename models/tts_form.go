package models

// TTSForm is the data structure for message form submissions
type TTSForm struct {
	Text string `form:"text"`
	Voice string `form:"voice"`
}
