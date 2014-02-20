package com.cem.aracYikama.strateji;

import com.cem.aracYikama.islem.DisTemizlik;
import com.cem.aracYikama.islem.Temizlik;

public class Dis {

	private static final Dis dis = new Dis();

	private Temizlik[] yapilacakTemizlikListesi = null;

	private Dis() {
		yapilacakTemizlikListesi = new Temizlik[1];
		DisTemizlik disTemizlik = new DisTemizlik();
		yapilacakTemizlikListesi[0] = disTemizlik;
	}

	public static final Temizlik ilkIslemiSoyle() {
		return dis.yapilacakTemizlikListesi[0];
	}
}