package queryController

import (
	"Pood/app/models/queryModel"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/url"
	"strings"
)

func GetQueries(c *fiber.Ctx) (queries queryModel.Query) {
	u, _ := url.Parse(c.OriginalURL())
	m, _ := url.ParseQuery(u.RawQuery)

	return GetQueryFromMap(m)
}

func GetQueryFromMap(m url.Values) queryModel.Query {
	replacer := strings.NewReplacer("[", " ", "]", " ", ",", " ")
	var queries queryModel.Query

	for n, v := range m {
		if len(v[0]) == 0 {
			continue
		}

		value := v[0]
		queryData := strings.Fields(replacer.Replace(n))
		valueData := strings.Fields(replacer.Replace(value))

		filter := ""
		order := ""

		switch queryData[0] {
		case "filter":
			switch queryData[2] {
			case "eq":
				filter = fmt.Sprintf("%s = '%s'", queryData[1], value)
			case "includes":
				filter = fmt.Sprintf("%s LIKE '%%%s%%'", queryData[1], value)
			case "exists":
				if value == "true" {
					filter = fmt.Sprintf("%s > 0", queryData[1])
				} else {
					filter = fmt.Sprintf("%s = 0", queryData[1])
				}
			case "in":
				data := strings.Fields(replacer.Replace(value))
				for i, d := range data {
					data[i] = fmt.Sprintf("'%s'", d)
				}
				dataStr := strings.Join(data, ",")
				filter = fmt.Sprintf("%s in (%s)", queryData[1], dataStr)
			case "gte":
				filter = fmt.Sprintf("%s >= '%s'", queryData[1], value)
			case "lte":
				filter = fmt.Sprintf("%s <= '%s'", queryData[1], value)
			}
		case "order":
			order = fmt.Sprintf("%s %s", valueData[0], strings.ToUpper(valueData[1]))
		case "deleted":
			filter = fmt.Sprintf("%s = %s", queryData[0], value)
		}

		if len(filter) != 0 {
			queries.Filters = append(queries.Filters, filter)
		}
		if len(order) != 0 {
			queries.Orders = append(queries.Orders, order)
		}

	}

	return queries
}

func ConfigurationDbQuery(db *gorm.DB, queries queryModel.Query) *gorm.DB {
	if queries.Filters != nil {
		for _, filter := range queries.Filters {
			db = db.Where(filter)
		}
	}

	if queries.Orders != nil {
		for _, order := range queries.Orders {
			db = db.Order(order)
		}
	}

	return db
}
