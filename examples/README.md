# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. Get the code:

    ```console
    $ go get google.golang.org/grpc/examples/todolist/todolist_client
    $ go get google.golang.org/grpc/examples/todolist/todolist_server
    ```

 2. Run the server:

    ```console
    $ $(go env GOPATH)/bin/todolist_server &
    ```

 3. Run the client:

    ```console
    $ $(go env GOPATH)/bin/todolist_client
    Greeting: Hello world
    ```

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart
