package active

import (
	"github.com/666ares/EasyEASM/pkg/active/alterx"
	"github.com/666ares/EasyEASM/pkg/active/dnsx"
	"github.com/666ares/EasyEASM/pkg/active/httpx"
)

type ActiveRunner struct {
	SeedDomains []string
	Results     int
	Subdomains  []string
}

func (r *ActiveRunner) RunActiveEnum(wordlist string, threads int) []string {
	return dnsx.RunDnsx(r.SeedDomains, wordlist, threads)
}

func (r *ActiveRunner) RunPermutationScan(threads int) []string {
	return alterx.RunAlterx(r.Subdomains, threads)
}

func (r *ActiveRunner) RunHttpx() {
	httpx.RunHttpx(r.Subdomains)
}
