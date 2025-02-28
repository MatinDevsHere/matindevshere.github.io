package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Post struct {
	Frontmatter map[string]interface{} `yaml:",inline"`
	Content     string
	Permalink   string
}

const (
	postsDir  = "posts"
	buildDir  = "build"
	delimiter = "---"
)

func main() {
	// Create build directory if it doesn't exist
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		log.Fatal(err)
	}

	// Read all markdown files
	files, err := ioutil.ReadDir(postsDir)
	if err != nil {
		log.Fatal(err)
	}

	var posts []Post
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			post, err := processPost(file.Name())
			if err != nil {
				log.Printf("Error processing %s: %v", file.Name(), err)
				continue
			}
			posts = append(posts, post)
		}
	}

	// Generate HTML files
	for _, post := range posts {
		if err := generateHTML(post); err != nil {
			log.Printf("Error generating HTML for %s: %v", post.Permalink, err)
		}
	}

	// Generate index.html
	if err := generateIndex(posts); err != nil {
		log.Fatal(err)
	}
}

func processPost(filename string) (Post, error) {
	content, err := ioutil.ReadFile(filepath.Join(postsDir, filename))
	if err != nil {
		return Post{}, err
	}

	// Split content into frontmatter and body
	parts := bytes.Split(content, []byte(delimiter))
	if len(parts) < 3 {
		return Post{}, fmt.Errorf("invalid post format: %s", filename)
	}

	var post Post
	post.Frontmatter = make(map[string]interface{})

	// Parse frontmatter
	if err := yaml.Unmarshal(bytes.TrimSpace(parts[1]), &post.Frontmatter); err != nil {
		return Post{}, err
	}

	// Set content
	post.Content = string(bytes.TrimSpace(parts[2]))

	// Set permalink from filename
	post.Permalink = strings.TrimSuffix(filename, ".md")

	return post, nil
}

func generateHTML(post Post) error {
	// Create post template
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Frontmatter.title}}</title>
    <link rel="stylesheet" href="/styles.css">
</head>
<body>
    <div class="container">
        <h1>{{.Frontmatter.title}}</h1>
        {{if .Frontmatter.date}}<p class="date">{{.Frontmatter.date}}</p>{{end}}
        <div class="content">
            {{.Content}}
        </div>
        <p><a href="/">Back to Home</a></p>
    </div>
</body>
</html>`

	t, err := template.New("post").Parse(tmpl)
	if err != nil {
		return err
	}

	// Create output file
	outPath := filepath.Join(buildDir, post.Permalink+".html")
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, post)
}

func generateIndex(posts []Post) error {
	// Create index template
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Blog</title>
    <link rel="stylesheet" href="/styles.css">
</head>
<body>
    <div class="container">
        <h1>Blog Posts</h1>
        <ul>
            {{range .}}
            <li>
                <a href="/{{.Permalink}}.html">{{.Frontmatter.title}}</a>
                {{if .Frontmatter.date}} - {{.Frontmatter.date}}{{end}}
            </li>
            {{end}}
        </ul>
    </div>
</body>
</html>`

	t, err := template.New("index").Parse(tmpl)
	if err != nil {
		return err
	}

	// Create index.html
	f, err := os.Create(filepath.Join(buildDir, "index.html"))
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, posts)
}
