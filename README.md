<a id="readme-top"></a>

[![project_license][license-shield]][license-url]
[![Go][Go]][Go-url]
<br />
<div align="center">

<h3 align="center">Pokedex</h3>

</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>


## About The Project

Pokedex is a REPL application written in Go for learning purposes of Go. It connects to pokeAPI and allows user to search locations for pokemons.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

You need to have Go installed on your local machine
* Go
  ```sh
  go version
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/MaciejSieradz/pokedexcli.git
   ```
3. Build with go
   ```sh
   go build
   ```
4. Run REPL application 
   ```sh
   ./pokedexcli
   ```
5. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin github_username/repo_name
   git remote -v # confirm the changes
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage

You can get list of available commands using `help` commmand

```sh
Pokedex > help

Welcome to the Pokedex!
Usage:

mapb: Return list of previous 20 locations
explore: Explore the location to list all the Pokemon
catch: Try to catch a Pokemon
inspect: Inspect caught Pokemon
pokedex: Display names of your caught Pokemons
help: Displays a help message
exit: Exit the Pokedex
map: Return list of next 20 locations

Pokedex > 
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License

Distributed under the MIT license. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[license-url]: https://github.com/MaciejSieradz/pokedexcli/blob/main/LICENSE
[license-shield]: https://img.shields.io/github/license/MaciejSieradz/pokedexcli.svg?style=for-the-badge
[Go-url]: https://go.dev/
[Go]: https://img.shields.io/badge/logo-Go-blue?logo=Go?style=for-the-badge
