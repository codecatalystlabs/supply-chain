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

// findTemplatesPath finds the templates directory by checking multiple possible locations
func findTemplatesPath() string {
	// Possible paths relative to different working directories
	possiblePaths := []string{
		"internals/web/views",                    // From project root
		"../internals/web/views",                 // From cmd/server
		"../../internals/web/views",              // From cmd/server/server
		"./internals/web/views",                 // Current directory
	}

	// Get current working directory
	wd, err := os.Getwd()
	if err == nil {
		// Also try absolute paths
		possiblePaths = append(possiblePaths,
			filepath.Join(wd, "internals/web/views"),
			filepath.Join(wd, "../internals/web/views"),
			filepath.Join(wd, "../../internals/web/views"),
		)
	}

	// Check each path
	for _, path := range possiblePaths {
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			absPath, err := filepath.Abs(path)
			if err == nil {
				return absPath
			}
			return path
		}
	}

	// Default fallback
	return "internals/web/views"
}

// GetTemplateService returns singleton instance of TemplateService
func GetTemplateService() *TemplateService {
	once.Do(func() {
		templatesPath := findTemplatesPath()
		templateService = &TemplateService{
			templatesPath: templatesPath,
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

	// Get all .tpl files recursively (including subdirectories)
	var matches []string
	err := filepath.Walk(ts.templatesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".tpl" {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		log.Printf("Error walking templates directory: %v\n", err)
		return
	}

	// Create a map to store each template separately
	// Map structure: key -> *template.Template
	// Also store template name mapping: key -> actual template name to execute
	ts.templates = make(map[string]*template.Template)

	// Parse each child template separately with base.tpl
	// This keeps each template's "content" block isolated
	for _, tplFile := range matches {
		filename := filepath.Base(tplFile)

		// Skip base and sidebar as they'll be included with each page
		if filename == "base.tpl" || filename == "sidebar_menu.tpl" {
			continue
		}

		// Get relative path from templates directory for subdirectory support
		relPath, err := filepath.Rel(ts.templatesPath, tplFile)
		if err != nil {
			relPath = filename // Fallback to filename if relative path fails
		}

		// Create a NEW template instance for each page
		// Parse base + sidebar + child template together
		// Use filename as the template name (this is what ExecuteTemplate will use)
		tmpl, err := template.New(filename).ParseFiles(basePath, sidebarPath, tplFile)
		if err != nil {
			log.Printf("Error loading template %s: %v\n", relPath, err)
		} else {
			// Store by filename (for backward compatibility) and by relative path
			// Both keys point to the same template, which uses 'filename' as its name
			ts.templates[filename] = tmpl
			if relPath != filename {
				ts.templates[relPath] = tmpl
			}
			log.Printf("✅ Loaded %s (template name: %s)\n", relPath, filename)
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

	// Get the actual template name (filename) for execution
	// Templates are created with filename as the name, regardless of the key used to store them
	actualTemplateName := filepath.Base(templateName)
	if actualTemplateName == "." || actualTemplateName == "" {
		actualTemplateName = templateName
	}

	// Execute the template (it will call base.tpl internally)
	// Use the actual template name (filename) that was used when creating the template
	err := tmpl.ExecuteTemplate(c.Writer, actualTemplateName, data)
	if err != nil {
		log.Printf("Template rendering error for '%s': %v\n", templateName, err)
		c.JSON(500, gin.H{"error": "failed to render template", "template": templateName, "details": err.Error()})
	}
}

// RenderWithLayout is deprecated - use RenderTemplate with Content field in ViewData instead
func (ts *TemplateService) RenderWithLayout(c *gin.Context, layoutName string, contentName string, data interface{}) {
	ts.RenderTemplate(c, layoutName, data)
}
