import java.util.Scanner;

public class ContaTerminal {
    public static void main(String[] args) throws Exception {
        
        Scanner terminal = new Scanner(System.in);
        
        System.out.println("Por favor, digite o número de conta desejado.");
        int numero = Integer.parseInt(terminal.nextLine());
        System.out.println("Iremos verificar a disponibilidade do número de conta escolhido.");

        System.out.println("Agora, digite o número da agência onde você gostaria de ter conta.");
        String agencia = terminal.nextLine();

        System.out.println("Digite seu nome completo.");
        String nomeCliente = terminal.nextLine();

        System.out.println("Para terminarmos, digite o valor que gostaria de transferir para a conta para ativá-la.");
        double saldo = terminal.nextDouble();

        System.out.printf("Olá %s, obrigado por criar uma conta em nosso banco. Sua agência é %s, conta %d e seu saldo %.2f já está disponível para saque.", nomeCliente, agencia, numero, saldo);

        terminal.close();
    }
}
