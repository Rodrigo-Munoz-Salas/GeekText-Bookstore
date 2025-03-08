package database

// Add other imports if necessary
import (
    "context"
    "github.com/google/uuid"
)

// GetCartSubtotal retrieves the subtotal price of all items in the user's cart
func (q *Queries) GetCartSubtotal(ctx context.Context, userID uuid.UUID) (float64, error) {
    var subtotal float64
    err := q.db.QueryRowContext(ctx, `
        SELECT SUM(price)
        FROM cart_items
        WHERE user_id = $1
    `, userID).Scan(&subtotal)
    if err != nil {
        return 0, err
    }
    return subtotal, nil
}