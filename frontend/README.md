# Vugu/Todo/Async Demo - Frontend

The Frontend is based on the Vugu Go Frontend development framework. It generates
a WASM binary which can be run in modern browsers. The Vugu toolchain includes
a development server which can serve the content locally.

It is assumed that a sane Go development environment is installed.

# Modifying the Frontend to point to the Backend

The location of the backend Websockets service is defined in the file `root.go`.
This file should be modified to contain the Websockets URI which was output when
the backend was configured.

# Running the Frontend

## Installing Vugu tools

If necessary, install the `vgrun` tool as follows:

```sh
go get -u github.com/vugu/vgrun
```

# Running application from the devserver

Do the following to run the Frontend:

```sh
vgrun devserver.go
```

Then browse to the running server: http://localhost:8844/

You should now be able to add simple Todos, change their completed state and
remove them and the information should be persisted in the DynamoDB tables.
