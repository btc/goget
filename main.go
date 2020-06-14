package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"

	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/MontFerret/ferret/pkg/drivers"
	fhttp "github.com/MontFerret/ferret/pkg/drivers/http"
	"github.com/MontFerret/ferret/pkg/runtime"
	"github.com/manifoldco/promptui"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	query := os.Args[len(os.Args)-1]
	u := "https://pkg.go.dev/search?q=" + url.QueryEscape(query)
	results, err := search(u)
	if err != nil {
		return err
	}
	if len(results) == 0 {
		return errors.New("no results")
	}
	prompt := promptui.Select{
		Label: "select a package",
		Items: results,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return err
	}
	cmd := exec.Command("go", "get", result)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func search(url string) ([]string, error) {
	query := `
	LET doc = DOCUMENT(@url)

	FOR e in ELEMENTS(doc, 'h2.SearchSnippet-header > a')
		RETURN e
	FOR e in ELEMENTS(doc, '.DetailsHeader-infoLabelModule')
		RETURN e
	`
	// follow and then return DetailsHeader-infoLabelModule
	p, err := compiler.New().Compile(query)
	if err != nil {
		return nil, err
	}
	out, err := p.Run(
		drivers.WithContext(
			context.Background(),
			fhttp.NewDriver(),
			drivers.AsDefault(),
		),
		runtime.WithParam("url", url),
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(out))
	var data []string
	if err := json.Unmarshal(out, &data); err != nil {
		return nil, err
	}
	return data, nil
}
