
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yanzay/tbot/v2"
)

func escapeMarkdown(s string) string {
	replacer := strings.NewReplacer(
		"\\", "\\\\",
		"_", "\\_",
		"*", "\\*",
		"`", "\\`",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"-", "\\-",
		".", "\\.",
		"!", "\\!",
	)
	return replacer.Replace(s)
}

func eventName(event string) string {
	names := map[string]string{
		"issue_comment":               "Комментарий к Issue",
		"issues":                      "Issue",
		"pull_request":                "Pull request",
		"pull_request_review_comment": "Комментарий к pull request",
		"push":                        "Изменения в коде",
		"project_card":                "Карточка проекта",
	}

	if name, ok := names[strings.ToLower(event)]; ok {
		return name
	}
	return event
}

func statusIcon(status string) string {
	icons := map[string]string{
		"failure":   "❗️",
		"cancelled": "⚠️",
		"success":   "✅",
	}

	if icon, ok := icons[strings.ToLower(status)]; ok {
		return icon
	}
	return "ℹ️"
}

func linkgen(repo, event, issueNumber string) string {
	switch strings.ToLower(event) {
	case "issue_comment", "issues":
		if issueNumber != "" {
			return fmt.Sprintf("https://github.com/%s/issues/%s", repo, issueNumber)
		}
		return fmt.Sprintf("https://github.com/%s/issues", repo)
	case "pull_request", "pull_request_review_comment":
		if issueNumber != "" {
			return fmt.Sprintf("https://github.com/%s/pull/%s", repo, issueNumber)
		}
		return fmt.Sprintf("https://github.com/%s/pulls", repo)
	case "push":
		return fmt.Sprintf("https://github.com/%s/commits", repo)
	default:
		return fmt.Sprintf("https://github.com/%s", repo)
	}
}

func main() {
	token := os.Getenv("INPUT_TOKEN")
	chat := os.Getenv("INPUT_CHAT")
	status := os.Getenv("INPUT_STATUS")
	event := os.Getenv("INPUT_EVENT")
	actor := os.Getenv("INPUT_ACTOR")

	workflow := os.Getenv("GITHUB_WORKFLOW")
	repo := os.Getenv("GITHUB_REPOSITORY")

	issueNumber := os.Getenv("ISSUE_NUMBER")
	issueTitle := os.Getenv("ISSUE_TITLE")

	if issueNumber == "" {
		issueNumber = os.Getenv("INPUT_ISSUE_NUMBER")
	}
	if issueTitle == "" {
		issueTitle = os.Getenv("INPUT_ISSUE_TITLE")
	}

	msg := statusIcon(status) + " *Событие GitHub: " +
		escapeMarkdown(eventName(event)) + "*\n"
	msg += "Репозиторий: " + escapeMarkdown(repo) + "\n"

	if issueNumber != "" {
		msg += "Issue №" + escapeMarkdown(issueNumber)
		if issueTitle != "" {
			msg += " — " + escapeMarkdown(issueTitle)
		}
		msg += "\n"
	}

	msg += "Автор события: " + escapeMarkdown(actor) + "\n"

	if workflow != "" {
		msg += "Workflow: " + escapeMarkdown(workflow) + "\n"
	}

	link := linkgen(repo, event, issueNumber)
	msg += "[Открыть на GitHub](" + link + ")"

	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")
	_, err := c.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("Не удалось отправить сообщение в Telegram: %v", err)
	}
}
