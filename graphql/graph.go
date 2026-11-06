package main

// type Server struct {
// 	accountClient *account.client
// 	catalogClient *catalog.client
// 	orderClient   *order.client
// }

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	catalogClient, err := catalog.NewClient(catalogUrl)
	if err != nil {
		accountClient.Close()
		return nil, err
	}

	orderClient, err := order.NewClient(orderUrl)
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return nil, err
	}

	return &Server {
		accountClient,
		catalogClient,
		orderClient,
	}, nil
}



func (s *Server) Mutation() Mutationresolver{
	return &mutationresolver{
		server: s, 
	}
}

func (s *Server) Query() QueryResolver{
	return &query
}