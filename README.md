

    AIM: To create backend in golang to solve sudoku with scalabilty in mind

    Assumption: 
       1. For now taking a 9x9 stand matrix and focused on backend archicture only
       2. Although exposed 1 API for frontend to consume, to check if the validity of a number in a specific position 
          is valid or not

    API contracts:
        1. Backend architeture has been written with scalabilty in mind by provided a reverse proxy feature where this current repo will act as entry point and decide where to forward the request

        2. We can further enhance this feature and add authorization checks or rate limiting for a user.
        
        curl https://sudoku-go-test.herokuapp.com/v1/sudoku --header 'Service: check'

        header Service is mandatory, as redirection logic is decided from Service header only (for now)
    

        Steps to run it locally
        Run these commands
        
        go mod tidy
        go mod vendor
        go run main.go

        1. curl --location --request GET 'http://localhost:8001/v1/solve' --header 'Service: solve'
        2. curl --location --request GET 'http://localhost:8001/v1/solve' --header 'Service: check' 

        Query parameter for 2nd request  (board = a 2D matrix)
        
        board
        horizontal_position 
        vertical_position
        value

