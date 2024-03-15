package method

import (
	"github.com/graphql-go/graphql"
	"log"
	"tiny_vote/model"
)

func Query(p graphql.ResolveParams) (interface{}, error) {
	name := p.Args["name"]
	log.Println(name.(string))
	votes := model.GetVotesByName(name.(string))
	return votes, nil
}
