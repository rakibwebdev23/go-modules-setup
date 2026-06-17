package main

import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  // Create a Gin router with default middleware (logger and recovery)
  r := gin.Default()

  // Define a simple GET endpoint
  r.GET("/ping", func(c *gin.Context) {
    // Return JSON response
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  // Start server on port 8080 (default)
  // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
  if err := r.Run(); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
} //go mod tidy ( to clean up go.mod and go.sum after edits to app/models/user.go)



// package main

// import (
// 	"github.com/olekukonko/tablewriter"
// 	"os"
// )

// func main() {
// 	data := [][]string{
// 		{"Package", "Version", "Status"},
// 		{"tablewriter", "v0.0.5", "legacy"},
// 		{"tablewriter", "v1.1.4", "latest"},
// 	}

// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader(data[0])
// 	for _, v := range data[1:] {
// 		table.Append(v)
// 	}
// 	table.Render()
// }