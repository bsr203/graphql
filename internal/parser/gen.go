// requires pigeon and goexports
// go get golang.org/x/tools/cmd/goimports github.com/PuerkitoBio/pigeon
//go:generate pigeon -o graphql.go ../../graphql.peg
//go:generate goimports -w graphql.go

package parser