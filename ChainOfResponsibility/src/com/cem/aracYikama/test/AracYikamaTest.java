package com.cem.aracYikama.test;

import com.cem.aracYikama.model.Araba;
import com.cem.aracYikama.model.Arac;
import com.cem.aracYikama.model.Motosiklet;
import com.cem.aracYikama.model.Otobus;

public class AracYikamaTest {

	public static void main(String args[]) {
		Arac otobus = new Otobus("MAN", "2005");
		Arac motosiklet = new Motosiklet("Mondial", "2009");
		Arac araba = new Araba("Porshe", "2014");

		otobus.ilkTemizlikIsleminiVer().basla(otobus);
		motosiklet.ilkTemizlikIsleminiVer().basla(motosiklet);
		araba.ilkTemizlikIsleminiVer().basla(araba);
	}
}
