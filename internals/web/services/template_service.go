package services

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
)

// TemplateService handles all template rendering
type TemplateService struct {
	templatesPath string
	templates     *template.Template
	mu            sync.RWMutex
}

var once sync.Once
var templateService *TemplateService

// GetTemplateService returns singleton instance of TemplateService
func GetTemplateService() *TemplateService {
	once.Do(func() {
		templateService = &TemplateService{
			templatesPath: "internals/web/views",
		}
		templateService.LoadTemplates()
	})
	return templateService
}

// LoadTemplates loads all template files
func (ts *TemplateService) LoadTemplates() {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	log.Printf("Loading templates from: %s\n", ts.templatesPath)

	// Verify templates path exists
	if info, err := os.Stat(ts.templatesPath); err != nil || !info.IsDir() {
		log.Printf("❌ Error: templates path does not exist: %s\n", ts.templatesPath)
		return
	}

	// Parse base templates first
	basePatterns := []string{
		filepath.Join(ts.templatesPath, "base.tpl"),
		filepath.Join(ts.templatesPath, "sidebar_menu.tpl"),
	}

	tmpl := template.New("")
	var err error
	for _, pattern := range basePatterns {
		if _, err := os.Stat(pattern); err == nil {
			tmpl, err = tmpl.ParseGlob(pattern)
			if err != nil {
				log.Printf("Warning loading base template %s: %v\n", pattern, err)
			}
		}
	}

	// Parse feature-specific templates
	featurePatterns := []string{
		filepath.Join(ts.templatesPath, "auth/*.tpl"),
		filepath.Join(ts.templatesPath, "core/*.tpl"),
		filepath.Join(ts.templatesPath, "roles/*.tpl"),
		filepath.Join(ts.templatesPath, "services/*.tpl"),
		filepath.Join(ts.templatesPath, "settings/*.tpl"),
		filepath.Join(ts.templatesPath, "users/*.tpl"),
	}

	for _, pattern := range featurePatterns {
		matches, _ := filepath.Glob(pattern)
		if len(matches) > 0 {
			tmpl, err = tmpl.ParseGlob(pattern)
			if err != nil {
				log.Printf("Warning loading template %s: %v\n", pattern, err)
			}
		}
	}

	ts.templates = tmpl
	log.Println("✅ Templates loaded successfully")
}

// RenderTemplate renders a template with data
func (ts *TemplateService) RenderTemplate(c *gin.Context, templateName string, data interface{}) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	if ts.templates == nil {
		log.Printf("❌ Error: templates not loaded. Templates path was: %s\n", ts.templatesPath)
		c.JSON(500, gin.H{"error": "templates not loaded", "path": ts.templatesPath})
		return
	}

	err := ts.templates.ExecuteTemplate(c.Writer, templateName, data)
	if err != nil {
		log.Printf("Template rendering error for '%s': %v\n", templateName, err)
		c.JSON(500, gin.H{"error": "failed to render template", "template": templateName, "details": err.Error()})
	}
}

// RenderWithLayout renders a template with base layout
func (ts *TemplateService) RenderWithLayout(c *gin.Context, layoutName string, contentName string, data interface{}) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	if ts.templates == nil {
		log.Printf("❌ Error: templates not loaded. Templates path was: %s\n", ts.templatesPath)
		c.JSON(500, gin.H{"error": "templates not loaded", "path": ts.templatesPath})
		return
	}

	// Merge data with layout context
	contextData := gin.H{}
	if dataMap, ok := data.(gin.H); ok {
		contextData = dataMap
	}
	contextData["Content"] = contentName

	err := ts.templates.ExecuteTemplate(c.Writer, layoutName, contextData)
	if err != nil {
		log.Printf("Template rendering error for '%s': %v\n", layoutName, err)
		c.JSON(500, gin.H{"error": "failed to render template", "template": layoutName, "details": err.Error()})
	}
}
