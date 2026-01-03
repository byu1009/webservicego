package bpjs

type AntrolBatalRequest struct {
	KodeBooking         string 	`json:"kodebooking" binding:"required"`
	Keterangan      	string 	`json:"keterangan" binding:"required"`
}
