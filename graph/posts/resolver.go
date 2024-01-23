package posts

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"necodeo.com/m/v2/helpers"
)

var PostsResolver = func(p graphql.ResolveParams) (interface{}, error) {
	var columns []string
	var scanArgs []interface{}

	for _, fieldAST := range p.Info.FieldASTs {
		if fieldAST.SelectionSet != nil {
			for _, selection := range fieldAST.SelectionSet.Selections {
				if field, ok := selection.(*ast.Field); ok {
					columns = append(columns, field.Name.Value)
				}
			}
		}
	}

	columns = helpers.RemoveDuplicates(columns)

	if len(columns) == 0 {
		return nil, fmt.Errorf("no columns selected")
	}

	sqlQuery, args, err := squirrel.Select(columns...).From("b_posts").ToSql()
	if err != nil {
		return nil, err
	}

	fmt.Println(sqlQuery, args)

	rows, err := helpers.DB.Query(sqlQuery, args...)

	if err != nil {
		return nil, err
	}

	var results []Post

	for rows.Next() {
		var result Post

		scanArgs = make([]interface{}, len(columns))

		for i, col := range columns {
			switch col {
			case "id":
				scanArgs[i] = &result.Id
			case "title":
				scanArgs[i] = &result.Title
			case "content":
				scanArgs[i] = &result.Content
			default:
				return nil, fmt.Errorf("unknown column: %s", col)
			}
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	rows.Close()

	return results, nil
}
