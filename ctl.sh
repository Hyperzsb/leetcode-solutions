#!/bin/bash

# Help function is used to print the help message of this script
help() {
    echo -e "CTL Usage: ./ctl.sh COMMAND [OPTION] [PROBLEM_NAME]\n"

    echo -e "COMMAND"
    echo -e "\t create \t create a new problem with the given language"
    echo -e "\t update \t modify an existing problem, 
             \t \t either updating a solution in an existing language, 
             \t \t or adding another solution in a new language"
    echo -e "\t info   \t print the stats info about the given problem"
    echo -e "\t help   \t print this help message and exist\n"

    echo -e "OPTION"
    echo -e "\t --c     \t use the C programming language"
    echo -e "\t --cpp   \t use the C++ programming language"
    echo -e "\t --go    \t use the Go programming language"
    echo -e "\t --java  \t use the Java programming language"
    echo -e "\t --js    \t use the JavaScript programming language"
    echo -e "\t --py    \t use the Python programming language"
    echo -e "\t --sql   \t use the SQL for database scripts\n"

    echo -e "PROBLEM_NAME"
    echo -e "\t Scheme:  <problem_id>-<problem_name (separated by dashes)>"
    echo -e "\t Example: 123-A-New-Problem\n"

    echo -e "EXAMPLES"
    echo -e "\t ./ctl.sh create --c 123-A-New-Problem"
    echo -e "\t ./ctl.sh update --go 123-A-New-Problem"
    echo -e "\t ./ctl.sh info 123"
    echo -e "\t ./ctl.sh help\n"
}

# Create function is used to create a new solution of a new problem,
# if and only if there is no previous problem with the same name as provided
create() {
    # Check whether there exists a problem with the same name as provided
    # If so, print the help message telling user yo use the `update` command instead, and exit
    if [[ -d "problems/$2" ]]; then
        echo -e "The problem \"$2\" already exists. If you want to update it, try:\n"
        echo -e "\t ./ctl.sh update --$1 $2"
        return
    fi

    # Create a new directory for the problem
    mkdir "problems/$2"
    # Create a new file for the solution
    touch "problems/$2/solution.$1"
    # Use Vim to edit the solution file
    vim "problems/$2/solution.$1"
    echo -e "The new solution of the problem $2 in $1 programming language has been created\n"

    # Prompt a message to receive the confirmation of user whether to commit the new problem
    read -rp "Do you want to commit this new solution? (y/n, default y) " confirm
    echo -e ""
    # If user does not confirm, discard all changes, and exit
    if [[ "$confirm" == "n" ]]; then
        rm -rf "problems/$2"
        echo -e "All changes have been discarded"
        return
    fi

    # Extract the id of the problem
    IFS='-' read -ra split_string <<< "$2"
    # Add and commit the problem to Git
    git add .
    git commit -S -s -m "feat(solution): add the solution in $1 to the problem ${split_string[0]}"
    echo -e "\nThe new solution of the problem $2 in $1 programming language has been committed"
}

# Update function is used to update an existing problem with either adding a new solution in a different language,
# or editing an old solution, if and only if there is a previous problem with the same name as provided
update() {
    # Check whether there exists a problem with the same name as provided
    # If not, print the help message telling user yo use the `create` command instead, and exit
    if [[ ! -d "problems/$2" ]]; then
        echo -e "The problem \"$2\" does not exist. If you want to create it, try:\n"
        echo -e "\t ./ctl.sh create --$1 $2"
        return
    fi

    # Check whether the updated solution is in a new language or not
    # If so, create a new file for the new language
    if [[ ! -f "problems/$2/solution.$1" ]]; then
        touch "problems/$2/solution.$1"
    fi
    vim "problems/$2/solution.$1"
    echo -e "The solution of the problem $2 in $1 programming language has been updated\n"

    # Prompt a message to receive the confirmation of user whether to commit the new problem
    read -rp "Do you want to commit this updated solution? (y/n, default y) " confirm
    echo -e ""
    # If user does not confirm, discard all changes, and exit
    if [[ "$confirm" == "n" ]]; then
        git reset --hard HEAD
        echo -e "All changes have been discarded"
        return
    fi

    # Extract the id of the problem
    IFS='-' read -ra split_string <<< "$2"
    # Add and commit the problem to Git
    git add .
    git commit -S -s -m "feat(solution): update the solution in $1 to the problem ${split_string[0]}"
    echo -e "\nThe solution of the problem $2 in $1 programming language has been committed"
}

info() {
    problem_name="$(ls problems/ | grep -E "^$1.*")"
    if [[ -z "$problem_name" ]]; then
        echo -e "No match found: $1"
    else
        echo -e "$problem_name\n"
        ls "problems/$problem_name"
    fi
}

# Main procedure starts

command="$1"
shift

# Parse command
case "$command" in
# If command `create` is specified
create)
    if [[ $# -ne 2 ]]; then
        echo -e "Invalid number of arguments. See './ctl.sh help'"
    fi

    opt="$1"
    shift
    case "$opt" in
    --c)
        language="c"
        ;;
    --cpp)
        language="cpp"
        ;;
    --go)
        language="go"
        ;;
    --java)
        language="java"
        ;;
    --js)
        language="js"
        ;;
    --py)
        language="py"
        ;;
    --sql)
        language="sql"
        ;;
    *)
        echo -e "Invalid option. See './ctl.sh help'"
        exit 1
        ;;
    esac

    name="$1"
    shift
    if [[ -z "$name" ]]; then
        echo -e "The problem name is required\n"
        help
        exit 2
    fi

    create "$language" "$name"
    ;;
# If command `update` is specified
update)
    if [[ $# -ne 2 ]]; then
        echo -e "Invalid number of arguments. See './ctl.sh help'"
    fi

    opt="$1"
    shift
    case "$opt" in
    --c)
        language="c"
        ;;
    --cpp)
        language="cpp"
        ;;
    --go)
        language="go"
        ;;
    --java)
        language="java"
        ;;
    --js)
        language="js"
        ;;
    --py)
        language="py"
        ;;
    --sql)
        language="sql"
        ;;
    *)
        echo -e "Invalid option. See './ctl.sh help'"
        exit 1
        ;;
    esac

    name="$1"
    shift
    if [[ -z "$name" ]]; then
        echo -e "The problem name is required\n"
        help
        exit 2
    fi

    update "$language" "$name"
    ;;
# If command `info` is specified
info)
    if [[ $# -ne 1 ]]; then
        echo -e "Invalid number of arguments. See './ctl.sh help'"
    fi

    pattern="$1"
    shift

    info "$pattern"
    ;;
# If command `help` is specified
help)
    help
    exit 0
    ;;
# If another undefined command is specified
*)
    help
    exit 1
    ;;
esac

# Main procedure ends
