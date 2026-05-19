# Todo CLI

A small Go command-line todo app.

## Clone the project

```bash
git clone <your-repo-url>
cd Basic_GO/todo
```

## Build the binary

```bash
go build -o todo ./cmd
```

This creates a `todo` executable in the current directory.

## Install it as a command

Move the binary into a directory that is already in your `PATH`:

```bash
sudo mv todo /usr/local/bin
```

After that, you can run it from anywhere with:

```bash
todo
```

## Run without installing

If you do not want to move it to `/usr/local/bin`, you can still run it locally:

```bash
./todo
```
