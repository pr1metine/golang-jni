public class VierBitAddierer{
	static {
		System.loadLibrary("VierBitAddierer");
    }

    public static class AddiererAusgabe {
        public boolean s,ue;

        @Override
        public String toString() {
            return "s: " + (this.s ? 1:0) + ", ue: " + (this.ue ? 1:0);
        }
    }

    public static native void booleanTestMethod(boolean a);
    public static native boolean nand(boolean a, boolean b);
    public static native boolean xor(boolean a, boolean b);
    public static native AddiererAusgabe ha(boolean a, boolean b);
    public static native boolean[] fa(boolean a, boolean b, boolean c);

	public static void main(String[] args){
        System.out.println("Hello");
        System.out.println(nand(true, true));
        System.out.println();
        booleanTestMethod(false);
        booleanTestMethod(true);
        System.out.println();
        booleanTestMethod(nand(false, false));
        booleanTestMethod(nand(false, true));
        booleanTestMethod(nand(true, false));
        booleanTestMethod(nand(true, true));
        System.out.println();
        booleanTestMethod(xor(false, false));
        booleanTestMethod(xor(false, true));
        booleanTestMethod(xor(true, false));
        booleanTestMethod(xor(true, true));
        System.out.println();
        System.out.println(ha(false, false));
        System.out.println(ha(false, true));
        System.out.println(ha(true, false));
        System.out.println(ha(true, true));
        System.out.println();
        System.out.println(fa(false, false, false));
        System.out.println(fa(false, true, false));
        System.out.println(fa(true, false, false));
        System.out.println(fa(true, true, false));
        System.out.println(fa(false, false, true));
        System.out.println(fa(false, true, true));
        System.out.println(fa(true, false, true));
        System.out.println(fa(true, true, true));
	}
}
