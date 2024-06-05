package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"transport-api/pkg/tokenizer"
)

func Auth(tkn tokenizer.Tokenizer) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerValue := c.GetHeader("Authorization")
		token := getTokenFromHeader(headerValue)
		if token == "" {
			c.Next()
		}

		data, err := tkn.DecryptToken(token)
		if err != nil {
			c.Next()
		}
		c.Set("EmployeeId", strconv.FormatUint(uint64(data.EmployeeId), 10))
		c.Set("CompanyId", strconv.FormatUint(uint64(data.CompanyId), 10))
		c.Next()
	}
}

func getTokenFromHeader(headerValue string) string {
	header := strings.Split(headerValue, " ")
	if len(header) < 2 || header[0] != "Bearer" {
		return ""
	}
	return header[1]
}

func GetAuthData(c *gin.Context) (uint64, uint64, bool) {
	employeeId, ok := GetIdData(c, "EmployeeId")
	if !ok {
		return 0, 0, false
	}

	companyId, ok := GetIdData(c, "CompanyId")
	if !ok {
		return 0, 0, false
	}

	return employeeId, companyId, true
}

func GetIdData(c *gin.Context, key string) (id uint64, ok bool) {
	idAny, ok := c.Get(key)
	if !ok {
		return
	}

	idStr, ok := idAny.(string)
	if !ok {
		return
	}

	var err error
	id, err = strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return
	}
	return id, true
}
