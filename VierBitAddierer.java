public class VierBitAddierer{
	static {
		System.loadLibrary("VierBitAddierer");
    }

    public static native void booleanTestMethod(boolean a);
    public static native boolean nand(boolean a, boolean b);
    public static native boolean xor(boolean a, boolean b);

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
	}
}
