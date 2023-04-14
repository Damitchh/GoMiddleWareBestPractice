package test

/*import (
	"Hacktiv10JWT/repositories"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)*/

/*func TestGetProductbyID(t *testing.T) {
	// Create a new Gin context for testing
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Set the ID parameter in the context
	ctx.Params = append(ctx.Params, gin.Param{Key: "ID", Value: "1"})

	// Mock the database and product repository
	db, mock, _ := gormmock.New()
	defer db.Close()
	productRepo := repositories.NewProductRepository(db)

	// Define the expected database query and result
	rows := sqlmock.NewRows([]string{"id", "title", "description", "created_at", "updated_at"}).
		AddRow(1, "Test Product", "A test product", time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM products WHERE id = (.+)$").WillReturnRows(rows)

	// Call the controller function
	GetProductbyID(ctx)

	// Check that the response status is 200 OK
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())

	// Check that the response body matches the expected result
	assert.JSONEq(t, `{
        "ID": 1,
        "Title": "Test Product",
        "Description": "A test product",
        "CreatedAt": "...",
        "UpdatedAt": "..."
    }`, ctx.Writer.Body.String())
}*/
