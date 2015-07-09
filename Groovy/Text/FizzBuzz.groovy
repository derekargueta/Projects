for(int x = 1; x < 101; x++) {
    if(x % 3 == 0 && x % 5 == 0) {
        println "FizzBuzz"
    } else if(x % 3 == 0) {
        println 'Fizz'
    } else if(x % 5 == 0) {
        println 'Buzz'
    } else {
        println x
    }
}