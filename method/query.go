package method

import (
	"github.com/graphql-go/graphql"
	"log"
	"tiny_vote/model/db"
)

func Query(p graphql.ResolveParams) (interface{}, error) {
	name := p.Args["name"]
	log.Println("query name is", name.(string))
	votes, err := db.GetVotesByName(name.(string))
	if err != nil {
		return 0, err
	}
	return votes, nil
}
