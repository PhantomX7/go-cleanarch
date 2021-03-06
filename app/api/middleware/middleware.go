package middleware

import (
	"github.com/PhantomX7/go-cleanarch/util/errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"strings"
)

type Config struct {
	// put middleware config here
}

type Middleware struct {
	config Config
}

func New(cfg Config) *Middleware {
	return &Middleware{
		config: cfg,
	}
}

func (m *Middleware) ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// Only run if there are some errors to handle
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				// Find out what type of error it is

				switch e.Type {
				case gin.ErrorTypePublic:
					// Only output public errors if nothing has been written yet
					if !c.Writer.Written() {
						// check if it is part of custom error
						if err, ok := e.Err.(errors.CustomError); ok {
							c.JSON(err.HTTPCode, gin.H{"error": err.Message})
						} else {
							c.JSON(c.Writer.Status(), gin.H{"error": e.Error()})
						}

					}
				case gin.ErrorTypeBind:
					errs := e.Err.(validator.ValidationErrors)
					list := make(map[string]string)
					for _, err := range errs {
						list[strings.ToLower(err.Field)] = validationErrorToText(err)
					}

					// Make sure we maintain the preset response status
					status := http.StatusBadRequest
					if c.Writer.Status() != http.StatusOK {
						status = c.Writer.Status()
					}
					c.JSON(status, gin.H{"errors": list})

					// below here is custom error
				case errors.ErrorTypeUnprocessableEntity:
					response := errors.ErrUnprocessableEntity
					response.Message = "something went wrong when processing data"
					c.JSON(response.HTTPCode, response)
				default:
					// Log all other errors
					//rollbar.RequestError(rollbar.ERR, c.Request, e.Err)
				}

			}
			// If there was no public or bind error, display default 500 message
			if !c.Writer.Written() {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": "something went wrong"})
			}
		}
	}
}
