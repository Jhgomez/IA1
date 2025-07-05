package main

import (
	"fmt"
	"github.com/snugml/go" // Importa el paquete ml donde está la lógica de LinearRegression
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func createPlot(X, Y, Ypredict []float64) {
	filename := "plot.png"
	p := plot.New()
	p.Title.Text = "Scatter & Prediction"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	data := make(plotter.XYs, len(X))
	pred := make(plotter.XYs, len(X))
	for i := range X {
		data[i].X, data[i].Y = X[i], Y[i]
		pred[i].X, pred[i].Y = X[i], Ypredict[i]
	}

	scatter, err := plotter.NewScatter(data)
	if err != nil {
		log.Fatalf("Scatter create error: %v", err)
	}
	line, err := plotter.NewLine(pred)
	if err != nil {
		log.Fatalf("Line create error: %v", err)
	}

	scatter.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // rojo
	line.LineStyle.Color = color.RGBA{B: 255, A: 255}                 // azul

	p.Add(scatter, line)
	p.Legend.Add("Train", scatter)
	p.Legend.Add("Predicted", line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, filename); err != nil {
		log.Fatalf("Saving file error: %v", err)
	}
}

func main() {
	// Datos de ejemplo (X e Y)
	X := []float64{2, 5, 1, 9, 6, 3, 4}
	y := []float64{6, 7, 5, 8, 9, 6, 5}

	// Instancia de LinearRegression
	model := ml.LinearRegression{}

	// Entrenamiento del modelo
	model.Fit(X, y)

	// Predicciones del modelo
	yPredict := model.Predict(X)

	// Calcular el MSE y R^2
	mse := model.MSE(y, yPredict)
	r2 := model.R2(y, yPredict)

	xNew := []float64{15}
	yNew := model.Predict(xNew)
	fmt.Printf("Predicción para x = 15: %.4f\n", yNew[0]) // Mostramos la predicción

	// Imprimir los resultados
	fmt.Println("X:", X)
	fmt.Println("y:", y)
	fmt.Println("yPredict:", yPredict)
	// MSE, mide en unidades de la variable de salida cuánto se equivocan en promedio (o en total) las predicciones.
	fmt.Printf("MSE/error cuadratico: %.4f\n", mse)
	// Nos indica que tan buen modelo es, el ideal es 1 si es abajo de 0.9 el modelo no es muy bueno
	fmt.Printf("R2/ Coeficiente de determinación: %.4f\n", r2)

	var sumX, sumY, sumXY, sumX2 float64
	n := float64(len(X))

	for i := 0; i < len(X); i++ {
		sumX += X[i]
		sumY += y[i]
		sumXY += X[i] * y[i]
		sumX2 += X[i] * X[i]
	}

	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY / n) - slope*(sumX/n)

	// este coeficiente indica cuánto cambia y por cada unidad que cambia x.
	fmt.Printf("Coeficiente de regresión (pendiente): %.4f\n", slope)

	// este numero indica valor de y cuando x = 0
	fmt.Printf("Intercepto: %.4f\n", intercept)

	createPlot(X, y, yPredict)
}
