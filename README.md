# matrix-veles (WIP)

Veles is a bot for Matrix chatrooms, which can compare hashes of images against a list of known
spam images and mutes (or bans, your choice) the users who post them.

In the long-term Veles will become a multipurpose bot for moderation, spam-busting and
statistical work.

## Help Translating

[![Crowdin](https://badges.crowdin.net/matrix-veles/localized.svg)](https://crowdin.com/project/matrix-veles)

If you have time and language skills, consider [helping to translate the docs](https://crwd.in/matrix-veles)!

## Building

### Building without WebUI

Simply run `go build` in the project directory.

### Building with WebUI

Make sure you have NodeJS and Yarn installed.

Then run the following commands:

1. `go generate ./...` - This will build the webui react project.
2. `go build -tags withUI` - This will build the binary with included UI.
