package ltests

import(
	"github.com/mutannejs/luof-go/pkg/lerror"
)

func GetMsgError(vErr lerror.ValueError) string {
	return vErr.GetErrors()[0].GetMessage()
}
