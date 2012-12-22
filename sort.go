package sort

func Ints(a []int) {
    var sorted bool
    for {
        sorted = true
        for i := 0; i < len(a)-1; i++ {
            if a[i] > a[i+1] {
                a[i], a[i+1] = a[i+1], a[i]
                sorted = false
            }
        }
        if sorted {
            return
        }
    }
}
