package practica1.ejercicio8;

import java.util.Iterator;
import java.util.NoSuchElementException;

public class Sequence<Item> implements Iterable<Item> {
	private int n; // number of elements on list
	private Node first; // first element
	private Node last; // last element

	// linked list node helper data type
	private class Node {
		private Item item;
		private Node next;

		Node(Item item, Node next) {
			this.item = item;
			this.next = next;
		}
	}

	public boolean isEmpty() {
		return first == null;
	}

	public int size() {
		return n;
	}

	// add item to the end of the list
	public void add(Item item) {
		Node x = new Node(item, null);
		if (isEmpty())
			first = x;
		else
			last.next = x;
		last = x;
		n++;
	}

	//// prepend item to the beginning of the list
	// public void addFirst(Item item) {
	// first = new Node(item, first);
	// if (isEmpty()) last = first;
	// n++;
	// }

	public Iterator<Item> iterator() {
		return new SeqIterator();
	}

	// an iterator, doesn't implement remove() since it's optional
	private class SeqIterator implements Iterator<Item> {
		private Node current = first;

		public boolean hasNext() {
			return current != null;
		}

		public void remove() {
			throw new UnsupportedOperationException();
		}

		public Item next() {
			if (!hasNext())
				throw new NoSuchElementException();
			Item item = current.item;
			current = current.next;
			return item;
		}
	}
}