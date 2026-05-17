package main

import "fmt"

func FormatAndPrintActivity(events []Event){
	if len(events)==0{
		fmt.Println("No recent activity found.")
		return
	}

	fmt.Println("Output:")
	for _,event:=range events{
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed %d commits to %s\n", len(event.Payload.Commits), event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- %s a new issue in %s\n", capitalize(event.Payload.Action), event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)
		case "CreateEvent":
			fmt.Printf("- Created a new repository/branch in %s\n", event.Repo.Name)
		default:
			// Fallback for unhandled event types
			fmt.Printf("- %s in %s\n", event.Type, event.Repo.Name)
		}
	}
}

func capitalize(s string) string {
	if len(s)==0{
		return s
	}
	return string(s[0]-32) + s[1:]
}
