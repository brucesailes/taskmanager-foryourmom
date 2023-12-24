# Task Manager 

Task Manager is a simple web-based task management system built using htmx and Go.

## Description

Task Manager allows users to add tasks dynamically without reloading the page. It utilizes htmx for seamless form submissions and Go for server-side processing.

## Installation

To run this project locally, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/brucesailes/task-manager.git
    ```

2. Navigate to the project directory:

    ```bash
    cd task-manager
    ```

3. Run the Go server:

    ```bash
    go run main.go
    ```

4. Access the application in your web browser at `http://localhost:8080`.

## Usage

1. Enter a task in the input field.
2. Click "Add Task" to submit the form.
3. The task will be added to the list dynamically without page reload.

## Screenshots

(![go termainl](https://github.com/brucesailes/taskmanager-foryourmom/assets/110871245/b323538a-2b50-4dfc-9d5c-ae44259b6988)
)
(![htmx go web](https://github.com/brucesailes/taskmanager-foryourmom/assets/110871245/92d6a1b0-9eee-4b14-9620-98d3afbd1d2a)
)

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/new-feature`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature/new-feature`).
6. Create a new Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

- [htmx](https://htmx.org/) - For providing seamless AJAX interactions
- [Go](https://golang.org/) - For server-side processing
