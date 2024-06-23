package tp5.ejercicio1b.Impl.AdjList;

import tp5.ejercicio1b.Edge;
import tp5.ejercicio1b.Vertex;

public class EdgeImpl<T> implements Edge<T> {

    private Vertex<T> target;
    private int weight;

    public EdgeImpl(Vertex<T> target, int weight) {
        this.target = target;
        this.weight = weight;
    }

    @Override
    public Vertex<T> target() {
        return target;
    }

    @Override
    public int getWeight() {
        return weight;
    }
}
