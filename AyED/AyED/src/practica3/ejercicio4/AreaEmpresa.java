package practica3.ejercicio4;

public class AreaEmpresa {
	
	private String Area;
	private int delay;
	
	public AreaEmpresa(){}
	
	public AreaEmpresa(String area, int delay) {
		super();
		Area = area;
		this.delay = delay;
	}

	public String getArea() {
		return Area;
	}

	public void setArea(String area) {
		Area = area;
	}

	public int getDelay() {
		return delay;
	}

	public void setDelay(int delay) {
		this.delay = delay;
	}
	
	
	
	
}
