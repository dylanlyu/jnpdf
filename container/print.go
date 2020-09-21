package container

import (
	"bytes"
	"image"
	"image/jpeg"
	"jnpdf/schema"
	"jnpdf/service"
	"time"

	"log"

	"github.com/astaxie/beego"
	"github.com/signintech/gopdf"
	"github.com/skip2/go-qrcode"
)

// Print is container print
func Print(id string) (reply schema.Reply) {
	reply.ID = id
	reply.IsError = false
	reply.Message = ""
	reply.UpdateDate = time.Now().Format("2006-01-02 15:04:05")

	certificate, err := service.GetPDFCertificate(id)
	if err != nil {
		reply.IsError = true
		reply.Message = err.Error()
		return reply
	}
	certificateX, err := GetAllX(certificate)
	if err != nil {
		reply.IsError = true
		reply.Message = err.Error()
		return reply
	}

	err = SedPrinter(certificateX, certificate)
	if err != nil {
		reply.IsError = true
		reply.Message = err.Error()
		return reply
	}

	return reply
}

// SedPrinter is exec printer
func SedPrinter(certificateX schema.InformationX, certificate schema.Information) (err error) {
	var pdfPath = beego.AppConfig.String("pdfPath")
	standardX := 147.0
	//QRCode
	qrPng, err := qrcode.Encode(certificate.QRCodeURL, qrcode.Medium, 256)
	if err != nil {
		beego.Error(err.Error())
		return err
	}
	qrJpg, _, _ := image.Decode(bytes.NewReader(qrPng))
	qrBuf := new(bytes.Buffer)
	jpeg.Encode(qrBuf, qrJpg, nil)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 610.1, H: 505.9}})
	pdf.AddPage()
	//pdf.Image("img/background.jpg", 0, 0, &gopdf.Rect{W: 612.2, H: 502})

	qrH, _ := gopdf.ImageHolderByBytes(qrBuf.Bytes())
	pdf.ImageByHolder(qrH, 255, 376, &gopdf.Rect{W: 50, H: 50})

	if certificate.DataImage != "" {
		img, err := service.GetImage(certificate.DataImage)
		if err != nil {
			beego.Error(err.Error())
			return err
		}

		imgH, _ := gopdf.ImageHolderByBytes(img)
		pdf.ImageByHolder(imgH, 26, 355, &gopdf.Rect{W: 100, H: 80})
	}

	err = pdf.AddTTFFont("HanWangHeiLight", "fonts/HanWangHeiLight.ttf")
	if err != nil {
		beego.Error(err.Error())
		return err
	}

	err = pdf.SetFont("HanWangHeiLight", "", 9)
	if err != nil {
		beego.Error(err.Error())
		return err
	}

	pdf.SetX(445)
	pdf.SetY(269)
	pdf.Cell(nil, certificate.ID)

	err = pdf.SetFont("HanWangHeiLight", "", 10)
	if err != nil {
		beego.Error(err.Error())
		return err
	}

	pdf.SetX(certificateX.EnglishConclusionX)
	pdf.SetY(294)
	pdf.Cell(nil, certificate.EnglishConclusion)

	err = pdf.SetFont("HanWangHeiLight", "", 16)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetX(certificateX.ChineseConclusionX)
	pdf.SetY(304)
	pdf.Cell(nil, certificate.ChineseConclusion)

	err = pdf.SetFont("HanWangHeiLight", "", 10)

	if err != nil {
		beego.Error(err.Error())
		return err
	}

	pdf.SetX(standardX)
	pdf.SetY(64)
	pdf.Cell(nil, points(standardX, certificateX.ColorX, certificateX.Point))
	pdf.SetX(certificateX.ColorX)
	pdf.SetY(67)
	pdf.Cell(nil, certificate.Color)

	pdf.SetX(standardX)
	pdf.SetY(82)
	pdf.Cell(nil, points(standardX, certificateX.TransparencyX, certificateX.Point))
	pdf.SetX(certificateX.TransparencyX)
	pdf.SetY(85)
	pdf.Cell(nil, certificate.Transparency)

	pdf.SetX(standardX)
	pdf.SetY(98)
	pdf.Cell(nil, points(standardX, certificateX.ShapeX, certificateX.Point))
	pdf.SetX(certificateX.ShapeX)
	pdf.SetY(101)
	pdf.Cell(nil, certificate.Shape)

	pdf.SetX(standardX)
	pdf.SetY(117)
	pdf.Cell(nil, points(standardX, certificateX.MeasurementX, certificateX.Point))
	pdf.SetX(certificateX.MeasurementX)
	pdf.SetY(120)
	pdf.Cell(nil, certificate.Measurement)

	pdf.SetX(standardX)
	pdf.SetY(135)
	pdf.Cell(nil, points(standardX, certificateX.WeightX, certificateX.Point))
	pdf.SetX(certificateX.WeightX)
	pdf.SetY(138)
	pdf.Cell(nil, certificate.Weight)

	pdf.SetX(standardX)
	pdf.SetY(151)
	pdf.Cell(nil, points(standardX, certificateX.ViewCommentFirstX, certificateX.Point))
	pdf.SetX(certificateX.ViewCommentFirstX)
	pdf.SetY(154)
	pdf.Cell(nil, certificate.ViewCommentFirst)

	pdf.SetX(27)
	pdf.SetY(170)
	pdf.Cell(nil, certificate.ViewCommentSecond)

	pdf.SetX(27)
	pdf.SetY(185)
	pdf.Cell(nil, certificate.ViewCommentThird)

	pdf.SetX(standardX)
	pdf.SetY(229)
	pdf.Cell(nil, points(standardX, certificateX.HardnessX, certificateX.Point))
	pdf.SetX(certificateX.HardnessX)
	pdf.SetY(232)
	pdf.Cell(nil, certificate.Hardness)

	pdf.SetX(standardX)
	pdf.SetY(246)
	pdf.Cell(nil, points(standardX, certificateX.SpecificGravityX, certificateX.Point))
	pdf.SetX(certificateX.SpecificGravityX)
	pdf.SetY(249)
	pdf.Cell(nil, certificate.SpecificGravity)

	pdf.SetX(standardX)
	pdf.SetY(263)
	pdf.Cell(nil, points(standardX, certificateX.RefractiveIndexX, certificateX.Point))
	pdf.SetX(certificateX.RefractiveIndexX)
	pdf.SetY(266)
	pdf.Cell(nil, certificate.RefractiveIndex)

	pdf.SetX(standardX)
	pdf.SetY(281)
	pdf.Cell(nil, points(standardX, certificateX.PolarisCopeX, certificateX.Point))
	pdf.SetX(certificateX.PolarisCopeX)
	pdf.SetY(284)
	pdf.Cell(nil, certificate.PolarisCope)

	pdf.SetX(standardX)
	pdf.SetY(298)
	pdf.Cell(nil, points(standardX, certificateX.FluorescenceX, certificateX.Point))
	pdf.SetX(certificateX.FluorescenceX)
	pdf.SetY(301)
	pdf.Cell(nil, certificate.Fluorescence)

	pdf.SetX(standardX)
	pdf.SetY(315)
	pdf.Cell(nil, points(standardX, certificateX.MagnificationX, certificateX.Point))
	pdf.SetX(certificateX.MagnificationX)
	pdf.SetY(318)
	pdf.Cell(nil, certificate.Magnification)

	pdf.SetX(27)
	pdf.SetY(333)
	pdf.Cell(nil, certificate.AnalysisComment)

	fileName := pdfPath + certificate.ID + ".pdf"

	pdf.WritePdf(fileName)

	//cmd := exec.Command("/bin/sh", "-c", `cat `+ fileName +` | pdftops -level3 -origpagesizes - - | cat - > /dev/usb/lp0`)
	//
	//if err := cmd.Run(); err != nil {
	//	return err
	//}
	//
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	cmd := exec.Command("rm", "-rf", fileName)
	//	cmd.Run()
	//}()

	return nil
}

