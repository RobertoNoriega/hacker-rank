import java.util.Arrays;
import java.util.Scanner;

public class CompareTheTriplets {

    public static void main(String... args) {
        Scanner in = new Scanner(System.in);
        String[] pointsAlice = in.nextLine().split(" ");
        String[] pointsBob = in.nextLine().split(" ");
        if (pointsAlice.length == pointsBob.length) {
            compareTriplets(pointsAlice, pointsBob);
        }
    }

    private static void compareTriplets(String[] pointsAlice, String[] pointsBob) {
        int i = 0;
        int scoreAlice = 0;
        int scoreBob = 0;
        while (i < pointsAlice.length) {
            try {
                final int alice = Integer.valueOf(pointsAlice[i]);
                final int bob = Integer.valueOf(pointsBob[i]);
                if (alice > bob) {
                    scoreAlice = scoreAlice + 1;
                } else if (alice < bob) {
                    scoreBob = scoreBob + 1;
                }
                i = i + 1;
            } catch (Exception ex) {
                System.err.println(ex.getMessage());
                break;
            }
        }
        System.out.println(scoreAlice + " " + scoreBob);
    }

}