package config
type Style struct {
	Theme string // "default", "dark", "ocean", "forest", "sunset", "rose", "mono"
	Primary      string
	PrimaryHover string
	OnPrimary    string
	Background   string
	Surface      string
	Text         string
	Muted        string
	Border       string
	Error        string
	Success      string
	Warning      string
	Link         string
	Focus        string
	// Shape / typography
	Radius string // e.g. "8px"
	Font   string // e.g. "Inter, system-ui, sans-serif"
	Shadow string // box-shadow value
}

type Path struct {

		Title string
		LoginPage string
		LoginAdminPage string
		WelcomePage string
		RegisterPage string
		ForgetPage string
		
		LoginBackend string
		WelcomeBackend string
		LoginAdminBackend string
		RegisterBackend string
		ForgetBackend string
		ResetBackend string
		Style *Style
		TokenReset string
}

func (c *Path) DefaultPath() {
	if c.Title == "" {
		c.Title = "Authorization"
	}
	if c.LoginPage == "" {
		c.LoginPage = "/login"
	}
	if c.LoginAdminPage == "" {
		c.LoginAdminPage = "/login/admin"
	}
	if c.RegisterPage == "" {
		c.RegisterPage = "/register"
	}
	if c.ForgetPage == "" {
		c.ForgetPage = "/forget"
	}
	if c.LoginBackend == "" {
		c.LoginBackend = "/api/login"
	}
	if c.LoginAdminBackend == "" {
		c.LoginAdminBackend = "/api/admin"
	}
	if c.RegisterBackend == "" {
		c.RegisterBackend = "/api/register"
	}
	if c.ForgetBackend == "" {
		c.ForgetBackend = "/api/forget"
	}

	if c.WelcomeBackend == "" {
		c.WelcomeBackend = "/api/welcome"
	}
	if c.WelcomePage == "" {
		c.WelcomePage = "/welcome"
	}
}
