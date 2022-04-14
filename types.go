package cherry

type DuckStat string

const (
	StatusOk   DuckStat = "ok"
	StatusFail DuckStat = "fail"
)

type Role int32

const (
	// Зарегистрированный пользователь
	Peasant Role = 0

	// Менеджер
	Vicegerent Role = 1

	// Разработчик
	Witcher Role = 2

	// Администратор
	BloodyBaron Role = 3

	// Главный администратор
	Vesemir Role = 4
)
