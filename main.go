package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Constant error messages
const errorMessageWrongParameter = "Provided parameters are not correct. Assure integer numbers!"
const errorMessageWrongAction = "Provided action is not correct. Please use basic math equation"
const errorDivisionByZero = "Division by zero is not allowed in this universe."
const errorRouterProblem = "Router was not able to start"

// Result struct to send as response
type Result struct {
	Action string `json:"action"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Answer int    `json:"answer"`
	Cached bool   `json:"cached"`
}

// Cache object with default expiration and cleanup interval
var c = cache.New(1*time.Minute, 10*time.Minute)

func main() {
	router := gin.Default()
	router.GET("/:action", handleRoute)
	err := router.Run("localhost:80")
	if err != nil {
		log.Println(errorRouterProblem)
	}
}

// Actions that are supported by this API
var actions = map[string]bool{
	"add":      true,
	"subtract": true,
	"multiply": true,
	"divide":   true,
}

// Function to handle incoming requests

func handleRoute(c *gin.Context) {
	// Getting the action from the URL
	action := c.Param("action")
	if !actions[action] {
		c.JSON(http.StatusNotFound, errorMessageWrongAction)
		return
	}
	// Getting the parameters from the URL
	xParam := c.Query("x")
	yParam := c.Query("y")
	x, errX := strconv.Atoi(xParam)
	y, errY := strconv.Atoi(yParam)

	// If the parameters are not valid integers, return error
	if (errX != nil) || (errY != nil) {
		c.JSON(http.StatusBadRequest, errorMessageWrongParameter)
		return
	}

	var result int
	var err error

	// Try to get the result from cache
	cachedResult, existsInCache := getCache(x, y, action)
	switch existsInCache {
	case true:
		result = cachedResult

	default:
		// If result is not in cache, perform the math operation
		result, err = doMath(x, y, action)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, errorDivisionByZero)
			return
		}
		// Store the result in cache
		errSetCache := setCache(x, y, action, result)
		if errSetCache != nil {
			log.Printf("error setting cache: %v", errSetCache)
		}
	}
	// Return the result
	c.JSON(http.StatusOK, Result{action, x, y, result, existsInCache})
}

// Function to set a value in cache
func setCache(x, y int, action string, result int) error {
	cacheKey := generateCacheKey(x, y, action)
	c.Set(cacheKey, result, cache.DefaultExpiration)
	return nil
}

// Function to get a value from cache
func getCache(x, y int, action string) (int, bool) {
	cacheKey := generateCacheKey(x, y, action)
	if v, found := c.Get(cacheKey); found {
		result := v.(int)
		return result, true
	}
	return 0, false
}

// Function to generate cache key
func generateCacheKey(x, y int, action string) string {
	return "x:" + strconv.Itoa(x) + "y:" + strconv.Itoa(y) + "action:" + action
}

// Function to perform math operations
func doMath(x, y int, action string) (int, error) {
	switch action {
	case "add":
		return x + y, nil
	case "subtract":
		return x - y, nil
	case "multiply":
		return x * y, nil
	case "divide":
		// Ensure we don't divide by zero. In reality this could be handled by Math.Inf but is not necessary for this challenge
		if y == 0 {
			return 0, errors.New("division_by_zero")
		}
		return x / y, nil
	}
	return 0, errors.New("invalid action")
}
