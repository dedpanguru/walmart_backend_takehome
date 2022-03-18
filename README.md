# walmart_backend_takehome
Movie Theater Seating Challenge 2020

## My assumptions and design
- When making this project, I assumed that the theater enforced 2 public safety guidelines: *a 1 row buffer between seats* and *a 3 seat buffer between seats*
- I then designed the seating arrangement to follow this assumption while also maximizing customer satisfaction
- Given that the overall dimensions of the theater are 10 rows by 20 columns, after enforcing the public safety guidelines, there were only 25 seats available in the whole theater
  - A 3 seat buffer would remove 15 seats from a row entirely
  - 1 one row buffer reduces the available amount of rows by half, only leaving 5 rows
  - 5 rows times 5 seats per row, leaves 25 seats in total
- Finally, I assumed that on an unfullfillable request , the program would not stop running (as one bad request shouldn't crash the whole system) but instead return a request rejection message. Only accepted requests will be outputted into the output file. 

## Priorities
- To maximize customer satisfaction, I implemented a first-come first-serve based seat allocation system, where back row and middle seats were granted to requests that came earlier than others. 
- I also implemented a group-based seating system, where if the request asked for multiple seats, the system would try to reserve the same request in the same row, so the whole group sits together.
- I would only reserve seats in columns 10,14,6,18,2 in rows B,D,F,H,J to get the best balance of public safety and customer satisfaction

## Testing
- I used a equivalence testing strategy, validating valid cases and invalidating invalid cases
- valid cases would reserve seats up to 25
- invalid cases would reserve seats above 25 or send invalid input
- tests are found in the testing directory
  - run each test by running `python3 reserve.py ./testing/<input_test_file>` and then taking a look at the generated output file or the console.  

## How to run
- This program takes **up to 2** command-line arguments: an *input file path* and a *output file path*
  - For example:  `python3 reserve.py input.txt output.txt`
- The output file path is optional. If it is not provided, the program will default to creating and outputting to a *output.txt* file in the local directory
  - However, **an input file is mandatory** (it's name does not have to be input.txt) 
- This program does not validate the input file as of yet, so please make sure that you are entering the input like so:
  ```
  R001 1
  R002 2
  R003 3
  and so on...
  ``` 
