package main

import (
  "fmt"
)

type order struct {
  id int 
  amount int
  status string
  customer
}

type customer struct {
  name string
}

func main () {

  myOrder := order{
    id: 1,
    amount: 100,
    status: "received",
  }

  newOrder := order {
    id: 2,
    amount: 200,
    status: "received",
    customer: customer{
      name: "John Doe",
    },
  }
  
  fmt.Println(myOrder)

  fmt.Println(newOrder.customer)
}