package main

import "fmt"
import "github.com/alecthomas/participle"
import "github.com/alecthomas/participle/lexer"
import "github.com/alecthomas/participle/lexer/ebnf"

// https://tools.ietf.org/html/rfc2254
// TODO: Parse attr := AttributeDescription according to RFC 2251 4.1.5

type FilterList struct {
	Filters []*Filter `@@+`
}

type Filter struct {
    Value *FilterComp `"(" @@ ")"`
}

type FilterComp struct {
    Item *Item `@@`
}

type Item struct {
    Simple  *Simple `@@`
}

type Simple struct {
    Attr string `@LdapString`
    FilterType string `@FilterType`
    Value string `@LdapString`
}

func main() {
    fmt.Println("hello world")
    var lexer = lexer.Must(ebnf.New(`
        FilterType = ( "=" | "~=" | ">=" | "<=" ) . 
		LdapString = { "a"…"z" } . 
	`))

    var parser = participle.MustBuild(
        &FilterList{},
		participle.Lexer(lexer),
    )

    fmt.Println("Parser built")

    s := "(foo=bar)"
    filterList := &FilterList{}
    err := parser.ParseString(s, filterList)

    fmt.Println("Parsed")
    fmt.Println(filterList)
    fmt.Println(err)
}
