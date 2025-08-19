package controller

import (
	"log"
	"net/http"

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

func Index(locations []model.Location) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := Render(c, http.StatusOK, view.Index(locations))	
		HandlerRenderError(err)
	}
}
