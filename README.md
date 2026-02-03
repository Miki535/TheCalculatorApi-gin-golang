# TheCalculatorApi-gin-golang

A simple REST API written in Go using the Gin framework.  
This API provides both basic and advanced mathematical calculations.

---

## 🚀 Features

- Basic arithmetic operations: `+`, `-`, `*`, `/`
- Advanced mathematical functions:
  - `sin`, `cos`, `tan`
  - `sqrt`, `pow`, `log`
  - `mod` (modulo)
- JSON-based requests and responses
- Simple and lightweight implementation

---

## ⚙️ Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Miki535/TheCalculatorApi-gin-golang
Open the project folder in terminal.

Run the server:
  go run main.go

📡 Endpoints
1. POST /getResponse

Used for advanced mathematical calculator functions (sin, cos, tan, sqrt, pow, log, modulo).

2. POST /getRequest

Used for classic calculator functions (addition, subtraction, multiplication, division).

3. GET /

Returns basic information about the API.

Postman screenshots:

Basic calculator (+, -, /, *)
<img width="906" height="546" alt="Postman-calc-ussing-screenshot" src="https://github.com/user-attachments/assets/84ec4fb1-2f7a-4509-8aca-800d62bcebe9" />


Mathematican calculator (sin, cos, tan, sqrt, pow, log, module)
<img width="905" height="478" alt="Postman-math-calc-ussing-screenshot" src="https://github.com/user-attachments/assets/dd0b3f21-8154-4c0d-8814-8f66036e35ac" />


Information about API
<img width="908" height="332" alt="Postman-api-info-screenshot" src="https://github.com/user-attachments/assets/ce007708-c73f-42a8-97a9-99916c0179f8" />



📌 Project Purpose

This project was created for learning and practicing:

Go (Golang)

Gin web framework

REST API design

Mathematical expression processing

It can be used as a backend service for calculator applications or as an educational example.
