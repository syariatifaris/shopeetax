package tax

//Tax table database structure
type Tax struct {
	ID          int64  `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Description string `json:"desc" db:"description"`
	Expression  string `json:"expression" db:"expression"`
}
