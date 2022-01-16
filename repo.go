package main

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func createRepo(name string, token string) {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	oauthClient := oauth2.NewClient(ctx, tokenSource)
	client := github.NewClient(oauthClient)

	isPrivate := true
	description := "This is a demo web service to be deployed on k8s clusters created as part of onboarding"

	repository := &github.Repository{Name: &name, Private: &isPrivate, Description: &description}
	repo, _, err := client.Repositories.Create(ctx, "", repository)
	if err != nil {
		log.Println(err)
	}
	log.Println("Suceessfully created new repo: ", repo.GetName())
}

func deleteRepo(name string, token string, owner string) {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	oauthClient := oauth2.NewClient(ctx, tokenSource)
	client := github.NewClient(oauthClient)
	_, err := client.Repositories.Delete(ctx, owner, name)
	if err != nil {
		log.Println(err)
	}
	log.Println("Suceessfully deleted repo: ", name)
}
