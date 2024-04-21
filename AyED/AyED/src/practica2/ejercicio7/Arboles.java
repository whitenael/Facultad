package practica2.ejercicio7;

import practica2.BinaryTree;

public class Arboles {

	private BinaryTree<Integer> tree;
	
	public Arboles() {}
	
	public Arboles(BinaryTree<Integer> tree) {
		this.tree = tree;
	}	

	public BinaryTree<Integer> getTree() {
		return tree;
	}

	public void setTree(BinaryTree<Integer> tree) {
		this.tree = tree;
	}
	
	public boolean isLeftTree(int num) {
		
		if (this.tree.isEmpty())
			return false;
		
		BinaryTree<Integer> nodo = buscarNodo(this.tree, num);
		
		if (nodo == null)
			return false;
		
		int lCount = count(nodo.getLeftChild());
		int rCount = count(nodo.getRightChild());
		
		System.out.println("Cantidad de arboles con un solo hijo en el hijo izquierdo: " + lCount);
		System.out.println("Cantidad de arboles con un solo hijo en el hijo derecho: " + rCount);
		
		return lCount > rCount;
		
	}
	
	public BinaryTree<Integer> buscarNodo(BinaryTree<Integer> nodo, int num) {
		
		// Si llegamos al fin del arbol y no encontramos el dato
		if (nodo == null || nodo.getData() == null)
			return null;
		
		// Si encontramos el dato
		if (nodo.getData() == num) 
			return nodo;
		
		// Si no lo encontramos en el nodo raiz, lo buscamos en el nodo izquierdo
		BinaryTree<Integer> nb = new BinaryTree<Integer>();
		if (nodo.hasLeftChild())
			nb = buscarNodo(nodo.getLeftChild(), num);
		
		// Caso contrario, lo continuamos buscando en el derecho
		if (nb == null && nodo.hasRightChild()) {
			nb = buscarNodo(nodo.getRightChild(), num);
		}
					
		return nb;		
	}		
	
	private int count(BinaryTree<Integer> node) {
		
		if (node == null)
			return 0;
		
		if (node.isLeaf())
			return 0;
				
		if (hasOneChildren(node)) {
			if (node.hasRightChild())
				return count(node.getRightChild()) + 1;
			if (node.hasLeftChild())
				return count(node.getLeftChild()) + 1;				
		}
		
		return count(node.getRightChild()) + count(node.getLeftChild());		
	}
	
	private boolean hasOneChildren(BinaryTree<Integer> node) {
		if (!node.isLeaf() && (!node.hasLeftChild() || !node.hasRightChild())) {
			return true;
		}
		
		return false;
	}	

	
}
