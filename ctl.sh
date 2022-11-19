#!/bin/bash

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
    echo -e "\t -c     \t use the C programming language"
    echo -e "\t -cpp   \t use the C++ programming language"
    echo -e "\t -go    \t use the Go programming language"
    echo -e "\t -js    \t use the JavaScript programming language"
    echo -e "\t -py    \t use the Python programming language\n"

    echo -e "EXAMPLES"
    echo -e "\t ./ctl.sh create --c 123-A-New-Problem"
    echo -e "\t ./ctl.sh update --go 123-A-New-Problem"
    echo -e "\t ./ctl.sh info 123\n"
}

create() {
    if [ -d "problems/$2" ]; then
        echo -e "The problem \"$2\" already exists. If you want to update it, try:\n"
        echo -e "\t ./ctl.sh update --$1 $2"

        return
    fi

    mkdir "problems/$2"
    touch "problems/$2/solution.$1"

    vim "problems/$2/solution.$1"
    echo -e "The new solution of the problem $2 in $1 programming language created.\n"

    IFS='-'
    read -ra split_string <<< "$2"

    git add .
    git commit -S -s -m "feat(solution): add the solution in $1 to the problem ${split_string[0]}"

    echo -e "\nThe new solution of the problem $2 in $1 programming language committed."
}

update() {
    if [ ! -d "problems/$2" ]; then
        echo -e "The problem \"$2\" does not exist. If you want to create it, try:\n"
        echo -e "\t ./ctl.sh create --$1 $2"

        return
    fi

    vim "problems/$2/solution.$1"
    echo -e "The solution of the problem $2 in $1 programming language updated\n"

    IFS='-'
    read -ra split_string <<< "$2"

    git add .
    git commit -S -s -m "feat(solution): update the solution in $1 to the problem ${split_string[0]}"

    echo -e "\nThe solution of the problem $2 in $1 programming language committed"
}

info() {
    echo "This feature has not been implemented yet."
}

command="$1"
shift

case "$command" in
create)
    language=""

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
    --js)
        language="js"
        ;;
    --py)
        language="py"
        ;;
    *)
        echo -e "Invalid option. See './ctl.sh help'"
        exit 1
        ;;
    esac

    name="$1"
    shift
    if [ -z "$name" ]; then
        echo -e "The problem name is required\n"
        help
        exit 2
    fi

    create "$language" "$name"
    ;;

update)
    language=""

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
    --js)
        language="js"
        ;;
    --py)
        language="py"
        ;;
    *)
        echo -e "Invalid option. See './ctl.sh help'"
        exit 1
        ;;
    esac

    name="$1"
    shift
    if [ -z "$name" ]; then
        echo -e "The problem name is required\n"
        help
        exit 2
    fi

    update "$language" "$name"
    ;;

info)
    info
    ;;

help)
    help
    exit 0
    ;;

*)
    help
    exit 1
    ;;
esac

