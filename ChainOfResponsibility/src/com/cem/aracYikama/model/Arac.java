package com.cem.aracYikama.model;

import com.cem.aracYikama.islem.Temizlik;

public abstract class Arac {

	String marka;
	String model;
	
	public abstract Temizlik ilkTemizlikIsleminiVer();

	public Arac(String marka, String model) {
		super();
		this.marka = marka;
		this.model = model;
	}

	public String getMarka() {
		return marka;
	}

	public void setMarka(String marka) {
		this.marka = marka;
	}

	public String getModel() {
		return model;
	}

	public void setModel(String model) {
		this.model = model;
	}

}
