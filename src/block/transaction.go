package block

type Transaction struct {
	from    string
	to      string
	account string
}

func GetFrom(t *Transaction) string {
	return t.from
}
func GetTo(t *Transaction) string {
	return t.to
}
func GetAccount(t *Transaction) string {
	return t.account
}
