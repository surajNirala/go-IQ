package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// Item represents the data structure
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Global variables to store items in memory and manage IDs
var (
	items  = make(map[int]Item)
	nextID = 1
	mutex  = &sync.Mutex{}
)

func main() {
	r := gin.Default()

	// Routes
	r.GET("/items", getItems)
	r.GET("/items/:id", getItem)
	r.POST("/items", createItem)
	r.PUT("/items/:id", updateItem)
	r.DELETE("/items/:id", deleteItem)

	// Run the server
	r.Run(":8080")
}

// Get all items
func getItems(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	// Convert map to slice
	result := make([]Item, 0, len(items))
	for _, item := range items {
		result = append(result, item)
	}

	c.JSON(http.StatusOK, result)
}

// Get an item by ID
func getItem(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	item, exists := items[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Create a new item
func createItem(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	newItem.ID = nextID
	nextID++
	items[newItem.ID] = newItem

	c.JSON(http.StatusCreated, newItem)
}

// Update an item by ID
func updateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var updatedItem Item
	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	item, exists := items[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	updatedItem.ID = item.ID
	items[id] = updatedItem

	c.JSON(http.StatusOK, updatedItem)
}

// Delete an item by ID
func deleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	_, exists := items[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	delete(items, id)
	c.Status(http.StatusNoContent)
}
