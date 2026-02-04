package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"math"
)

type CalculatorRequest struct {
	FirstNumber  float64 `json:"firstnumber"`
	SecondNumber float64 `json:"secondnumber"`
	Sign         string  `json:"sign"`
}

type CalculatorResponse struct {
	Result float64 `json:"result"`
}

type MathematicanRequest struct {
	Number    float64 `json:"number"`
	Type      string  `json:"type"`
	SecondNum float64 `json:"secondnum"` // second number for Pow(number, secondNum)
}

type MathematicanResponse struct {
	Response float64 `json:"result"`
}

func main() {
	r := gin.Default()

	//just for my js frontend:)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	//information about API
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Information about api": "Simple REST API. 1. Basic calculator functionality 2. Mathematican calculator functionality",
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
		var result float64

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
	r.POST("/getRequest", func(c *gin.Context) {
		var mathRequest MathematicanRequest

		if err := c.ShouldBindJSON(&mathRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bad JSON request",
			})
			return
		}
		var result float64
		//cheking type (sin, cos, tan, sqrt, pow, log, module) // add pow functional tommorow!!! 5 february!
		switch mathRequest.Type {
		case "sin":
			result = math.Sin(mathRequest.Number * math.Pi / 180)
		case "cos":
			result = math.Cos(mathRequest.Number * math.Pi / 180)
		case "tan":
			result = math.Tan(mathRequest.Number * math.Pi / 180)
		case "sqrt":
			result = math.Sqrt(mathRequest.Number)
		case "pow":
			result = math.Pow(mathRequest.Number, mathRequest.SecondNum)
		case "log":
			result = math.Log(mathRequest.Number)
		case "module":
			result = math.Abs(mathRequest.Number)
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

//FIX PLEASE NOT WORKING API TOMMOROW!!!!!!!
