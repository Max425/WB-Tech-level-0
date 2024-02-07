package core

type Order struct {
	ID       int    `db:"id"`
	OrderUID string `db:"order_uid"`
	Data     []byte `db:"data"`
}
