package dto

type VideoDTO struct {
	VideoURL  string `json:"video_url" form:"video_url"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
}
