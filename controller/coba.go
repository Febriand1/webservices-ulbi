package controller

import (
	inimodel "github.com/Febriand1/Nilai/Model"
	inimodul "github.com/Febriand1/Nilai/Module"
	"github.com/Febriand1/webservices-ulbi/config"
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/Febriand1/webservices-ulbi",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensiold(c *fiber.Ctx) error {
	nl := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(nl)
}

func GetPresensi(c *fiber.Ctx) error {
	nl := inimodul.GetPresensiFromMahasiswa("budiman", config.Ulbimongoconn, "presensi")
	return c.JSON(nl)
}

func GetMatakuliah(c *fiber.Ctx) error {
	nl := inimodul.GetMatakuliahFromJadwal("07.00", config.Ulbimongoconn, "matakuliah")
	return c.JSON(nl)
}

func GetGrade(c *fiber.Ctx) error {
	nl := inimodul.GetGradeFromMahasiswa(121395, config.Ulbimongoconn, "grade")
	return c.JSON(nl)
}

func GetAll(c *fiber.Ctx) error {
	nl := inimodul.GetAllNilaiFromNamaMahasiswa("budiman", config.Ulbimongoconn, "nilai")
	return c.JSON(nl)
}

func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var nilai inimodel.Nilai
	if err := c.BodyParser(&nilai); err != nil {
		return err
	}
	insertedID := inimodul.InsertNilai(db, "nilai",
		nilai.All_Tugas,
		nilai.UTS,
		nilai.UAS,
		nilai.Grade,
		nilai.Kategori,
		nilai.Biodata)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertIdentitas(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var mahasiswa inimodel.Mahasiswa
	if err := c.BodyParser(&mahasiswa); err != nil {
		return err
	}
	insertedID := inimodul.InsertMahasiswa(db, "mahasiswa",
		mahasiswa.Nama,
		mahasiswa.NPM,
		mahasiswa.Phone_Number)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Identitas berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertMatkul(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var matakuliah inimodel.Matakuliah
	if err := c.BodyParser(&matakuliah); err != nil {
		return err
	}
	insertedID := inimodul.InsertMatakuliah(db, "matakuliah",
		matakuliah.Nama_MK,
		matakuliah.SKS,
		matakuliah.Jadwal,
		matakuliah.Pengampu)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Matkul berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDosen(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var dosen inimodel.Dosen
	if err := c.BodyParser(&dosen); err != nil {
		return err
	}
	insertedID := inimodul.InsertDosen(db, "dosen",
		dosen.Nama_Dosen,
		dosen.NIK,
		dosen.Phone_NumberD)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Dosen berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
