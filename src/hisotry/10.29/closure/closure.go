package closure
import (
	"fmt"
	"string"
)

func makeSuffix(suffix string) func(fileName string) string {
	return func(fileName string) {
		if !string.HasSuffix(fileName, suffix) {
			return fileName
		} else {
			return fileName + suffix
		}
	}
}