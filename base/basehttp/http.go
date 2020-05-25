package basehttp

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/goees/yago"
	"github.com/goees/yago/coms/logger"
	"github.com/goees/yago/libs/validator"
	"github.com/sirupsen/logrus"
)

type BaseHttp struct {
}

func init() {
	binding.Validator = &defaultValidator{}
}

func (h *BaseHttp) Rules() []validator.Rule {
	return nil
}

func (h *BaseHttp) Labels() validator.Label {
	return nil
}

func (h *BaseHttp) BeforeAction(c *yago.Ctx) yago.Err {
	return yago.OK
}

func (h *BaseHttp) AfterAction(c *yago.Ctx) {
	resp, ok := c.GetResponse()
	if !ok {
		return
	}

	if resp.ErrNo != 0 {
		logger.Ins().Category("http.biz.error").WithFields(logrus.Fields{
			"url":     c.Request.URL.Path,
			"params":  c.Keys,
			"header":  c.Request.Header,
			"user_ip": c.ClientIP(),
		}).Error(resp.ErrMsg)

	} else {
		logger.Ins().Category("http.biz.info").WithFields(logrus.Fields{
			"url":     c.Request.URL.Path,
			"params":  c.Keys,
			"header":  c.Request.Header,
			"user_ip": c.ClientIP(),
			"resp":    resp,
		}).Debug()

	}
}
