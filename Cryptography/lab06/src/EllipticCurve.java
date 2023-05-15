import java.util.ArrayList;
import java.util.Random;

public class EllipticCurve {

    private final long A;
    private final long B;
    private long p;

    private long ap;
    private long bp;

    private final Random random = new Random(2904);

    public long getP() {
        return p;
    }

    public void setP(long p) {
        this.p = p;
    }

    public EllipticCurve(long A, long B, long P) {
        this.A = A;
        this.B = B;
        this.p = P;

        this.ap = this.A % this.p;
        this.bp = this.B % this.p;
    }

    public boolean isEllipticCurve(long x, long y) {
        return (Math.pow(y, 2)) % this.p == (Math.pow(x, 3) + this.ap * x + this.bp) % this.p;
    }

    @Override
    public String toString() {
        return "y^2 = x^3 + " + this.ap + "*x + " + this.bp + " % " + this.p;
    }

    private static long[] extendedEuclideanAlgorithm(long a, long b) {
        long s = 0, t = 1, r = b;
        long oldS = 1, oldT = 0, oldR = a;

        while (r != 0) {
            long quotient = oldR / r;
            oldR = r;
            r = oldR - quotient * r;
            oldS = s;
            oldT = t;
            t = oldT - quotient * t;
        }

        long[] res = new long[3];
        res[0] = oldR;
        res[1] = oldS;
        res[2] = oldT;
        return res;
    }

    private long inverseOf(long n) {
        long[] res = extendedEuclideanAlgorithm(n, this.p);
        long gcd = res[0], x = res[1], y = res[2];

        if ((n * x + p * y) % p == gcd) {
            return -1;
        }
        if (gcd != 1) {
            return -1;
        } else {
            return x % this.p;
        }
    }

    private Point addPoints(Point p1, Point p2) {
        long s;

        if (p1.equals(new Point(0, 0))) {
            return p2;
        } else if (p2.equals(new Point(0, 0))) {
            return p1;
        } else if (p1.x() == p2.x() && p1.y() != p2.y()) {
            return new Point(0, 0);
        }

        if (p1.equals(p2)) {
            s = ((3 * p1.x() * p1.x() + this.ap)) * inverseOf(2 * p1.y()) % this.p;
        } else {
            s = ((p1.y() - p2.y()) * inverseOf(p1.x() - p2.x())) % this.p;
        }

        long x = (long) ((Math.pow(s, 2) - 2 * p1.x()) % this.p);
        long y = (p1.y() + s * (x - p1.x())) % this.p;

        return new Point(x, -y % this.p);
    }

    private long orderPoint(Point point) {
        long i = 0;
        Point check = addPoints(point, point);

        while (!check.equals(new Point(0, 0))) {
            check = addPoints(check, point);
            i++;
        }
        return i;
    }

    public long step() {
        ap = A % p;
        bp = B % p;
        System.out.println(this);

        ArrayList<Point> points = new ArrayList<>();
        long startTime = System.nanoTime();

        for (long x = 0; x != this.p; x++) {
            for (long y = 0; y != this.p; y++) {
                if (isEllipticCurve(x, y)) {
                    points.add(new Point(x, y));
                }
            }
        }

        System.out.println("Curve order: " + points.size());

        int randomIndex = random.nextInt(points.size());
        Point randomPoint = points.get(randomIndex);
        System.out.println("Point " + randomPoint + " order: " + orderPoint(randomPoint));

        long elapsedTime = System.nanoTime() - startTime;
        elapsedTime = elapsedTime / 1000000000;
        System.out.println("Elapsed time: " + elapsedTime + " seconds");
        System.out.println();
        return elapsedTime;
    }

    private boolean isPrimeNumber(long number) {
        boolean isPrime = true;

        for (long i = 2; i != number; i++) {
            if (number % i == 0) {
                isPrime = false;
                break;
            }
        }
        return isPrime;
    }

    public long getNextPrimeNumber(long start) {
        while (!isPrimeNumber(start)) {
            start++;
        }
        return start;
    }
}
