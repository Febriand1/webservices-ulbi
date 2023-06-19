package controller

type Mahasiswa struct {
	// ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789" `
	Nama         string `bson:"nama,omitempty" json:"nama,omitempty" example:"Tes Swagger"`
	NPM          int    `bson:"npm,omitempty" json:"npm,omitempty" example:"1234567"`
	Phone_Number string `bson:"phonenumber,omitempty" json:"phonenumber,omitempty" example:"08123456789"`
}

type Matakuliah struct {
	// ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Nama_MK  string `bson:"nama_mk,omitempty" json:"nama_mk,omitempty" example:"Pemrograman III"`
	SKS      int    `bson:"sks,omitempty" json:"sks,omitempty" example:"1"`
	Jadwal   Waktu  `bson:"jadwal,omitempty" json:"jadwal,omitempty"`
	Pengampu Dosen  `bson:"pengampu,omitempty" json:"pengampu,omitempty"`
}

type Waktu struct {
	Jam_Masuk string `bson:"jammasuk,omitempty" json:"jammasuk,omitempty" example:"08:00"`
	Jam_Keluar string   `bson:"jamkeluar,omitempty" json:"jamkeluar,omitempty" example:"16:00"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
}

type Dosen struct {
	// ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Dosen    string `bson:"namadosen,omitempty" json:"namadosen,omitempty" example:"Rizkyria"`
	NIK           string `bson:"nik,omitempty" json:"nik,omitempty" example:"123.456.789"`
	Phone_NumberD string `bson:"phonenumberd,omitempty" json:"phonenumberd,omitempty" example:"089876543210"`
}

type Tugas struct {
	Tugas1 int `bson:"tugas1,omitempty" json:"tugas1,omitempty" example:"123"`
	Tugas2 int `bson:"tugas2,omitempty" json:"tugas2,omitempty" example:"123"`
	Tugas3 int `bson:"tugas3,omitempty" json:"tugas3,omitempty" example:"123"`
	Tugas4 int `bson:"tugas4,omitempty" json:"tugas4,omitempty" example:"123"`
	Tugas5 int `bson:"tugas5,omitempty" json:"tugas5,omitempty" example:"123"`
}

type Nilai struct {
	// ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	All_Tugas Tugas      `bson:"alltugas,omitempty" json:"alltugas,omitempty"`
	UTS       int        `bson:"uts,omitempty" json:"uts,omitempty" example:"123"`
	UAS       int        `bson:"uas,omitempty" json:"uas,omitempty" example:"123"`
	Grade     Grade      `bson:"grade,omitempty" json:"grade,omitempty"`
	Kategori  Matakuliah `bson:"kategori,omitempty" json:"kategori,omitempty"`
	Absensi   Presensi   `bson:"absensi,omitempty" json:"absensi,omitempty"`
}

type Grade struct {
	Nama_Grade string `bson:"namagrade,omitempty" json:"namagrade,omitempty" example:"A"`
	Skala      string `bson:"skala,omitempty" json:"skala,omitempty" example:"75.87"`
}

type Presensi struct {
	// ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Jumlah_Kehadiran int       `bson:"jumlahkehadiran,omitempty" json:"jumlahkehadiran,omitempty" example:"1"`
	Biodata          Mahasiswa `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
