package functions
import(
	"fmt"
)

func HelloWorld(){
	str := "world"
	_, err := fmt.Printf("Type: %T\n", str)
	if err != nil {
	}
}