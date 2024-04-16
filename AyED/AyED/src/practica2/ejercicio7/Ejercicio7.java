package practica2.ejercicio7;

import java.util.LinkedList;

import resources.TreeTraversal;
import resources.Util;

public class Ejercicio7 {
	
	public static void main(String[] args) {
				
		TreeTraversal<Integer> tt = new TreeTraversal<Integer>();
		LinkedList<Integer> numList = new LinkedList<Integer>();
		
		Arboles arbol = new Arboles(Util.generarArbolAleatorio(6, numList));
			
		tt.visualize(arbol.getTree());
		System.out.println();
		System.out.println(arbol.isLeftTree()); 
		
	}
}
