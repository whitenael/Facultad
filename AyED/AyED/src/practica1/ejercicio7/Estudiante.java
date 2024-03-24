package practica1.ejercicio7;
import static practica1.ejercicio7.Util.generarString;

import java.util.Objects;

public class Estudiante {
	
	private String nombre;
	
	private String apellido;
	
	private String comision;
	
	public Estudiante() {
		this.nombre = generarString();
		this.apellido = generarString();
		this.comision = generarString();
	}
	
	public Estudiante(String nombre, String apellido, String comision) {
		this.nombre = nombre;
		this.apellido = apellido;
		this.comision = comision;
	}

	public String getNombre() {
		return nombre;
	}

	public void setNombre(String nombre) {
		this.nombre = nombre;
	}

	public String getApellido() {
		return apellido;
	}

	public void setApellido(String apellido) {
		this.apellido = apellido;
	}

	public String getComision() {
		return comision;
	}

	public void setComision(String comision) {
		this.comision = comision;
	}

	@Override
	public String toString() {
		return "Estudiante [nombre=" + nombre + ", apellido=" + apellido + ", comision=" + comision + "]";
	}

	@Override
	public int hashCode() {
		return Objects.hash(apellido, comision, nombre);
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		Estudiante other = (Estudiante) obj;
		return Objects.equals(apellido, other.apellido) && Objects.equals(comision, other.comision)
				&& Objects.equals(nombre, other.nombre);
	}		
}