func GetAllX(certificate schema.Information) (x schema.InformationX, err error) {
	standardX := 300.0
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 610.1, H: 505.9}})
	pdf.AddPage()

	err = pdf.AddTTFFont("HanWangHeiLight", "fonts/HanWangHeiLight.ttf")
	if err != nil {
		beego.Error(err.Error())
		return x, err
	}

	err = pdf.SetFont("HanWangHeiLight", "", 10)
	if err != nil {
		beego.Error(err.Error())
		return x, err
	}

	pdf.SetX(0)
	pdf.Cell(nil, certificate.EnglishConclusion)
	englishConclusionX := pdf.GetX()

	err = pdf.SetFont("HanWangHeiLight", "", 16)
	if err != nil {
		log.Print(err.Error())
		return x, err
	}

	pdf.SetX(0)
	pdf.Cell(nil, certificate.ChineseConclusion)
	chineseConclusionX := pdf.GetX()
	if englishConclusionX < chineseConclusionX {
		x.ChineseConclusionX = 380 + (200-chineseConclusionX)/2
		x.EnglishConclusionX = 380 + (200-chineseConclusionX)/2 + (chineseConclusionX-englishConclusionX)/2
	} else {
		x.EnglishConclusionX = 380 + (200-englishConclusionX)/2
		x.ChineseConclusionX = 380 + (200-englishConclusionX)/2 + (englishConclusionX-chineseConclusionX)/2
	}

	err = pdf.SetFont("HanWangHeiLight", "", 10)

	if err != nil {
		beego.Error(err.Error())
		return x, err
	}

	pdf.SetX(0)
	pdf.Cell(nil, ".")
	x.Point = pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.Color)
	x.ColorX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.Transparency)
	x.TransparencyX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.Shape)
	x.ShapeX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.Measurement)
	x.MeasurementX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.Weight)
	x.WeightX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.ViewCommentFirst)
	x.ViewCommentFirstX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.Hardness)
	x.HardnessX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.SpecificGravity)
	x.SpecificGravityX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.RefractiveIndex)
	x.RefractiveIndexX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.PolarisCope)
	x.PolarisCopeX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.Fluorescence)
	x.FluorescenceX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.Magnification)
	x.MagnificationX = standardX - pdf.GetX()

	pdf.SetX(0)
	pdf.Cell(nil, certificate.AnalysisComment)
	x.AnalysisCommentX = standardX - pdf.GetX()

	pdf.Close()

	return x, nil
}

func points(startX, endX, offsetX float64) (point string) {
	var start, end, offset int
	start = int(startX)
	end = int(endX) - 3
	offset = int(offsetX)
	for i := start; i < end; i = i + offset {
		point = point + "."
	}

	return point
}
