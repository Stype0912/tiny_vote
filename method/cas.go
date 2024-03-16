package method

import (
	"github.com/graphql-go/graphql"
	"tiny_vote/model/redis"
)

func Cas(p graphql.ResolveParams) (interface{}, error) {
	currentTicket, err := redis.GetTicket()
	if err != nil {
		return false, err
	}
	return currentTicket, nil
}
