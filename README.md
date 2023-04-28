# copier-run

> `copier-run` is a simple command line tool, that provides a list of available [copier](https://copier.readthedocs.io/en/stable/) templates from a GitHub user and allows to execute them.

![Go version](https://img.shields.io/github/go-mod/go-version/brpaz/copier-run?style=for-the-badge)
[![Latest Release](https://img.shields.io/github/v/release/brpaz/copier-run?style=for-the-badge](https://github.com/brpaz/copier-run/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/brpaz/copier-run/CI?style=for-the-badge)](https://github.com/brpaz/copier-run/actions/CI)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](LICENSE)

## Motivation

[copier](https://copier.readthedocs.io/en/stable/) is one of my favorite prject scaffold tools, but I was having an hard time remembering the names of my templates and the full command to run copier with any of these templates.

That´s when this tool was born. `copier-run` fetches all the copier templates from my GitHub account (repostiories that contain the topic `copier-template`), and provides a list, where I can easily execute any of my templates.
## Built With

- [Cobra](https://cobra.dev/)
- [Viper](https://github.com/spf13/viper)

## Getting started

### Pre-requisites

- [copier](https://copier.readthedocs.io/en/stable/)
- `GITHUB_TOKEN` environment variable that will be used to authenticate on GitHub.

### Installation

Check [Releases page](https://github.com/brpaz/copier-run/releases/latest) and download the latest release in the most appropriate format for your Operating System.

## Usage

Open a terminal window and run `copier-run`. The tool expects

```bash
copier-run
```

A list will be shown, with all the repositories with topic `copier-template` in your GitHub user account.

You can select the template you want to use and the destination path. After submitting all the prompts, `copier` will be executed with the specified arguments.



## Contributing

All contributions are welcome. Please check [Contributing guide](CONTRIBUTING.md) for instructions howe to contribute to this project.

## Author

👤 **Bruno Paz**

- Website: [brunopaz.dev](https://brunopaz.dev)
- Github: [@brpaz](https://github.com/brpaz)


## 📝 License

Copyright [Bruno Paz](https://github.com/brpaz).

This project is [MIT](https://opensource.org/licenses/MIT) licensed.


