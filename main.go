package main

import (
	"context"
	"fmt"
	"go-pdf/sample_data"
	"html/template"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	filename, err := generateHTML()
	if err != nil {
		return
	}

	HtmlFilePath := getFilePath(filename)

	generatePDF(HtmlFilePath)
	deleteOutputHTML(HtmlFilePath)
}

func generateHTML() (string, error) {
	tmpl := template.Must(template.ParseFiles("./templates/bank-statement.templ"))

	outputFile, err := os.Create("./templates/output.html")
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, sample_data.GetBankAccountData())
	if err != nil {
		return "", err
	}

	println("HTML Generated", outputFile.Name())

	return outputFile.Name(), nil
}

func generatePDF(HtmlFilePath string) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.ExecPath("/usr/bin/brave-browser"),
		chromedp.Headless,
	}

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer allocCancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	url := "file:///" + HtmlFilePath
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(newCtx context.Context) error {
			time.Sleep(300 * time.Millisecond)

			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperHeight(11.69).
				WithPaperWidth(8.26).
				Do(newCtx)
			if err != nil {
				return err
			}

			pdf, err := os.Create("output.pdf")
			if err != nil {
				return err
			}
			defer pdf.Close()

			if _, err = pdf.Write(buf); err != nil {
				return err
			}

			return nil
		}),
	)
	if err != nil {
		fmt.Println("error navigating the website:", err)
		return
	}

	fmt.Println("PDF file generated")
}

func deleteOutputHTML(HtmlFilePath string) {
	err := os.Remove(HtmlFilePath)
	if err != nil {
		fmt.Println("Failed to delete HTML output file", err)
	}
}

func getFilePath(filename string) string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current directory")
		panic(err)
	}

	return cwd + "/" + filename
}
