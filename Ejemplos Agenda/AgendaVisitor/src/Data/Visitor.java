package Data;

public interface Visitor<T> {
    public T visitContactoFamiliar(ContactoFamiliar ctx);
    public T visitEventoReunion(EventoReunion ctx);
    //cualquier otra clase nueva en la jerarquía debe visitarse
}
