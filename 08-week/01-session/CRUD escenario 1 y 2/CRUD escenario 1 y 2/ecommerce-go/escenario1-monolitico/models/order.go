package models

type Order struct {
    ID        uint    `json:"id"`
    UserID    uint    `json:"userId"`
    ProductID uint    `json:"productId"`
    Quantity  int     `json:"quantity"`
    Total     float64 `json:"total"`
}
