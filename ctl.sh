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
        echo -e "\t./ctl.sh update --$1 $2"

        return
    fi

    mkdir "problems/$2"
    touch "problems/$2/solution.$1"

    vim "problems/$2/solution.$1"
    echo "Problem $2 solution in $1 language created."

    IFS='-'
    read -ra split_string <<< "$2"

    git add .
    git commit -S -s -m "feat(solution): add the solution in $1 to the problem ${split_string[0]}"

    echo "Problem $2 solution in $1 language committed."
}

update() {
    echo "This feature has not been implemented yet."
}

info() {
    echo "This feature has not been implemented yet."
}

command="$1"
shift

case "$command" in
create)
    language="c"

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
        help
        exit 1
        ;;
    esac

    name="$1"
    shift
    if [ -z "$name" ]; then
        echo -e "Problem name required\n"
        help
        exit 1
    fi

    create "$language" "$name"
    ;;

update)
    update
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

