package test

import (
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nunonano/hacktiv-final-project/model/entity"
	repository "github.com/nunonano/hacktiv-final-project/repository/order"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// type OrderRepositoryMock struct{}

// MODELS
type Orders []entity.Order

// CONTROLLER
type OrderController struct{}

// REPOSITORY
type OrderRepository struct{}
type OrderRepositoryInterface interface {
	GetAllOrder() Orders
	GetOrder(id int) entity.Order
}

type OrderRepositoryMock struct {
	Mock mock.Mock
}

type OrderServiceMock struct {
	OrderRepository repository.OrderRepository
}

func (ctrl OrderController) GetAllOrder(c *gin.Context, repository OrderRepositoryInterface) {
	c.JSON(200, repository.GetAllOrder())
}

func (ctrl OrderController) GetOrder(c *gin.Context, repository OrderRepositoryInterface) {

	id := c.Param("id")

	idConv, _ := strconv.Atoi(id)

	c.JSON(200, repository.GetOrder(idConv))
}

func (r OrderRepositoryMock) GetAllOrder() Orders {
	orders := Orders{
		{Name: "Wilson"},
		{Name: "Panda"},
	}

	return orders
}

func (r OrderRepositoryMock) GetOrder(id int) entity.Order {
	orders := Orders{
		{Name: "Wilson"},
		{Name: "Panda"},
	}

	return orders[0]
	// return orders[id-1]
}

// TESTING CONTROLLER FUNCTIONS
func TestControllerGetAll(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	repo := OrderRepositoryMock{}
	ctrl := OrderController{}

	ctrl.GetAllOrder(c, repo)

	assert.Equal(t, 200, w.Code, "Success")
}

func TestControllerGet(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	repo := OrderRepositoryMock{}
	ctrl := OrderController{}

	ctrl.GetOrder(c, repo)

	// assert.Equal(t, 200, w.Code, "Success")
}

// TESTING REPOSITORY FUNCTIONS
func TestRepoGetAll(t *testing.T) {

	userRepo := OrderRepositoryMock{}

	amountUsers := len(userRepo.GetAllOrder())
	t.Log(amountUsers)
	e := 2
	if amountUsers != e {
		t.Errorf("Expected %d, Actual %d", e, amountUsers)
	}
}

func TestRepoGet(t *testing.T) {

	expectedUser := struct {
		Name string
	}{
		"Wilson",
	}

	userRepo := OrderRepositoryMock{}

	user := userRepo.GetOrder(1)
	e := expectedUser.Name
	if user.Name != e {
		t.Errorf("Expected %s, Actual %s", e, user.Name)
	}
}
