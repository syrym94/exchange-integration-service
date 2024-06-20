package exchanges

type Trader interface {
	// метод первончальной инициализации бирж
	Setup(name string, settings struct{}) error

	// метод получения списка валют с биржи
	GetCurrencies() ([]string, error)

	// метод получения списка пар с биржи
	GetPairs() ([]string, error)

	// получение ордербуков и отправка их в канал
	GetOrderBooks(struct{})

	// получение баланса
	GetBalance(account, currency string) (float64, error)

	// получение активных ордеров
	GetActiveOrders(account string, pairs []string) ([]interface{}, error)

	// метод выставления ордера
	// orderType: buy/sell
	Order(account, pair, orderType string, amount, price float64) (interface{}, error)

	// метод отмены активного ордера
	CancelOrder(account string, orderId string) error

	// метод вывода (отправки) монет на адрес
	Withdrawal(account, currency string, amount float64, address string, paymentId string) error

	// метод проверки состояния биржи
	HealthCheck() bool

	//Shutdown - завершение работы
	Shutdown()
}
