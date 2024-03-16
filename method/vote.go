package method

import (
	"errors"
	"github.com/graphql-go/graphql"
	"tiny_vote/model/db"
	"tiny_vote/model/redis"
)

func Vote(p graphql.ResolveParams) (interface{}, error) {
	var err error

	currentTicket, err := redis.GetTicket()
	if err != nil {
		return false, err
	}

	ticket := p.Args["ticket"].(string)
	if ticket != currentTicket {
		return false, errors.New("your ticket has been expired")
	}

	// Assume the names are all legal and skip the validation.
	names := p.Args["name"].([]interface{})
	namesStr := []string{}
	for _, name := range names {
		namesStr = append(namesStr, name.(string))
	}
	_, err = db.UpdateVotesByNames(namesStr)
	if err != nil {
		return false, err
	}
	return true, nil
}
