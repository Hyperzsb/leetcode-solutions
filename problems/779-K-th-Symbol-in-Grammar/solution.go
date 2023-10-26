func recursion(n, k int) int {
    if n == 0 && k == 0 {
        return 0
    }

    if prev := recursion(n-1, k/2); prev == 0 {
        if k%2 == 0 {
            return 0
        } else {
            return 1
        }
    } else {
        if k%2 == 0 {
            return 1
        } else {
            return 0
        }
    }
}

func kthGrammar(n int, k int) int {
    return recursion(n-1, k-1)
}


