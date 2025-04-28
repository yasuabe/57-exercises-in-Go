// ## Ex53: Todo List
// - Command-line todo list app.
// - Store data in Redis.
//
// ### Commands:
// - add: <task>  – Add a new task (empty tasks are not allowed)
// - list         – Show all tasks
// - remove: <id> – Remove completed task by ID
// - exit         – Quit the app

package main

import (
	"bufio"
	"context" // Use the standard library's context package
	"fmt"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

const redisURL = "redis://127.0.0.1:6379/"
const redisKey = "ex53:tasks"

var ctx = context.Background()

func main() {
	// Connect to Redis
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		fmt.Println("Error parsing Redis URL:", err)
		return
	}
	client := redis.NewClient(options)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Todo List App (type 'exit' to quit)")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		parts := strings.SplitN(input, " ", 2)
		command := parts[0]

		switch command {
		case "add:":
			if len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
				fmt.Println("Error: Task cannot be empty.")
				continue
			}
			task := parts[1]
			err := addTask(client, task)
			if err != nil {
				fmt.Println("Error adding task:", err)
			} else {
				fmt.Println("Task added.")
			}

		case "list":
			err := listTasks(client)
			if err != nil {
				fmt.Println("Error listing tasks:", err)
			}

		case "remove:":
			if len(parts) < 2 {
				fmt.Println("Error: Task ID is required.")
				continue
			}
			id := parts[1]
			err := removeTask(client, id)
			if err != nil {
				fmt.Println("Error removing task:", err)
			} else {
				fmt.Println("Task removed.")
			}

		default:
			fmt.Println("Unknown command. Available commands: add, list, remove, exit.")
		}
	}
}

func addTask(client *redis.Client, task string) error {
	// Increment the task ID
	id, err := client.Incr(ctx, "ex53:task").Result()
	if err != nil {
		return err
	}

	// Add the task to the hash
	_, err = client.HSet(ctx, redisKey, fmt.Sprintf("%d", id), task).Result()
	return err
}

func listTasks(client *redis.Client) error {
	// Get all tasks from the hash
	tasks, err := client.HGetAll(ctx, redisKey).Result()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	fmt.Println("Tasks:")
	for id, task := range tasks {
		fmt.Printf("%s: %s\n", id, task)
	}
	return nil
}

func removeTask(client *redis.Client, id string) error {
	// Remove the task from the hash
	_, err := client.HDel(ctx, redisKey, id).Result()
	return err
}
