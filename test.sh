#!/bin/bash

#String
text="Welcome@to@GeeksforGeeks@!!"

# Set @ as the delimiter
IFS='@'

# Read the split words into an array 
# based on space delimiter
read -ra newarr <<< "$text"


# Print each value of the array by 
# using the loop
echo ${newarr[0]}