package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct{}

type response interface {
	// Message response http success no data [200]
	Message(c *gin.Context, message interface{})
	// Success response http success with data [200]
	Success(c *gin.Context, message interface{}, data interface{})
	// Created response http create [201]
	Created(c *gin.Context, message interface{})
	// Failed response http failed [400]
	Failed(c *gin.Context, message interface{})
	// Forbidden response http forbidden [403]
	Forbidden(c *gin.Context, message interface{})
	// NotFound response http not found [404]
	NotFound(c *gin.Context, message interface{})
	// ValidForm response http request param failed [422]
	ValidForm(c *gin.Context, message interface{})
	// TooManyRequest response http too many request [429]
	TooManyRequest(c *gin.Context, message interface{})
	// Internal response http internal error [500]
	Internal(c *gin.Context, message interface{})
}

//	@method Message
//	@description: response http success no data [200]
//	@receiver r
//	@param c *gin.Context
//	@param message interface{}
func (r *Response) Message(c *gin.Context, message interface{}) {
	if message == nil {
		message = "success"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
	})

	return
}

//	@method Success
//	@description: response http success with data [200]
//	@receiver r
//	@param c *gin.Context
//	@param message interface{}
//	@param data interface{}
func (r *Response) Success(c *gin.Context, message interface{}, data interface{}) {
	if message == nil {
		message = "success"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	})
	return
}

//	@method Created
//	@description: response http create [201]
//	@receiver r
//	@param c *gin.Context
//	@param message interface{}
func (r *Response) Created(c *gin.Context, message interface{}) {
	if message == nil {
		message = "created"
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": message,
	})
	return
}

//	@method Failed
//	@description: response http failed [400]
//	@receiver r
//	@param c *gin.Context
//	@param message interface{}
func (r *Response) Failed(c *gin.Context, message interface{}) {
	if message == nil {
		message = "failed"
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": message,
	})

	return
}

//	@method Forbidden
//	@description: response http forbidden [403]
//	@receiver r
//	@param c *gin.Context
//	@param message interface{}
func (r *Response) Forbidden(c *gin.Context, message interface{}) {
	if message == nil {
		message = "forbidden"
	}

	c.JSON(http.StatusForbidden, gin.H{
		"code":    http.StatusForbidden,
		"message": message,
	})

	return
}

//	@method NotFound
//	@description: response http not found [404]
//	@receiver r
//	@param c *gin.Context
//	@param message interface{}
func (r *Response) NotFound(c *gin.Context, message interface{}) {
	if message == nil {
		message = "not found"
	}

	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"message": message,
	})

	return
}

//	@method ValidForm
//	@description: response http request param failed [422]
//	@receiver r
//	@param c *gin.Context
//	@param message interface{}
func (r *Response) ValidForm(c *gin.Context, message interface{}) {
	if message == nil {
		message = "validate failed"
	}

	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"code":    http.StatusUnprocessableEntity,
		"message": message,
	})

	return
}

//	@method TooManyRequest
//	@description: response http too many request [429]
//	@receiver r
//	@param c *gin.Context
//	@param message interface{}
func (r *Response) TooManyRequest(c *gin.Context, message interface{}) {
	if message == nil {
		message = "too many request"
	}

	c.JSON(http.StatusTooManyRequests, gin.H{
		"code":    http.StatusTooManyRequests,
		"message": message,
	})
	return
}

//	@method Internal
//	@description: response http internal error [500]
//	@receiver r
//	@param c *gin.Context
//	@param message interface{}
func (r *Response) Internal(c *gin.Context, message interface{}) {
	var msg string

	if message == nil {
		msg = "server error"
	} else {
		switch message.(type) {
		case error:
			msg = message.(error).Error()
		case string:
			msg = message.(string)
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": msg,
	})

	return
}
