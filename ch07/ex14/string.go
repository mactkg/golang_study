package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%s%s", string(u.op), u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("%s %s %s", b.x, string(b.op), b.y)
}

func (c call) String() string {
	args := ""
	for _, v := range c.args {
		args += v.String() + ", "
	}
	args = strings.TrimRight(args, ", ")

	return fmt.Sprintf("%s(%s)", c.fn, args)
}

func (c comment) String() string {
	// return fmt.Sprintln(c.str)
	return fmt.Sprintf("%s#%s", c.expr, string(c.str))
}

func (n nop) String() string {
	return ""
}
