package opt_test

import (
	"testing"

	"github.com/grimdork/opt"
)

func TestTags(t *testing.T) {
	var o struct {
		opt.DefaultHelp
		One      bool     `short:"1" long:"one" help:"An example bool." group:"Group A"`
		Two      string   `short:"2" placeholder:"TWO" help:"An example string argument." default:"two" group:"Group A"`
		Three    float64  `short:"3" placeholder:"FLOAT64" help:"An example float64 argument." group:"Group B"`
		Four     int      `placeholder:"INT" opt:"required" help:"Int argument with choices." group:"Group B" choices:"0,1,2"`
		CmdOne   struct{} `command:"one" help:"Command one." group:"Command group 1"`
		CmdTwo   struct{} `command:"two" help:"Command two." group:"Command group 1"`
		CmdThree struct{} `command:"three" help:"Command three." group:"Command group 2"`
	}

	a := opt.Parse(&o)
	a.Usage()
}
