package processing

func reverse(s []string) []string {
    a := make([]string, len(s))
    copy(a, s)

    for i := len(a)/2 - 1; i >= 0; i-- {
        opp := len(a) - 1 - i
        a[i], a[opp] = a[opp], a[i]
    }

    return a
}
