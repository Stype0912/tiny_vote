package method

import (
	"errors"
	"github.com/graphql-go/graphql"
	"log"
	"tiny_vote/model"
)

func Vote(p graphql.ResolveParams) (interface{}, error) {
	ticket := p.Args["ticket"].(string)
	if ticket != CurrentTicket.ticketId {
		return false, errors.New("your ticket has been expired")
	}
	// Assume the names are all legal and skip the validation.
	log.Println(p.Args["name"].([]interface{}))
	names := p.Args["name"].([]interface{})
	namesStr := []string{}
	for _, name := range names {
		namesStr = append(namesStr, name.(string))
	}
	_, err := model.UpdateVotesByNames(namesStr)
	if err != nil {
		return false, err
	}
	return true, nil
}
