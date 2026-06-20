package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/pucora/lura/v2/proxy"
)

func Render(c *gin.Context, response *proxy.Response) {
	status := c.Writer.Status()
	if response == nil {
		c.String(status, "")
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	_ = json.NewEncoder(c.Writer).Encode(response.Data)
}
