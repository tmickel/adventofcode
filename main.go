package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

const Year = "2021"

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("please use aoc day number as first arg")
	}
	day := args[0]
	log.Printf("generating day %s", day)

	cookieBytes, err := os.ReadFile("./cookie.txt")
	if err != nil {
		log.Fatal("please put session cookie in cookie.txt")
	}
	cookie := strings.TrimSpace(string(cookieBytes))

	if err := mkdir(day); err != nil {
		log.Printf("mkdir: %v (skipping)", err)
	} else {
		log.Printf("🎄 created q%s", day)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	u, err := url.Parse("https://adventofcode.com")
	if err != nil {
		log.Fatal(err)
	}

	jar.SetCookies(u, []*http.Cookie{
		{
			Name:  "session",
			Value: cookie,
		},
	})

	client := &http.Client{
		Jar: jar,
	}

	if err := fetchAndWriteExample(client, day); err != nil {
		log.Printf("failed to write example: %v (skipping)", err)
	} else {
		log.Printf("🎄 wrote q%v/example.txt", day)
	}

	if err := fetchAndWriteInput(client, day); err != nil {
		log.Printf("failed to write input: %v (skipping)", err)
	} else {
		log.Printf("🎄 wrote q%v/input.txt", day)
	}

	if err := writeTemplate(day); err != nil {
		log.Printf("failed to write main.go: %v (skipping)", err)
	} else {
		log.Printf("🎄 wrote q%v/main.go", day)
	}

	execCount := 1
	partOne := "0"
	partTwo := "0"
	partOne, partTwo, err = exe(day, execCount)
	if err != nil {
		log.Printf("execution failed with err: %v", err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					execCount++
					partOne, partTwo, err = exe(day, execCount)
					if err != nil {
						log.Printf("execution failed with err: %v", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fmt.Sprintf("./q%s", day))
	if err != nil {
		log.Fatal(err)
	}

	input := make(chan int)
	go func() {
		for {
			var u int
			_, err := fmt.Scanf("%d\n", &u)
			if err != nil {
				log.Println("??? allowed commands: 1, 2 to submit each part respectively")
			}
			input <- u
		}
	}()

	for {
		select {
		case <-done:
			log.Println("finished")
			break
		case cmd := <-input:
			if cmd == 1 {
				submit(client, day, "1", partOne)
				continue
			}
			if cmd == 2 {
				submit(client, day, "2", partTwo)
				continue
			}
			log.Println("??? allowed commands: 1, 2 to submit each part respectively")
		}
	}
}

func mkdir(day string) error {
	if err := os.Mkdir(fmt.Sprintf("q%s", day), 0755); err != nil {
		return err
	}
	return nil
}

func fetchAndWriteExample(client *http.Client, day string) error {
	path := fmt.Sprintf("q%s/example.txt", day)
	if _, err := os.Stat(path); err == nil {
		return errors.New("example.txt already exists")
	}
	promptResp, err := client.Get(fmt.Sprintf("https://adventofcode.com/%s/day/%s", Year, day))
	promptBytes, err := io.ReadAll(promptResp.Body)
	if err != nil {
		return err
	}
	prompt := string(promptBytes)
	p := strings.Split(prompt, "<pre><code>")
	if len(p) < 2 {
		return errors.New("could not find example input")
	}
	p2 := strings.Split(p[1], "</code></pre>")
	if len(p) < 2 {
		return errors.New("could not find example input")
	}
	example := p2[0]
	if err := os.WriteFile(path, []byte(example), 0755); err != nil {
		return fmt.Errorf("could not write example.txt: %v", err)
	}
	return nil
}

func fetchAndWriteInput(client *http.Client, day string) error {
	path := fmt.Sprintf("q%s/input.txt", day)
	if _, err := os.Stat(path); err == nil {
		return errors.New("input.txt already exists")
	}
	resp, err := client.Get(fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", Year, day))
	inputBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, inputBytes, 0755); err != nil {
		return fmt.Errorf("could not write input.txt: %v", err)
	}
	return nil
}

func writeTemplate(day string) error {
	path := fmt.Sprintf("q%s/main.go", day)
	if _, err := os.Stat(path); err == nil {
		return errors.New("main.go already exists")
	}
	template := `package main
import (
	"fmt"
	"log"

	"github.com/tmickel/advent2021/fileutil"
)
	
func main() {
	partOne()
	partTwo()
}

func partOne() {
	input, err := fileutil.ScanStrings("example.txt")
	// input, err := fileutil.ScanStrings("input.txt")
	_ = input
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	fmt.Println(result)
}

func partTwo() {
	input, err := fileutil.ScanStrings("example.txt")
	// input, err := fileutil.ScanStrings("input.txt")
	_ = input
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	fmt.Println(result)	
}
	`
	if err := os.WriteFile(path, []byte(template), 0755); err != nil {
		return fmt.Errorf("could not write main.go: %v", err)
	}
	return nil
}

func exe(day string, execCount int) (string, string, error) {
	fmt.Printf("------- exec #%v -------\n", execCount)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = fmt.Sprintf("./q%s", day)
	start := time.Now()
	out, err := cmd.CombinedOutput()
	duration := time.Since(start)
	if err != nil {
		return "", "", err
	}
	results := string(out)
	parts := strings.Split(results, "\n")
	fmt.Printf("\033[31mpart 1:\033[0m %s\n", parts[0])
	fmt.Printf("\033[32mpart 2:\033[0m %s\n", parts[1])
	fmt.Printf("------- 🎄 finished in %s. save main.go for another; press 1+enter to submit part 1 (or 2 for 2) -------\n", duration)
	return parts[0], parts[1], nil
}

func submit(client *http.Client, day string, part string, answer string) error {
	fmt.Printf("submitting day %s, part %s with answer: %s\n", day, part, answer)
	resp, err := client.PostForm(fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", Year, day),
		url.Values{
			"level":  []string{part},
			"answer": []string{answer},
		})
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := string(body)
	if strings.Contains(result, "not the right answer") {
		fmt.Println("\033[33mwrong answer :(\033[0m")
		return nil
	}
	if strings.Contains(result, "gave an answer too recently") {
		fmt.Println("\033[32mrate limited :(\033[0m")
		return nil
	}
	if strings.Contains(result, "That's the right answer") {
		fmt.Println("\033[32m🎄🎄correct! :)🎄🎄\033[0m")
		return nil
	}
	fmt.Println(result)
	return nil
}
