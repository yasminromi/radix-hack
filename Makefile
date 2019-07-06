run:
	go run main.go

mock-ons:
	curl

commit-deploy:
	git add --all
	git commit -am '$(COMMIT_MSG)'
	git push
	git push heroku master 

