# nz-authui
module login UI   for [Gin](https://github.com/gin-gonic/gin), 
built with [templ](https://templ.guide) and [htmx](https://htmx.org).

## Requirements
- Go 1.21+ 
- optional : htmx  if you want butter spa 
- optional  : [templ](https://templ.guide)
- templ generate   // if you edit the templ file

## Quick start
```go
package main

import (
	"github.com/gin-gonic/gin"
	"nz.com/authui"
	"nz.com/authui/config"
)

func main() {
	router := gin.Default()

	a := authui.Module{ Router: router }
	a.Init() 

	router.Run(":8080")
}
```

## Configuration Example `DefaultPath()`:
### `config.Path`
| Field               | Purpose                                  | Default             |
|---------------------|------------------------------------------|---------------------|
| `LoginPage`         | URL of the login page                    | `/login`            |
| `RegisterPage`      | URL of the register page                 | `/register`         |
| `LoginAdminPage`    | URL of the admin login page              | `/login/admin`      |
| `WelcomeBack`       | url                                      | `/welcome`          |
| `LoginBackend`      | `POST` endpoint for login                | `/api/login`        |
| `LoginAdminBackend` | `POST` endpoint for admin login          | `/api/login/admin`  |
| `RegisterBackend`   | `POST` endpoint for register             | `/api/register`     |

| `Style`             | Theme and style overrides (`*Style`)     | `nil` (default theme) |

## Pre-define Themes
Set with `Style.Theme`. Names are case-insensitive. Unknown names fall back to `default`.

| Theme     | Look                                         |
|-----------|----------------------------------------------|
| `default` | Light gray page, dark navy accent            |
| `light`   | Light gray page, blue accent                 |
| `dark`    | Dark gray (not pure black), blue accent      |
| `navy`    | Light page, deep navy accent                 |
| `slate`   | Understated gray accent                      |
| `green`   | Green accent                                 |
| `red`     | Red accent                                   |
| `classic` | Serif font, square corners, no shadow        |
| `flat`    | Border only, no shadow                       |
| `mono`    | Black and white, monospace, no shadow        |

### `config.Style`
```go
type Style struct {
	Theme string // preset name, see "Themes"
	// Colors (any valid CSS color)
	Primary, PrimaryHover, OnPrimary string
	Background, Surface              string
	Text, Muted, Border              string
	Error, Success, Warning          string
	Link, Focus                      string
	Radius string // e.g. "4px"
	Font   string // e.g. "system-ui, sans-serif"
	Shadow string // CSS box-shadow value, or "none"
}
```
---
## CSS variables
The theme is applied as CSS variables on `<body>`, so you can use them in your own CSS:
```
--primary  --primary-hover  --on-primary
--bg       --surface        --text       --muted    --border
--error    --success        --warning    --link     --focus
--radius   --font           --shadow
```

```css
.my-banner {
	background: var(--surface);
	border: 1px solid var(--border);
	border-radius: var(--radius);
}
```

---
## Security notes
- Add **CSRF protection** to the forms before production use.
- Always validate and rate-limit the login and register endpoints on the server. The browser-side `required`, `minlength` and `pattern` checks are only for convenience.
- Style values are filtered, but still avoid passing untrusted user input into `Style`.
- Serve over HTTPS and set secure cookie flags in your login handler.

