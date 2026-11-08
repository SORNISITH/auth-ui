package ux

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/SORNISITH/auth-ui/config"
	"github.com/a-h/templ"
)

// ai generate stlye shit fk off ai slop
var presets = map[string]map[string]string{
	"default": {
		"primary": "#1B0F45", "primary-hover": "#2d1b69", "on-primary": "#ffffff",
		"bg": "#f5f5f7", "surface": "#ffffff", "text": "#111827", "muted": "#6b7280",
		"border": "#d1d5db", "error": "#dc2626", "success": "#16a34a", "warning": "#d97706",
		"link": "#1B0F45", "focus": "#1B0F45",
		"radius": "4px", "font": "system-ui, -apple-system, Segoe UI, Roboto, sans-serif",
		"shadow": "0 1px 1px rgba(0,0,0,.10)",
	},
	"light": {
		"primary": "#2563eb", "primary-hover": "#1d4ed8", "on-primary": "#ffffff",
		"bg": "#f3f4f6", "surface": "#ffffff", "text": "#111827", "muted": "#6b7280",
		"border": "#d1d5db", "error": "#dc2626", "success": "#16a34a", "warning": "#d97706",
		"link": "#2563eb", "focus": "#2563eb",
		"radius": "4px", "font": "system-ui, -apple-system, Segoe UI, Roboto, sans-serif",
		"shadow": "0 1px 1px rgba(0,0,0,.10)",
	},
	"dark": {
		"primary": "#3b82f6", "primary-hover": "#60a5fa", "on-primary": "#ffffff",
		"bg": "#111827", "surface": "#1f2937", "text": "#f3f4f6", "muted": "#9ca3af",
		"border": "#374151", "error": "#f87171", "success": "#4ade80", "warning": "#fbbf24",
		"link": "#60a5fa", "focus": "#3b82f6",
		"radius": "4px", "font": "system-ui, -apple-system, Segoe UI, Roboto, sans-serif",
		"shadow": "0 1px 1px rgba(0,0,0,.40)",
	},
	"navy": {
		"primary": "#1e3a8a", "primary-hover": "#1e40af", "on-primary": "#ffffff",
		"bg": "#f1f5f9", "surface": "#ffffff", "text": "#0f172a", "muted": "#64748b",
		"border": "#cbd5e1", "error": "#dc2626", "success": "#15803d", "warning": "#b45309",
		"link": "#1e3a8a", "focus": "#1e3a8a",
		"radius": "3px", "font": "system-ui, -apple-system, Segoe UI, Roboto, sans-serif",
		"shadow": "0 1px 2px rgba(0,0,0,.12)",
	},
	"slate": {
		"primary": "#334155", "primary-hover": "#1e293b", "on-primary": "#ffffff",
		"bg": "#f8fafc", "surface": "#ffffff", "text": "#0f172a", "muted": "#64748b",
		"border": "#e2e8f0", "error": "#dc2626", "success": "#16a34a", "warning": "#d97706",
		"link": "#334155", "focus": "#334155",
		"radius": "4px", "font": "system-ui, -apple-system, Segoe UI, Roboto, sans-serif",
		"shadow": "0 1px 2px rgba(0,0,0,.08)",
	},
	"green": {
		"primary": "#15803d", "primary-hover": "#166534", "on-primary": "#ffffff",
		"bg": "#f3f4f6", "surface": "#ffffff", "text": "#111827", "muted": "#6b7280",
		"border": "#d1d5db", "error": "#dc2626", "success": "#15803d", "warning": "#d97706",
		"link": "#15803d", "focus": "#15803d",
		"radius": "4px", "font": "system-ui, -apple-system, Segoe UI, Roboto, sans-serif",
		"shadow": "0 1px 3px rgba(0,0,0,.10)",
	},
	"red": {
		"primary": "#b91c1c", "primary-hover": "#991b1b", "on-primary": "#ffffff",
		"bg": "#f3f4f6", "surface": "#ffffff", "text": "#111827", "muted": "#6b7280",
		"border": "#d1d5db", "error": "#b91c1c", "success": "#16a34a", "warning": "#d97706",
		"link": "#b91c1c", "focus": "#b91c1c",
		"radius": "4px", "font": "system-ui, -apple-system, Segoe UI, Roboto, sans-serif",
		"shadow": "0 1px 1px rgba(0,0,0,.10)",
	},
	"classic": {
		"primary": "#1f2937", "primary-hover": "#111827", "on-primary": "#ffffff",
		"bg": "#fafaf9", "surface": "#ffffff", "text": "#1c1917", "muted": "#78716c",
		"border": "#d6d3d1", "error": "#b91c1c", "success": "#15803d", "warning": "#b45309",
		"link": "#1f2937", "focus": "#1f2937",
		"radius": "0px", "font": "Georgia, 'Times New Roman', serif",
		"shadow": "none",
	},
	"flat": {
		"primary": "#2563eb", "primary-hover": "#1d4ed8", "on-primary": "#ffffff",
		"bg": "#ffffff", "surface": "#ffffff", "text": "#111827", "muted": "#6b7280",
		"border": "#d1d5db", "error": "#dc2626", "success": "#16a34a", "warning": "#d97706",
		"link": "#2563eb", "focus": "#2563eb",
		"radius": "2px", "font": "system-ui, -apple-system, Segoe UI, Roboto, sans-serif",
		"shadow": "none",
	},
	"mono": {
		"primary": "#111111", "primary-hover": "#333333", "on-primary": "#ffffff",
		"bg": "#fafafa", "surface": "#ffffff", "text": "#111111", "muted": "#6b7280",
		"border": "#e5e5e5", "error": "#dc2626", "success": "#16a34a", "warning": "#d97706",
		"link": "#111111", "focus": "#111111",
		"radius": "0px", "font": "ui-monospace, SFMono-Regular, Menlo, monospace",
		"shadow": "none",
	},
}
var styleOrder = []string{
	"primary", "primary-hover", "on-primary", "bg", "surface", "text", "muted",
	"border", "error", "success", "warning", "link", "focus",
	"radius", "font", "shadow",
}

