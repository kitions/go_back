package model

import "go_back/entity"

type FeedbackResp struct {
	entity.Feedback
	Pictures []entity.Picture
}
