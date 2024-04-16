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
	
	public boolean isLeftTree() {
		
		if (this.tree.isEmpty())
			return false;
		
		int lCount = count(tree.getLeftChild());
		int rCount = count(tree.getRightChild());
		
		System.out.println("Cantidad de arboles con un solo hijo en el hijo izquierdo: " + lCount);
		System.out.println("Cantidad de arboles con un solo hijo en el hijo derecho: " + rCount);
		
		return lCount > rCount;
		
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
