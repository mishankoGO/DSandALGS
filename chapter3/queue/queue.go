package queue

// Queue of Orders
type Queue []*Order

// Order class
type Order struct {
	Priority     int
	Quantity     int
	Product      string
	CustomerName string
}

// New method inits the order
func (order *Order) New(priority, quantity int, product, customerName string) {
	order.Priority = priority
	order.Quantity = quantity
	order.Product = product
	order.CustomerName = customerName
}

// Add method
func (q *Queue) Add(order *Order) {
	if len(*q) == 0 {
		*q = append(*q, order)
	} else {
		var appended = false
		for i, addedOrder := range *q {
			if order.Priority > addedOrder.Priority {
				*q = append((*q)[:i], append(Queue{order}, (*q)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*q = append(*q, order)
		}
	}
}

// Remove method
func (q *Queue) Remove() *Order {
	return (*q)[len(*q)-1]
}
