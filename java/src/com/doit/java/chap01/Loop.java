package com.doit.java.chap01;



public class Loop {
    public static void NineByNine() {
//        do {
//            System.out.print("숫자 1~9 : ");
//            n = sc.nextInt();
//        } while (n >= 1 || n <= 9);
        String result = "";
        String top = " |";
        String line = "-+-";
        for(int i=1;i <= 9; i++) {
            top +=  i + " ";
            line += "--";
            result += i + "|";
            for (int j=1; j<=9; j++) {
                result += i*j + " ";
            }
            result += "\n";
        }
        System.out.println(top+"\n" + line + "\n" + result);
    }

    /**
     * *
     * **
     * ***
     * ****
     * *****
     */
    public static void triangleRB() {
        String star = "*****";
        for(int i=0; i<5; i++) {
            System.out.println(star.substring(0, i+1));
        }
    }
    /**
     *     *
     *    **
     *   ***
     *  ****
     * *****
     */
    public static void triangleLB() {
        StringBuilder star = new StringBuilder("    *");
        System.out.println(star);
        for(int i=3; i >= 0; i--) {
            star.replace(i, i+1, "*");
            System.out.println(star);
        }
    }
    /**
     *     *
     *    * *
     *   * * *
     *  * * * *
     * * * * * *
     */
    public static void triangleTree() {
        StringBuilder star = new StringBuilder("    *");
        System.out.println(star);
        for(int i=3; i >= 0; i--) {
            star.replace(i, i+1, "* ");
            System.out.println(star);
        }
    }
    /**
     * *****
     * ****
     * ***
     * **
     * *
     */
    public static void triangleLU() {
        StringBuilder star = new StringBuilder("*****");
        System.out.println(star);
        for(int i=4; i > 0; i--) {
            star.replace(i, i+1, " ");
            System.out.println(star);
        }
    }
    /**
     * *****
     *  ****
     *   ***
     *    **
     *     *
     */
    public static void triangleRU() {
        StringBuilder star = new StringBuilder("*****");
        System.out.println(star);
        for(int i=0; i <4; i++) {
            star.replace(i, i+1, " ");
            System.out.println(star);
        }
    }

    /**
     *    *
     *   ***
     *  *****
     * *******
     */
    public static void spira() {
        StringBuilder star = new StringBuilder("   *");
        System.out.println(star);
        for(int i=2; i >=0; i--) {
            star.replace(i, i+1, "**");
            System.out.println(star);
        }
    }
    /**
     *    1
     *   222
     *  33333
     * 4444444
     */
    public static void npira() {
        StringBuilder star = new StringBuilder("    ");
        for(int i=1 ;i<=4; i++) {
            System.out.print(star.replace(4-i, 5-i, "").toString());
            for(int j=1; j<=2*i-1; j++) {
                System.out.print(i);
            }
            System.out.println();
        }
    }


    public static void main(String[] args) {
        NineByNine();
        triangleRB();
        triangleLB();
        triangleTree();
        triangleLU();
        triangleRU();
        spira();
        npira();

    }
}
