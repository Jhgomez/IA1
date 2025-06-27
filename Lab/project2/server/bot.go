package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	tg_key       = os.Getenv("TG_API_KEY")
	trello_key   = os.Getenv("TRELLO_API_KEY")
	trello_token = os.Getenv("TRELLO_TOKEN")
	board_id     = os.Getenv("TRELLO_BOARD_ID")
)

func main() {
	bot, err := tgbotapi.NewBotAPI(tg_key)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// if for any reason a webhook was already set up you have to delete it here before you
	// can start polling
	if _, err := bot.Request(tgbotapi.DeleteWebhookConfig{}); err != nil {
		log.Fatalf("couldn't delete webhook: %v", err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /historia and /status."
		case "historia":

			// Build the query parameters
			params := url.Values{}
			params.Set("key", trello_key)
			params.Set("token", trello_token)

			urlWithParams := fmt.Sprintf("https://api.trello.com/1/boards/%s/lists?%s", board_id, params.Encode())

			// urlWithParams := fmt.Sprintf("%s?%s", endpoint, params.Encode())

			resp, err := http.Get(urlWithParams)

			if err != nil {
				log.Fatalf("request failed: %v", err)
			}

			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()

			if err != nil {
				fmt.Printf("error reading body %v", err)
			}

			var data1 []map[string]interface{}

			if err := json.Unmarshal(body, &data1); err != nil {
				fmt.Println(urlWithParams)
				fmt.Printf("error unmarshalling board lists JSON %v\n", err)
			}

			var list string

			for _, board := range data1 {
				if board["name"] == "todo" {
					list = fmt.Sprintf("%s", board["id"])
				}
			}

			if list == "" {
				fmt.Println(urlWithParams)
				fmt.Println("no list with name todo")
				msg.Text = "Unable to create history"
				break
			}

			parts := strings.Split(update.Message.Text, "\n")

			// fmt.Println(strings.TrimSpace(parts[0][9:]))
			// fmt.Println(strings.TrimSpace(parts[1]))
			fmt.Println(strings.TrimSpace(parts[1]))
			fmt.Println(strings.TrimSpace(parts[2]))

			params.Set("idList", list)
			params.Set("name", strings.TrimSpace(parts[1]))
			params.Set("desc", strings.TrimSpace(parts[2]))

			// Construct the full URL with query string
			urlWithParams = fmt.Sprintf("https://api.trello.com/1/cards?%s", params.Encode())
			// urlWithParams = fmt.Sprintf("%s?%s", endpoint, params.Encode())

			// Create a POST request (with no body, since we're sending everything in the URL)
			req, err := http.NewRequest(http.MethodPost, urlWithParams, nil)
			if err != nil {
				log.Fatalf("Error building request: %v", err)
			}

			// (Optional) Set a custom User-Agent
			req.Header.Set("User-Agent", "my-go-trello-client/1.0")

			// Perform the request
			resp, err = http.DefaultClient.Do(req)
			if err != nil {
				log.Fatalf("Request failed: %v", err)
			}
			defer resp.Body.Close()

			// Check for non-200 status
			if resp.StatusCode != http.StatusOK {
				log.Fatalf("Trello API returned %s", resp.Status)
			}

			body, err = io.ReadAll(resp.Body)
			defer resp.Body.Close()

			if err != nil {
				fmt.Printf("error reading body %v", err)
			}

			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err != nil {
				fmt.Printf("error unmarshalling JSON %v", err)
			}

			// fmt.Printf("✅ Created card!\n  ID:   %s\n  Name: %s\n  URL:  %s\n", data["id"], data["name"], data["url"])
			msg.Text = fmt.Sprintf("✅ Card Created!, copy and paste next message to move to progress. URL: %s", data["url"])

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

			msg.Text = fmt.Sprintf("/move\nlist:\n<in_progress or done>\nID:\n%s\nName:\n%s", data["id"], data["name"])

		case "move":
			// Build the query parameters
			params := url.Values{}
			params.Set("key", trello_key)
			params.Set("token", trello_token)

			urlWithParams := fmt.Sprintf("https://api.trello.com/1/boards/%s/lists?%s", board_id, params.Encode())

			// urlWithParams := fmt.Sprintf("%s?%s", endpoint, params.Encode())

			resp, err := http.Get(urlWithParams)

			if err != nil {
				log.Fatalf("request failed: %v", err)
			}

			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()

			if err != nil {
				fmt.Printf("error reading body %v", err)
			}

			var data1 []map[string]interface{}

			if err := json.Unmarshal(body, &data1); err != nil {
				fmt.Println(urlWithParams)
				fmt.Printf("error unmarshalling board lists JSON %v\n", err)
			}

			var list string

			parts := strings.Split(update.Message.Text, "\n")

			for _, board := range data1 {
				if board["name"] == parts[2] {
					list = fmt.Sprintf("%s", board["id"])
				}
			}

			if list == "" {
				fmt.Println(urlWithParams)
				msg.Text = "Unable to move card, invalid list name"
				break
			}

			params.Set("idList", list)

			if len(parts) < 7 {
				update.Message.Text = "Please enter all the command information, card ID and card name"
			}

			urlWithParams = fmt.Sprintf("https://api.trello.com/1/cards/%s?%s", parts[4], params.Encode())

			// Create a POST request (with no body, since we're sending everything in the URL)
			req, err := http.NewRequest(http.MethodPut, urlWithParams, nil)
			if err != nil {
				log.Fatalf("Error building request: %v", err)
			}

			// (Optional) Set a custom User-Agent
			req.Header.Set("User-Agent", "my-go-trello-client/1.0")

			// Perform the request
			resp, err = http.DefaultClient.Do(req)
			if err != nil {
				log.Fatalf("Request to move to in progress failed: %v", err)
			}
			defer resp.Body.Close()

			// Check for non-200 status
			if resp.StatusCode != http.StatusOK {
				log.Fatalf("Trello API returned %s", resp.Status)
			}

			update.Message.Text = fmt.Sprintf("Card \"%s\" is in progress now", parts[4])

			body, err = io.ReadAll(resp.Body)
			defer resp.Body.Close()

			if err != nil {
				fmt.Printf("error reading body %v", err)
			}

			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err != nil {
				fmt.Printf("error unmarshalling JSON %v", err)
			}

			if data["name"] != parts[6] {
				msg.Text = fmt.Sprintf("Actual name of card is \"%s\" and its now in %s ✅ ", data["name"], parts[2])
			} else {
				msg.Text = fmt.Sprintf("Card \"%s\" is now in progress ✅", parts[6])
			}

		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
