package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mitchellh/go-mruby"
)

var mrb *mruby.Mrb

const coreScript = `
def new_hash
  Hash.new
end
`

func loadFile(name string) string {
	bytes, _ := ioutil.ReadFile(name)
	return string(bytes)
}

func main() {
	rb2 := loadFile("test.rb")
	mrb = mruby.NewMrb()
	defer mrb.Close()
	parser := mruby.NewParser(mrb)
	defer parser.Close()

	parser.Parse(rb2, nil)
	v := parser.GenerateCode()

	mrb.LoadString(coreScript)
	hv, _ := mrb.TopSelf().Call("new_hash")
	h := hv.Hash()
	h.Set(mrb.StringValue("test"), mrb.StringValue("1"))
	fmt.Printf("out: %s\n", hv.String())

	mrb.Run(v, nil)
	o2, r2 := v.Call("test3", hv)
	fmt.Printf("%s %s\n", o2, r2)
}
