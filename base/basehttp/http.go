package basehttp

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/goees/yago"
	"github.com/goees/yago/coms/logger"
	"github.com/sirupsen/logrus"
)

type BaseHttp struct {
}

func init() {
	binding.Validator = &defaultValidator{}
}

func (h *BaseHttp) BeforeAction(c *yago.Ctx) error {
	return nil
}

func (h *BaseHttp) AfterAction(c *yago.Ctx) {
	resp, ok := c.GetResponse()
	if !ok {
		return
	}

	if !yago.Config.GetBool("app.http_bizlog_on") {
		return
	}

	params := c.GetString(yago.CtxParamsKey)

	if resp.ErrNo != 0 {
		logger.Ins().Category("http.biz.error").WithFields(logrus.Fields{
			"url":     c.Request.URL.String(),
			"params":  params,
			"header":  c.Request.Header,
			"user_ip": c.ClientIP(),
		}).Error(c.GetError())
	} else {
		logger.Ins().Category("http.biz.info").WithFields(logrus.Fields{
			"url":     c.Request.URL.String(),
			"params":  params,
			"header":  c.Request.Header,
			"user_ip": c.ClientIP(),
			"resp":    resp,
		}).Debug()
	}
}
