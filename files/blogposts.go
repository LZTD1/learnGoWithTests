package files

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title"
	descriptionSeparator = "Description"
	tagsSeparator        = "Tags"
)

func newPostsFromFs(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		if isMDFile(f.Name()) {
			p, err := getPost(fileSystem, f.Name())
			if err != nil {
				return nil, err
			}
			posts = append(posts, p)
		}
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	t, err := scanLine(titleSeparator, scanner)
	if err != nil {
		return Post{}, err
	}
	d, err := scanLine(descriptionSeparator, scanner)
	if err != nil {
		return Post{}, err
	}
	tags, err := scanLine(tagsSeparator, scanner)
	if err != nil {
		return Post{}, err
	}

	return Post{
		Title:       t,
		Description: d,
		Tags:        strings.Split(tags, ", "),
		Body:        readBody(scanner),
	}, nil
}

func scanLine(s string, scanner *bufio.Scanner) (string, error) {
	scanner.Scan()
	val, err := getValue(scanner.Text(), s)
	if err != nil {
		return "", err
	}
	return val, nil
}

func getValue(s, d string) (string, error) {
	if !strings.HasPrefix(s, d) {
		got := s[:strings.Index(s, ":")]
		return "", fmt.Errorf("Wanted %s, but got %s", d, got)
	}
	return strings.TrimSpace(s[strings.Index(s, ":")+1:]), nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	b := strings.TrimSuffix(buf.String(), "\n")
	return b
}

func isMDFile(name string) bool {
	if strings.HasSuffix(name, ".md") {
		return true
	}
	return false
}
