## slp-ranked-bot
A slippi ranked leaderboard for discord

### Description
Still pretty rough around the edges but the following things work:
1. `!rankadd <connect code>` to add a user to the leaderboard
2. `!board` to display the list of users, sorted by rating in descending order

### TODO
The following things still need work, roughly sorted by priority:
* [infra] Where should we host this? I've only ever run this locally for testing, but that won't work if we want it on the SRG discord
* [improvement] Write tests
* [improvement] Change `fetchUserData` such that it can detect when a given connect code doesn't exist. I haven't tested thoroughly and this would involve inspecting the response data from the graphql endpoint and returning an error if the user doesn't exist.
* [refactor] I don't like that each command handler instantiates its own `FileBasedUserManager`, this should probably only happen once at startup and then the same instance should be passed around via something like a [context](https://pkg.go.dev/context).
* [improvement] I made a `UserManager` interface, but only have a local file implementation. This means that wherever we host this will need to support a persistent file system. We might consider using an external db. (I think google sheets could be an option)
* [improvement] Add pagination to `!board`: user should be able to specify which page of the board they want to see
* [improvement] Add a new command: `!rankremove <user>`
* [improvement] Make the board prettier, have the bot write markdown?

### How to run
1. clone the repo: `git clone git@github.com:trstruth/slp-ranked-bot.git`
2. get the token from me and write it to a file in the root of the repo called `token.txt`
3. `go run .`
