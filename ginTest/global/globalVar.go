package global

import (
	"ginTest/config"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
)

var (
	Lg       *zap.Logger
	Trans    ut.Translator
	Settings config.ServerConfig
)