var safeValue = regexp.MustCompile(`^[a-zA-Z0-9#%.,()\s'"/_-]+$`) // safe css
func clean(v string) string {
	v = strings.TrimSpace(v)
	if v == "" || !safeValue.MatchString(v) {
		return ""
	}
	return v
}

func styleVars(s *config.Style) templ.SafeCSS {
	if s == nil {
		s = &config.Style{}
	}

	base, ok := presets[strings.ToLower(s.Theme)]
	if !ok {
		base = presets["default"]
	}

	vars := make(map[string]string, len(base))
	for k, v := range base {
		vars[k] = v
	}

	override := map[string]string{
		"primary": s.Primary, "primary-hover": s.PrimaryHover, "on-primary": s.OnPrimary,
		"bg": s.Background, "surface": s.Surface, "text": s.Text, "muted": s.Muted,
		"border": s.Border, "error": s.Error, "success": s.Success, "warning": s.Warning,
		"link": s.Link, "focus": s.Focus,
		"radius": s.Radius, "font": s.Font, "shadow": s.Shadow,
	}
	for k, v := range override {
		if c := clean(v); c != "" {
			vars[k] = c
		}
	}

	var b strings.Builder
	for _, k := range styleOrder {
		fmt.Fprintf(&b, "--%s:%s;", k, vars[k])
	}
	return templ.SafeCSS(b.String())
}

func ThemeNames() []string {
	names := make([]string, 0, len(presets))
	for k := range presets {
		names = append(names, k)
	}
	return names
}

func propForget(defaultRoute string) string {
	return "/forget"
}

