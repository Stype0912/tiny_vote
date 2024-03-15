package method

import "github.com/graphql-go/graphql"

func Cas(p graphql.ResolveParams) (interface{}, error) {
	CurrentTicket.mutex.Lock()
	defer CurrentTicket.mutex.Unlock()
	return CurrentTicket.ticketId, nil
}
