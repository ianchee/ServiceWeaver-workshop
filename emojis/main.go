package main

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), run); err != nil {
		panic(err)
	}
}

// app is the main component of our application.
type app struct {
	weaver.Implements[weaver.Main]
	searcher weaver.Ref[Searcher]
}

// run implements the application main.
func run(ctx context.Context, a *app) error {
	emojis, err := a.searcher.Get().Search(ctx, "pig")
	if err != nil {
		return nil
	}
	fmt.Println(emojis)
	return nil
}
