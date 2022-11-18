#!/bin/bash

help() {
    echo -e "CTL Usage: ./ctl.sh <create|update|help> <--c|--cpp|--go|--js|--py> <name>\n"
    echo -e "COMMAND"
    echo -e "\t create \t create a new problem with a given language\n"
    echo -e "\t update \t modify an existing problem, 
             \t \t either updating a solution in an existing language, 
             \t \t or adding another solution in a new language\n"
    echo -e "\t help   \t print this help message\n"
    echo -e "OPTION"
    echo -e "\t -c   \t use C language\n"
    echo -e "\t -cpp \t use C++ language\n"
    echo -e "\t -go  \t use Go language\n"
    echo -e "\t -js  \t use JavaScript language\n"
    echo -e "\t -py  \t use Python language\n"
}

create() {
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
help)
    help
    exit 0
    ;;
*)
    help
    exit 1
    ;;
esac

