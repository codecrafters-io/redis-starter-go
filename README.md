This is a starting point for Golang solutions to the CodeCrafters
Redis Challenge.

In this challenge, you'll build a toy Redis clone that's capable of handling
basic commands like `PING`, `SET` and `GET`. Along the way we'll learn about
event loops, the Redis protocol and more. 

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to start the challenge.

# Usage

1. Ensure you have `go` installed locally
1. Run `make run_local_server` to run your Redis server, which is implemented in
   `app/server.go`.
1. Commit your changes and run `git push origin master` to submit your solution
   to CodeCrafters. Test output will be streamed to your terminal.
1. Once you pass a stage, increment the `current_stage` value in
   `.codecrafters.yml`, and run `git push origin master` to advance to the next
   stage.
   


# Passing the first stage

CodeCrafters runs tests when you do a `git push`. Make an empty commit and push
your solution to see the first stage fail.
   
``` sh
git commit --allow-empty -m "Running tests"
git push origin master
```

You should see a failure message that says it wasn't able to connect to port
`6379`.

Go to `app/server.go` and uncomment the server implementation. Commit and
push your changes, and you'll now see the first stage pass.

Time to move on to the next stage! Bump the `current_stage` value in
`.codecrafters.yml` and run `git push origin master` again.