// Base CSS (uses the variables from styleVars)
const baseCSS = `
body{margin:0;padding:0;background:var(--bg);color:var(--text);font-family:var(--font);line-height:1.5;}
.ux-page,.ux-page *{box-sizing:border-box;}
.ux-page{display:flex;align-items:center;justify-content:center;width:100%;min-height:100vh;padding:1rem;background:var(--bg);overflow-x:hidden;}
.ux-card{width:100%;max-width:28rem;padding:2rem;background:var(--surface);border:1px solid var(--border);border-radius:var(--radius);box-shadow:var(--shadow);overflow:hidden;}
@media (max-width:639px){
.ux-page{padding:0;}
.ux-card{max-width:none;padding:1.5rem;border-width:0;border-radius:0;box-shadow:none;}
}
.ux-title-row{display:flex;align-items:center;justify-content:flex-start;width:100%;margin-bottom:.5rem;}
.ux-title{margin:0;font-size:1.5rem;font-weight:700;color:var(--text);}
.ux-sub{margin:.25rem 0 1.25rem 0;font-size:.875rem;color:var(--muted);}
.ux-form{display:flex;flex-direction:column;width:100%;}
.ux-field{width:100%;margin-bottom:1rem;}
.ux-label{display:flex;align-items:center;width:100%;padding:.625rem .75rem;background:var(--surface);border:1px solid var(--border);border-radius:var(--radius);}
.ux-label:focus-within{border-color:var(--focus);box-shadow:0 0 0 2px var(--focus);}
.ux-icon{flex-shrink:0;width:1.25rem;height:1.25rem;margin-right:.5rem;opacity:.5;}
.ux-input{width:100%;border:none;outline:none;background:transparent;color:var(--text);font:inherit;font-size:1rem;}
.ux-hint{display:none;margin:.25rem 0 0 0;font-size:.75rem;color:var(--error);}
.ux-label:has(.ux-input:user-invalid) + .ux-hint{display:block;}
.ux-msg{width:100%;margin-top:.5rem;text-align:center;font-size:.875rem;color:var(--error);}
.ux-btn{position:relative;display:flex;align-items:center;justify-content:center;width:100%;padding:.75rem 1rem;border-radius:var(--radius);font:inherit;font-size:1rem;font-weight:600;text-align:center;text-decoration:none;cursor:pointer;transition:background .15s;}
.ux-btn-primary{background:var(--primary);color:var(--on-primary);border:1px solid var(--primary);}
.ux-btn-primary:hover{background:var(--primary-hover);border-color:var(--primary-hover);}
.ux-btn-outline{gap:.5rem;padding:.75rem;background:var(--surface);color:var(--text);border:1px solid var(--border);font-size:.875rem;font-weight:500;}
.ux-btn-outline:hover{border-color:var(--primary);}
.ux-spinner{position:absolute;left:1rem;opacity:0;}
.ux-stack{display:flex;flex-direction:column;gap:.75rem;width:100%;}
.ux-or{display:flex;align-items:center;margin:1.25rem 0;font-size:.875rem;text-align:center;color:var(--muted);}
.ux-or-line{flex:1;border-bottom:1px solid var(--border);}
.ux-or-text{padding:0 .75rem;}
.ux-divider{width:100%;margin:1rem 0;border-bottom:1px solid var(--border);}
.ux-row{display:flex;align-items:center;justify-content:space-between;width:100%;margin:1rem 0 .5rem 0;}
.ux-row-center{display:flex;justify-content:center;width:100%;margin-top:1rem;}
.ux-link{color:var(--link);font-size:.875rem;text-decoration:underline;cursor:pointer;}
.ux-link-strong{font-weight:500;}
.ux-link-block{display:block;padding:.5rem;text-align:center;}
.ux-welcome{display:flex;flex-direction:column;align-items:center;gap:.75rem;width:100%;padding:.5rem 0;}
.ux-welcome-text{margin:0;font-weight:500;color:var(--text);}
.ux-welcome-line{width:100%;margin:.5rem 0;border-bottom:1px solid var(--border);}
.ux-avatar{display:flex;align-items:center;justify-content:center;width:5rem;height:5rem;overflow:hidden;border:1px solid var(--border);border-radius:var(--radius);}
.ux-avatar img{width:100%;height:100%;object-fit:cover;}
`

func styles() templ.Component {
	return templ.Raw("<style>" + baseCSS + "</style>")
}

