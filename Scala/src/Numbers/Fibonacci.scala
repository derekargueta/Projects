import io.Source.stdin

object Main {

  def fib(a: Int): Int = {
    if(a == 0) return 0
    else if(a == 1) return 1
    else {
      return fib(a - 1) + fib(a - 2)
    }
  }

  def main(args: Array[String]) {
    for(ln <- stdin.getLines()) {
      val maxNum = ln.toString.toInt
      println(fib(maxNum))
    }
  }
}
