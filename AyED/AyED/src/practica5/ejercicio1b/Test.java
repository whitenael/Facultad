package tp5.ejercicio1b;

import tp5.ejercicio1b.Impl.AdjList.GraphImpl;
import tp5.ejercicio1b.Impl.AdjList.VertexImpl;

public class Test {
    public static void main(String[] args){
        GraphImpl<String> graph = new GraphImpl<>();
        Vertex<String> a;
        Vertex<String> b;
        Vertex<String> c;
        Vertex<String> d;
        Vertex<String> e;

        a = graph.createVertex("A");
        b = graph.createVertex("B");
        c = graph.createVertex("C");
        d = graph.createVertex("D");
        e = graph.createVertex("E");

        graph.connect(a, b);
        graph.connect(a, c);

        graph.connect(b, d);

        graph.connect(c, d);
        graph.connect(c, e);

        System.out.println(graph.toString());

    }
}
