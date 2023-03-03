package main

import (
	"gihub.com/marcelowis/linuxtips-github-tags/git"
)

func main() {
	repo := "kubernetes"
	donoDoRepositorio := "kubernetes"
	tag := "v1.26.2"

	// var b = new(git.BuscaGitTag)
	b := git.BuscaGitTag{
		DonoDoRepositorio: donoDoRepositorio,
		Repo:              repo,
	}
	b.BuscaGitTag(tag)
}
