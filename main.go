package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "math"
    "net/http"
)

const Pi = 3.14

func main() {
    http.HandleFunc("/sum", sumHandler)
    http.HandleFunc("/sub", subHandler)
    http.HandleFunc("/mul", mulHandler)
    http.HandleFunc("/div", divHandler)
    http.HandleFunc("/mod", modHandler)
    http.HandleFunc("/pow", powHandler)
    http.HandleFunc("/sqrt", sqrtHandler)
    http.HandleFunc("/sin", sinHandler)
    http.HandleFunc("/cos", cosHandler)
    http.HandleFunc("/tan", tanHandler)
    http.HandleFunc("/log", logHandler)
    http.HandleFunc("/abs", absHandler)

	// Health check endpoint
	http.HandleFunc("/health", healthCheckHandler)

    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
    handleTwoOperands(w, r, sum)
}

func subHandler(w http.ResponseWriter, r *http.Request) {
    handleTwoOperands(w, r, sub)
}

func mulHandler(w http.ResponseWriter, r *http.Request) {
    handleTwoOperands(w, r, mul)
}

func divHandler(w http.ResponseWriter, r *http.Request) {
    handleTwoOperandsWithError(w, r, div)
}

func modHandler(w http.ResponseWriter, r *http.Request) {
    handleTwoOperandsWithError(w, r, mod)
}

func powHandler(w http.ResponseWriter, r *http.Request) {
    handleTwoOperands(w, r, pow)
}

func sqrtHandler(w http.ResponseWriter, r *http.Request) {
    handleOneOperand(w, r, sqrt)
}

// Handlers for additional functions

func sinHandler(w http.ResponseWriter, r *http.Request) {
    handleOneOperand(w, r, sin)
}

func cosHandler(w http.ResponseWriter, r *http.Request) {
    handleOneOperand(w, r, cos)
}

func tanHandler(w http.ResponseWriter, r *http.Request) {
    handleOneOperand(w, r, tan)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
    handleOneOperandWithError(w, r, log)
}

func absHandler(w http.ResponseWriter, r *http.Request) {
    handleOneOperand(w, r, abs)
}

// Utility functions for handling request and response

func handleTwoOperands(w http.ResponseWriter, r *http.Request, f func(int, int) int) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    a, b, err := getVariables(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result := f(a, b)
    fmt.Fprintf(w, "Result: %d", result)
}

func handleTwoOperandsWithError(w http.ResponseWriter, r *http.Request, f func(int, int) (int, error)) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    a, b, err := getVariables(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := f(a, b)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Result: %d", result)
}

func handleOneOperand(w http.ResponseWriter, r *http.Request, f func(int) int) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    a, _, err := getVariables(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result := f(a)
    fmt.Fprintf(w, "Result: %d", result)
}

func handleOneOperandWithError(w http.ResponseWriter, r *http.Request, f func(int) (int, error)) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    a, _, err := getVariables(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := f(a)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Result: %d", result)
}

// Get Variables from request body

func getVariables(r *http.Request) (int, int, error) {
    vars := make(map[string]int)
    err := json.NewDecoder(r.Body).Decode(&vars)
    if err != nil {
        return 0, 0, errors.New("Invalid request body")
    }

    a, ok := vars["a"]
    if !ok {
        return 0, 0, errors.New("Missing value for 'a'")
    }

    b, ok := vars["b"]
    if !ok {
        return 0, 0, errors.New("Missing value for 'b'")
    }

    return a, b, nil
}

// Math functions

func sum(a, b int) int {
    return a + b
}

func sub(a, b int) int {
    return a - b
}

func mul(a, b int) int {
    return a * b
}

func div(a, b int) (int, error) {
    if b != 0 {
        return a / b, nil
    } else {
        return 0, errors.New("Error: Division by zero")
    }
}

func mod(a, b int) (int, error) {
    if b != 0 {
        return a % b, nil
    } else {
        return 0, errors.New("Error: Division by zero")
    }
}

func pow(a, b int) int {
    return int(math.Pow(float64(a), float64(b)))
}

func sqrt(a int) int {
    return int(math.Sqrt(float64(a)))
}

// Input in degrees

func sin(a int) int {
    return int(math.Sin(float64(a) * math.Pi / 180)) 
}

func cos(a int) int {
    return int(math.Cos(float64(a) * math.Pi / 180)) 
}

func tan(a int) int {
    return int(math.Tan(float64(a) * math.Pi / 180)) 
}

func log(a int) (int, error) {
    if a > 0 {
        return int(math.Log(float64(a))), nil
    } else {
        return 0, errors.New("Error: Logarithm of non-positive number")
    }
}

func abs(a int) int {
    return int(math.Abs(float64(a)))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running")
}