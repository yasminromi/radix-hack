run:
	go run main.go

mock-ons:
	curl

commit-deploy:
	git commit -am '$(COMMIT_MSG)'
	git push
	git remote add heroku 

