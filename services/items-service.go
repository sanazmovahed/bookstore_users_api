package services

var (
	ItemsService itemServiceInterface = &itemsService{}
)

type itemsService struct {
}

type itemServiceInterface interface {
	Get()
	Create()
}

func (i *itemsService) Get() {

}

func (i *itemsService) Create() {

}
