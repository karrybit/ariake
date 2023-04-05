# ariake

ariake is utility tool to support development for google cloud pubsub in local environment

# Requirement

- Python
- JDK
- Google Cloud CLI

# Installation

```sh
gcloud components install pubsub-emulator
gcloud components update
go install github.com/karrybit/ariake@latest
```

# Help

```sh
ariake is utility tool to support development for google cloud pubsub in local environment

Usage:
  ariake [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create a topic to register the emulator
  delete      Delete the topic in the emulator
  help        Help about any command
  print       Print the list of the topics registered in the emulator
  publish     Publish a topic to the emulator
  reset       Reset all topic and subscription
  subscribe   Subscribe to topics in the emulator

Flags:
  -h, --help   help for ariake

Use "ariake [command] --help" for more information about a command.
```

# Example

```sh
# terminal 1
gcloud beta emulators pubsub start --project=dummy
```

```sh
# terminal 2
ariake create -t hgoe
ariake subscribe -t hoge
```

```sh
# terminal 3
ariake publish -t hoge -m '{"hello":"world"}'
```

```sh
# output in terminal 2
{
  "hello": "world"
}
```
