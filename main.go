package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalculatorRequest struct {
	FirstNumber  int    `json:"firstnumber"`
	SecondNumber int    `json:"secondnumber"`
	Sign         string `json:"sign"`
}

type CalculatorResponse struct {
	Result int `json:"result"`
}

func main() {
	r := gin.Default()

	r.POST("/getRequest", func(c *gin.Context) {
		var calcRequest CalculatorRequest

		if err := c.ShouldBindJSON(&calcRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bad JSON request",
			})
			return
		}
		var result int

		switch calcRequest.Sign {
		case "+":
			result = calcRequest.FirstNumber + calcRequest.SecondNumber
		case "-":
			result = calcRequest.FirstNumber - calcRequest.SecondNumber
		case "*":
			result = calcRequest.FirstNumber * calcRequest.SecondNumber
		case "/":
			if calcRequest.SecondNumber == 0 {
				c.JSON(400, gin.H{"error": "You cannot divide by zero"})
				return
			} else {
				result = calcRequest.FirstNumber / calcRequest.SecondNumber
			}
		default:
			c.JSON(400, gin.H{
				"error": "unknown sign!",
			})
			return
		}

		resp := CalculatorResponse{
			Result: result,
		}

		c.JSON(200, resp)
	})

	r.Run()
}
