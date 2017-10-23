package u3d
import (
	"fmt"
)

func strf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
