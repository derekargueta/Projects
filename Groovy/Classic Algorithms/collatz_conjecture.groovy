def collatz = { int i, count=0 ->
    while(i != 1) {
        if(i % 2 == 0) i /= 2
        else i = (i * 3) + 1
        count++
    }
    return count
}

println(collatz(27))
