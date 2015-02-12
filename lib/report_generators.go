package fcompute

type ProfitNLossReport struct {
  Name string     `json:"name"`
  Data []float64  `json:"data"`
}
func (report *ProfitNLossReport) Generate(computed ComputedFData) {
  report.Name = "pnl_report"
  report.Data = computed.PNLData()
}


type NetReturnsReport struct {
  Name string     `json:"name"`
  Data float64    `json:"data"`
}
func (report *NetReturnsReport) Generate(computed ComputedFData) {
  report.Name = "net_returns_report"
  report.Data = computed.NetReturns()
}


type AnnualizedReturnsReport struct {
  Name string     `json:"name"`
  Data float64    `json:"data"`
}
func (report *AnnualizedReturnsReport) Generate(computed ComputedFData) {
  report.Name = "annualized_returns_report"
  report.Data = computed.AnnualizedReturns()
}


type MaximumDrawdownReport struct {
  Name string     `json:"name"`
  Data float64    `json:"data"`
}
func (report *MaximumDrawdownReport) Generate(computed ComputedFData) {
  report.Name = "max_drawdown_report"
  report.Data = computed.MaximumDrawdown()
}
