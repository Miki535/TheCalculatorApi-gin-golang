package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"math"
)

type CalculatorRequest struct {
	FirstNumber  int    `json:"firstnumber"`
	SecondNumber int    `json:"secondnumber"`
	Sign         string `json:"sign"`
}

type CalculatorResponse struct {
	Result int `json:"result"`
}

type MathematicanRequest struct {
	Number    float64 `json:"number"`
	Type      string  `json:"type"`
	secondNum float64 `json:"secondNum"` // second number for Pow(number, secondNum)
}

type MathematicanResponse struct {
	Response float64 `json:"result"`
}

func main() {
	r := gin.Default()

	//information about API
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Information about api": "info", // put info zavtra!
		})
	})

	//get calculator response
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

		response := CalculatorResponse{
			Result: result,
		}

		c.JSON(200, response)
	})

	// get mathematican response
	r.POST("/getResponse", func(c *gin.Context) {
		var mathResponse MathematicanRequest

		if err := c.ShouldBindJSON(&mathResponse); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bad JSON request",
			})
			return
		}
		var result float64
		//cheking type (sin, cos, tan, sqrt, pow, log, module)
		switch mathResponse.Type {
		case "sin":
			result = math.Sin(mathResponse.Number * math.Pi / 180)
		case "cos":
			result = math.Cos(mathResponse.Number * math.Pi / 180)
		case "tan":
			result = math.Tan(mathResponse.Number * math.Pi / 180)
		case "sqrt":
			result = math.Sqrt(mathResponse.Number)
		case "pow":
			result = math.Pow(mathResponse.Number, mathResponse.secondNum)
		case "log":
			result = math.Log(mathResponse.Number)
		case "module":
			result = math.Abs(mathResponse.Number)
		default:
			c.JSON(400, gin.H{
				"error": "unknown type!",
			})
			return
		}

		response := MathematicanResponse{
			Response: result,
		}

		c.JSON(200, response)
	})

	fmt.Println("Running on http://localhost:8080/")
	r.Run()
}
