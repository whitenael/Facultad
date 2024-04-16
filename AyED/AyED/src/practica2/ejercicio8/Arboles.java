package practica2.ejercicio8;

import practica2.BinaryTree;

public class Arboles {

	public boolean esPrefijo(BinaryTree<Integer> a1, BinaryTree<Integer> a2){
		
		if ((a1 == null || a2 == null) || a1.isEmpty() || a2.isEmpty())
			return true;					
		
		if ( a1.hasLeftChild() && !a2.hasLeftChild() || !a1.hasLeftChild() && a2.hasLeftChild() &&
				a1.hasRightChild() && !a2.hasRightChild() || !a1.hasRightChild() && a2.hasRightChild())
			return false;
				
		if (a1.getData() == a2.getData()){
			return esPrefijo(a1.getLeftChild(), a2.getLeftChild()) && esPrefijo(a1.getRightChild(), a2.getRightChild()) && true;
		}
		
		else {
			return false;
		}			
	}

}
