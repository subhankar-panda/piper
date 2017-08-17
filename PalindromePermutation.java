/*
 * Runtime Complexity : O(n), n is the length of string
 * Space Complexity : O(n)
 */

public class PalindromePermutation {

    private static boolean palindromePerm (String input) {

        int[] array = buildArray(input);
        boolean foundOdd = false;
        for (int i : array) {
            if (i % 2 == 1) {
                if (!foundOdd) {
                    foundOdd = true;
                }
                else {
                    return false;
                }
            }
        }
        return true;
    }

    private static int getCharacterValue (char c) {
        int a = Character.getNumericValue('a');
        int z = Character.getNumericValue('z');
        int input = Character.getNumericValue(c);

        if (input <= z && input >= a) {
            return input - a;
        }
        return -1;
    }

    private static int[] buildArray(String str) {
                
        int[] result = new int[Character.getNumericValue('z') - Character.getNumericValue('a') + 1]; 
        for (char c : str.toCharArray()) {

            int val = getCharacterValue(c);
            if (val != -1) {
                result[val]++;
            }
        }
        return result;
    }

    public static void main(String[] args) {
       System.out.println(palindromePerm(args[0]));
    }
}
