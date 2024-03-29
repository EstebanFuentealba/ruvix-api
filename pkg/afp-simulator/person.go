package simulator

import (
	"math"
)

type fund struct {
	Fund       string `json:"fund" valid:"-"`
	Balance    int    `json:"balance" valid:"-"`
	Percentage int    `json:"percentage" valid:"-"` //0 -> 100
}

/*
	Account categories

	* mandatory-contributions-account
  * voluntary-contributions-account
  * account-two
  * agreed-contributions-account
*/
type account struct {
	//TODO: check this name
	Name                string  `json:"name" valid:"-"`
	MonthlyContribution float64 `json:"monthly_contribution" valid:"-"`
	Funds               []fund  `json:"funds" valid:"-"`
}

// Person ...
type Person struct {
	CurrentAge    int       `json:"current_age" valid:"-"`
	RetirementAge int       `json:"retirement_age" valid:"-"`
	Gender        string    `json:"gender" valid:"-"`
	Accounts      []account `json:"accounts" valid:"-"`
}

// PersonParams ...
type PersonParams struct {
	Person        Person        `json:"person" valid:"-"`
	InterestFunds InterestFunds `json:"interest_funds" valid:"-"`
}

// InterestFunds All arrays contain three elements 0-> best_case, 1-> argv_case, 2-> worst_case
type InterestFunds struct {
	FundA []float64 `json:"A" valid:"-"`
	FundB []float64 `json:"B" valid:"-"`
	FundC []float64 `json:"C" valid:"-"`
	FundD []float64 `json:"D" valid:"-"`
	FundE []float64 `json:"E" valid:"-"`
}

// IsWoman return true if person is a woman
func (p *Person) IsWoman() bool {
	return p.Gender == "f" || p.Gender == "F"
}

// IsMan return true if person is a man
func (p *Person) IsMan() bool {
	return p.Gender == "m" || p.Gender == "M"
}

// Normalize Function to normalize MonthlyContribution with MAX contributions 74.3 UF or 194182
func (p *Person) Normalize() {
	for ix := range p.Accounts {
		if p.Accounts[ix].MonthlyContribution > 194182 {
			p.Accounts[ix].MonthlyContribution = 194182
		}
	}
}

// Balance ...
func (p *Person) Balance(funds map[string]float64) (float64, Histogram) {

	missingYears := (p.RetirementAge - p.CurrentAge)
	hist := Histogram{Buckets: make([]Bucket, missingYears)}
	var generalBalance float64
	var accountBalance float64
	var fundBalance float64
	var per float64

	for i := 0; i < missingYears; i++ {
		hist.Buckets[i].Year = i
		hist.Buckets[i].Amount = 0.0
	}

	for _, account := range p.Accounts {
		accountBalance = 0.0

		for _, fund := range account.Funds {
			per = float64(fund.Percentage) / 100.0
			fundBalance = float64(fund.Balance)

			for i := 0; i < missingYears; i++ {
				fundBalance = (fundBalance + (float64(12*account.MonthlyContribution) * per)) * (1 + funds[fund.Fund])
				hist.Buckets[i].Amount += fundBalance
			}
			accountBalance += fundBalance
		}
		generalBalance += accountBalance
	}

	return generalBalance, hist
}

// Retirement ...
func (p *Person) Retirement(funds map[string]float64) float64 {
	balance, _ := p.Balance(funds)
	return balance / (12.0 * p.cnu())
}

/************************/
/********  CNU  *********/
/************************/

func (p *Person) cnu() float64 {
	const rf = 0.0312
	T := 110 - p.RetirementAge
	sum := 0.0
	for t := 0; t <= T; t++ {
		sum += p.lx(p.RetirementAge, t) / math.Pow(1+rf, float64(t))
	}

	return sum - 11/24
}

func (p *Person) lx(age int, t int) float64 {
	sum := 1.0

	for i := 0; i <= t; i++ {
		sum *= (1 - p.qx(age+i))
	}
	return sum
}

func (p *Person) qx(age int) float64 {
	var mortalityTable map[int]mortalityPair

	if p.IsWoman() {
		mortalityTable = rvBWomen
	} else if p.IsMan() {
		mortalityTable = rvBMen
	}

	return mortalityTable[age].qx * math.Pow(1-mortalityTable[age].aa, float64(age-p.RetirementAge))
}

type mortalityPair struct {
	qx float64
	aa float64
}

