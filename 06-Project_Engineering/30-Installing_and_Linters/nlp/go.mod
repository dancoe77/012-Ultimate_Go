module github.com/dancoe77/nlp

go 1.25.5

// Risks of using 3rd party packages
// - Security (ignorance, intentional)
// - Bugs
// - Compatibility (API Changes); Semantic Versioning
// - Legal (license)
// - Package and dependencies might be missing; go mod vendor

require (
	github.com/BurntSushi/toml v1.6.0
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20231108232855-2478ac86f678 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)

tool honnef.co/go/tools/cmd/staticcheck
