package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"webservicego/app/models"

	"gorm.io/gorm"
)

type Umur struct {
	Tahun int
	Bulan int
	Hari  int
}

func ValidTanggal(tanggal string) bool {
    layout := "2006-01-02"
    t, err := time.Parse(layout, tanggal)
    if err != nil {
        return false
    }
    return t.Format(layout) == tanggal
}

func HitungUmur(tanggalLahir string) (string, error) {
	// parse tanggal lahir
	layout := "2006-01-02"
	lahir, err := time.Parse(layout, tanggalLahir)
	if err != nil {
		return "", err
	}

	now := time.Now()

	// selisih tahun, bulan, hari
	years := now.Year() - lahir.Year()
	months := int(now.Month() - lahir.Month())
	days := now.Day() - lahir.Day()

	// adjust jika bulan/hari negatif
	if days < 0 {
		months--
		days += daysInMonth(now.AddDate(0, -1, 0))
	}
	if months < 0 {
		years--
		months += 12
	}

	// return sesuai logika PHP
	if years >= 1 {
		return fmt.Sprintf("%d Th", years), nil
	} else if months >= 1 {
		return fmt.Sprintf("%d Bl", months), nil
	} else {
		return fmt.Sprintf("%d Hr", days), nil
	}
}

func HitungUmur1(tanggalLahir, tanggalDaftar string) Umur {
	// jika tanggal lahir kosong atau 0000-00-00
	if tanggalLahir == "" || tanggalLahir == "0000-00-00" {
		return Umur{0, 0, 0}
	}

	layout := "2006-01-02"

	dob, err := time.Parse(layout, tanggalLahir)
	if err != nil {
		return Umur{0, 0, 0}
	}

	daftar, err := time.Parse(layout, tanggalDaftar)
	if err != nil {
		return Umur{0, 0, 0}
	}

	years := daftar.Year() - dob.Year()
	months := int(daftar.Month() - dob.Month())
	days := daftar.Day() - dob.Day()

	// adjust jika bulan/hari negatif
	if days < 0 {
		months--
		days += daysInMonth(daftar.AddDate(0, -1, 0))
	}
	if months < 0 {
		years--
		months += 12
	}

	return Umur{
		Tahun: years,
		Bulan: months,
		Hari:  days,
	}
}

func daysInMonth(t time.Time) int {
	year, month := t.Year(), t.Month()
	// ambil tanggal 0 bulan berikutnya → otomatis terakhir bulan ini
	firstOfNextMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, t.Location())
	return firstOfNextMonth.Day()
}

func FormatUmur(tanggalLahir, tanggalDaftar string) string {
	u := HitungUmur1(tanggalLahir, tanggalDaftar)
	return fmt.Sprintf("%d Th %d Bl %d Hr", u.Tahun, u.Bulan, u.Hari)
}

func ToYMD(datetime string) (string, error) {
	t, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return "", err
	}

	return t.Format("2006-01-02"), nil
}

func TebakHari(tanggal string) (string, error) {
	// parse tanggal (Y-m-d atau RFC3339 juga aman)
	t, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		// fallback jika format ISO seperti 1981-03-02T00:00:00Z
		t, err = time.Parse(time.RFC3339, tanggal)
		if err != nil {
			return "", err
		}
	}

	hariIndonesia := map[time.Weekday]string{
		time.Sunday:    "Minggu",
		time.Monday:    "Senin",
		time.Tuesday:   "Selasa",
		time.Wednesday: "Rabu",
		time.Thursday:  "Kamis",
		time.Friday:    "Jumat",
		time.Saturday:  "Sabtu",
	}

	namaHari := hariIndonesia[t.Weekday()]

	return strings.ToUpper(namaHari), nil
}

func HitungEstimasiLayanan(
	tglRegistrasi string,
	jamMulai string,
	noReg string,
	estimasiMenit int,
) (int64, error) {

	// gabungkan tanggal + jam
	layout := "2006-01-02 15:04:05"
	datetimeStr := fmt.Sprintf("%s %s", tglRegistrasi, jamMulai)

	jamMulaiTime, err := time.ParseInLocation(layout, datetimeStr, time.Local)
	if err != nil {
		return 0, fmt.Errorf("format tanggal/jam tidak valid")
	}

	// no_reg bisa "001", "002", dll
	noRegInt, err := strconv.Atoi(strings.TrimLeft(noReg, "0"))
	if err != nil {
		return 0, fmt.Errorf("no_reg tidak valid")
	}

	// total estimasi menit
	totalMenit := estimasiMenit * noRegInt

	// tambah menit ke jam mulai
	estimasiTime := jamMulaiTime.Add(time.Duration(totalMenit) * time.Minute)

	// unix timestamp dalam milidetik
	return estimasiTime.UnixMilli(), nil
}

func GenerateNobooking(db *gorm.DB, tglPeriksa string) (string, error) {
    // Parse string ke time.Time
    date, err := time.Parse("2006-01-02", tglPeriksa)
    if err != nil {
        return "", fmt.Errorf("format tanggal tidak valid: %v", err)
    }

    var lastNobooking string
    err = db.Model(&models.ReferensiMobilejknBpjs{}).
        Select("nobooking").
        Where("tanggalperiksa = ?", tglPeriksa).
        Order("nobooking DESC").
        Limit(1).
        Scan(&lastNobooking).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return "", err
    }

    // Ambil angka terakhir dari 6 digit terakhir
    lastNum := 0
    if len(lastNobooking) >= 14 { // YYYYMMDD + 6 digit
        numPart := lastNobooking[8:] // ambil 6 digit terakhir
        if n, e := strconv.Atoi(numPart); e == nil {
            lastNum = n
        }
    }

    newNum := lastNum + 1
    noBooking := fmt.Sprintf("%s%06d", date.Format("20060102"), newNum)
    return noBooking, nil
}
