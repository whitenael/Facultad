package practica3.ejercicio6;

import practica3.GeneralTree;
import resources.Util;

public class Ejercicio6 {

	public static void main(String[] args) {
		
		GeneralTree<Character> red = Util.generarArbolGeneral_Character(4, 2);
		red.visualize(3);
		
		RedAguaPotable red_agua = new RedAguaPotable(red);
		System.out.println(red_agua.minimoCaudal(1000));
		
	}

}
