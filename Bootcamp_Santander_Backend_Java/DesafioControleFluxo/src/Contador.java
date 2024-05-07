import java.util.Scanner;

public class Contador {

    public static void main(String[] args) {

        Scanner terminal = new Scanner(System.in);

        int parametroUm = terminal.nextInt();
        int parametroDois = terminal.nextInt();

        try {
            
            contar(parametroUm, parametroDois);

        } catch (ParametroInvalidoException exception) {
            
            System.out.println("O segundo parâmetro deve ser maior que o primeiro.");
        }
        terminal.close();
    }

    static void contar(int parametroUm, int parametroDois) throws ParametroInvalidoException {

        if (parametroUm < parametroDois) {

            int contagem = parametroDois - parametroUm;

            for (int ocorrencias = 1; ocorrencias <= contagem; ocorrencias++) {
                System.out.println("Imprimindo o número " + ocorrencias);
            }

        } else {

            throw new ParametroInvalidoException();
        }
    }
}
