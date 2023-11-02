// Clase base Contacto
abstract class Contacto {
    private String nombre;
    private String apellido;

    // Constructor, getters y setters
    public Contacto(String nombre, String apellido) {
        this.nombre = nombre;
        this.apellido = apellido;
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

    // Método toString para representar el contacto como un String
    @Override
    public String toString() {
        return "Contacto{" + "nombre='" + nombre + '\'' + ", apellido='" + apellido + '\'' + '}';
    }
}

// Especialización de la clase Contacto para ContactoT1
class ContactoT1 extends Contacto {
    private String email;

    public ContactoT1(String nombre, String apellido, String email) {
        super(nombre, apellido);
        this.email = email;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    @Override
    public String toString() {
        return super.toString() + ", Email: " + email;
    }
}

// Especialización de la clase Contacto para ContactoT2
class ContactoT2 extends Contacto {
    private String telefono;

    public ContactoT2(String nombre, String apellido, String telefono) {
        super(nombre, apellido);
        this.telefono = telefono;
    }

    public String getTelefono() {
        return telefono;
    }

    public void setTelefono(String telefono) {
        this.telefono = telefono;
    }

    @Override
    public String toString() {
        return super.toString() + ", Teléfono: " + telefono;
    }
}

// Clase base para Evento
abstract class Evento {
    private String titulo;
    private String fecha;

    public Evento(String titulo, String fecha) {
        this.titulo = titulo;
        this.fecha = fecha;
    }

    public String getTitulo() {
        return titulo;
    }

    public void setTitulo(String titulo) {
        this.titulo = titulo;
    }

    public String getFecha() {
        return fecha;
    }

    public void setFecha(String fecha) {
        this.fecha = fecha;
    }

    @Override
    public String toString() {
        return "Evento{" + "titulo='" + titulo + '\'' + ", fecha='" + fecha + '\'' + '}';
    }
}

// Especialización de la clase Evento para un tipo de evento específico
class EventoT1 extends Evento {
    private String ubicacion;

    public EventoT1(String titulo, String fecha, String ubicacion) {
        super(titulo, fecha);
        this.ubicacion = ubicacion;
    }

    public String getUbicacion() {
        return ubicacion;
    }

    public void setUbicacion(String ubicacion) {
        this.ubicacion = ubicacion;
    }

    @Override
    public String toString() {
        return super.toString() + ", Ubicación: " + ubicacion;
    }
}

// Clase Agenda que contiene contactos y eventos
class Agenda {
    private List<Object> entradas; // Podría ser List<Contacto> si solo fueran contactos

    public Agenda() {
        entradas = new ArrayList<>();
    }

    public void agregarEntrada(Object entrada) {
        entradas.add(entrada);
    }

    public boolean eliminarEntrada(Object entrada) {
        return entradas.remove(entrada);
    }

    public void mostrarAgenda() {
        for (Object entrada : entradas) {
            System.out.println(entrada);
        }
    }
}

public class Main {
    public static void main(String[] args) {
        // Crear la primera agenda con contactos de tipo 1
        Agenda agenda1 = new Agenda();
        agenda1.agregarEntrada(new ContactoT1("Juan", "Pérez", "juan.perez@example.com"));
        agenda1.agregarEntrada(new ContactoT1("Ana", "García", "ana.garcia@example.com"));

        // Crear la segunda agenda con contactos de tipo 2 y eventos
        Agenda agenda2 = new Agenda();
        agenda2.agregarEntrada(new ContactoT2("Luis", "Martínez", "123456789"));
        agenda2.agregarEntrada(new ContactoT2("Carmen", "Lopez", "987654321"));
        agenda2.agregarEntrada(new EventoT1("Reunión de trabajo", "2023-10-15", "Oficina Central"));

        // Mostrar ambas agendas
        System.out.println("Agenda 1:");
        agenda1.mostrarAgenda();
        System.out.println("\nAgenda 2:");
        agenda2.mostrarAgenda();
    }
}