var rvBWomen = map[int]mortalityPair{
	20:  {0.00014697, 0.0216},
	21:  {0.00014834, 0.0216},
	22:  {0.00014994, 0.0216},
	23:  {0.00015244, 0.0216},
	24:  {0.00015595, 0.0216},
	25:  {0.00015701, 0.0255},
	26:  {0.00016123, 0.0255},
	27:  {0.00016800, 0.0255},
	28:  {0.00017809, 0.0255},
	29:  {0.00019099, 0.0255},
	30:  {0.00020182, 0.0308},
	31:  {0.00021777, 0.0308},
	32:  {0.00023368, 0.0308},
	33:  {0.00024883, 0.0308},
	34:  {0.00026424, 0.0308},
	35:  {0.00028251, 0.0301},
	36:  {0.00030354, 0.0301},
	37:  {0.00032823, 0.0301},
	38:  {0.00035730, 0.0301},
	39:  {0.00039088, 0.0301},
	40:  {0.00043346, 0.0277},
	41:  {0.00047663, 0.0277},
	42:  {0.00052397, 0.0277},
	43:  {0.00057551, 0.0277},
	44:  {0.00063222, 0.0277},
	45:  {0.00070640, 0.0241},
	46:  {0.00077939, 0.0241},
	47:  {0.00086039, 0.0241},
	48:  {0.00095007, 0.0241},
	49:  {0.00104953, 0.0241},
	50:  {0.00115484, 0.0256},
	51:  {0.00127943, 0.0256},
	52:  {0.00141455, 0.0256},
	53:  {0.00156014, 0.0256},
	54:  {0.00171951, 0.0256},
	55:  {0.00190027, 0.0250},
	56:  {0.00209962, 0.0250},
	57:  {0.00232771, 0.0250},
	58:  {0.00258931, 0.0250},
	59:  {0.00288425, 0.0250},
	60:  {0.00326841, 0.0208},
	61:  {0.00356919, 0.0208},
	62:  {0.00383856, 0.0208},
	63:  {0.00425648, 0.0208},
	64:  {0.00479267, 0.0208},
	65:  {0.00540752, 0.0209},
	66:  {0.00605835, 0.0209},
	67:  {0.00672943, 0.0209},
	68:  {0.00739141, 0.0209},
	69:  {0.00807320, 0.0209},
	70:  {0.00882683, 0.0221},
	71:  {0.00978717, 0.0221},
	72:  {0.01093797, 0.0221},
	73:  {0.01228757, 0.0221},
	74:  {0.01382224, 0.0221},
	75:  {0.01568726, 0.0202},
	76:  {0.01775670, 0.0202},
	77:  {0.02024799, 0.0202},
	78:  {0.02319710, 0.0202},
	79:  {0.02660428, 0.0202},
	80:  {0.03097353, 0.0161},
	81:  {0.03531713, 0.0161},
	82:  {0.04002585, 0.0161},
	83:  {0.04501651, 0.0161},
	84:  {0.05021994, 0.0161},
	85:  {0.05652689, 0.0121},
	86:  {0.06226553, 0.0121},
	87:  {0.06837755, 0.0121},
	88:  {0.07501311, 0.0121},
	89:  {0.08232981, 0.0121},
	90:  {0.09195911, 0.0081},
	91:  {0.10124796, 0.0081},
	92:  {0.11166539, 0.0081},
	93:  {0.12330934, 0.0081},
	94:  {0.13624260, 0.0081},
	95:  {0.15299191, 0.0040},
	96:  {0.16879145, 0.0040},
	97:  {0.18034215, 0.0040},
	98:  {0.19690391, 0.0040},
	99:  {0.21462541, 0.0040},
	100: {0.23728995, 0.0000},
	101: {0.25766910, 0.0000},
	102: {0.27921041, 0.0000},
	103: {0.30187501, 0.0000},
	104: {0.32560493, 0.0000},
	105: {0.35032245, 0.0000},
	106: {0.37593004, 0.0000},
	107: {0.40231103, 0.0000},
	108: {0.42933114, 0.0000},
	109: {0.45684066, 0.0000},
	110: {1.00000000, 0.0000},
}

