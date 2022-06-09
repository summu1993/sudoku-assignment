

    AIM: To create backend in golang to solve sudoku with scalabilty in mind

    Assumption: 
       1. For now taking a 9x9 stand matrix and focused on backend archicture only
       2. Although exposed 1 API for frontend to consume, to check if the validity of a number in a specific position 
          is valid or not

    API contracts:
        1. Backend architeture has been written with scalabilty in mind by provided a reverse proxy feature where this current repo will act as entry   
           point and decide where to forward the request

        2. We can further enhance this feature and add authorization checks or rate limiting for a user.
        
        curl https://sudoku-go-test.herokuapp.com/v1/sudoku --header 'Service: check'
        
        possible values of Service 
          i)  solve (solve header will route to /v1/solve/sudoku )
          ii) check (check header will route to /v1/check/sudoku ) 
          
          /v1/check/sudoku - is the endpoint consumed by Frontend
          
        header Service is mandatory, as redirection logic is decided from Service header only (for now)
        
        ** Possible value of Service header are **
        1. solve
        2. check

    
        `Steps to run it locally`
        Run these commands
        
        go mod tidy
        go mod vendor
        go run main.go
        
        and change SOLVE_SUDOKU_PROXY_URL key to http://localhost:8001 in config.yml to run it locally


     # Changes made to make it more scalable


       1. API's will not accept dimention as well with user input comma seaparted sudoku board
          

       --- 
       I could have made this API a POST verb, so make url parameter more elegant, but we not creating any resource on server/database as of now so kept it GET request for now. 

        **Purpose of this API to solve a given Sudoku**

       `curl --location --request GET 'http://localhost:8001/v1/sudoku?board="5,3,0,0,7,0,0,0,0,6,0,0,1,9,5,0,0,0,0,9,8,0,0,0,0,6,0,8,0,0,0,6,0,0,0,3,4,0,0,8,0,3,0,0,1,7,0,0,0,2,0,0,0,6,0,6,0,0,0,0,2,8,0,0,0,0,4,1,9,0,0,5,0,0,0,0,8,0,0,7,9"&dimension=9' \
       --header 'Service: solve'`

       Query parameters
       `board` and `dimension`

       Mandatory header
       `Service : solve`

       	---
 
         **Purpose of this API to check the validity of a number entered by a player at a specific position (row and column value) **

        `curl --location --request GET 'http://localhost:8001/v1/sudoku?board="5,3,0,0,7,0,0,0,0,6,0,0,1,9,5,0,0,0,0,9,8,0,0,0,0,6,0,8,0,0,0,6,0,0,0,3,4,0,0,8,0,3,0,0,1,7,0,0,0,2,0,0,0,6,0,6,0,0,0,0,2,8,0,0,0,0,4,1,9,0,0,5,0,0,0,0,8,0,0,7,9"&horizontalPosition=1&verticalPosition=2&value=2&dimension=9' \
        --header 'Service: check'`


        Query parameters
       `board` , `dimension` , `horizontalPosition` , `verticalPosition` , `value`

       Mandatory header
       `Service : check`





       	Command to run nested test cases 
    
        go test ./...
