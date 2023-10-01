package handler

import (
	"assignment-2/database"
	"assignment-2/entity"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type orderHandlerImpl struct{}

func NewOrderHandlerImpl() OrderHandler {
	return &orderHandlerImpl{}
}

func (s *orderHandlerImpl) GetAllOrder(c *gin.Context) {
	var db = database.GetDB()

	var orders []entity.Orders
	err := db.Find(&orders).Error

	if err != nil {
		fmt.Println("Error getting order datas : ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (s *orderHandlerImpl) GetOneOrder(c *gin.Context) {
	var db = database.GetDB()

	var orderOne entity.Orders

	err := db.First(&orderOne, "order_id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": orderOne})
}

func (s *orderHandlerImpl) CreateOrder(c *gin.Context) {
	var db = database.GetDB()

	var input entity.Orders
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var existingItem entity.Items
	if err := db.Where("item_code = ?", input.Items[0].Item_Code).First(&existingItem).Error; err != nil {
		errorMessage := "Terdapat item code yang sama"
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
	}

	now := time.Now()

	items := []entity.Items{
		{
			Item_Code:   input.Items[0].Item_Code,
			Description: input.Items[0].Description,
			Quantity:    input.Items[0].Quantity,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	orderInput := entity.Orders{
		OrderedAt:    input.OrderedAt,
		CustomerName: input.CustomerName,
		CreatedAt:    now,
		UpdatedAt:    now,
		Items:        items,
	}
	db.Create(&orderInput)

	c.JSON(http.StatusOK, gin.H{"data": orderInput})
}

func (s *orderHandlerImpl) UpdateOrder(c *gin.Context) {
	var db = database.GetDB()

	var order entity.Orders
	err := db.First(&order, "order_id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	now := time.Now()

	items := []entity.Items{}
	if len(order.Items) > 0 {
		items = append(items, entity.Items{
			Item_Code:   order.Items[0].Item_Code,
			Description: order.Items[0].Description,
			Quantity:    order.Items[0].Quantity,
			UpdatedAt:   now,
		})
	}

	input := entity.Orders{
		CustomerName: order.CustomerName,
		UpdatedAt:    now,
		Items:        items,
	}
	db.Model(&order).Updates(input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func (s *orderHandlerImpl) DeleteOrder(c *gin.Context) {
	var db = database.GetDB()

	var orderDelete entity.Orders
	err := db.First(&orderDelete, "order_id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&orderDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
