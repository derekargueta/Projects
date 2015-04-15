/**
 * @author Derek Argueta
 *
 * TODO - getSize()?
 */
public class LinkedList<T> {

    private Node head;
    private int size;

    public LinkedList() {

    }


    public LinkedList(T t) {
        head = new Node(t);
    }


    public Node getListHead() {
        return head;
    }


    public int getSize() {
        return size;
    }


    public Node getLastNode() {
        Node current = head;
        while(current.getNext() != null) {
            current = current.getNext();
        }
        return current;
    }


    public void insert(T t) {
        if(head == null) {
            head = new Node(t);
        } else {
            Node current = head;
            while(current.getNext() != null) {
                current = current.getNext();
            }
            Node newNode = new Node(t);
            current.setNext(newNode);
            newNode.setPrevious(current);
        }
        size++;
    }


    public boolean remove(T t) {
        if(head.getValue().equals(t)) {
            head = head.getNext();
            size--;
            return true;
        } else {
            Node current = head;
            while(current.getNext() != null) {
                current = current.getNext();
                if(current.getValue().equals(t)) {
                    current.getPrevious().setNext(current.getNext());
                    size--;
                    return true;
                }
            }
        }
        return false;
    }


    public void printList() {
        Node current = head;
        while(current.getNext() != null) {
            System.out.print(current.getValue() + " ");
            current = current.getNext();
        }
        System.out.println(current.getValue());
    }


    @Override
    public String toString() {
        StringBuilder result = new StringBuilder();
        Node current = head;
        while(current.getNext() != null) {
            result.append(current.getValue() + " ");
            current = current.getNext();
        }
        result.append(current.getValue());
        return result.toString();
    }
}


/**
 * @author Derek Argueta
 */
class Node<T> {
    private T value;
    private Node next;
    private Node previous;

    public Node(T t) {
        value = t;
        next = null;
        previous = null;
    }


    public Node getNext() {
        return next;
    }


    public void setNext(Node n) {
        next = n;
    }


    public Node getPrevious() {
        return previous;
    }


    public void setPrevious(Node p) {
        previous = p;
    }


    public T getValue() {
        return value;
    }


    @Override
    public String toString() {
        return value.toString();
    }
}
