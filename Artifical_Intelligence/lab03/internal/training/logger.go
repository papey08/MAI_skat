package training

import (
	"ai_lab3/internal/activation"
	"ai_lab3/internal/loss"
	"ai_lab3/internal/network"
	"ai_lab3/internal/util"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Logger struct {
	w *tabwriter.Writer
}

func NewLogger() *Logger {
	return &Logger{tabwriter.NewWriter(os.Stdout, 16, 0, 3, ' ', 0)}
}

func (p *Logger) Init(n *network.Network) {
	_, _ = fmt.Fprintf(p.w, "Epochs\tElapsed\tLoss %s\t", n.Config.Loss)
	if n.Config.Mode == activation.MultiClass {
		_, _ = fmt.Fprintf(p.w, "Accuracy\t\n---\t---\t---\t---\t\n")
	} else {
		_, _ = fmt.Fprintf(p.w, "\n---\t---\t---\t\n")
	}
}

func crossValidate(n *network.Network, v Pairs) float64 {
	predictions, responses := make([][]float64, len(v)), make([][]float64, len(v))
	for i := range v {
		predictions[i] = n.Predict(v[i].Input)
		responses[i] = v[i].Response
	}
	f, _ := loss.GetLoss(n.Config.Loss)
	return f(predictions, responses)
}

func formatAccuracy(n *network.Network, v Pairs) string {
	if n.Config.Mode == activation.MultiClass {
		correct := 0
		for _, elem := range v {
			est := n.Predict(elem.Input)
			if util.ArgMax(elem.Response) == util.ArgMax(est) {
				correct++
			}
		}
		return fmt.Sprintf("%.2f\t", float64(correct)/float64(len(v)))
	} else {
		return ""
	}
}

func (p *Logger) WriteLog(n *network.Network, v Pairs, d time.Duration, i int) {
	_, _ = fmt.Fprintf(p.w, "%d\t%s\t%.4f\t%s\n",
		i,
		d.String(),
		crossValidate(n, v),
		formatAccuracy(n, v))
	_ = p.w.Flush()
}
