package com.cem.aracYikama.strateji;

import com.cem.aracYikama.islem.DisTemizlik;
import com.cem.aracYikama.islem.IcTemizlik;
import com.cem.aracYikama.islem.Temizlik;

public class IcDis {
	private static final IcDis icDis = new IcDis();

	private Temizlik[] yapilacakTemizlikListesi = null;

	private IcDis() {
		yapilacakTemizlikListesi = new Temizlik[2];

		DisTemizlik disTemizlik = new DisTemizlik();
		IcTemizlik icTemizlik = new IcTemizlik();

		yapilacakTemizlikListesi[0] = disTemizlik;
		yapilacakTemizlikListesi[1] = icTemizlik;

		disTemizlik.siradakiIslemAta(icTemizlik);
	}

	public static final Temizlik ilkIslemiSoyle() {
		return icDis.yapilacakTemizlikListesi[0];
	}

}
