package git

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Buscador interface {
	BuscaGitTag(tag string, donoDoRepositorio string, repo string)
}

type BuscaGitTag struct {
	Repo              string
	DonoDoRepositorio string
}

func (b *BuscaGitTag) BuscaGitTag(tag string) {

	url := fmt.Sprintf("https://api.github.com/repos/" + b.DonoDoRepositorio + "/" + b.Repo + "/git/refs/tags/" + tag)
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Buscando tag", tag, "no repositório", b.Repo, "do dono", b.DonoDoRepositorio)

	defer resp.Body.Close()

	r, _ := httputil.DumpResponse(resp, true)

	fmt.Println(string(r))

}
