package practica1.ejercicio7;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.Collections;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.Random;

public class Ejercicio7 {
	public static void main(String[] args){
		
	}
	
	public static Integer[] generarSecuenciaNumerica(int a, int b) {		
		Random random = new Random();		
		int size = random.ints(a, b).findFirst().getAsInt();
		
		Integer[] output = new Integer[size];
		
		for (int i = 0; i < size; i++) {
			output[i] = random.ints(a,b).findFirst().getAsInt(); 
		}
				
		return output;		
	}
	
	public static ArrayList<Integer> generarListaNumerica(int a, int b){
		Random random = new Random();		
		int size = random.ints(a, b).findFirst().getAsInt();
		
		ArrayList<Integer> output = new ArrayList<>();
		
		for (int i = 0; i < size; i++) {
			output.add(random.ints(a,b).findFirst().getAsInt()); 
		}
				
		return output;
	}

	public static void IncisoABC() {
		Integer[] nums = generarSecuenciaNumerica(100, 1000);
		
		double start = System.currentTimeMillis();
		
		ArrayList<Integer> arr = new ArrayList<>(Arrays.asList(nums));
		
		for (Integer integer : arr) {
			System.out.println(integer);
		}
		
		double finish = System.currentTimeMillis();
		double arrayListTime = finish - start;
						
		start = System.currentTimeMillis();
		
		LinkedList<Integer> list = new LinkedList(Arrays.asList(nums));
		
		for (Integer integer : list) {
			System.out.println(integer);
		}
		
		finish = System.currentTimeMillis();
		double linkedListTime = finish - start;
		
		System.out.println("LinkedList time: " + linkedListTime);
		System.out.println("ArrayList time: " + arrayListTime);
	}
	
	public static ArrayList<Estudiante> IncisoD() {
	
		// Original
		ArrayList<Estudiante> estudiantes = new ArrayList();
		estudiantes.add(new Estudiante());
		estudiantes.add(new Estudiante());
		estudiantes.add(new Estudiante());
		
		return estudiantes;
		
//		// Copia 1
//		ArrayList<Estudiante> estudiante2 = new ArrayList(estudiantes);
//		
//		// Copia 2
//		ArrayList<Estudiante> estudiante3 = new ArrayList<Estudiante>(); 
//		estudiante3.addAll(estudiantes);
//		
//		System.out.println("Elementos 1: ");
//		for (Estudiante estudiante : estudiantes) {
//			System.out.println(estudiante);
//		}
//		
//		System.out.println("Elementos 2: ");
//		for (Estudiante estudiante : estudiante2) {
//			System.out.println(estudiante);
//		}
//		
//		System.out.println("Elementos 3: ");
//		for (Estudiante estudiante : estudiante3) {
//			System.out.println(estudiante);
//		}
//		
//		System.out.println("Modificando elementos...: ");
//		for (Estudiante estudiante : estudiantes) {
//			estudiante.setNombre("Juan");
//			estudiante.setApellido("Dominguez");
//			estudiante.setComision("B");
//		}
//		
//		System.out.println("Elementos 1: ");
//		for (Estudiante estudiante : estudiantes) {
//			System.out.println(estudiante);
//		}
//		
//		System.out.println("Elementos 2: ");
//		for (Estudiante estudiante : estudiante2) {
//			System.out.println(estudiante);
//		}
//		
//		System.out.println("Elementos 3: ");
//		for (Estudiante estudiante : estudiante3) {
//			System.out.println(estudiante);
//		}		
	}
	
	public static boolean IncisoE(Estudiante estudiante, ArrayList<Estudiante> es) {

		if (es.contains(estudiante) == false) {
			es.add(estudiante);
			return true;
		}
		
		return false;
		
	}
	
	public static boolean IncisoF(ArrayList<Integer> l) {
		
		ArrayList<Integer> p1 = new ArrayList<>();
		ArrayList<Integer> p2 = new ArrayList<>();
		int middle = 0;
		boolean ok = true;
		
		for (int i = 0; i < l.size(); i++) {
			if ((i < l.size() / 2) && (i != ((double)l.size() / 2) + .5 )) // add only if is the first part but NOT if it is the middle part 
				p1.add(l.get(i));
			else if(i == ((double)l.size() / 2) - 0.5)
				middle = l.get(i);
			else
				p2.add(l.get(i));
		}		
		
		// revertimos la lista 2
		Collections.reverse(p2);
		
		for (int i = 0; i < p1.size(); i++) {
			if (p1.get(i) != p2.get(i))
			{
				ok = false;			
				break;
			}
		}
		
		return ok;
	}
}
