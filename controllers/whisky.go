package controllers

import (
	"github.com/Massad/gin-boilerplate/forms"
	"github.com/gin-gonic/gin"
)

// WhiskyController ...
type WhiskyController struct{}

// var WhiskyModel = new(models.)
var WhiskyForm = new(forms.WhiskyForm)

/*
// Create ...

	func (ctrl WhiskyController) Create(c *gin.Context) {
		userID := getUserID(c)

		var form forms.CreateWhiskyForm

		if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
			message := WhiskyForm.Create(validationErr)
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
			return
		}

		id, err := WhiskyModel.Create(userID, form)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Whisky could not be created"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Whisky created", "id": id})
	}

// All ...

	func (ctrl WhiskyController) All(c *gin.Context) {
		userID := getUserID(c)

		results, err := WhiskyModel.All(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get Whiskys"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"results": results})
	}

// One ...

	func (ctrl WhiskyController) One(c *gin.Context) {
		userID := getUserID(c)

		id := c.Param("id")

		getID, err := strconv.ParseInt(id, 10, 64)
		if getID == 0 || err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
			return
		}

		data, err := WhiskyModel.One(userID, getID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Whisky not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}

// Update ...

	func (ctrl WhiskyController) Update(c *gin.Context) {
		userID := getUserID(c)

		id := c.Param("id")

		getID, err := strconv.ParseInt(id, 10, 64)
		if getID == 0 || err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
			return
		}

		var form forms.CreateWhiskyForm

		if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
			message := WhiskyForm.Create(validationErr)
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
			return
		}

		err = WhiskyModel.Update(userID, getID, form)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Whisky could not be updated"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Whisky updated"})
	}

// Delete ...

	func (ctrl WhiskyController) Delete(c *gin.Context) {
		userID := getUserID(c)

		id := c.Param("id")

		getID, err := strconv.ParseInt(id, 10, 64)
		if getID == 0 || err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
			return
		}

		err = WhiskyModel.Delete(userID, getID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Whisky could not be deleted"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Whisky deleted"})
	}
*/
func (ctrl WhiskyController) Pong(c *gin.Context) {
	c.String(200, "pong")
	//c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
