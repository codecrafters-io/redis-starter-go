This is a starting point for Golang solutions to the CodeCrafters
Redis Challenge.

In this challenge, you'll build a toy Redis clone that's capable of handling
basic commands like `PING`, `SET` and `GET`. Along the way we'll learn about
event loops, the Redis protocol and more. 

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to start the challenge.

# Usage

1. Ensure you have `go` installed locally
2. Run `make install`, which'll install the required dependencies.
3. Run `make run_local_server` to run your Redis server, which is implemented in
   `app/server.go`.
3. Run `make test` to run local tests, which are located in `app/server_test.go`
4. Commit your changes and run `git push origin master` to submit your solution
   to CodeCrafters. Test output will be streamed to your terminal.
5. Bump the `current_stage` value in `.codecrafters.yml`, and run `git push
   origin master` to advance to the next stage.

# Passing the first stage

CodeCrafters runs tests when you do a `git push`. Make an empty commit and push
your solution to see the first stage fail.
   
``` sh
git commit --allow-empty -m "Running tests"
git push origin master
```

You should see a failure message that says it wasn't able to connect to port
`6379`.

These tests take a while to run, so we've built a basic local testing setup for
you at `app/server_test.go`, where you can write your own tests and run them.

We've already added a local test for the first stage, let's invoke it by running
`make test`. You should see a failure message that looks like this: 
   
```sh
============== FAILURES =============
______ test_can_connect_to_6379 _____

    def test_can_connect_to_6379():
        with spawn_server():
            # Shouldn't throw an error
>           socket.create_connection(("localhost", 6379))

    ...
    ...
    ...

>               sock.connect(sa)
E               ConnectionRefusedError: [Errno 111] Connection refused

/usr/lib64/python3.8/socket.py:796: ConnectionRefusedError
```
   
Let's fix the test by implementing a socket in `app/server.go`
 
```diff
--- app/main.py	2020-01-20 10:29:06.254902363 +0530
+++ app/main.py	2020-01-20 10:23:36.160486553 +0530
@@ -1,3 +1,4 @@
+import socket
 import time
 
 
@@ -5,6 +6,9 @@
     # Implement your server here
     print("hey")
 
+    s = socket.create_server(("localhost", 6379))
+    s.accept()  # Wait for a new connection
+
 
 if __name__ == "__main__":
     main()
```

Run `make test` again, and the tests should now pass!
   
```
$ make test
pipenv run pytest tests
============================= test session starts ====================
platform linux -- Python 3.8.1, pytest-5.3.3, py-1.8.1, pluggy-0.13.1
rootdir: /redis-solution-starter-py
collected 1 item                                                                                                                                                                                                                              

tests/test_main.py .                                                                                                                                                                                                                    [100%]

============================== 1 passed in 0.13s =====================
```

Now it's time to submit our result to CodeCrafters! Commit your changes and run
`git push origin master`.

Time to move on to the next stage! Bump the `current_stage` value in
`.codecrafters.yml` and run `git push origin master` again.

