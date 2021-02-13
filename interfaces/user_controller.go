package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justpoypoy/go-base/usecases"
)

// A UserController belong to the interface layer.
type UserController struct {
	UserIteractor usecases.UserInteractor
	Logger        usecases.Logger
}

// NewUserController returns the resource of users.
func NewUserController(sqlHandler SQLHandler, logger usecases.Logger) *UserController {
	return &UserController{
		UserIteractor: usecases.UserInteractor{
			UserRepository: &UserRepository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

// Index return response which contain the resource of a user.
func (uc *UserController) Index(c *gin.Context) {
	uc.Logger.LogAccess("%s %s %s\n", c.Request.RemoteAddr, c.Request.Method, c.Request.URL)

	users, err := uc.UserIteractor.Index()
	if err != nil {
		uc.Logger.LogError("%s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"lists": users,
	})
}

// Store return response which contain after create resource of a user.
func (uc *UserController) Store(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user store",
	})
}

// Show return response which contain the specified resource of a user.
func (uc *UserController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user show",
	})
}

// Update return response which contain after update resource of a user.
func (uc *UserController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user update",
	})
}

// Destroy return response which contain after removed resource of a user.
func (uc *UserController) Destroy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user destroy",
	})
}
