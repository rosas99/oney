package v1

import (
	"github.com/rosas/onex/internal/nightwatch/biz"
	"github.com/rosas/onex/internal/nightwatch/validation"
)

type NightWatchService struct {
	valid *validation.Validator
	biz   biz.IBiz
}

func NewNightWatchService(valid *validation.Validator, biz biz.IBiz) *NightWatchService {
	return &NightWatchService{valid: valid, biz: biz}
}
