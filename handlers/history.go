package handlers

import (
	"database/sql"

	"github.com/Mosteben/calculator-api/db"
	"github.com/gin-gonic/gin"
)

type HistoryItem struct {
	A         int     `json:"a"`
	B         int     `json:"b"`
	Operation string  `json:"operation"`
	Result    float64 `json:"result"`
}

func GetHistory(c *gin.Context) {

	operation := c.Query("operation")

	var rows *sql.Rows
	var err error

	if operation != "" {
		rows, err = db.DB.Query(
			"SELECT a, b, operation, result FROM history WHERE operation = $1",
			operation,
		)
	} else {
		rows, err = db.DB.Query(
			"SELECT a, b, operation, result FROM history",
		)
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var history []HistoryItem

	for rows.Next() {

		var h HistoryItem

		rows.Scan(&h.A, &h.B, &h.Operation, &h.Result)

		history = append(history, h)
	}

	c.JSON(200, gin.H{
		"history": history,
	})
}