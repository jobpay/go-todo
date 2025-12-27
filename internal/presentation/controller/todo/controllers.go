package todo

type Controllers struct {
	Show   *ShowController
	List   *ListController
	Store  *StoreController
	Update *UpdateController
	Delete *DeleteController
}

func NewControllers(
	show *ShowController,
	list *ListController,
	store *StoreController,
	update *UpdateController,
	delete *DeleteController,
) *Controllers {
	return &Controllers{
		Show:   show,
		List:   list,
		Store:  store,
		Update: update,
		Delete: delete,
	}
}

