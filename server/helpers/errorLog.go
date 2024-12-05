package helpers

import (
	"fmt"
)

func CheckError(err error,task string) bool {
	if err != nil {
		fmt.Printf("Error during %s : %v\n",task,err);
		return true
	}
	return false
}


// func CheckErrorSockets(err error,) bool {
// 	if err != nil {
// 		fmt.Printf("Error during %s : %v\n",task,err);
// 		return true
// 	}
// 	return false
// }
