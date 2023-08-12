package abstractfactory

// OrderMainDAO is order main data access object
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO is order detail data access object
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory is data access object factory
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// RDBMainDAO is RDB main data access object implementation
type RDBMainDAO struct{}

// SaveOrderMain save order main data
func (*RDBMainDAO) SaveOrderMain() {
	println("rdb main save")
}

// RDBDetailDAO is RDB detail data access object implementation
type RDBDetailDAO struct{}

// SaveOrderDetail save order detail data
func (*RDBDetailDAO) SaveOrderDetail() {
	println("rdb detail save")
}

// RDBDAOFactory is RDB data access object factory implementation
type RDBDAOFactory struct{}

// CreateOrderMainDAO create order main data access object
func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

// CreateOrderDetailDAO create order detail data access object
func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// XMLMainDAO is XML main data access object implementation
type XMLMainDAO struct{}

// SaveOrderMain save order main data
func (*XMLMainDAO) SaveOrderMain() {
	println("xml main save")
}

// XMLDetailDAO is XML detail data access object implementation
type XMLDetailDAO struct{}

// SaveOrderDetail save order detail data
func (*XMLDetailDAO) SaveOrderDetail() {
	println("xml detail save")
}

// XMLDAOFactory is XML data access object factory implementation
type XMLDAOFactory struct{}

// CreateOrderMainDAO create order main data access object
func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

// CreateOrderDetailDAO create order detail data access object
func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
