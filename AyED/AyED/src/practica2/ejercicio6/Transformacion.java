package practica2.ejercicio6;

import practica2.BinaryTree;

public class Transformacion {
	public BinaryTree<Integer> tree;
	
	public Transformacion(){}
	
	public Transformacion(BinaryTree<Integer> tree){
		this.tree = tree;
	}
	
	public BinaryTree<Integer> suma(BinaryTree<Integer> tree){
				
		if (tree == null)
			return null;
		
		BinaryTree<Integer> node = new BinaryTree<Integer>(getSumSubTrees(tree));
		node.addLeftChild(suma(tree.getLeftChild()));		
		node.addRightChild(suma(tree.getRightChild())); 
		
		return node;
	}
	
	public int getSumSubTrees(BinaryTree<Integer> node) {
		
		if (node == null)
	        return 0;

	    int leftSum = getSumSubTrees(node.getLeftChild());
	    int rightSum = getSumSubTrees(node.getRightChild());
	    
	    // obtenemos la suma de los hijos excluyendo la del nodo principal
	    return leftSum + rightSum + (node.getLeftChild() != null ? node.getLeftChild().getData() : 0)
	                                 + (node.getRightChild() != null ? node.getRightChild().getData() : 0);
	}
}