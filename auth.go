package authui

import (
	"github.com/gin-gonic/gin"
	"github.com/a-h/templ"
	"nz.com/authui/config"
	"nz.com/authui/ux"
)
type Config struct {
		Router 	gin.IRouter
		Path  config.Path
		Style config.Style
}
func (c *Config) Init()  {
	p := c.Path
	p.DefaultPath() 
	p.Style = &c.Style 
	c.Router.GET(p.LoginPage, page(&p, ux.Page_Login))
	c.Router.GET(p.RegisterPage, page(&p, ux.Page_Register))
	c.Router.GET(p.LoginAdminPage, page(&p, ux.Page_LoginAdmin))
	c.Router.GET(p.WelcomePage, page(&p, ux.Page_Welcome))
	c.Router.GET(p.ForgetPage, page(&p, ux.Page_Forget))
}

func page(p *config.Path, comp func(*config.Path) templ.Component) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		if err := comp(p).Render(c.Request.Context(), c.Writer); err != nil {
			c.Error(err)
		}
	}

}
