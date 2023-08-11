package rest

import (
	"loanapp"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleUserFindById godoc
// @Summary Get single user by ID
// @Description Returns the user whose ID matches the given ID.
// @Tags users
// @Accept json
// @Produce json
// @Param                  id      path            string  true    "search by id"
// @Success 200
// @Router /users/{id} [get]
func (h *Handler) handleUserFindById(c *gin.Context) {

	id := c.Param("id")

	user, _ := h.UserService.FindUserByID(c, id)

	c.IndentedJSON(http.StatusOK, user)
}

// handleUserFind godoc
// @Summary user ping
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Router /users [get]
func (h *Handler) handleUserFindAll(c *gin.Context) {
	users, _, _ := h.UserService.FindUsers(c, loanapp.UserFilter{})
	c.IndentedJSON(http.StatusOK, users)

}

// handleUserCreate godoc
// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Success 200	{object}	loanapp.User
// @Param                  user        body            loanapp.User true    "user JSON"
// @Router /users [POST]
func (h *Handler) handleUserCreate(c *gin.Context) {

	var newUser loanapp.User

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, _ := h.UserService.CreateUser(c, &newUser)
	c.IndentedJSON(http.StatusCreated, user)

}

// handleUserUpdate godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param                  id      path            string  true    "user id"
// @Param                  userUpdate        body            loanapp.UserUpdate true    "user update JSON"
// @Success 200
// @Router /users/{id} [PUT]
func (h *Handler) handleUserUpdate(c *gin.Context) {
	var userUpdate loanapp.UserUpdate
	id := c.Param("id")

	if err := c.BindJSON(&userUpdate); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, _ := h.UserService.UpdateUser(c, id, userUpdate)
	c.IndentedJSON(http.StatusCreated, user)

}

// handleUserDelete godoc
// @Summary Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param                  id      path            string  true    "user id"
// @Success 200
// @Router /users/{id} [DELETE]
func (h *Handler) handleUserDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.UserService.DeleteUser(c, id); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(http.StatusOK, 0)
}

// handleUserPing godoc
// @Summary user ping
// @Schemes
// @Description do ping
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /users/ping [get]
func (h *Handler) handleUserPing(c *gin.Context) {

	result, _ := h.UserService.Ping(c, "hello world")

	c.JSON(http.StatusOK, result)
}
