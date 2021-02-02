package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	goquery "github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type Client struct {
}

func Get(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	return resp.Body, nil
}

func Download(url string, outfile string) error {
	body, err := Get(url)
	if err != nil {
		return err
	}

	defer body.Close()

	// Create the file
	out, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, body)

	return err
}

func Body(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "div" {
			for _, attr := range node.Attr {
				if attr.Key == "class" && attr.Val == "problem-wrapper" {
					body = node
					return
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func DownloadToMd(url string, outfile string) error {
	body, err := Get(url)
	if err != nil {
		return err
	}

	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return err
	}

	sel := doc.Find("div.problem-wrapper")
	converter := md.NewConverter("open.kattis.com", true, &md.Options{})

	// The default commonmark text formatting rule escapes backslashes, which breaks
	// the nice tex2jax math tags in some problems.  So we add our own rule which
	// cuts the escapes again
	converter.AddRules(md.Rule{Filter: []string{"span"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			// If the span element has not the classname `tex2jax_process` return nil.
			// That way the next rules will apply. In this case the commonmark rules.
			if !selec.HasClass("tex2jax_process") {
				return nil
			}
			content = strings.TrimSpace(content)
			content = strings.ReplaceAll(content, "\\\\", "\\")
			fmt.Println(content)
			return md.String(content)
		}})

	markdown := converter.Convert(sel)
	if err != nil {
		log.Fatal(err)
	}

	// Create the file
	out, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = fmt.Fprint(out, markdown)

	return err

}

func CreateMain(problem string) error {
	file, err := os.Create(fmt.Sprintf("%s.go", problem))
	if err != nil {
		return err
	}

	defer file.Close()
	fmt.Fprintf(file, `package main

func main() {

}`)

	return nil
}

func main() {
	//Get problem id from args
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("usage: kattis-init <problem id>")
		os.Exit(1)
	}
	problem := flag.Arg(0)
	fmt.Println("Initializing problem", problem)

	//Create new module
	err := os.Mkdir(problem, 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatalln("Error creating directory", err)
	}
	err = os.Chdir(problem)
	if err != nil {
		log.Fatalln("Error changing directory", err)
	}

	err = exec.Command("go", "mod", "init", problem).Run()
	if err != nil {
		log.Println("Error initializing module", err)
	}

	err = CreateMain(problem)
	if err != nil {
		log.Println("Error initializing module", err)
	}

	//Download description
	err = Download(fmt.Sprintf("https://open.kattis.com/problems/%s", problem), "problem.html")
	if err != nil {
		log.Println("Error downloading description")
	}
	err = DownloadToMd(fmt.Sprintf("https://open.kattis.com/problems/%s", problem), "problem.md")
	if err != nil {
		log.Println("Error downloading description")
	}

	//Download examples
	err = Download(fmt.Sprintf("https://open.kattis.com/problems/%s/file/statement/samples.zip", problem), "samples.zip")
	if err != nil {
		log.Println("Error downloading samples")
	} else {
		err = exec.Command("unzip", "-n", "samples.zip").Run()
		if err != nil {
			log.Fatalln("Error unzipping samples", err)
		}
		err = exec.Command("rm", "samples.zip").Run()
		if err != nil {
			log.Fatalln("Error removing samples", err)
		}
	}

}
