public class Product {
    private String name;
    private String description;
    private double saleAmount;
    private double taxPercentage;

    public Product(String name, String description, double saleAmount, double taxPercentage) {
        this.name = name;
        this.description = description;
        this.saleAmount = saleAmount;
        this.taxPercentage = taxPercentage;
    }

    public double getSaleAmount() {
        return saleAmount;
    }

    public double getTaxPercentage() {
        return taxPercentage;
    }
}

public class Invoice {
    private List<Product> products;
    private double taxThreshold; // El umbral para determinar si se cobra impuesto o no

    public Invoice(double taxThreshold) {
        this.products = new ArrayList<>();
        this.taxThreshold = taxThreshold;
    }

    public void addProduct(Product product) {
        products.add(product);
    }

    public double calculateTotalAmountWithoutTax() {
        return products.stream().mapToDouble(Product::getSaleAmount).sum();
    }

    public double calculateTax() {
        return products.stream()
                .filter(p -> p.getSaleAmount() > taxThreshold)
                .mapToDouble(p -> p.getSaleAmount() * (p.getTaxPercentage() / 100))
                .sum();
    }

    public double calculateTotalAmount() {
        double totalWithoutTax = calculateTotalAmountWithoutTax();
        double totalTax = calculateTax();
        return totalWithoutTax + totalTax;
    }
}

public class Main {
    public static void main(String[] args) {
        // Datos quemados para la demostración
        double taxThreshold = 100.0; // Monto a partir del cual se paga impuestos
        Invoice invoice = new Invoice(taxThreshold);

        // Añadiendo productos a la factura
        invoice.addProduct(new Product("Laptop", "Laptop Description", 1200.0, 15.0));
        invoice.addProduct(new Product("Phone", "Phone Description", 800.0, 10.0));
        invoice.addProduct(new Product("Book", "Book Description", 30.0, 0)); // Este producto no pagará impuestos

        // Cálculo de montos
        double totalWithoutTax = invoice.calculateTotalAmountWithoutTax();
        double tax = invoice.calculateTax();
        double totalWithTax = invoice.calculateTotalAmount();

        // Impresión de resultados
        System.out.println("Total sin impuestos: " + totalWithoutTax);
        System.out.println("Impuestos: " + tax);
        System.out.println("Total con impuestos: " + totalWithTax);
    }
}
