package main

import (
	"flag"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/colin404/kubelistcheck/pkg/analyzer"
)

func main() {
	flag.Bool("unsafeptr", false, "")

	singlechecker.Main(analyzer.NewAnalyzer(false))
}
