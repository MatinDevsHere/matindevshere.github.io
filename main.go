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
	Content     template.HTML
	Permalink   string
}

const (
	postsDir  = "posts"
	buildDir  = "build"
	delimiter = "---"
)

func copyStaticFiles() error {
	// Copy profanity.js
	jsContent, err := ioutil.ReadFile("build/profanity.js")
	if err != nil {
		return fmt.Errorf("failed to read profanity.js: %v", err)
	}

	err = ioutil.WriteFile(filepath.Join(buildDir, "profanity.js"), jsContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to copy profanity.js: %v", err)
	}

	// Copy styles.css
	cssContent, err := ioutil.ReadFile("build/styles.css")
	if err != nil {
		return fmt.Errorf("failed to read styles.css: %v", err)
	}

	err = ioutil.WriteFile(filepath.Join(buildDir, "styles.css"), cssContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to copy styles.css: %v", err)
	}

	return nil
}

func main() {
	// Create build directory if it doesn't exist
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		log.Fatal(err)
	}

	// Copy static files
	if err := copyStaticFiles(); err != nil {
		log.Printf("Warning: Failed to copy static files: %v", err)
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

func processProfanityContent(content string) template.HTML {
	// Find all instances of ||text1||text2|| pattern
	parts := strings.Split(content, "||")
	var result strings.Builder

	for i := 0; i < len(parts); i++ {
		if i == 0 {
			result.WriteString(parts[i])
			continue
		}

		// Every third part is the end of a profanity block
		switch i % 3 {
		case 1: // NSFW content
			result.WriteString(`<span class="nsfw">`)
			result.WriteString(parts[i])
			result.WriteString(`</span>`)
		case 2: // SFW content
			result.WriteString(`<span class="sfw">`)
			result.WriteString(parts[i])
			result.WriteString(`</span>`)
		case 0: // Regular content
			result.WriteString(parts[i])
		}
	}

	return template.HTML(result.String())
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

	// Process content including profanity handling
	post.Content = processProfanityContent(string(bytes.TrimSpace(parts[2])))

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
    <link rel="stylesheet" href="styles.css">
    <script src="profanity.js"></script>
</head>
<body>
    <button id="profanityToggle" onclick="toggleProfanity()">Enable Adult Content</button>
    <div class="container">
        <h1>{{.Frontmatter.title}}</h1>
        {{if .Frontmatter.date}}<p class="date">{{.Frontmatter.date}}</p>{{end}}
        <div class="content">
            {{.Content}}
        </div>
        <p><a href="index.html">Back to Home</a></p>
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
    <link rel="stylesheet" href="styles.css">
    <script src="profanity.js"></script>
</head>
<body>
    <button id="profanityToggle" onclick="toggleProfanity()">Enable Adult Content</button>
    <div class="container">
        <h1>Blog Posts</h1>
        <ul>
            {{range .}}
            <li>
                <a href="{{.Permalink}}.html">{{.Frontmatter.title}}</a>
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
