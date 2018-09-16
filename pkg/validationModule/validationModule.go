package validationModule

import (
	"fmt"

	"../dbConnectModule"
	"../structModule"
)

func CheckAPIToken(valAuth structModule.ValAPIKey) bool {
	dbConnectModule.Sqlite_Open()
	tempCount := dbConnectModule.GetAPIKeyValiation(valAuth)
	dbConnectModule.Sqlite_Close()
	fmt.Print("tempCount = ")
	fmt.Println(tempCount)
	if len(tempCount) == 1 {
		return true
	} else {
		return false
	}
}
