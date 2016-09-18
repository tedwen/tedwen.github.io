import scala.io.Source
import scala.collection.mutable.Map

object CharFreq extends App {
    val filename = if (args.length > 0) args(0) else "page1.txt";
    
    val cf = Map[Char,Int]()
    
    try {
        for (line <- Source.fromFile(filename).getLines()) {
            line.toList foreach { (x) => cf += x -> (cf.getOrElse(x, 0) + 1) }
        }
        
        val cfv = cf.toSeq.sortBy(-_._2)
        
        cfv foreach { (cp) => println(s"${cp._1}: ${cp._2}") }
    } catch {
        case ex: Exception => println("Cannot open file")
    }
}
