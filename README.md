# Testing Go

This repo supports a blog post on testing in go.

It demonstrates two approaches to testing with third parties in go. The first is
using an interface where a mocked and real implementation are used
interchangably. Tests that don't really care about the behaviour of the third
party can use the mock and forget about it. This approach can be found in
posts/posts_test.go

The second approach is for when you want to test your third party client
directly. It makes use of httptest.NewServer to create a test server to make
requests against. This approach can be found in auditor/auditior_test.go

You can run the app with `go run main.go` or the tests with `go test ./...`
