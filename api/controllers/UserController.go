package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasanmarzaq87/goscholling/api/models"
)

func (server *Server) Users(c *gin.Context) {
	// render := render.New(render.Options{
	// 	Layout: "layout",
	// })
	// fmt.Println(r)
	userModel := models.User{}
	users, err := userModel.FindAllUsers(server.DBSQL)
	if err != nil {
		return
	}

	// _ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
	// 	"products": products,
	// })

	// var products []entities.Product
	// database.Instance.Find(&products)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(products)
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": users,
	})
}
