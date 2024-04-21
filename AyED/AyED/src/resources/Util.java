package resources;
import java.util.LinkedList;

import practica2.BinaryTree;
import practica3.GeneralTree;

public class Util {
	
	
	public static GeneralTree<Integer> generarArbolGeneral(int length, int grado) 
	{
		if (length == 0)
			return null;
		
		GeneralTree<Integer> tree = new GeneralTree<Integer>(getRandomNumber(1, 10));
		for (int i = 0; i < grado; i++) {
			tree.addChild(generarArbolGeneral(length-1, grado));
		}		
		
		return tree;		
	}
	
	// genera arbol de longitud 10 desde la raiz
	@SuppressWarnings({ "unchecked", "rawtypes" })
	public static BinaryTree<Integer> generarArbolEnteros(int length){			
		
		if (length == 0)
			return new BinaryTree<Integer>();
						
		BinaryTree<Integer> br = new BinaryTree(getRandomNumber(1, 10));
		br.addLeftChild(generarArbolEnteros(length-1));
		br.addRightChild(generarArbolEnteros(length-1));
		
		return br;
		
	}
	
	public static BinaryTree<Integer> generarArbolAleatorio(int length, LinkedList<Integer> numList){
		
		LinkedList<Integer> nums = new LinkedList<Integer>();
		
		if (length == 0)
			return null;
		
		int rng = getRandomNumber(1, 30);
		
		while(nums.contains(rng)) {
			rng = getRandomNumber(1, 30);
		}
					
		BinaryTree<Integer> br = new BinaryTree<Integer>(rng);
		if (getRandomNumber(1, 6) != 4)
			br.addLeftChild(generarArbolAleatorio(length - 1, nums));
		if (getRandomNumber(1, 6) != 4)
			br.addRightChild(generarArbolAleatorio(length - 1, nums));
		
		return br;
		
	}
	
	public static BinaryTree<Integer> generarArbolAleatorio(int length, LinkedList<Integer> numList, int rb)
	{		
		LinkedList<Integer> nums = new LinkedList<Integer>();
		
		if (length == 0)
			return null;
		
		int rng = getRandomNumber(1, rb);
		
		while(nums.contains(rng)) {
			rng = getRandomNumber(1, rb);
		}
		
		numList.add(rng);
					
		BinaryTree<Integer> br = new BinaryTree<Integer>(rng);
		if (getRandomNumber(1, 6) != 4)
			br.addLeftChild(generarArbolAleatorio(length - 1, nums, rb));
		if (getRandomNumber(1, 6) != 4)
			br.addRightChild(generarArbolAleatorio(length - 1, nums, rb));
		
		return br;
		
	}
	
	public static int getRandomNumber(int min, int max) {
	    return (int) ((Math.random() * (max - min)) + min);
	}
	
}
