# Introduction to Go

[![license](https://img.shields.io/badge/license-CC--BY--NC--SA--4.0-yellow)](https://github.com/jub0bs/go-course-beginner/raw/main/LICENSE)

## What is this?

This course is an introduction to Go for people
who already have some programming experience.
Its purpose is to teach you not only the syntax and semantics of Go
but also the relevant idioms and best practices,
through many practical exercises and a long-running project.

The course material is intended as a handy (if not authoritative) reference
that you can rely on as you cut your teeth on the language.
The slides are chock-full of links to external learning resources
that you can peruse at your own leisure.
I update the course material regularly, as new versions of Go get released.

## How to take this course?

I typically give the course remotely to paying customers over three days.
If that would be of interest to you or your company,
feel free to DM me on [Mastodon][mastodon] or [Bluesky][bluesky].

However, with enough time and motivation,
you should be able to follow this course all the way to the end on your own.
If you find it useful, consider [sponsoring me on GitHub][sponsor].

## Disclaimer

I don't version this course material.

I reserve the right to rewrite the history of this repo at any stage.

## Instructions for running the slide deck

The course material consists of slides
(broken down into five parts, for easier navigation)
some of which contain editable and executable code samples.
You can run the slide deck on your machine after following a few simple steps:

1. If you haven't already installed Go on your machine, do so by
    following the official [installation instructions][install].

2. Important: make sure that the directory where Go installs binaries
    is in your `PATH` environment variable.
    The directory in question is usually given by the following command:

    ```shell
    echo `go env GOBIN`
    ```

3. Install the [`present`][present] tool:

    ```shell
    go install golang.org/x/tools/cmd/present@latest
    ```

4. Clone this repository and `cd` to the clone:

    ```shell
    git clone https://github.com/jub0bs/go-course-beginner
    cd go-course-beginner
    ```

5. Run the present tool:

    ```shell
    present
    ```

    The output should contain some local URL (like `http://127.0.0.1:3999`).

6. Visit the URL from step 5 in your browser,
    then click on "00_prelude.slide" to run the first series of slides.

[bluesky]: https://bsky.app/profile/jub0bs.com
[install]: https://go.dev/doc/install
[mastodon]: https://infosec.exchange/@jub0bs
[present]: https://pkg.go.dev/golang.org/x/tools/cmd/present
[sponsor]: https://github.com/sponsors/jub0bs
