package main

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

//Creating our possible query structure
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"str": &graphql.Field{ //The query we must have as input "str"
			Type: graphql.String, //The type of the answer which is string
			//the function which returns the result about our query
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return "This is the answer about the Query !", nil //The response of our grapgql endpoint
			},
		},
	},
})

var shema, _ = graphql.NewSchema(graphql.SchemaConfig{ //Defining our graphql shema which contains our Query structure
	Query: rootQuery,
})

type QueryS struct { //the struct which contains our string request
	Query string
}

func main() {
	//creating a graphql endpoint which listens ar /graphql
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var q QueryS
		if r.Method == "POST" { //if we have a POST request

			json.NewDecoder(r.Body).Decode(&q) //we extract the query and save it in a string request
			//We excute the query
			result := graphql.Do(graphql.Params{
				Schema:        shema,
				RequestString: q.Query, //our string request
			})
			//We encode the result and send it back
			json.NewEncoder(w).Encode(result)
		} else { //if it's a GET or other method
			q.Query = r.URL.Query().Get("query") //we extract our string request
			//We excute the query
			result := graphql.Do(graphql.Params{
				Schema:        shema,
				RequestString: q.Query, //our string request
			})
			//We encode the result and send it back
			json.NewEncoder(w).Encode(result)
		}
	})

	http.ListenAndServe(":400", nil) //Set the server and listen on the port 400

}
