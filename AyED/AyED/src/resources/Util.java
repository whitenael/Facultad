package resources;
import java.util.Random;

import practica2.BinaryTree;

public class Util {
	
	
	// genera arbol de longitud 10 desde la raiz
	@SuppressWarnings({ "unchecked", "rawtypes" })
	public static BinaryTree<Integer> generarArbolEnteros(int length){			
		
		if (length == 0)
			return null;
						
		BinaryTree<Integer> br = new BinaryTree(getRandomNumber(1, 10));
		br.addLeftChild(generarArbolEnteros(length-1));
		br.addRightChild(generarArbolEnteros(length-1));
		
		return br;
		
	}
	
	public static int getRandomNumber(int min, int max) {
	    return (int) ((Math.random() * (max - min)) + min);
	}
	
}
