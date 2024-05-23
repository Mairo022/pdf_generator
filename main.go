package main

import (
	"context"

	"github.com/chromedp/chromedp"
)

func main() {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.ExecPath("/usr/bin/brave-browser"),
	}

	_, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
}
