# Using mods - IA

## Requirements

Here we define some guidelines to use [mods by charmbracelet](https://github.com/charmbracelet/mods).

Installing mods:

```bash
brew install charmbracelet/tap/mods
```

Once it is installed, you can do:
```bash
mods -v
```

We need a OpenAI API KEY for this on this link [https://platform.openai.com/api-keys](https://platform.openai.com/api-keys).


In this example, it is: `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`.

With this API key, we need to export to a env variable:
```bash
export OPENAI_API_KEY=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
```

## Examples

1. Create a minimal Go project

```bash
go mod init get-time-cli-cobra

mods "Give me the main.go for a CLI in Golang, using Cobra, that has a command 'time' that returns the current hour, minute, and second. Just Golang code." > main.go

go mod tidy

cat main.go | mods "Create a Dockerfile for this Golang App. Just the Dockerfile content" > Dockerfile

ls | mods "Given this folder structure, create a Makefile to build a Docker image with the title get-time cli, and a command in the Makefile to run the built image. To run the command line we need to pass the 'time' argument" > Makefile
```

2. Create a git diff

```bash
git diff --cached | mods "Create a short but descriptive commit title, starting in lowercase, for the following diff. Please use conventional commits format."

git diff --cached | mods -f "create a Pull Request body"
```


3. Create documentation for the current folder

```bash
cat * | mods -f "create a README for all these files" > README.md
```




