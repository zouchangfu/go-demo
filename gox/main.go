package main

import (
	"go/token"
	"os"

	"github.com/goplus/gox"
)

func ctxRef(pkg *gox.Package, name string) gox.Ref {
	return pkg.CB().Scope().Lookup(name)
}

var UserColumn = struct {
	ID        string
	Username  string
	Password  string
	Address   string
	Age       string
	Phone     string
	Score     string
	Dept      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Username:  "username",
	Password:  "password",
	Address:   "address",
	Age:       "age",
	Phone:     "phone",
	Score:     "score",
	Dept:      "dept",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func main() {
	pkg := gox.NewPackage("", "main", nil)
	//fmt := pkg.Import("fmt")
	//v := pkg.NewParam(token.NoPos, "v", types.Typ[types.String]) // v string

	pkg.NewVar(token.NoPos, nil, "User")
	pkg.NewVarStart(token.NoPos, nil, "User1", "User2").Val(3).EndInit(2)
	//pkg.NewFunc(nil, "main", nil, nil, false).BodyStart(pkg).
	//	DefineVarStart(token.NoPos, "a", "b").Val("Hi").Val(3).EndInit(2). // a, b := "Hi", 3
	//	NewVarStart(nil, "c").Val(ctxRef(pkg, "b")).EndInit(1).            // var c = b
	//	NewVar(gox.TyEmptyInterface, "x", "y").                            // var x, y interface{}
	//	Val(fmt.Ref("Println")).
	//	/**/ Val(ctxRef(pkg, "a")).Val(ctxRef(pkg, "b")).Val(ctxRef(pkg, "c")). // fmt.Println(a, b, c)
	//	/**/ Call(3).EndStmt().
	//	NewClosure(gox.NewTuple(v), nil, false).BodyStart(pkg).
	//	/**/ Val(fmt.Ref("Println")).Val(v).Call(1).EndStmt(). // fmt.Println(v)
	//	/**/ End().
	//	Val("Hello").Call(1).EndStmt(). // func(v string) { ... } ("Hello")
	//	End()

	pkg.WriteTo(os.Stdout)
}