var rvBMen = map[int]mortalityPair{
	0:   {0.00527215, 0.0437},
	1:   {0.00026525, 0.0437},
	2:   {0.00022360, 0.0437},
	3:   {0.00019462, 0.0437},
	4:   {0.00014259, 0.0437},
	5:   {0.00011414, 0.0416},
	6:   {0.00010688, 0.0416},
	7:   {0.00010098, 0.0416},
	8:   {0.00009377, 0.0416},
	9:   {0.00008580, 0.0416},
	10:  {0.00008242, 0.0374},
	11:  {0.00008778, 0.0374},
	12:  {0.00011181, 0.0374},
	13:  {0.00015942, 0.0374},
	14:  {0.00022468, 0.0374},
	15:  {0.00032342, 0.0168},
	16:  {0.00039999, 0.0168},
	17:  {0.00047245, 0.0168},
	18:  {0.00053602, 0.0168},
	19:  {0.00059158, 0.0168},
	20:  {0.00063810, 0.0207},
	21:  {0.00069289, 0.0207},
	22:  {0.00073413, 0.0207},
	23:  {0.00075783, 0.0207},
	24:  {0.00076864, 0.0207},
	25:  {0.00076600, 0.0236},
	26:  {0.00077479, 0.0236},
	27:  {0.00078607, 0.0236},
	28:  {0.00080265, 0.0236},
	29:  {0.00082433, 0.0236},
	30:  {0.00084034, 0.0256},
	31:  {0.00086515, 0.0256},
	32:  {0.00089755, 0.0256},
	33:  {0.00093947, 0.0256},
	34:  {0.00098969, 0.0256},
	35:  {0.00104253, 0.0269},
	36:  {0.00110494, 0.0269},
	37:  {0.00116753, 0.0269},
	38:  {0.00122813, 0.0269},
	39:  {0.00129039, 0.0269},
	40:  {0.00134303, 0.0297},
	41:  {0.00142290, 0.0297},
	42:  {0.00151948, 0.0297},
	43:  {0.00163658, 0.0297},
	44:  {0.00177235, 0.0297},
	45:  {0.00191307, 0.0313},
	46:  {0.00207717, 0.0313},
	47:  {0.00224793, 0.0313},
	48:  {0.00242250, 0.0313},
	49:  {0.00260650, 0.0313},
	50:  {0.00283049, 0.0298},
	51:  {0.00306631, 0.0298},
	52:  {0.00321990, 0.0298},
	53:  {0.00349669, 0.0298},
	54:  {0.00386465, 0.0298},
	55:  {0.00432423, 0.0287},
	56:  {0.00480206, 0.0287},
	57:  {0.00531015, 0.0287},
	58:  {0.00586043, 0.0287},
	59:  {0.00645862, 0.0287},
	60:  {0.00725702, 0.0234},
	61:  {0.00794351, 0.0234},
	62:  {0.00864756, 0.0234},
	63:  {0.00938077, 0.0234},
	64:  {0.01019303, 0.0234},
	65:  {0.01132571, 0.0197},
	66:  {0.01251170, 0.0197},
	67:  {0.01392091, 0.0197},
	68:  {0.01552920, 0.0197},
	69:  {0.01730257, 0.0197},
	70:  {0.01926206, 0.0193},
	71:  {0.02137767, 0.0193},
	72:  {0.02373374, 0.0193},
	73:  {0.02638194, 0.0193},
	74:  {0.02934671, 0.0193},
	75:  {0.03319344, 0.0150},
	76:  {0.03680136, 0.0150},
	77:  {0.04066463, 0.0150},
	78:  {0.04479283, 0.0150},
	79:  {0.04923599, 0.0150},
	80:  {0.05473812, 0.0120},
	81:  {0.06014039, 0.0120},
	82:  {0.06614873, 0.0120},
	83:  {0.07286029, 0.0120},
	84:  {0.08036190, 0.0120},
	85:  {0.08980874, 0.0090},
	86:  {0.09920256, 0.0090},
	87:  {0.10955595, 0.0090},
	88:  {0.12087779, 0.0090},
	89:  {0.13315283, 0.0090},
	90:  {0.14812221, 0.0060},
	91:  {0.16233538, 0.0060},
	92:  {0.17732952, 0.0060},
	93:  {0.19300355, 0.0060},
	94:  {0.20924055, 0.0060},
	95:  {0.22835692, 0.0030},
	96:  {0.24670781, 0.0030},
	97:  {0.26583060, 0.0030},
	98:  {0.28567680, 0.0030},
	99:  {0.30618872, 0.0030},
	100: {0.33125700, 0.0000},
	101: {0.35315391, 0.0000},
	102: {0.37549806, 0.0000},
	103: {0.39819949, 0.0000},
	104: {0.42116281, 0.0000},
	105: {0.44428855, 0.0000},
	106: {0.46747463, 0.0000},
	107: {0.49061790, 0.0000},
	108: {0.51361575, 0.0000},
	109: {0.53636771, 0.0000},
	110: {1.00000000, 0.0000},
}
