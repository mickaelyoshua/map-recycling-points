package controller

import (
	"log"
	"net/http"
	"sort"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/map-recycling-points/model"
	"github.com/mickaelyoshua/map-recycling-points/view"
)

func Render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}
func HandlerRenderError(err error) {
	if err != nil {
		log.Printf("Error rendering template: %v\n", err)
	}
}

func Index(allLocations []model.Location) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract unique categories
		categoriesMap := make(map[string]bool)
		for _, loc := range allLocations {
			if loc.Category != "" {
				categoriesMap[loc.Category] = true
			}
		}
		var categories []string
		for cat := range categoriesMap {
			categories = append(categories, cat)
		}
		sort.Strings(categories)

		activeCategory := c.Query("category")
		if activeCategory == "" {
			activeCategory = "all" // Default to "all" if no category is specified
		}

		var filteredLocations []model.Location
		if activeCategory == "all" {
			filteredLocations = allLocations
		} else {
			for _, loc := range allLocations {
				if loc.Category == activeCategory {
					filteredLocations = append(filteredLocations, loc)
				}
			}
		}

		err := Render(c, http.StatusOK, view.Index(filteredLocations, categories, activeCategory))
		HandlerRenderError(err)
	}
}
