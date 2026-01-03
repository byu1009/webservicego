package bpjs

type AntrolUpdateWaktuRequest struct {
	KodeBooking         string 	`json:"kodebooking" binding:"required"`
	TaskId         		string 	`json:"taskid" binding:"required"`
	Waktu         		int64 	`json:"waktu" binding:"required"`
	JenisResep         	string 	`json:"jenisresep"`
}