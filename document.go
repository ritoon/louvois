package louvois

import (
	"os"

	"code.google.com/p/gofpdf"
)

// GeneratePaySlip permet de générer un pdf à partir des informations de la Paie
func GeneratePaySlip(s Slip) {
	pdf := gofpdf.New("P", "mm", "A4", "../font")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Bulletin de paie")
	pdf.Output(os.Stdout)
}
