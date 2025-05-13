package subfinder

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/g0ldencybersec/EasyEASM/pkg/utils"
)

func RunSubfinder(seedDomain string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Runing Subfinder on %s\n", seedDomain)
	binPath := utils.GetGoEnvPathBin()
	cmd := exec.Command(binPath+"/subfinder", "-d", seedDomain, "-silent")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	for _, domain := range strings.Split(out.String(), "\n") {
		if strings.Contains(domain, seedDomain) && len(domain) != 0 {
			results <- domain
		}
	}
	fmt.Printf("Subfinder run completed for %s\n", seedDomain)
}
