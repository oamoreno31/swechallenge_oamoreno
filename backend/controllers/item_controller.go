package controllers

import (
	"fmt"
	"net/http"
	"sweapi/config"
	"sweapi/models"

	"github.com/gin-gonic/gin"
)

// GetProducts - Obtener todos los productos
func GetProducts(c *gin.Context) {
	FN := "GetProducts"
	rows, err := config.DB.Query(
		c,
		"SELECT id, ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time FROM items",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []models.Item

	for rows.Next() {
		var product models.Item
		err := rows.Scan(
			&product.ID,
			&product.Ticker,
			&product.Target_from,
			&product.Target_to,
			&product.Company,
			&product.Action,
			&product.Brokerage,
			&product.Rating_from,
			&product.Rating_to,
			&product.Time,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println("Error ", FN, " ", err.Error())
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

// GetProduct - Obtener un producto por ID
func GetProduct(c *gin.Context) {
	FN := "GetProduct"
	id := c.Param("id")

	var item models.Item
	err := config.DB.QueryRow(
		c,
		"SELECT id, ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time FROM items WHERE id = $1",
		id,
	).Scan(
		&item.ID,
		&item.Ticker,
		&item.Target_from,
		&item.Target_to,
		&item.Company,
		&item.Action,
		&item.Brokerage,
		&item.Rating_from,
		&item.Rating_to,
		&item.Time,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		fmt.Println("Error ", FN, " ", err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// CreateProduct - Crear un nuevo producto
func CreateProduct(c *gin.Context) {
	FN := "CreateProduct"
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Error ", FN, " ", err.Error())
		return
	}

	//var idInt int
	err := config.DB.QueryRow(
		c,
		"INSERT INTO items (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, time",
		item.Ticker,
		item.Target_from,
		item.Target_to,
		item.Company,
		item.Action,
		item.Brokerage,
		item.Rating_from,
		item.Rating_to,
		item.Time,
	).Scan(&item.ID, &item.Time)

	// Convert int to string
	//item.ID = fmt.Sprintf("%d", idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Error ", FN, " ", err.Error())
		return
	}

	c.JSON(http.StatusCreated, item)
}

// UpdateProduct - Actualizar un producto
func UpdateProduct(c *gin.Context) {
	FN := "UpdateProduct"
	id := c.Param("id")
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := config.DB.Exec(
		c,
		"UPDATE items SET ticker = $1, target_from = $2, target_to = $3, company = $4, action = $5, brokerage = $6, rating_from = $7, rating_to = $8, time = $9 WHERE id = $10",
		item.Ticker,
		item.Target_from,
		item.Target_to,
		item.Company,
		item.Action,
		item.Brokerage,
		item.Rating_from,
		item.Rating_to,
		item.Time,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Error ", FN, " ", err.Error())
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		fmt.Println("Error ", FN, " ", err.Error())
		return
	}

	// Keep the ID consistent with the model type (string)
	//item.ID = id
	c.JSON(http.StatusOK, item)
}

// DeleteProduct - Eliminar un producto
func DeleteProduct(c *gin.Context) {
	FN := "DeleteProduct"
	id := c.Param("id")

	result, err := config.DB.Exec(c, "DELETE FROM items WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		fmt.Println("Error ", FN, " ", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado exitosamente"})
}
