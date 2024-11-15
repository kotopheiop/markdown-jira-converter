package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

// Конвертация Markdown в Jira-синтаксис
func convertMdToJira(content string) string {
	replacements := []struct {
		regex *regexp.Regexp
		repl  string
	}{
		{regexp.MustCompile(`######\s*(.+)`), "h6. $1"},
		{regexp.MustCompile(`#####\s*(.+)`), "h5. $1"},
		{regexp.MustCompile(`####\s*(.+)`), "h4. $1"},
		{regexp.MustCompile(`###\s*(.+)`), "h3. $1"},
		{regexp.MustCompile(`##\s*(.+)`), "h2. $1"},
		{regexp.MustCompile(`#\s*(.+)`), "h1. $1"},
		{regexp.MustCompile(`\*\*(.+?)\*\*`), "*$1*"},
		{regexp.MustCompile(`_(.+?)_`), "_$1_"},
		{regexp.MustCompile(`~~(.+?)~~`), "-$1-"},
		{regexp.MustCompile(`\[(.+?)\]\((.+?)\)`), "[$1|$2]"},
		{regexp.MustCompile(`^\*\s+`), "- "},
		{regexp.MustCompile("`(.+?)`"), "{{$1}}"},
	}

	for _, r := range replacements {
		content = r.regex.ReplaceAllString(content, r.repl)
	}
	return content
}

// Конвертация Jira-синтаксиса в Markdown
func convertJiraToMd(content string) string {
	replacements := []struct {
		regex *regexp.Regexp
		repl  string
	}{
		{regexp.MustCompile(`h6\.\s*(.+)`), "###### $1"},
		{regexp.MustCompile(`h5\.\s*(.+)`), "##### $1"},
		{regexp.MustCompile(`h4\.\s*(.+)`), "#### $1"},
		{regexp.MustCompile(`h3\.\s*(.+)`), "### $1"},
		{regexp.MustCompile(`h2\.\s*(.+)`), "## $1"},
		{regexp.MustCompile(`h1\.\s*(.+)`), "# $1"},
		{regexp.MustCompile(`\*(.+?)\*`), "**$1**"},
		{regexp.MustCompile(`_(.+?)_`), "_$1_"},
		{regexp.MustCompile(`-(.+?)-`), "~~$1~~"},
		{regexp.MustCompile(`\[(.+?)\|(.+?)\]`), "[$1]($2)"},
		{regexp.MustCompile(`^- `), "* "},
		{regexp.MustCompile(`\{\{(.+?)\}\}`), "`$1`"},
	}

	for _, r := range replacements {
		content = r.regex.ReplaceAllString(content, r.repl)
	}
	return content
}

// Конвертация многострочных блоков кода
func convertCodeBlocks(content string, toJira bool) string {
	re := regexp.MustCompile("(?s)```(\\w+)?\\n(.*?)\\n```")
	return re.ReplaceAllStringFunc(content, func(match string) string {
		matches := re.FindStringSubmatch(match)
		lang := matches[1]
		code := matches[2]
		if toJira {
			if lang != "" {
				return fmt.Sprintf("{code:%s}\n%s\n{code}", lang, code)
			}
			return fmt.Sprintf("{code}\n%s\n{code}", code)
		} else {
			if lang != "" {
				return fmt.Sprintf("```%s\n%s\n```", lang, code)
			}
			return fmt.Sprintf("```\n%s\n```", code)
		}
	})
}

func main() {
	inputFile := flag.String("input", "", "Путь к входному файлу")
	outputFile := flag.String("output", "", "Путь к выходному файлу")
	from := flag.String("from", "md", "Формат исходного файла: md или jira")
	to := flag.String("to", "jira", "Формат выходного файла: md или jira")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		fmt.Println("⛔ Ошибка: укажите пути к входному и выходному файлам с помощью параметров -input и -output.")
		flag.Usage()
		os.Exit(1)
	}

	// Чтение входного файла
	content, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("⛔ Ошибка чтения файла: %v\n", err)
		os.Exit(1)
	}

	// Преобразование содержимого
	var result string
	contentStr := string(content)
	contentStr = convertCodeBlocks(contentStr, *to == "jira")

	if *from == "md" && *to == "jira" {
		result = convertMdToJira(contentStr)
	} else if *from == "jira" && *to == "md" {
		result = convertJiraToMd(contentStr)
	} else {
		fmt.Println("⛔ Ошибка: поддерживаются только конверсии md-to-jira или jira-to-md.")
		os.Exit(1)
	}

	// Запись в выходной файл
	err = os.WriteFile(*outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("⛔ Ошибка записи файла: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Конвертация успешно завершена 🎉🎉🎉")
}
