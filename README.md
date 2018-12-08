# golangTest

Création d'un serveur qui expose un noeud graphql point qui retourne "This is the answer about the Query !" lorsqu'il est interrogé.

# Installation

Ajouter le package Graphql dans Golang en utilisant la commande :
go get "github.com/graphql-go/graphql"

# Lancer le serveur

En utilisant la commande :
go run main.go

# L'envoie d'une requete

On utilise le noeud point /graphQL
En utilisant une requete POST:

Comme notre serveur est limité, la requete doit etre comme ceci :

{
	'query: {str}
}

En utilisant une requete GET:

la requete doit etre comme ceci :
localhost:400/graphql?query={str}
