package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func parseLinkText(n *html.Node) string {
	var text string

	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += parseLinkText(c)
	}

	return strings.Join(strings.Fields(text), " ")
}

func buildLinks(n *html.Node) Link {
	var link Link

	for _, attribute := range n.Attr {
		if attribute.Key == "href" {
			link.Href = attribute.Val
			break
		}
	}

	link.Text = parseLinkText(n)

	return link
}

func linkNodes(n *html.Node) []*html.Node {
	var nodes []*html.Node

	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}

	return nodes
}

func Parse(r io.Reader) ([]Link, error) {
	var links []Link

	document, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	nodes := linkNodes(document)

	for _, node := range nodes {
		links = append(links, buildLinks(node))
	}

	return links, nil
}
