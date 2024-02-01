# gopt-renamer

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/robin-collins/gopt-renamer">
    <img src="images/github_logo.png" alt="Logo" width="150" height="150">
  </a>

  <h3 align="center">gopt-renamer</h3>

  <p align="center">
    A Golang command-line application that auto-renames images based on content by sending to a GPT-vision API.
  </p>
  <p align="center">
    <img src="https://raw.githubusercontent.com/robin-collins/gopt-renamer-media/main/media/in_action.gif" alt="Watch as gopt-renamer quickly renames screenshots" >
  </p>
  <p align="center">
    <br />
    <a href="https://github.com/robin-collins/gopt-renamer"><strong>Explore the docs »</strong></a>
    <br />
    ·
    <a href="https://github.com/robin-collins/gopt-renamer/issues">Report Bug</a>
    ·
    <a href="https://github.com/robin-collins/gopt-renamer/issues">Request Feature</a>
  </p>
</p>

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
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

gopt-renamer is a tool that leverages the power of GPT-vision API to automatically generate descriptive and informative file names for images. It is designed to streamline the process of organizing and searching for images by providing meaningful names that reflect the content of the images.

### Built With

* [Golang](https://golang.org/)

## Installation

Download a release package for your platform of choice from the releases. 

### Linux Installation

```sh
sudo dpkg -i gopt-renamer-ubuntu-latest.deb
```

### Mac Installation

Click the .pkg file and follow the onscreen prompts

### Windows Installation

Click the .exe file and follow the onscreen prompts

<!-- GETTING STARTED TO BUILD-->
## Getting Started to build manually

To get a local copy up and running follow these simple steps.

### Prerequisites

* Golang installed on your machine. You can download it from [here](https://golang.org/dl/).

### Build Installation

1. Clone the repo

```sh
git clone https://github.com/robin-collins/gopt-renamer.git
```

2. Build the project

```sh
go build 
```

<!-- USAGE EXAMPLES -->
## Usage

To use gopt-renamer, simply provide the path to the image file you wish to rename:

```sh
./gopt-renamer --image="path/to/image.jpg"
```

_For more examples, please refer to the [Documentation](https://github.com/robin-collins/gopt-renamer)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

* [x] Functional Windows, Ubuntu and Mac builds
* [x] Ollama (llava and baklavva) tested and found to be not suitable.
* [x] OpenAI API Key saves to either environment variable OR .conf file
* [ ] Create Docs and configure deployment
* [ ] CI/CD ready for release
* [ ] Configurable API endpoint URL via system environment
* [ ] Better Error Handling
* [x] Windows Installer
* [ ] Mac package
* [ ] Ubuntu package
* [ ] Support Ubuntu and Mac context menu / gui integration

See the [open issues](https://github.com/robin-collins/gopt-renamer/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Robin Collins - [@robin_collins](https://twitter.com/RobinFCollins) - <robin.f.collins+gopt@Outlook.com>

Project Link: [https://github.com/robin-collins/gopt-renamer](https://github.com/robin-collins/gopt-renamer)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [OpenAI](https://openai.com/)
* [othneildrew](https://github.com/othneildrew/Best-README-Template)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[license-shield]: https://img.shields.io/github/license/robin-collins/gopt-renamer.svg?style=for-the-badge
[license-url]: https://github.com/robin-collins/gopt-renamer/blob/master/LICENSE
