package generate

//go:generate gotemplate "github.com/ncw/gotemplate/sort" "SortGt(Generate, func(a, b Generate) bool { return a.Bar < b.Bar })"

type Generate struct {
	Foo string
	Bar int
}
