// Code generated by qtc from "signup.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/signup.qtpl:1
package template

//line template/signup.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/signup.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/signup.qtpl:2
// SignupValues ...
type SignupValues struct {
	Name    string
	Company string
}

//line template/signup.qtpl:9
func StreamSignup(qw422016 *qt422016.Writer, t SignupValues) {
//line template/signup.qtpl:9
	qw422016.N().S(`
Welcome to `)
//line template/signup.qtpl:10
	qw422016.E().S(t.Company)
//line template/signup.qtpl:10
	qw422016.N().S(` `)
//line template/signup.qtpl:10
	qw422016.E().S(t.Name)
//line template/signup.qtpl:10
	qw422016.N().S(`!

Thank you very much for registering.

Best regards,
`)
//line template/signup.qtpl:15
	qw422016.E().S(t.Company)
//line template/signup.qtpl:15
	qw422016.N().S(` Team.
`)
//line template/signup.qtpl:16
}

//line template/signup.qtpl:16
func WriteSignup(qq422016 qtio422016.Writer, t SignupValues) {
//line template/signup.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/signup.qtpl:16
	StreamSignup(qw422016, t)
//line template/signup.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line template/signup.qtpl:16
}

//line template/signup.qtpl:16
func Signup(t SignupValues) string {
//line template/signup.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line template/signup.qtpl:16
	WriteSignup(qb422016, t)
//line template/signup.qtpl:16
	qs422016 := string(qb422016.B)
//line template/signup.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line template/signup.qtpl:16
	return qs422016
//line template/signup.qtpl:16
}
