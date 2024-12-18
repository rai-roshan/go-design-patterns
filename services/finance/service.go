package finance

type FinanceService interface {
	Add(a int, b int) (c int)
	Sub(a int, b int) (c int)
}

type financeService struct {

}

func (fs *financeService) Add(a int, b int) (c int) {
	return a + b
}

func (fs *financeService) Sub(a int, b int) (c int) {
	return a - b
}

func NewFinanceService() FinanceService {
	return &financeService{}
}