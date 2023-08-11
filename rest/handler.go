package rest

import (
	"loanapp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService loanapp.UserService
}

func NewHandler(r *gin.Engine) *Handler {

	h := &Handler{}

	r.GET("/example/helloworld", h.Helloworld)
	r.GET("/example/ping", h.GetPing)

	r.GET("/users", h.handleUserFindAll)
	r.GET("/users/:id", h.handleUserFindById)
	r.POST("/users", h.handleUserCreate)
	r.PUT("/users/:id", h.handleUserUpdate)
	r.DELETE("/users/:id", h.handleUserDelete)

	r.GET("/users/ping", h.handleUserPing)

	return h
}

// GetPing  godoc
//
//	@Summary                Ping example
//	@Description    		Do a ping
//	@Tags					example
//	@Accept 				json
//	@Produce                json
//	@Router                 /example/ping  [get]
func (h *Handler) GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (h *Handler) Helloworld(c *gin.Context) {
	c.JSON(200, "hello world")
}
