/**
 * @author Derek Argueta
 */
public class Main {

    public static void main(String[] args) {
        LinkedList ll = new LinkedList();
        ll.insert(4);
        ll.insert(1);
        ll.insert(3);
        ll.insert(10);
        ll.insert(9);
        ll.printList();

        if(!ll.getListHead().getValue().equals(4)) {
            System.err.println("FAILED getListHead()");
        } else {
            System.out.println("Passed getListHead()");
        }

        if(!ll.getLastNode().getValue().equals(9)) {
            System.err.println("FAILED getLastNode()");
        } else {
            System.out.println("Passed getLastNode()");
        }

        Node test = ll.getListHead();
        test = test.getNext();
        if(!test.getValue().equals(1)) {
            System.err.println("FAILED getNext() 1st iteration");
        } else {
            System.out.println("Passed getNext() 1st iteration");
        }

        test = test.getNext();
        if(!test.getValue().equals(3)) {
            System.err.println("FAILED getNext() 2nd iteration");
        } else {
            System.out.println("Passed getNext() 2nd iteration");
        }

        test = test.getPrevious();
        if(!test.getValue().equals(1)) {
            System.err.println("FAILED getPrevious() 1st iteration");
        } else {
            System.out.println("Passed getPrevious() 1st iteration");
        }

        if(!ll.remove(3)) {
            System.err.println("FAILED remove() 1st iteration");
        } else {
            System.out.println("Passed remove() 1st iteration");
        }

        if(ll.remove(1000)) {
            System.err.println("FAILED remove() 1st iteration");
        } else {
            System.out.println("Passed remove() 1st iteration");
        }

        ll.printList();
    }
}
