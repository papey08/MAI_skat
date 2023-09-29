import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        long a, b, p, timeToCalculate;

        Scanner in = new Scanner(System.in);
        System.out.println("Enter a:");
        a = in.nextLong();
        System.out.println("Enter b:");
        b = in.nextLong();
        System.out.println("Enter p:");
        p = in.nextLong();
        System.out.println("Enter time in seconds:");
        timeToCalculate = in.nextLong();
        System.out.println();

        EllipticCurve ec = new EllipticCurve(a, b, p);

        long timePassed = 0;
        long iter = 1;

        while (timePassed < timeToCalculate) {
            System.out.println("Iteration " + iter);
            ec.setP(ec.getNextPrimeNumber(ec.getP() + iter * 3000));
            timePassed = ec.step();
            iter++;
        }

        System.out.println("Resulting elliptic curve: " + ec);
    }
}
