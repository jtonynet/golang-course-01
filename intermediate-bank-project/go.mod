module golang-course-01

go 1.16

require (
	github.com/jtonynet/golang-course-01/clientes v0.0.0-00010101000000-000000000000
	github.com/jtonynet/golang-course-01/contas v0.0.0
)

replace github.com/jtonynet/golang-course-01/contas => ./contas

replace github.com/jtonynet/golang-course-01/clientes => ./clientes
