package link

import (
	"main-service/internal/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for link operations
type Handler struct {
	service *Service
}

// NewHandler creates new Handler instance
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Create godoc
// @Summary Create pseudo link
// @Description Create pseudo name for your url-link
// @Tags Link
// @Accept json
// @Produce json
// @Param origin_link path string true "URL-Link"
// @Router /link/create [post]
func (h *Handler) Create(ctx *gin.Context) {
	var jsonBody struct {
		OriginLink string `json:"origin_link" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&jsonBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format: " + err.Error(),
		})
		return
	}

	pseudoLink, err := h.service.Create(jsonBody.OriginLink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create link: " + err.Error(),
		})
		return
	}
	link := models.Link{
		OriginLink: jsonBody.OriginLink,
		PseudoLink: pseudoLink,
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Link created successfully",
		"data":    link,
	})
}

// GetPseudo generates pseudo link for testing purposes.
// Get godoc
// @Summary Get pseudo link
// @Description Return your struct link on url
// @Tags Link
// @Accept json
// @Produce json
// @Param origin_link path string true "URL-Link"
// @Router /link/get [get]
func (h *Handler) GetPseudo(ctx *gin.Context) {
	originLink := ctx.Query("origin_link")
	if originLink == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "origin_link parameter is required",
		})
		return
	}
	link, err := h.service.GetLink(originLink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get link: " + err.Error(),
		})
		return
	}

	if link == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Link not found",
		})
		return
	}

	// Возвращаем найденную ссылку
	ctx.JSON(http.StatusOK, gin.H{
		"data": link,
	})
}

// Redirect redirects short URL to original link.
// Redirect godoc
// @Summary Redirect url 302
// @Description Redirect on url
// @Tags Link
// @Accept json
// @Produce json
// @Param shortID path string true "URL-Link"
// @Router /{shortID}} [get]
func (h *Handler) Redirect(ctx *gin.Context) {
	linkID := ctx.Param("shortID")
	if linkID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "linkID parameter is required",
		})
		return
	}
	link, err := h.service.GetLink(ctx.Request.URL.Path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get link: " + err.Error(),
		})
		return
	}

	if link == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Link not found",
		})
		return
	}
	ctx.Redirect(302, link.OriginLink)
}

// ParselinkID extracts link ID from URL.
func ParselinkID(link string) string {
	return link[strings.LastIndex(link, "/")+1:]
}

// Delete godoc
// @Summary Delete link
// @Description Delete your struct link on url
// @Tags Link
// @Accept json
// @Produce json
// @Param origin_link path string true "URL-Link"
// @Router /link/delete [delete]
func (h *Handler) Delete(ctx *gin.Context) {
	originLink := ctx.Query("origin_link")
	if originLink == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "origin_link parameter is required",
		})
		return
	}
	if err := h.service.DeleteLink(originLink); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete link: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Link deleted successfully",
	})

}
