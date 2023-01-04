# Getting Captain working with `go test`

1. 🧪 Ensure `go test` produces JSON output

`go test ./... -json` will produce Captain-compatible JSON output, but it will squash any other output. We recommend
installing [gotestsum](https://github.com/gotestyourself/gotestsum) and using it to produce `jsonfile` output.

```sh
go install gotest.tools/gotestsum@latest
gotestsum --jsonfile results.json
```

2. 🔐 Create an Access Token

Create an Access Token for your organization within [Captain][captain] ([more documentation here][create-access-token]).

Add the new token as an action secret to your repository. Conventionally, we call this secret `RWX_ACCESS_TOKEN`.

3. 💌 Install the Captain CLI, and then call it when running tests

See the [full documentation on test suite integration][test-suite-integration].

```yaml
  - run: |
      captain run --suite-id captain-examples-go-test --test-results results.json -- \
        gotestsum --jsonfile results.json
    env:
      RWX_ACCESS_TOKEN: ${{ secrets.RWX_ACCESS_TOKEN }}
```

4. 🎉 See your test results in Captain!

Take a look at the [final workflow!][workflow-with-captain]

[captain]: https://account.rwx.com/deep_link/manage/access_tokens
[create-access-token]: https://www.rwx.com/docs/access-tokens
[workflow-with-captain]: https://github.com/captain-examples/go-test/blob/main/.github/workflows/ci.yml
[test-suite-integration]: https://www.rwx.com/captain/docs/test-suite-integration
