package practica2.ejercicio8;

import java.util.LinkedList;

import practica2.BinaryTree;
import resources.TreeTraversal;
import resources.Util;

public class Ejercicio8 {
	public static void main(String[] args){
		Arboles ar = new Arboles();
		BinaryTree<Integer> root = new BinaryTree<Integer>(1);
		root.addLeftChild(new BinaryTree<Integer>(2));
		root.getLeftChild().addLeftChild(new BinaryTree<Integer>(4));
		root.getLeftChild().addRightChild(new BinaryTree<Integer>(5));
		root.addRightChild(new BinaryTree<Integer>(3));
		//root.getRightChild().addLeftChild(new BinaryTree<Integer>(7));
		root.getRightChild().addRightChild(new BinaryTree<Integer>(10));
			
		BinaryTree<Integer> root2 = new BinaryTree<Integer>(1);
		root2.addLeftChild(new BinaryTree<Integer>(2));
		root2.getLeftChild().addLeftChild(new BinaryTree<Integer>(4));
		root2.getLeftChild().addRightChild(new BinaryTree<Integer>(5));
		root2.addRightChild(new BinaryTree<Integer>(3));
		//root2.getRightChild().addLeftChild(new BinaryTree<Integer>(8));
	
		TreeTraversal<Integer> tt = new TreeTraversal<Integer>();
		
		System.out.println("Root 1: ");
		tt.visualize(root);
		System.out.println("----------------");
		System.out.println("Root 2: ");
		tt.visualize(root2);
		
		System.out.println(ar.esPrefijo(root, root2)); 
		
	}
}
