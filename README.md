# How to use
## Run app
Call in a warking directory `make run` and follow the instructions. 
The app will be started in a docker environment.

## Run tests
Call in a warking directory `make test`. It will start tests locally.
You need to have go 1.14 or greater installed for running the tests. 

# Complexity
Complexity: time complexity O(n)=2N+M, space complexity: not grater than 2N+M
Where N is amount of age values and M is max age == 250
Time complexity explanation:
 
Step 1: We use an array for storing persons' age. Saving all ages takes N time.

Step 2: We iterate from minYear to MaxYear in worst case it takes MaxYears iterations.
We can get extra iteration for each age value, in worst case it's N - amount of persons' age

Step 3: Parse data into string, is linear

[Algorithm implementation is available here](app/processor/processor.go)

# Assumptions
- Age must be in the rage [0, 250)
- Amount of values for ages always must be equal and more than 0
            
    