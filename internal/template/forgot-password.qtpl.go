// Code generated by qtc from "forgot-password.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/forgot-password.qtpl:1
package template

//line template/forgot-password.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/forgot-password.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/forgot-password.qtpl:2
// ForgotPasswordValues ...
type ForgotPasswordValues struct {
	Email      string
	TokenURL   string
	ExpireTime string
	Company    string
}

//line template/forgot-password.qtpl:11
func StreamForgotPassword(qw422016 *qt422016.Writer, t ForgotPasswordValues) {
//line template/forgot-password.qtpl:11
	qw422016.N().S(`
You've asked to reset your password for the `)
//line template/forgot-password.qtpl:12
	qw422016.E().S(t.Company)
//line template/forgot-password.qtpl:12
	qw422016.N().S(`  account associated with this email address (`)
//line template/forgot-password.qtpl:12
	qw422016.E().S(t.Email)
//line template/forgot-password.qtpl:12
	qw422016.N().S(`). Please click on the following link:

`)
//line template/forgot-password.qtpl:14
	qw422016.E().S(t.TokenURL)
//line template/forgot-password.qtpl:14
	qw422016.N().S(`

This password change code will expire `)
//line template/forgot-password.qtpl:16
	qw422016.E().S(t.ExpireTime)
//line template/forgot-password.qtpl:16
	qw422016.N().S(` from the time this email was sent. You can also copy and paste the link above into a new browser window.


If you didn't make the request, you can ignore this email and do nothing. Another user likely entered your email address by mistake while trying to reset a password.

Replies to this email are not monitored or answered.


Thank you for using `)
//line template/forgot-password.qtpl:24
	qw422016.E().S(t.Company)
//line template/forgot-password.qtpl:24
	qw422016.N().S(` .
The `)
//line template/forgot-password.qtpl:25
	qw422016.E().S(t.Company)
//line template/forgot-password.qtpl:25
	qw422016.N().S(`  Team


***This is an automatic notification. Replies to this email are not monitored or answered.
`)
//line template/forgot-password.qtpl:29
}

//line template/forgot-password.qtpl:29
func WriteForgotPassword(qq422016 qtio422016.Writer, t ForgotPasswordValues) {
//line template/forgot-password.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/forgot-password.qtpl:29
	StreamForgotPassword(qw422016, t)
//line template/forgot-password.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line template/forgot-password.qtpl:29
}

//line template/forgot-password.qtpl:29
func ForgotPassword(t ForgotPasswordValues) string {
//line template/forgot-password.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
//line template/forgot-password.qtpl:29
	WriteForgotPassword(qb422016, t)
//line template/forgot-password.qtpl:29
	qs422016 := string(qb422016.B)
//line template/forgot-password.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
//line template/forgot-password.qtpl:29
	return qs422016
//line template/forgot-password.qtpl:29
}
