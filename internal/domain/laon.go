package domain


type Loan struct {
    ID     string  `json:"id"`
    Amount float64 `json:"amount"`
    Term   int     `json:"term"`
    Status string  `json:"status"`
}
