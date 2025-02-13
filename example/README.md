# Example Service

This directory contains an example "Hello, World!" service that demonstrates how to use the `goservice` library.

## Prerequisites

-   Go installed
-   [Taskfile](https://taskfile.dev/) installed

## Building and Running

1.  Navigate to this directory:

    ```bash
    cd example
    ```

2.  Build the example service:

    ```bash
    task build
    ```

3.  Install the service (requires sudo):

    ```bash
    sudo task install
    ```

4.  Run the service in immediate mode:

    ```bash
    task run
    ```

5.  Uninstall the service (requires sudo):

    ```bash
    sudo task uninstall
    ```

## Taskfile

This example uses [Taskfile](https://taskfile.dev/) to define build and run tasks. See `Taskfile.yml` for details.
