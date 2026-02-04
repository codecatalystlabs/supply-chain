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
	templates     map[string]*template.Template // Map of template name to template instance
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
// Each page template is parsed separately with base.tpl to maintain isolated content blocks
func (ts *TemplateService) LoadTemplates() {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	log.Printf("Loading templates from: %s\n", ts.templatesPath)

	// Verify templates path exists
	if info, err := os.Stat(ts.templatesPath); err != nil || !info.IsDir() {
		log.Printf("❌ Error: templates path does not exist: %s\n", ts.templatesPath)
		return
	}

	basePath := filepath.Join(ts.templatesPath, "base.tpl")
	sidebarPath := filepath.Join(ts.templatesPath, "sidebar_menu.tpl")

	// Get all .tpl files
	allTPLFiles := filepath.Join(ts.templatesPath, "*.tpl")
	matches, _ := filepath.Glob(allTPLFiles)

	// Create a map to store each template separately
	ts.templates = make(map[string]*template.Template)

	// Parse each child template separately with base.tpl
	// This keeps each template's "content" block isolated
	for _, tplFile := range matches {
		filename := filepath.Base(tplFile)

		// Skip base and sidebar as they'll be included with each page
		if filename == "base.tpl" || filename == "sidebar_menu.tpl" {
			continue
		}

		// Create a NEW template instance for each page
		// Parse base + sidebar + child template together
		tmpl, err := template.New(filename).ParseFiles(basePath, sidebarPath, tplFile)
		if err != nil {
			log.Printf("Error loading template %s: %v\n", filename, err)
		} else {
			ts.templates[filename] = tmpl
			log.Printf("✅ Loaded %s\n", filename)
		}
	}

	log.Printf("✅ All templates loaded successfully")
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

	// Get the specific template instance for this page
	tmpl, exists := ts.templates[templateName]
	if !exists {
		log.Printf("❌ Error: template '%s' not found\n", templateName)
		c.JSON(500, gin.H{"error": "template not found", "template": templateName})
		return
	}

	// Execute the template (it will call base.tpl internally)
	err := tmpl.ExecuteTemplate(c.Writer, templateName, data)
	if err != nil {
		log.Printf("Template rendering error for '%s': %v\n", templateName, err)
		c.JSON(500, gin.H{"error": "failed to render template", "template": templateName, "details": err.Error()})
	}
}

// RenderWithLayout is deprecated - use RenderTemplate with Content field in ViewData instead
func (ts *TemplateService) RenderWithLayout(c *gin.Context, layoutName string, contentName string, data interface{}) {
	ts.RenderTemplate(c, layoutName, data)
}
