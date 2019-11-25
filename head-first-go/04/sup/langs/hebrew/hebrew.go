package hebrew

import (
	"fmt"

	"github.com/barucharky/go-shenanigans/stringutil"
)

func Shalom() {
	fmt.Println(stringutil.Reverse("שלום."))
}

func Mah() {
	fmt.Println(stringutil.Reverse("מה שלומך?"))
}
