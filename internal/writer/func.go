package writer

import (
	"errors"
	"fmt"
	"github.com/skhoroshavin/automap/internal/core"
	"io"
)

func writeFunc(out io.Writer, fn *core.FuncBody) error {
	if fn.Result == "" {
		return errors.New("function has empty return statement")
	}
	
	for _, v := range fn.Vars {
		_, err := fmt.Fprintf(out, "\t%s := %s\n", v.Name, v.Value)
		if err != nil {
			return err
		}
	}

	_, err := fmt.Fprintf(out, "\treturn %s\n", fn.Result)
	return err
}
