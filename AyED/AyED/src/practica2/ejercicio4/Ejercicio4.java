package practica2.ejercicio4;

import java.util.LinkedList;
import java.util.Queue;

import javax.swing.tree.TreeNode;

import practica2.BinaryTree;
import resources.TreeTraversal;
import resources.Util;

public class Ejercicio4 {

	public static void main(String[] args) {
		BinaryTree<Integer> tree = Util.generarArbolEnteros(3);
		TreeTraversal<Integer> tt = new TreeTraversal<Integer>();
		
		tt.visualize(tree);
		
		System.out.println(retardoReenvio(tree)); 
	}
	
	public static int retardoReenvio(BinaryTree<Integer> root){
		
		if (root == null)
			return 0;
		
		// recorremos en profundidad pre-orden		
		int max = Math.max(retardoReenvio(root.getLeftChild()), retardoReenvio(root.getRightChild()));		
		int maxDelay = root.getData() + max;
		
		return maxDelay;
	}

}
