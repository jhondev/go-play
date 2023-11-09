# Rain CSV Parser

This parser aims to standardize data from files coming from different employers to create a list of eligible employees to sign up for Rain.

Although files need to have key pieces of data, Rain does not control the overall structure of the files. For example, column names and order can be different from file to file.

The parser read these files and standardize the data to be processed later. 4 file examples are present in the testdata folder, additional file structures are allowed.

## Requirements

1.Parse input files one at a time and generate two files as output, one for correct data, one for bad data

2.Minimal validations (you are free to create additional validations): 
    
    a. Required data:
        i.Employee name
        ii.Employee salary
        iii.Employee email
        iv.Employee ID
    b. Employee email must be unique
    c. Employee ID must be unique

3.Output a summary of the processing steps in the console

4.Create unit tests

5.Project must be implemented in Golang 1.16+

6.Project must have a clear folder structure, scalable and simple (not a single file script) 

7.In the README file, you should cover at least:
- How to run the project
- Explain your chosen architecture. Why do you think it is a good fit for this problem?
- How you would evolve your submitted code