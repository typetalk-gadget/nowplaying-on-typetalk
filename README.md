# nowplaying-on-typetalk

A command line tool to display the song currently playing on Spotify in Typetalk status.

## Motivation

We want to share the track we are listening to at work in a discreet and easy way.
If we post to a topic in Typetalk, we have to open that topic and that post is noise.
If it's a status, we can just hover over any topic, and it's not noise.

## Installation

### GoBinaries

You can easily install it by making a curl request to [gobinaries.com](http://gobinaries.com/). You don't have to install Go.

```sh
$ curl -sf https://gobinaries.com/typetalk-gadget/nowplaying-on-typetalk | sh
```

### Go

If you have the Go(go1.14+) installed, you can also install it with go get command.

```sh
$ go get github.com/typetalk-gadget/nowplaying-on-typetalk
```

## Synopsis

```sh
$ nowplaying-on-typetalk [flags]
```

## Flags

```sh
  -c, --config string                   config file path (default "~/.nowplaying-on-typetalk/config.yml")
      --debug                           debug mode
  -h, --help                            help for nowplaying-on-typetalk
      --spotify_client_id string        spotify client id [SPOTIFY_CLIENT_ID]
      --spotify_client_secret string    spotify client secret [SPOTIFY_CLIENT_SECRET]
      --status_emoji string             typetalk status emoji [STATUS_EMOJI] (default ":musical_note:")
      --typetalk_client_id string       typetalk client id [TYPETALK_CLIENT_ID]
      --typetalk_client_secret string   typetalk client secret [TYPETALK_CLIENT_SECRET]
      --typetalk_space_key string       typetalk space key [TYPETALK_SPACE_KEY]
  -v, --version                         version for nowplaying-on-typetalk
```

## Config File

### YAML

```yaml
debug: true
typetalk_client_id: deadbeef
typetalk_client_secre: deadcode
typetalk_space_key: foo
spotify_client_id: deadbeef
spotify_client_secret: deadcode
status_emoji: ":musical_note:"
```
## Bugs and Feedback

For bugs, questions and discussions please use the GitHub Issues.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)

## Author

* [vvatanabe](https://github.com/vvatanabe)