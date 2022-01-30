Notes for go:
    1. Variable names must be used
    2. Imported package must be used
    3. if you decalare a constant then it is not necessary to use that constant name

Go is a statically typed language
    you'll need to tell compiler, the data type when declaring variable else you'll get error
    Go can imply the type based on value
    Values will be validated at compile time only, lets say if you are using negative for unsigned int then it will show error
    if you want to define type explicitly or constant then in that case := will not work

Take input from  user
    fmt.Scan(userName) : save user input in variable called userName. but if we run the program with this statement program will not wait for user input so you'll need to add &(pointer) before variable name (fmt.scan(userName))

Pointer:
    A pointer is a variable that points to the memory address of another variable

Arrays and slices: 
    Array: Fixed in size, only same type of element

range : range allows us to iterate over element for different data structure (not only arrays and slices)

strings.Fields():
    split the string with white space as separator
    and return a slice with the split element
    e.g name := "kalyan jonwal"
        strings.Fields(name) it will return ["kalyan","jonwal"]
loop:
    for loop:
        infinite loop:
            for {

            }
    for-each loop:
        for index, ele := range Array/slice/DifferentdataStructureName{

        }