package handlers

import (
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Group Request format
type CreateGroupRequest struct {
	Name      string   `json:"name"`
	CreatedBy string   `json:"created_by"`
	Members   []string `json:"members"`
}

// Split format
type SplitInput struct {
	UserId string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

// Expense Request format
type CreateExpenseRequest struct {
	GroupId     string       `json:"group_id"`
	PaidBy      string       `json:"paid_by"`
	Amount      float64      `json:"amount"`
	Description string       `json:"description"`
	Splits      []SplitInput `json:"splits"`
}

func CreateGroup(ctx *gin.Context) {
	var grpReq CreateGroupRequest

	// Bind to struct
	if err := ctx.ShouldBindJSON(&grpReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// call the service
	grpId, err := service.CreateGroup(grpReq.Name, grpReq.CreatedBy, grpReq.Members)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create group"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"group_id": grpId})
}

func GetBalances(ctx *gin.Context) {
	grpId := ctx.Param("group_id")

	balances, err := service.CalculateBalances(grpId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"balances": balances,
	})
}

func AddExpense(ctx *gin.Context) {
	var expReq CreateExpenseRequest

	// Bind to struct
	if err := ctx.ShouldBindJSON(&expReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Convert Slice -> Map
	splitMap := make(map[string]float64)
	for _, s := range expReq.Splits {
		splitMap[s.UserId] = s.Amount
	}

	// Get the Expense Id
	_, err := service.AddExpense(expReq.GroupId, expReq.PaidBy, expReq.Amount, expReq.Description, splitMap)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "expense added"})
}

func GetSettlements(ctx *gin.Context) {

	groupID := ctx.Param("group_id")

	balances, err := service.CalculateBalances(groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	settlements := service.ClarifyDebts(balances)

	ctx.JSON(http.StatusOK, settlements)
}
