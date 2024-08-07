package tp5.ejercicio1;

import tp5.ejercicio1.adjList.AdjListGraph;

public class GraphMain {
    public static void main(String[] args){
        AdjListGraph<String> grafo;
        Vertex<String> a;
        Vertex<String> b;
        Vertex<String> c;
        Vertex<String> d;
        Vertex<String> e;

        grafo = new AdjListGraph<>();

        a = grafo.createVertex("A");
        b = grafo.createVertex("B");
        c = grafo.createVertex("C");
        d = grafo.createVertex("D");
        e = grafo.createVertex("E");

        grafo.connect(a, b);
        grafo.connect(a, c);

        grafo.connect(b, d);

        grafo.connect(c, d);
        grafo.connect(c, e);

        grafo.connect(d, e);


        System.out.println(grafo.toString());
    }
}
